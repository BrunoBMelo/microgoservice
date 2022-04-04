run:
	echo starting docker-compose ...
	docker-compose -f ./docker/consortium-compose.yaml down
	docker-compose -f ./docker/consortium-compose.yaml up
	sleep 10

	echo creating table in aws-localstack dynamodb
	aws dynamodb create-table --cli-input-json file://./docker/aws-create-table.json --endpoint-url http://localhost:4566 > ./docker/aws-dynamodb.txt
	sleep 10

	echo add datas in table
	aws --endpoint-url=http://localhost:4566 dynamodb put-item \
	--table-name ConsortiumOffers \
	--item "{\"Id\":{\"S\":\"6de4c121-a5d8-4cfa-8a40-b06fcc2b2a33\"},\"value_offer\":{\"S\":\"4436.01\"},\"tax_offer\":{\"N\":\"0.089\"},\"quota_offer\":{\"N\":\"35\"}}" \
	> ./docker/aws-dynamodb-result.txt
