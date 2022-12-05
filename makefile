GOSOURCEFILE="./main.go"
SWAGDOCS="./docs/swagger"

r:
	nodemon --exec go run main.go --signal SIGTERM

run:
	echo "Update Swagger Docs"
	swag init -g ./$(GOSOURCEFILE) -o $(SWAGDOCS)
	nodemon --exec go run $(GOSOURCEFILE)  --signal SIGTERM

swag:
	echo "Create Swagger files"
	swag init -g ./$(GOSOURCEFILE) -o $(SWAGDOCS)