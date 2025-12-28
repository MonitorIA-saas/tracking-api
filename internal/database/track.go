package database

import "go.mongodb.org/mongo-driver/mongo"

type TrackRepository struct {
	Collection mongo.Collection
}

func NewTrackRepository(database mongo.Database, collection string) *TrackRepository {
	return &TrackRepository{
		Collection: *database.Collection(collection),
	}
}
