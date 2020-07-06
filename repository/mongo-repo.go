package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jlamb1/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository interface {
	Save(blogPost *model.BlogPost)
	FindAll() []*model.BlogPost
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "gograph"
	COLLECTION = "blog"
)

func New() BlogRepository {

	MONGODB := "mongodb://root:example@localhost:27017"

	clientOptions := options.Client().ApplyURI(MONGODB)
	clientOptions = clientOptions.SetMaxPoolSize(50)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to MongoDB")

	return &database{
		client: dbClient,
	}
}

func (db *database) Save(blogPost *model.BlogPost) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), blogPost)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) FindAll() []*model.BlogPost {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.BlogPost
	for cursor.Next(context.TODO()) {
		var v *model.BlogPost
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}
