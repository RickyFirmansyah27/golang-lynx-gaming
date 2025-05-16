package config

import (
	"database/sql"
	"errors"
	"go-fiber-vercel/models"
)

func GetUserByEmail(email string) (models.User, error) {
	query := "SELECT id, gameId, serverId, name, nickname, email, password FROM lynx.user WHERE email = $1"

	row, err := ExecuteSQLWithParams(query, email)
	if err != nil {
		return models.User{}, nil
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.GameID, &user.ServerID, &user.Name, &user.Nickname, &user.Email, &user.Password); err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, nil
}

func GetUserByGameID(gameID, serverID string) (models.User, error) {
	query := "SELECT id, gameId, serverId, name, nickname, email, password FROM lynx.user WHERE gameId = $1 AND serverId = $2"

	row, err := ExecuteSQLWithParams(query, gameID, serverID)
	if err != nil {
		return models.User{}, nil
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.GameID, &user.ServerID, &user.Name, &user.Nickname, &user.Email, &user.Password); err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, nil
}

func CreateUser(user models.User) (models.User, error) {
	insertQuery := `INSERT INTO lynx.user (gameId, serverId, name, nickname, email, password) 
					VALUES ($1, $2, $3, $4, $5, $6) 
					RETURNING id, gameId, serverId, name, nickname, email, password`

	row, err := ExecuteSQLWithParams(insertQuery, user.GameID, user.ServerID, user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var newUser models.User
	if !row.Next() {
		return models.User{}, errors.New("failed to create user: no rows returned")
	}

	if err := row.Scan(&newUser.ID, &newUser.GameID, &newUser.ServerID, &newUser.Name, &newUser.Nickname, &newUser.Email, &newUser.Password); err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func UpdateUser(id uint, user models.User) (models.User, error) {
	query := `UPDATE lynx.user 
	          SET name = $1, email = $2
	          WHERE id = $3 
	          RETURNING id, gameId, serverId, name, email`

	row, err := ExecuteSQLWithParams(query, user.Name, user.Email, id)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var updatedUser models.User
	if row.Next() {
		if err := row.Scan(&updatedUser.ID, &updatedUser.GameID, &updatedUser.ServerID, &updatedUser.Name, &updatedUser.Nickname, &updatedUser.Email); err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, sql.ErrNoRows
	}

	return updatedUser, nil
}

func UpdatePassword(id uint, newPassword string) error {
	query := "UPDATE lynx.user SET password = $1 WHERE id = $2"

	_, err := ExecuteSQLWithParams(query, newPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id uint) error {
	query := "DELETE FROM lynx.user WHERE id = $1"

	_, err := ExecuteSQLWithParams(query, id)
	if err != nil {
		return err
	}

	return nil
}
