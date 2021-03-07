package users

import (
	"fmt"

	"github.com/backendLearningProjects/bookstore_users-api/utils/date"
	"github.com/backendLearningProjects/bookstore_users-api/utils/errors"
)

// for now let's use DB as a map only. later on will connect it with real database
// userID is used as the primary key here
var (
	usersDB = make(map[int64]*User)
)

// Get ... get user info from database based on the userID
func (user *User) Get() *errors.RestError {
	// we are modifying the user struct on which this method is called
	result := usersDB[user.ID]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Create ... add the user in the database
func (user *User) Create() *errors.RestError {
	// check if user already exists or email passed has already been registered or not
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	user.DateCreated = date.GetNowString()
	usersDB[user.ID] = user
	return nil
}

/*
	DAO means "data access object". this file will contain all the logic of interaction with the database for a domain (in this case user).
	basically these type of files will handle all the CRUD operations for a table in the database. it's helpful to handle the database logic in seperte
	files because in future we might change the way in which we are interacting to the database and it would be easier this way if that logic is
	seperate in files like this.
*/
