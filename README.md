# Elastic on Docker (v7)

This is a small example taken from [this](https://github.com/olivere/elastic-with-docker), 
which shows how to use [Elastic](https://github.com/olivere/elastic-with-docker) for 
Elasticsearch 7.x with [Docker](https://docs.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).

The example is for running Elasticsearch and a test application inside containers with Docker Compose.

## Tech stack
I have tested the example in
- Docker: 19.03.8
- Docker Compose: 1.25.4
- Go: 1.14.2
- Elastic: v7.0.15
- Elasticsearch: 7.6.2

## How to use
To start everything:
```shell script
docker-compose up --build
```
To stop everything:
```shell script
docker-compose down
```

## LICENSE
MIT