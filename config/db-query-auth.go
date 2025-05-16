package config

import (
	"go-fiber-vercel/models" // Import the models package
)

func CreateUser(user models.User) (models.User, error) {
	query := `INSERT INTO lynx.users (gameId, serverId, name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, gameId, serverId, name, email, password`
	row, err := ExecuteSQLWithParams(query, user.GameID, user.ServerID, user.Name, user.Email, user.Password)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var newUser models.User
	if row.Next() {
		if err := row.Scan(&newUser.ID, &newUser.GameID, &newUser.ServerID, &newUser.Name, &newUser.Email, &newUser.Password); err != nil {
			return models.User{}, err
		}
	}

	return newUser, nil
}

func GetUserByEmail(email string) (models.User, error) {
	query := `SELECT id, gameId, serverId, name, email, password FROM lynx.users WHERE email = $1`
	row, err := ExecuteSQLWithParams(query, email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.GameID, &user.ServerID, &user.Name, &user.Email, &user.Password); err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, nil
}
