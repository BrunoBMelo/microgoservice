#Variables
#Consortium
TAG_VERSION=0.0.0
export TAG_VERSION

set-env:
	echo 'setting variables'

run-dev: compose-up
	docker run -d -p 8081:8081 \
	  -e AWS_DEFAULT_REGION=us-east-1 \
      -e AWS_ACCESS_KEY_ID=FAKE \
      -e AWS_SECRET_ACCESS_KEY=FAKE \
	  --network consortium_app bbmello/consortium:$(tag)

put-data:
	echo 'add raw data in table create ...'
	aws --endpoint-url=http://localhost:4566 dynamodb put-item \
	--table-name consortium-offers \
	--item "{\"customerid\":{\"S\":\"6de4c121-a5d8-4cfa-8a40-b06fcc2b2a33\"},\"availabe\":{\"S\":\"4436.01\"},\"tax\":{\"N\":\"0.089\"},\"quota\":{\"N\":\"35\"}}" \
    >> add-data-to-db.json

compose-up:
	echo 'starting docker-compose ...'
	docker-compose -f ./docker/consortium-compose.yaml down --remove-orphans
	docker-compose -f ./docker/consortium-compose.yaml up -d --remove-orphans

compose-down:
	echo 'shutdown docker-compose ...'
	docker-compose -f ./docker/consortium-compose.yaml down

del-table:
	echo 'delete tables'
	aws --endpoint-url http://localhost:4566 dynamodb delete-table --table-name consortium-offers 

build-image:
	echo 'building new image docker ...'
	echo 'TAG version $(TAG_VERSION)'
	docker build -t bbmello/consortium:$(tag) .

build-container:
	echo 'building container...'
	echo 'version image $(TAG_VERSION)'