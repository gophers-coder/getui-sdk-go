PROJECTNAME="getui-sdk-go"



default:
	@echo "Hello World" ${PROJECTNAME}


test:
	@go test -cpu=1,2,4 -v -tags integration ./...

version:
	@go version

login:
	@echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin


.PHONY: default test version login
