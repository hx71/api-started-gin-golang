GOSOURCEFILE="./main.go"
SWAGDOCS="./docs/swagger"

r:
	go run main.go

run:
	echo "Update Swagger Docs"
	swag init -g ./$(GOSOURCEFILE) -o $(SWAGDOCS)
	nodemon --exec go run $(GOSOURCEFILE)  --signal SIGTERM

swag:
	echo "Create Swagger files"
	swag init -g ./$(GOSOURCEFILE) -o $(SWAGDOCS)