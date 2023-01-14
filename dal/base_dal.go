package dal

import (
	"context"
	"fmt"
	"log"

	"github.com/Justin-Willson/BarApp/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_CONNECTION_STRING = "mongodb://root:rootpassword@localhost:27017"
const DATABASE_NAME = "BarApp"

func MakeClient() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(MONGO_CONNECTION_STRING)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func InsertOne[T domain.DatabaseObject](databaseObject T, tableName string) {
	client := MakeClient()
	collection := client.Database(DATABASE_NAME).Collection(tableName)
	insertResult, err := collection.InsertOne(context.TODO(), databaseObject)
	fmt.Println(databaseObject)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

func GetAll[T domain.DatabaseObject](tableName string) []*T {
	client := MakeClient()
	collection := client.Database(DATABASE_NAME).Collection(tableName)

	// Pass these options to the Find method
	findOptions := options.Find()

	var objects []*T

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem T
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		objects = append(objects, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", objects)
	return objects
}

func DeleteById(tableName string, idString string) *mongo.DeleteResult {
	client := MakeClient()
	database := client.Database(DATABASE_NAME)
	collection := database.Collection(tableName)
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		log.Fatal(err)
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return result
}
