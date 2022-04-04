local: ramp-up-container

add-fake-data:
	echo 'add raw data in table create ...'
	aws --endpoint-url=http://localhost:4566 dynamodb put-item \
	--table-name consortium-offers \
	--item "{\"customerid\":{\"S\":\"6de4c121-a5d8-4cfa-8a40-b06fcc2b2a33\"},\"availabe\":{\"S\":\"4436.01\"},\"tax\":{\"N\":\"0.089\"},\"quota\":{\"N\":\"35\"}}" \
    >> add-data-to-db.json

ramp-up-container:
	echo 'starting docker-compose ...'
	docker-compose -f ./docker/consortium-compose.yaml down
	docker-compose -f ./docker/consortium-compose.yaml up -d --remove-orphans

delete-table:
	echo 'delete tables'
	aws --endpoint-url http://localhost:4566 dynamodb delete-table --table-name consortium-offers 