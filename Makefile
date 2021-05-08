.PHONY: swagger swagger-install db-check-local db-create-local db-run-local db-new db-migrate db-unmigrate

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

db-check-local :
	which migrate || (echo "need to install migrate!" && exit 1)
	which psql || (echo "need to install postgresql!" && exit 1)
	which docker || (echo "need to install docker!" && exit 1)

db-create-local : db-check-local
	psql -h localhost -U postgres -w -c "create database turnkey;"

db-run-local : db-check-local
	psql -h localhost -U postgres -w turnkey -c "$(CMD)"

db-start-local : db-check-local
	sudo docker run --rm -ti --network host -e POSTGRES_PASSWORD=password postgres

db-new :
	migrate create -ext sql -dir db/migrations -seq $(SEQ_NAME)

db-migrate :
	migrate -database $(POSTGRESQL_URL) -path db/migrations up $(ARGS)

db-unmigrate :
	migrate -database $(POSTGRESQL_URL) -path db/migrations down $(ARGS)
