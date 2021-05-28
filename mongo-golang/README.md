# Mongo-Golang

This is a project to get familiar with the golang mongo driver, gin, golang packages and integration testing a rest api in golang.
There is lots to improve here e.g. preparing the database with some preset documents to query for the integration test or decoupling the http handler of performing database queries

## Prerequisites

* golang
* docker

## Setup


Download dependencies:
> go mod download

## Run

### Run service

> docker-compose up

After this you can access the api at `localhost:8080/api/v1/books`

If you made changes to code and want to see a difference when running via compose first you have to run
> docker-compose build

and then you can run 

> docker-compose up

### Run tests:

First you need to run the mongo database. For this i just commented out the api service part in the docker-compose file. And then run

> docker-compose up

After this run

> go test -v

You should see the test results in the console
