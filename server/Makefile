# Application name, version, build number
APP=cubicasa_test
VERSION=1.0.000.1
# Using git commit hash to identify the build number
BUILD=`git rev-parse HEAD` 
BUILDDATE=`date +%FT%T%z`
REGISTRY_URL=datnguyen.cubicasa.test
IMAGE_NAME=${REGISTRY_URL}/cubicasa_test
IMAGE_FILE=${APP}_${VERSION}_${BUILD}

# Build flags
LDFLAGS=-ldflags "-X main.AppName=${APP} -X main.Version=${VERSION} -X main.BuildNum=${BUILD} -X main.BuildDate=${BUILDDATE}"

# Default target
.DEFAULT_GOAL: ${APP}

# Build application
${APP}: clean
	go build ${LDFLAGS} -o dist/${APP}

# Build application for linux arch
linux: clean
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o dist/${APP}

# Build application for windows arch
windows: clean
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o dist/${APP}.exe

# Build application for docker container
docker:
	[ -f dist/${APP} ] || exit 1
	docker build --pull -t ${IMAGE_NAME}:latest .

docker_save:
	docker save -o dist/${IMAGE_FILE} ${IMAGE_NAME}:latest

docker_load:
	docker load < dist/${IMAGE_FILE}

docker_push:
	docker tag ${IMAGE_NAME}:${VERSION}-${BUILD} ${IMAGE_NAME}:${VERSION}
	docker tag ${IMAGE_NAME}:${VERSION}-${BUILD} ${IMAGE_NAME}:latest
	docker push ${IMAGE_NAME}:${VERSION}-${BUILD}
	docker push ${IMAGE_NAME}:${VERSION}
	docker push ${IMAGE_NAME}:latest

# Install application
install:
	go install ${LDFLAGS}

# Clean application
clean:
	go clean
	if [ -f dist/${APP} ]; then rm dist/${APP}; fi
	if [ -f dist/${APP}.exe ]; then rm dist/${APP}.exe; fi

.PHONY: clean install
