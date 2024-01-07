# music-app
1. Setup App Server Port
   "APP_SERVER_PORT":"8080"

2. SetUp Mongo DB.
   Add ENV Vars below in launch.json 
   "APP_MONGO_MUSIC_DB_NAME":"music-app"
   "APP_MONGO_DB_USERNAME":"**************"
   "APP_MONGO_DB_PASSWORD":"**************"

3. Setup spotify
   Add ENV Vars below in launch.json
   "APP_SPOTIFY_AUTH_BASE_URL":"https://accounts.spotify.com",
   "APP_SPOTIFY_AUTH_USERNAME":"***************",
   "APP_SPOTIFY_AUTH_PASSWORD":"***************",
   "APP_SPOTIFY_API_BASE_URL":"https://api.spotify.com"

4. Open CMD and run commands below
   Run - go mod tidy
   Run - go run /main.go
 
5. Run below API Endpoints
    Create Music Metadata
       [POST] localhost:8080/api/v1/metadata?isrc=USWB11403680
    Find track By ISRC
       [GET] localhost:8080/api/v1/track/{isrc}
    Find tracks by artist
       [GET] localhost:8080/api/v1/artist/track?artist=The Beatles
