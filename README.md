# go-barebones

A barebones Golang web server.

## Local setup & run

1. Clone repo
2. Create .env in working directory
3. Add env variables listed in the `Config` struct in /pkg/config/config.go

### ... With Docker

1. ```docker build -t go-barebones:latest .```
2. ```docker run -p ${YOUR_PORT}:${YOUR_PORT} go-barebones:latest```

### ... Without Docker

1. ```echo "export ENVIRONMENT=DEV" >> ~/.bash_profile```
2. ```go run main.go```

## Test it out

Head to localhost:${YOUR_PORT} in your browser.. voila :)
