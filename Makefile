IMAGE_NAME=mseiwald/cgroup-metrics-reporter
IMAGE_TAG=$(TRAVIS_BUILD_NUMBER)

build-image:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

push-image:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push  $(IMAGE_NAME):$(IMAGE_TAG)
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):latest

.ONESHELL:
