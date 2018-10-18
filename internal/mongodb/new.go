package mongodb

import (
	"fmt"
	"todo/internal/app/config"

	"gopkg.in/mgo.v2"
)

func New() *mgo.Database {

	var (
		session *mgo.Session
		db      *mgo.Database
		err     error
	)

	if session, err = mgo.Dial(
		config.GetDef("MONGO_DB_URL", "http://localhost:7070"),
	); err != nil {
		panic(fmt.Errorf("Error in establishing database connection %+v", err))
	}

	db = session.DB(config.GetDef("DB_NAME", "to-do"))

	return db
}
