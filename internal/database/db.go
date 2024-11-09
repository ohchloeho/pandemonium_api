package database

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// DB struct to hold the MongoDB client instance
type DB struct {
    Client *mongo.Client
    Database *mongo.Database
}

// NewDB initializes a new MongoDB connection
func NewDB() (*DB, error) {
    // Define context for timeout on connection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://100.127.215.78:27017")

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }

    // Ping the database to verify connection
    if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }

    log.Println("Connected to MongoDB successfully")
    
    // Return DB struct containing client and database instance
    return &DB{
        Client:   client,
        Database: client.Database("Pandemonium"),
    }, nil
}

// Close closes the MongoDB connection
func (db *DB) Close() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return db.Client.Disconnect(ctx)
}
