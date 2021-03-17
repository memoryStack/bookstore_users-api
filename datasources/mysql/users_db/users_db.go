package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// the params required to start the DB will be stored in the OS as env variables for now
const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

// this function will be called only when we import this file somewhere
// at that time we want to make a connection to the database
func init() {

	fmt.Println(username, password, host, schema)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName) // it rerquires "deriverName" and the "dataSourceName"
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}

/*
	this reason for having this kind of the structure is that in the "datasources" folder we can have
	different different types of the datasources for a microservies like "mysql", "nosql", "cassaderra" etc.
	and under a database we can have different tabels like "users" "e-mails" data etc etc.
	so this pattern is scalable.
*/
