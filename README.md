# Url Shortener
Simple url shortener service.

#### Technologies used:
* [Golang](https://go.dev)
* [gRPC Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) 
* [Redis](https://redis.io/)
* [Swagger](https://swagger.io)
---

## Getting started
This is an example of how you may give instructions on setting up your project locally. To get a local copy up and running follow these example steps.

### Installation
1. Clone the repository:
```shell
git clone https://github.com/zenorachi/url-shortener
```
2. Setup environment variables (create .env file in the project's root):
```dotenv
export REDIS_HOST=
export REDIS_PORT=
export REDIS_PASSWORD=
export REDIS_INDEX=
```
> **Note:** if you build the project using Docker, setup *DB_HOST=redis* (as the container name)
3. Compile and run the project:
```shell
make
```