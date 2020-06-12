package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type m bson.M

// SuperHero ...
type SuperHero struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Ability   string             `bson:"ability"`
	Power     int                `bson:"power"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// RemoveAllHeroes ...
func RemoveAllHeroes(db *mongo.Database) interface{} {
	col := db.Collection("heroes")
	filter := bson.M{}
	result, err := col.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatalln("Error on remove all super hero", err)
	}
	return result.DeletedCount
}

// GetAllHeroes ...
func GetAllHeroes(db *mongo.Database) []*SuperHero {
	var heroes []*SuperHero
	col := db.Collection("heroes")
	filter := bson.M{}
	result, err := col.Find(context.TODO(), filter)
	if err != nil {
		log.Fatalln("Error on remove all super hero", err)
	}

	// Iterate through the returned cursor.
	for result.Next(context.TODO()) {
		var hero SuperHero
		err = result.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes
}

// FindOne ...
func FindOne(db *mongo.Database, filter bson.M) SuperHero {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.Collection("heroes")
	result := col.FindOne(ctx, filter)

	var superHero SuperHero
	err := result.Decode(&superHero)
	if err != nil {
		log.Fatalln("Error on Decoding the document", err)
	}

	return superHero
}

// Aggregate ...
func Aggregate(db *mongo.Database, hero SuperHero) []*SuperHero {

	var heroes []*SuperHero
	col := db.Collection("heroes")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// convert string to objectId
	// id, _ := primitive.ObjectIDFromHex("5e3d5aab10417f8d971f0282")
	// pipeLine := []m{
	//   m{"$match": m{"_id": id}},
	// }

	pipeLine := []m{
		m{"$match": m{"_id": hero.ID}},
	}
	result, err := col.Aggregate(ctx, pipeLine)
	if err != nil {
		log.Fatalln("Error on aggregate super hero", err)
	}

	// Iterate through the returned cursor.
	for result.Next(context.TODO()) {
		var hero SuperHero
		err = result.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}

	return heroes
}

// InsertOneHero ...
func InsertOneHero(db *mongo.Database, hero SuperHero) interface{} {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.Collection("heroes")
	result, err := col.InsertOne(ctx, hero)
	if err != nil {
		log.Fatalln("Error insert new super hero", err)
	}
	return result.InsertedID
}

// UpdateOneHero ...
func UpdateOneHero(db *mongo.Database, filter bson.M, hero *SuperHero) interface{} {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.Collection("heroes")
	updateData := bson.M{"$set": hero}
	result, err := col.UpdateOne(ctx, filter, updateData)
	if err != nil {
		log.Fatalln("Error update new super hero", err)
	}
	return result.ModifiedCount
}

// DeleteOneHero ...
func DeleteOneHero(db *mongo.Database, filter bson.M) interface{} {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.Collection("heroes")
	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalln("Error update new super hero", err)
	}
	return result.DeletedCount
}

// GetDatabase ....
func GetDatabase() *mongo.Database {
	options := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	errPing := client.Ping(context.Background(), readpref.Primary())
	if errPing != nil {
		log.Fatal("Couldn't connect to the database", errPing)
	} else {
		log.Println("Connected!")
	}

	db := client.Database("exampleGoDB")
	return db
}

// main function
func main() {
	db := GetDatabase()

	deletedCount := RemoveAllHeroes(db)
	log.Println(deletedCount)

	// style 1
	hero := SuperHero{ID: primitive.NewObjectID(), Name: "Hulk", Ability: "Hammer of God", Power: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	insertedID := InsertOneHero(db, hero)
	log.Println(insertedID)

	hero = SuperHero{ID: primitive.NewObjectID(), Name: "Iron Man", Ability: "Hi-Tech weapon", Power: 90, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	insertedID = InsertOneHero(db, hero)
	log.Println(insertedID)

	hero = SuperHero{ID: primitive.NewObjectID(), Name: "Hawkeye", Ability: "super arrow", Power: 90, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	insertedID = InsertOneHero(db, hero)
	log.Println(insertedID)

	// style 2
	hero = SuperHero{}
	hero.ID = primitive.NewObjectID()
	hero.Name = "Superman"
	hero.Ability = "Eye of laser"
	hero.Power = 100
	hero.CreatedAt = time.Now()
	hero.UpdatedAt = time.Now()
	insertedID = InsertOneHero(db, hero)
	log.Println(insertedID)

	heroUpdate := &hero
	heroUpdate.Name = "Spiderman"
	heroUpdate.Ability = "super sense"

	// example create ObjectID form hex
	id, err := primitive.ObjectIDFromHex("5c71ce5c6e6d43eb6e2e93be")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("example create ObjectID form hex", id)

	// Update
	filter := bson.M{"_id": heroUpdate.ID}
	modifiedCount := UpdateOneHero(db, filter, heroUpdate)
	log.Println(modifiedCount)

	// Delete
	filter = bson.M{"name": "Hawkeye"}
	deletedCount = DeleteOneHero(db, filter)
	log.Println(deletedCount)

	// Find One
	filter = bson.M{"name": "Spiderman"}
	superHero := FindOne(db, filter)
	log.Println(superHero)

	// Aggregate
	aggregateRows := Aggregate(db, superHero)
	buff, _ := json.Marshal(&aggregateRows)
	log.Println(string(buff))

}
