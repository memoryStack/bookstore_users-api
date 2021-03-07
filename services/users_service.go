package services

import (
	"github.com/backendLearningProjects/bookstore_users-api/domain/users"
	"github.com/backendLearningProjects/bookstore_users-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestError) {
	// get the user from database
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return &user, nil
}

/*
	controllers responsibility is to just extract infos from the request and then pass over the required info to the service handlers.
	and then it's the responsibility to do all the business logic for returning the response for the request.
*/
