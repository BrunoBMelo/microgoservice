compose-up:
	echo 'starting composes ...'
	docker-compose -f ./deploy/aws-localstack-compose.yaml up -d --remove-orphans