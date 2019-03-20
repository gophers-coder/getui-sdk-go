PROJECTNAME="getui-sdk"



default:
	@echo "Hello World" ${PROJECTNAME}


test:
	@go test -cpu=1,2,4 -v -tags integration ./...



.PHONY: default test
