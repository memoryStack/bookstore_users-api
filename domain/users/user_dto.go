package users

import (
	"strings"

	"github.com/backendLearningProjects/bookstore_users-api/utils/errors"
)

// User ... struct for user
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// this is a method on the type "User". It's not a normal function
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("invalid email address")
	}
	return nil
}

/*

	the DTO stands for "data transfer object". this file will define the structure of the data transfered by the database
	functions in the DAO file. And we can also define some functions or methods in this files over the struct defined here so that they can
	be reused at other places as well.

*/
