package handlers

import (
	"encoding/json"
	"log"
	"music-app/models"
	"music-app/spotify"
	"music-app/store"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	spotifyClient spotify.Context
	storeCtx      store.Context
}

func New(mainDB *mongo.Client) Context {
	spotifyCtx := spotify.New()
	storeCtx := store.New(mainDB)
	return Context{storeCtx: storeCtx, spotifyClient: spotifyCtx}
}

// CreateMusicData godoc
// @Summary Create music tracks metadata
// @Description Store a new track with the input ISRC
// @Tags tracks
// @Accept  json
// @Produce  json
// @Param isrc query string true "ISRC"
// @Success 200
// @Router /api/v1/metadata [post]
func (ctx Context) CreateMusicData(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Metadata")
	isrc := r.URL.Query().Get("isrc")

	isExist, err := ctx.storeCtx.IsTrackAlreadyExists(isrc)
	if err != nil {
		log.Println("Error while fetching tracks from DB", err)
		respondWithError(w, err, http.StatusBadRequest)
		return
	}

	if isExist {
		log.Println("Track with isrc already exists in DB", isrc)
		err := models.ErrorDetails{Code: "ALREADY_EXIST", Desc: "Track already exist in DB with isrc " + isrc}
		respondWith(w, err, http.StatusBadRequest)
		return
	}

	tracks, err := ctx.spotifyClient.FetchTrack(isrc)
	if err != nil {
		log.Println("Error while fetching tracks", err)
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	item := models.Item{}
	if tracks != nil && len(tracks.TrackDetails.Items) > 0 {
		item = GetHighestPopularityResult(tracks.TrackDetails.Items)
	} else {
		log.Println("Error while fetching tracks", err)
		err := models.ErrorDetails{Code: "NOT_FOUND", Desc: "No new Tracks found with isrc :  " + isrc}
		respondWith(w, err, http.StatusBadRequest)
		return
	}

	err = ctx.storeCtx.InsertTrack(GetTrackDetails(item))
	if err != nil {
		log.Println("Error while inserting tracks", err)
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	respondWith(w, item, http.StatusOK)
}

// Fetch Track By ISRC from Mongo collection
// @Summary Get details of all music tracks
// @Description Get details of all tracks
// @Tags tracks
// @Accept  json
// @Produce  json
// @Param isrc path string true "ISRC"
// @Success 200
// @Router /api/v1/track/{isrc} [get]
func (ctx Context) FetchMusicByIsrc(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Data By ISRC")
	isrc := mux.Vars(r)["isrc"]
	if isrc == "" {
		log.Println("ISRC is blank in path params : ", isrc)
		err := models.ErrorDetails{Code: "INVALID_VALUE", Desc: "Please provide valid ISRC value in path param :" + isrc}
		respondWith(w, err, http.StatusBadRequest)
		return

	}

	track, err := ctx.storeCtx.FetchTrackByISRCFromDB(isrc)
	if err != nil {
		log.Println("Error while fetching tracks from DB", err)
		if strings.Contains(err.Error(), "no documents in result") {
			err := models.ErrorDetails{Code: "NOT_FOUND", Desc: "No tracks found for this isrc :" + isrc}
			respondWith(w, err, http.StatusBadRequest)
			return
		}
		respondWithError(w, err, http.StatusInternalServerError)
		return

	}
	log.Println("Sending success response")
	respondWith(w, track, http.StatusOK)
}

// Search Tracks By Artist from DB
// @Summary Get details of all music tracks by artist
// @Description Get details of all tracks
// @Tags tracks
// @Accept  json
// @Produce  json
// @Param artist query string true "Artist"
// @Success 200
// @Router /api/v1/artist/track [get]
func (ctx Context) FetchMusicByArtist(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get Data By Artist")
	artist := r.URL.Query().Get("artist")
	if artist == "" {
		log.Println("Artist is blank in path params : ", artist)
		err := models.ErrorDetails{Code: "INVALID_VALUE", Desc: "Please provide valid Artist value in path param :" + artist}
		respondWith(w, err, http.StatusBadRequest)
		return

	}
	tracks, err := ctx.storeCtx.FetchTracksByArtistFromDB(artist)
	if err != nil {
		log.Println("Error while fetching tracks from DB", err)
		respondWithError(w, err, http.StatusInternalServerError)
		return

	}
	respondWith(w, tracks, http.StatusOK)
}

func GetHighestPopularityResult(items []models.Item) models.Item {
	item := items[0]
	for _, it := range items {
		if it.Popularity > item.Popularity {
			item = it
		}
	}
	return item
}

func GetTrackDetails(item models.Item) models.Track {
	track := models.Track{}
	track.Title = item.Name
	for _, art := range item.Artists {
		track.Artists = append(track.Artists, art.Name)
	}
	if len(item.Album.Images) > 0 {
		track.ImageURI = item.Album.Images[0].URL
	}
	track.ISRC = item.ExternalIDs.ISRC
	return track
}

func respondWith(w http.ResponseWriter, res interface{}, status int) {
	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Println("Error while marshalling response", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(resBytes))
}

func respondWithError(w http.ResponseWriter, res interface{}, status int) {
	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Println("Error while marshalling response", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(resBytes))
}
