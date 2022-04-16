#Variables
TAG_VERSION=0.0.0
export TAG_VERSION

local: compose-up

put-data:
	echo 'add raw data in table create ...'
	aws --endpoint-url=http://localhost:4566 dynamodb put-item \
	--table-name financing-offers \
	--item "{\"customerid\":{\"S\":\"6de4c121-a5d8-4cfa-8a40-b06fcc2b2a33\"},\"availabe\":{\"S\":\"7000.00\"},\"tax\":{\"N\":\"0.089\"},\"quota\":{\"N\":\"35\"}}" \
    >> add-data-to-db.json

compose-up:
	echo 'starting docker-compose ...'
	docker-compose -f ./docker/financing-compose.yaml down
	docker-compose -f ./docker/financing-compose.yaml up -d --remove-orphans

compose-down:
	echo 'shutdown docker-compose ...'
	docker-compose -f ./docker/financing-compose.yaml down

del-table:
	echo 'delete tables'
	aws --endpoint-url http://localhost:4566 dynamodb delete-table --table-name financing-offers 

build-image:
	echo 'building new image docker ...'
	echo 'TAG version $(TAG_VERSION)'
	docker build -t bbmello/financing:$TAG_VERSION_CONSORTIUM .

build-container:
	echo 'building container...'
	echo 'version image $(TAG_VERSION)'