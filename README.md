# Suspect Recall
A game for testing your memory and detective skills!

<p align="center">
  <img src="frontend/public/screenshots/1.png" height="200" style="margin-right: 10px;"  alt="Screenshot 1">
  <img src="frontend/public/screenshots/2.png" height="200" style="margin-right: 10px;" alt="Screenshot 2">
  <img src="frontend/public/screenshots/3.png" height="200" alt="Screenshot 3">
</p>

## Build

```
cd frontend
npm install
npm run build
```

## Run locally 

```
add `REACT_APP_MODE=development` to .env file
cd frontend
npm run start

(another terminal window)
cd ..
go run .
go to http://localhost:3000

(you can install air https://github.com/air-verse/air for live-reload)
```

## Deploy to Google App Engine

```
gcloud app deploy
```

## Tests

```
go test ./...
```