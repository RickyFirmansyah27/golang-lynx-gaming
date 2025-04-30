package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var dbPool *pgxpool.Pool
var contextLogger = "[Postgres DB - connection]"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable"
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal("Unable to parse connection string: ", err)
	}

	dbPool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
}

func GetDBPool() *pgxpool.Pool {
	return dbPool
}

func DBConnection() error {
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("unable to acquire connection: %v", err)
	}
	defer conn.Release()

	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		return fmt.Errorf("database connection test failed: %v", err)
	}

	log.Printf("%s | Database connection successfully | Version: %s", contextLogger, version)
	return nil
}

func ExecuteSQLWithParams(sql string, params ...interface{}) (pgx.Rows, error) {
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to acquire connection: %v", err)
	}
	defer conn.Release()

	log.Printf("%s | Executing SQL Query: %s", contextLogger, sql)
	if len(params) > 0 {
		log.Printf("%s | Query Parameters: %v", contextLogger, params)
	}

	rows, err := conn.Query(context.Background(), sql, params...)
	if err != nil {
		log.Printf("%s | Query Failed: %v", contextLogger, err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return rows, nil
}

func StartTransaction(ctx context.Context) (pgx.Tx, error) {
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to acquire connection: %v", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	return tx, nil
}
func ExecuteSQLTransaction(ctx context.Context, tx pgx.Tx, sql string, params ...interface{}) (pgx.Rows, error) {
	rows, err := tx.Query(ctx, sql, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query in transaction: %v", err)
	}

	return rows, nil
}
func RollbackTransaction(ctx context.Context, tx pgx.Tx) error {
	err := tx.Rollback(ctx)
	if err != nil {
		return fmt.Errorf("failed to rollback transaction: %v", err)
	}
	return nil
}
func CommitTransaction(ctx context.Context, tx pgx.Tx) error {
	err := tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}
func Close() {
	if dbPool != nil {
		dbPool.Close()
	}
}
