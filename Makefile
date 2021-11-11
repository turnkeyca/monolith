swagger-check :
	which swagger || (echo "need to install swagger!" && exit 1)

swagger : swagger-check
	swagger generate spec -o ./swagger.yml --scan-models

run : monolith
	./monolith

monolith : clean install test
	go build

install : go.mod go.sum
	# go get

test : 
	# go test ./...

clean :
	# rm -f monolith
	# go clean -cache -modcache
	