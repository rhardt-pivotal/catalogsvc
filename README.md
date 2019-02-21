# Catalog Service

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
├── entrypoint
│   └── docker-entrypoint.sh
├── go.mod
├── go.sum
├── images
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
    
       export CATALOG_HOST=0.0.0.0
       export CATALOG_PORT=8082

7. Also, export ENV variables related to the database

    ```
    export CATALOG_DB_USERNAME=mongoadmin
    export CATALOG_DB_PASSWORD=secret
    export CATALOG_DB_HOST=0.0.0.0
    ```

8. Run the catalog service

   ```./bin/catalog```


## API

> **Returns list of all catalog items or creates a new product item (POST)**
   
   **'/products' methods=['GET', 'POST']**
   
      For 'GET' Method
      
      Expected JSON Response
   
      {
         "data": [
         {
            "id": "5c61f497e5fdadefe84ff9b9",
            "name": "Yoga Mat",
            "shortDescription": "Limited Edition Mat",
            "description": "Limited edition yoga mat",
            "imageUrl1": "/static/images/yogamat_square.jpg",
            "imageUrl2": "/static/images/yogamat_thumb2.jpg",
            "imageUrl3": "/static/images/yogamat_thumb3.jpg",
            "price": 62.5,
            "tags": [
                "mat"
            ]
          },
          {
            "id": "5c61f497e5fdadefe84ff9ba",
            "name": "Water Bottle",
            "shortDescription": "Best water bottle ever",
            "description": "For all those athletes out there, a perfect bottle to enrich you",
            "imageUrl1": "/static/images/bottle_square.jpg",
            "imageUrl2": "/static/images/bottle_thumb2.jpg",
            "imageUrl3": "/static/images/bottle_thumb3.jpg",
            "price": 34.99,
            "tags": [
                "bottle"
               ]
          }
         ]}
         
         For 'POST' Method
         
         Expected JSON Body with Request - The image(s) should be placed under the images directory as shown above in the tree          structure
         
         {
            "name": "Tracker",
            "shortDescription": "Limited Edition Tracker",
            "description": "Limited edition Tracker with longer description",
            "imageurl1": "/static/images/tracker_square.jpg",
            "imageurl2": "/static/images/tracker_thumb2.jpg",
            "imageurl3": "/static/images/tracker_thumb3.jpg",
            "price": 149.99,
            "tags": [
                "tracker"
             ]

          }
          
          Expected JSON Response 
          
          {
                "message": "Product created successfully!",
                "resourceId": {
                    "id": "5c61f8f81d41c8e94ecaf25f",
                    "name": "Tracker",
                    "shortDescription": "Limited Edition Tracker",
                    "description": "Limited edition Tracker with longer description",
                    "imageUrl1": "/static/images/tracker_square.jpg",
                    "imageUrl2": "/static/images/tracker_thumb2.jpg",
                    "imageUrl3": "/static/images/tracker_thumb3.jpg",
                    "price": 149.99,
                    "tags": [
                        "tracker"
                    ]
                },
                "status": 201
           }
   
   
> **Returns details about a specific product id**

   **'/products/:id' methods=['GET']**
   
      Expected JSON Response
      
      {
       "data": {
           "id": "5c61f497e5fdadefe84ff9b9",
           "name": "Yoga Mat",
           "shortDescription": "Limited Edition Mat",
           "description": "Limited edition yoga mat",
           "imageUrl1": "/static/images/yogamat_square.jpg",
           "imageUrl2": "/static/images/yogamat_square.jpg",
           "imageUrl3": "/static/images/bottle_square.jpg",
           "price": 62.5,
           "tags": [
             "mat"
           ]
         },
       "status": 200
     }
   
  > Retrieve specific image
  
   **'/static/images/:imageName' methods=['GET']**
   
   Expected response is the image
      
  
   
   
