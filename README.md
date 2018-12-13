# catalogsvc

## Getting Started

These instructions will allow you to run catalog service

## Requirements

Go (golang) : 1.11+

mongodb 

zipkin

## Instructions

1. Clone this repository 


2. You will notice the following directory structure

├── catalog
│   ├── db.go
│   └── service.go
├── images
│   ├── catsocks_1.jpg
│   ├── cross_1.jpeg
│   ├── puma_1.jpeg
│   ├── weave1.jpg
│   └── youtube_1.jpeg
├── main.go
├── mongo.json
└── README.md

3. Set GOPATH appropriately as per the documentation - https://github.com/golang/go/wiki/SettingGOPATH

4. Build the go application from the root of the folder

``` go build -o catalogsvc ```





### Additional Info