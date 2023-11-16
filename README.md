# Proshop V2 with Golang Backend

Kind of a fork of the [ProshopV2](https://github.com/bradtraversy/proshop-v2) of Traversy Media.

Here i am using his frontend and recreating the backend using Go.

The goal here (not funny) is to create all the features that alreay exist in the original project and after that will be a playground for me to add new technologies that i want to apply in a small size project. I am doing this using a microservices approach.

### Work in Early stage

There are a lot of problems in the code, i am aware of it, i just didn't have the time yet. Trying to get thing up and running as fast as i can and than I'll fix the problems.

## Technologies so far

- GO
- Gin Gonic
- PostgreSQL
- Kong
- Tilt for dev
- Docker

<br>

## In the future i want to add

- gRPC -> Communication between microservices
- RabbitMQ with asynchronous requests

## How to run this locally

For the backend you need to have Go and Tilt installed.

To Start:

```bash
$ make start
```

To stop:

```bash
$ make stop
```

for the Frontend you will need to run the following:

```bash
$ cd ./frontend && npm i && npm start
```
