package user

import (
	"github.com/MaxKudIT/messkudi/internal/domain/user"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (us *userStorage) GetUserById(id uuid.UUID) (user.User, error) {
	userp := user.User{}

	const GET_USER_QUERY = "SELECT name, lastname, password, phonenumber FROM users WHERE id = $1"
	if err := us.db.QueryRow(GET_USER_QUERY, id).Scan(&userp.Name, &userp.LastName, &userp.Password, &userp.PhoneNumber); err != nil {
		return user.User{}, err
	}
	return userp, nil
}

func (us *userStorage) SaveUser(userp user.User) error {

	const CREATE_USER_QUERY = "INSERT INTO users (name, lastname, password, phonenumber, id, createdat, updatedat, refreshtoken) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	if _, err := us.db.Exec(CREATE_USER_QUERY, userp.Name, userp.LastName, userp.Password, userp.PhoneNumber, userp.Id, userp.CreatedAt, userp.UpdatedAt, userp.Token.RefreshToken); err != nil {
		return err
	}

	return nil
}

func (us *userStorage) DeleteUser(id uuid.UUID) error {
	const DELETE_USER_QUERY = "DELETE FROM users WHERE id = $1"
	if _, err := us.db.Exec(DELETE_USER_QUERY, id); err != nil {
		return err
	}
	return nil
}

// func (userRep userStorage) UpdateUser(userUpd user.UserUpdate) {
//
//		const UPDATE_USER_QUERY = "UPDATE users SET name = $1, lastname = $2, password = $3, phonenumber = $4"
//		if _, err := userRep.db.Exec(UPDATE_USER_QUERY, userUpd.Name, userUpd.LastName, userUpd.Password, userUpd.PhoneNumber); err != nil {
//			log.Print(exceptions.UpdateUserExc().Error())
//		}
//	}
