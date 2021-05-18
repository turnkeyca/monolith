.PHONY: db-local-check db-local db-local-cmd db-local-file

swagger-check :
	which swagger || (echo "need to install swagger!" && exit 1)

swagger : swagger-check
	swagger generate spec -o ./swagger.yml --scan-models

run : monolith
	./monolith

monolith : install test
	go build

install : go.mod go.sum
	go get

test : 
	go test ./...

db-local-check :
	which migrate || (echo "need to install migrate!" && exit 1)
	which psql || (echo "need to install postgresql!" && exit 1)
	which docker || (echo "need to install docker!" && exit 1)

db-local-cmd : db-local-check
	psql -h localhost -U postgres -w turnkey -c "$(CMD)"

db-local : db-local-check
	sudo docker run --rm -ti --network host -e POSTGRES_PASSWORD=password postgres

db-local-file : db-local-check
	psql -h localhost -U postgres -w turnkey -f $(FILE)
