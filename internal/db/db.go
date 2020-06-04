package db

import (
	"context"
	"fmt"

	"github.com/vmwarecloudadvocacy/catalogsvc/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

var (

	// Mongo stores the mongodb connection string information
	//mongo *mgo.DialInfo
	//
	//db *mgo.Database
	//
	//Collection *mgo.Collection
	//collection *mongo.Collection
	dbName string
	collectionName string
	client *mongo.Client
	//Context context.Context

)

// GetEnv accepts the ENV as key and a default string
// If the lookup returns false then it uses the default string else it leverages the value set in ENV variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	logger.Logger.Info("Setting default values for ENV variable " + key)
	return fallback
}

// ConnectDB accepts name of database and collection as a string
func ConnectDB(_dbName string, _collectionName string)  {

	dbName = _dbName
	collectionName = _collectionName

	dbUsername := os.Getenv("CATALOG_DB_USERNAME")
	dbSecret := os.Getenv("CATALOG_DB_PASSWORD")

	// Get ENV variable or set to default value
	dbIP := GetEnv("CATALOG_DB_HOST", "0.0.0.0")
	dbPort := GetEnv("CATALOG_DB_PORT", "27017")

	mongoDBUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", dbUsername, dbSecret, dbIP, dbPort)

	ctx := context.Background()
	_client, error := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBUrl))

	//Session, error := mgo.Dial(mongoDBUrl)

	if error != nil {
		fmt.Printf(error.Error())
		logger.Logger.Fatalf(error.Error())
		os.Exit(1)

	}

	client = _client
	//db = Session.DB(dbName)

	//error = db.Session.Ping()
	error = client.Ping(ctx, readpref.Primary())
	if error != nil {
		logger.Logger.Errorf("Unable to connect to database %s", dbName)
	}

	//Collection = db.C(collectionName)
	//Collection = client.Database(dbName).Collection(collectionName)
	logger.Logger.Info("Connected to database and the collection")

	//return Collection
}

func Collection() *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

// CloseDB accepst Session as input to close Connection to the database
//func CloseDB(ctx context.Context) {
//
//	defer s.Close()
//	logger.Logger.Info("Closed connection to db")
//}
