set-env:
	echo 'setting variables'

run-dev:
	docker run -d -p 8082:8082 \
	  -e AWS_DEFAULT_REGION=us-east-1 \
      -e AWS_ACCESS_KEY_ID=FAKE \
      -e AWS_SECRET_ACCESS_KEY=FAKE \
	  --network awsinfralocal_app bbmello/financing:$(tag)

build-image:
	echo 'building new image docker ...'
	echo 'TAG version $(tag)'
	docker build -t bbmello/financing:$(tag) .

build-container:
	echo 'building container...'
	echo 'version image $(tag)'