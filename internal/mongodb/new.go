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

	mongoDBURL := config.GetDef("MONGO_DB_URL", "mongodb://127.0.0.1:27017")
	if session, err = mgo.Dial(
		mongoDBURL,
	); err != nil {
		panic(fmt.Errorf("Error in establishing database connection on %v , error : %+v",
			mongoDBURL,
			err,
		))
	}

	db = session.DB(config.GetDef("DB_NAME", "to-do"))

	return db
}
