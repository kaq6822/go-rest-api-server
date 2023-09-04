package repository

import (
	"go-rest-api-server/domain"
)

var userDB = map[int]domain.User{}

func SaveUser(user domain.User) domain.User {
	userDB[user.ID] = user
	return user
}

func FindUserByID(ID int) domain.User {
	return userDB[ID]
}

func DeleteUserByID(ID int) {
	delete(userDB, ID)
}

func UpdateUserByID(ID int, user domain.User) domain.User {
	userDB[ID] = user
	return user
}
