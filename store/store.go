package store

import (
	"context"
	"log"
	"music-app/models"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	dbClient *mongo.Collection
}

const (
	collName = "tracks"
)

// Returns Store context for mongo connection
func New(mainDB *mongo.Client) Context {
	dbName := os.Getenv("APP_MONGO_MUSIC_DB_NAME")
	coll := mainDB.Database(dbName).Collection(collName)
	return Context{dbClient: coll}
}

// Insert track in mongo collection for a ISRC
func (ctx Context) InsertTrack(track models.Track) error {
	id, err := ctx.dbClient.InsertOne(context.Background(), track)
	if err != nil {
		return err
	}
	log.Println("Track inserted : Id ", id)
	return nil
}

// Check if track by isrc already exist in mongo collection
func (ctx Context) IsTrackAlreadyExists(isrc string) (bool, error) {
	filter := bson.M{"isrc": isrc}
	count, err := ctx.dbClient.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Find track by ISRC from mongoDB collection
func (ctx Context) FetchTrackByISRCFromDB(isrc string) (models.Track, error) {
	filter := bson.M{"isrc": isrc}
	var track models.Track
	res := ctx.dbClient.FindOne(context.Background(), filter)
	if res.Err() != nil {
		log.Println("Error while fetching track from DB by isrc", res.Err(), isrc)
		return track, res.Err()
	}
	err := res.Decode(&track)
	if err != nil {
		return track, err
	}
	return track, nil
}

// Find tracks by artist from mongoDB collection
func (ctx Context) FetchTracksByArtistFromDB(artist string) ([]models.Track, error) {
	tracks := []models.Track{}
	//filter := bson.M{"artists": artist}
	filter := bson.D{{"artists", bson.D{{"$all", bson.A{artist}}}}}
	cur, err := ctx.dbClient.Find(context.Background(), filter)
	if err != nil {
		log.Println("Error while fetching track from DB by artist", err, artist)
		return nil, err
	}

	err = cur.All(context.TODO(), &tracks)
	if err != nil {
		log.Println("Error while fetching tracks from Cursor", err, artist)
		return nil, err
	}
	return tracks, nil
}
