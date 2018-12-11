package catalog

import (
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

var (

	// Mongo stores the mongodb connection string information
	mongo *mgo.DialInfo

	db *mgo.Database

	collection *mgo.Collection
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the database.
	MongoDBUrl = "mongodb://mongoadmin:secret@0.0.0.0:27017/?authSource=admin"
)

// ConnectDB accepts name of database and collection as a string
func ConnectDB(dbName string, collectionName string, logger *logrus.Logger) *mgo.Session {

	Session, error := mgo.Dial(MongoDBUrl)

	if error != nil {
		logger.Fatalf("Could not connect to database %s", dbName)
	}

	db = Session.DB(dbName)

	collection = db.C(collectionName)

	return Session
}

// CloseDB accepst Session as input to close Connection to the database
func CloseDB(s *mgo.Session, logger *logrus.Logger) {

	defer s.Close()
	logger.Info("Closed connection to db")
}
