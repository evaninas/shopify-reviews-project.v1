# Omnisend's Shopify Reviews scraper v1.

Stores two collections: All reviews collection and Three-word reviews collection with Docker service which can be accessed via API.
## Description

Scrapes reviews of Omnisend app in Shopify (https://apps.shopify.com/omnisend). Creats Docker service, writes data into two collections: All reviews collection and Three-word reviews collection. Has option to access collections via API

## Steps to get a project running
Project uses new go modules (go v1.18).

### Software

* [Docker Desktop](https://docs.docker.com/get-docker/)
* [MongoDB Compass](https://www.mongodb.com/atlas/database) 

### Installing

First clone the repository from (URL) to your desired folder outside of GOPATH.

```
ssh:
git clone git@github.com:evaninas/shopify-reviews-project.git

https:
git clone https://github.com/evaninas/shopify-reviews-project.git
```
### Executing program

* Open Docker software
* Open terminal
* Change directory to main project folder
* Run code, that will initiate docker container, server for database, server for apk, and will start scraping for all reviews
```
docker-compose up --build
```

## Connect to Mongo database with MongoDB Atlass

* New connection
* Advanced Connection Options
* Authentication 

```
URI: mongodb://{username}:{password}@localhost:27017/?authMechanism=DEFAULT
```

<img width="1049" alt="image" src="https://user-images.githubusercontent.com/24540505/202558595-4fbff0ee-e31b-4023-b963-76b576d0c0cf.png">


## Query params
To set limit, skip or sortBy you need to pass those values through query params.

Max number of results returned set to 50.
To set limit, skip or sortBy you need to pass those values through query params.

```
limit  int                        (limits how many results returned)
skip   int                        (skips as many results as provided in skip param)
sortBy string ratingDes/ratingAsc (sorts results by rating in descending or ascending order)
```

Without any params results returned by newest date. Params can be used individually.

Examples

```
localhost:8080
localhost:8080?skip=15
localhost:8080?limit=10&skip=15
localhost:8080?sortBy=ratingDes&skip=15&limit=10
```

## Used third party packages

[Colly](https://github.com/gocolly/colly) - Scraping Framework.

[Mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB supported driver for Go.

[HttpRouter](https://github.com/julienschmidt/httprouter) - HTTP request router (mux).

## TODO to improve project

Refactor bigger functions, unit tests, dependecy injections. 
Get all customer support agent names list, gather frequency in the last 7 and 30 days, 12 month. Use React to implement front-end within scraped and cleaned data.
