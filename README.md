# Library Management API
## Learning Go
Just a basic REST API using Gin and GORM

### Run App
You'll need to have go installed in your system to run the app with this method.
1. `cd` to the app directory
2. Run following command on your terminal 
```bash
go run main.go
```
#### OR alternatively,
```bash
go build -o main . && ./main
```

### Run App with Docker
Alternatively, you can also run the app using docker. The `Dockerfile` is already present in the project.
```bash
docker run --rm -it $(docker build -q .)
```

#### Alternative, you can also build an image with a tag (to run later without building again)

```bash
docker build -t <TAG_NAME> .
```
and run it 
```bash
docker run -d <TAG_NAME>
```
Don't forget to replace `<TAG_NAME>` with the name of your choice.