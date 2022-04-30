check_install:
	which swagger ||  GO111MODULE=off go get -u github.com/go-swagger/goswagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models