package spotify

import (
	"encoding/json"
	"errors"
	"log"
	"music-app/models"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Context struct {
	AuthBaseURL  string
	AuthUsername string
	AuthPassword string
	ApiBaseURL   string
}

// Return Spotify Context for spotify connection
func New() Context {
	authUrl := os.Getenv("APP_SPOTIFY_AUTH_BASE_URL")
	user := os.Getenv("APP_SPOTIFY_AUTH_USERNAME")
	pwd := os.Getenv("APP_SPOTIFY_AUTH_PASSWORD")
	apiUrl := os.Getenv("APP_SPOTIFY_API_BASE_URL")

	return Context{AuthBaseURL: authUrl, AuthUsername: user, AuthPassword: pwd, ApiBaseURL: apiUrl}
}

type Token struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

type Error struct {
	Code string `json:"error,omitempty"`
	Desc string `json:"error_description,omitempty"`
}

// Fetch Tracks by ISRC from Spotify
func (ctx Context) FetchTrack(isrc string) (*models.MusicData, error) {
	var musicData models.MusicData

	token, err := ctx.GetToken()
	if err != nil || token == "" {
		log.Println("Error while fetching token from Spotify")
		return nil, err
	}
	baseUrl := ctx.ApiBaseURL
	endpoint := "/v1/search?q=isrc:" + isrc + "&type=track"
	req, err := http.NewRequest(http.MethodGet, baseUrl+endpoint, nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error while fetching tracks from Spotify")
		return nil, err
	}

	if response.StatusCode == 200 {
		err = json.NewDecoder(response.Body).Decode(&musicData)
		if err != nil {
			log.Println("Error while parsing music")
			return nil, err
		}
	} else if response.StatusCode == 400 {
		var errS Error
		err = json.NewDecoder(response.Body).Decode(&errS)
		if err != nil {
			log.Println("Error while parsing music", err)
			return nil, errors.New(errS.Code + " - " + errS.Desc)
		}
	} else {
		return nil, err
	}

	return &musicData, nil
}

// Get Auth Token from Spotify for tacks API access
func (ctx Context) GetToken() (string, error) {
	user := ctx.AuthUsername
	password := ctx.AuthPassword
	baseUrl := ctx.AuthBaseURL
	path := "/api/token"
	var token Token
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, baseUrl+path, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.SetBasicAuth(user, password)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error while fetching tracks from Spotify")
		return "", err
	}

	if response.StatusCode == 200 {
		err = json.NewDecoder(response.Body).Decode(&token)
		if err != nil {
			log.Println("Error while parsing music", err)
			return "", err
		}
	} else if response.StatusCode == 400 {
		var errS Error
		err = json.NewDecoder(response.Body).Decode(&errS)
		if err != nil {
			log.Println("Error while parsing music", err)
			return "", errors.New(errS.Code + " - " + errS.Desc)
		}
	} else {
		return "", err
	}

	return token.AccessToken, nil
}
