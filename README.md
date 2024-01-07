# music-app
1. Setup App Server Port <br/>
   "APP_SERVER_PORT":"8080"

2. SetUp Mongo DB. <br/>
   Add ENV Vars below in launch.json <br/>
   "APP_MONGO_MUSIC_DB_NAME":"music-app" <br/>
   "APP_MONGO_DB_USERNAME":"**************" <br/>
   "APP_MONGO_DB_PASSWORD":"**************" <br/>

3. Setup spotify <br/>
   Add ENV Vars below in launch.json <br/>
   "APP_SPOTIFY_AUTH_BASE_URL":"https://accounts.spotify.com" <br/>
   "APP_SPOTIFY_AUTH_USERNAME":"***************" <br/>
   "APP_SPOTIFY_AUTH_PASSWORD":"***************" <br/>
   "APP_SPOTIFY_API_BASE_URL":"https://api.spotify.com" <br/>

4. Setup Basic Auth creds for API authentication
    "APP_MUSIC_API_AUTH_USERNAME":"mayank90cse" <br/>
    "APP_MUSIC_API_AUTH_PASSWORD":"mayank123456" <br/>

5. Open CMD and run commands below <br/>
   Run - go mod tidy <br/>
   Run - go run /main.go <br/>

6. Swagger initialization Run below commands <br />
   Run - swag init -g main.go <br/>
   Run - go run /main.go <br/>
   Open in Browser - http://localhost:8080/swagger/index.html

7. Run below API Endpoints using Basic Auth <br/>
    Create Music Metadata <br/>
       [POST] localhost:8080/api/v1/metadata?isrc=USWB11403680 <br/>
    Find track By ISRC <br/>
       [GET] localhost:8080/api/v1/track/{isrc} <br/>
    Find tracks by artist <br/>
       [GET] localhost:8080/api/v1/artist/track?artist=The Beatles  <br/>
