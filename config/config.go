package config

import (
	"github.com/globalsign/mgo/"
)

//DBconfig stores the configuration for connecting to the database

type DBconfig struct {
	URL      string
	Database string
}

//DB is a global variable for the MongoDB session

var DB *mgo.session

//Connect connects to MongoDB instance

func (c *DBconfig) Connect() {
	session, err := mgo.Dial(c.URL)
	if err != nil {
		panic(err)
	}
	DB = session
}
