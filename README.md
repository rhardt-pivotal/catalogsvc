# catalogsvc

## Getting Started

These instructions will allow you to run catalog service

## Requirements

Go (golang) : 1.11.2

mongodb as docker container

zipkin as docker container (optional)

## Instructions

1. Clone this repository 

2. You will notice the following directory structure

``` 
├── catalog-db
│   ├── Dockerfile
│   ├── products.json
│   └── seed.js
├── db.go
├── Dockerfile
├── go.mod
├── go.sum
├── images
├── log.info
├── main.go
├── README.md
└── service.go

```

3. Set GOPATH appropriately as per the documentation - https://github.com/golang/go/wiki/SettingGOPATH
   Also, run ``` export GO111MODULE=on ```

4. Build the go application from the root of the folder

``` go build -o bin/catalog ```

5. Run a mongodb docker container

```sudo docker run -d -p 27017:27017 --name mgo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret -e MONGO_INITDB_DATABASE=acmefit gcr.io/vmwarecloudadvocacy/acmeshop-catalog-db```

6. Export CATALOG_HOST/CATALOG_PORT (port and ip) as ENV variable. You may choose any used port as per your environment setup.
    
    ```export CATALOG_HOST=0.0.0.0```
    ```export CATALOG_PORT=8082```

7. Also, export ENV variables related to the database

    ```
    export CATALOG_DB_USERNAME=mongoadmin
    export CATALOG_DB_PASSWORD=secret
    export CATALOG_DB_HOST=0.0.0.0
    ```

8. Run the catalog service

```./bin/catalog```


### Additional Info
