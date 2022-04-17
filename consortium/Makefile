#Variables
#Consortium
TAG_VERSION=0.0.0
export TAG_VERSION

set-env:
	echo 'setting variables'

run-dev:
	docker run -d -p 8081:8081 \
	  -e AWS_DEFAULT_REGION=us-east-1 \
      -e AWS_ACCESS_KEY_ID=FAKE \
      -e AWS_SECRET_ACCESS_KEY=FAKE \
	  --network awsinfralocal_app bbmello/consortium:$(tag)

build-image:
	echo 'building new image docker ...'
	echo 'TAG version $(TAG_VERSION)'
	docker build -t bbmello/consortium:$(tag) .

build-container:
	echo 'building container...'
	echo 'version image $(tag)'