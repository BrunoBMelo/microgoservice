echo 'creating tables ...'
echo 'in case fail check the file create-table-throughput.txt inside container'
	aws dynamodb create-table --endpoint-url http://localhost:4566 \
    --table-name consortium-offers \
    --attribute-definitions \
        AttributeName=customerid,AttributeType=S \
    --key-schema AttributeName=customerid,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
    >> create-table-throughput.log

echo 'add raw data in table create ...'
    aws --endpoint-url=http://localhost:4566 dynamodb put-item \
	--table-name consortium-offers \
	--item "{\"customerid\":{\"S\":\"6de4c121-a5d8-4cfa-8a40-b06fcc2b2a33\"},\"available\":{\"S\":\"4436.01\"},\"tax\":{\"N\":\"0.089\"},\"quota\":{\"N\":\"35\"}}" \
    >> add-data-to-db.json