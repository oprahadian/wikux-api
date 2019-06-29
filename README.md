# Installing go dep ?
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# How to consume API ?
Wikux.postman_collection.json

# Config ?
cp config/config.go config/config.local.go

# Run env ?
docker-compose up

# Run project ?
cd /go/src/wikux-api && go run .
