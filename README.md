# Installing go dep ?
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# How to consume API ?
Wikux.postman_collection.json

# Run env ?
docker-compose -f docker-compose.local.yml up

# Run project ?
cd /go/src/wikux-api && go run .

# Expose local http ?
~/Downloads/ngrok http 10080