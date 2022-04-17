run-compose:
	echo 'starting composes ...'
	docker-compose -f ./docker/aws-localstack-compose.yaml down --remove-orphans
	docker-compose -f ./docker/aws-localstack-compose.yaml up -d