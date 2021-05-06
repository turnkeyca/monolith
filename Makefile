swagger : swagger-install
	swagger generate spec -o ./swagger.yml --scan-models

swagger-install :
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

run : monolith
	./monolith

monolith : install
	go build

install : go.mod go.sum
	go get