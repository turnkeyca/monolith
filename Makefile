swagger-check :
	which swagger || (echo "need to install swagger!" && exit 1)

swagger : swagger-check
	swagger generate spec -o ./swagger.yml --scan-models

swagger-client : swagger-check
	rm -rf integration
	mkdir integration
	swagger generate client -f ./swagger.yml --default-scheme=http -t integration

run : monolith
	./monolith

monolith : clean install
	go build

install : go.mod go.sum
	go get

test : 
	go test ./...

clean :
	rm -f monolith
	# go clean -cache -modcache
	