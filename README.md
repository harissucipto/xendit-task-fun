# Software Engineer - Technical Assessment

This repo was created for completing the Software Engineer -
Technical Assessment at Xendit.

The programming language used to develop the service is Golang, and the backend with Postgres and containerized with Docker.

## API Service

In this repo provide APIs to do the following things:

- POST requests to `/orgs/<org-name>/comments` should allow the user to persist comments (in a MongoDB collection or Postgres table) against a given github organization.
- GET requests to `/orgs/<org-name>/comments/` should return an array of all the comments that have been registered against the organization.
- DELETE requests to `/orgs/<org-name>/comments` should soft delete all comments associated with a particular organization. We define a "soft delete" to mean that deleted items should not be returned in GET calls, but should remain in the database for emergency retrieval and audit purposes
- GET requests to `/orgs/<org-name>/members/` should return an array of members of an organization (with their login, avatar url, the numbers of followers they have, and the number of people they're following), sorted in descending order by the number of followers.

## Running the Service with Docker compose

make sure you have the following environment variables set:

```.env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@postgres:5432/xendit?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
GITHUB_ENDPOINT=https://api.github.com
GITHUB_TOKEN=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
```

after that you can run the service with

```cmd

docker compose up

```

## Setup local development

For the local development make sure to changes the following environment variables especially for the DB_SOURCE because we want to point to the our server with the same database endpoint.

```.env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
GITHUB_ENDPOINT=https://api.github.com
GITHUB_TOKEN=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
```

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)
- [Gomock](https://github.com/golang/mock)
- [Makefile](https://www.gnu.org/software/make/manual/make.html)

### Setup infrastructure

- Create the xendit-network

  ```bash
  make network
  ```

- Start postgres container:

  ```bash
  make postgres
  ```

- Create xendit database:

  ```bash
  make createdb
  ```

- Run db migration up all versions:

  ```bash
  make migrateup
  ```

- Run db migration down all versions:

  ```bash
  make migratedown
  ```

### How to generate code

- Generate SQL CRUD with sqlc:

  ```bash
  make sqlc
  ```

- Create a new db migration:

  ```bash
  migrate create -ext sql -dir db/migration -seq <migration_name>
  ```

### How to run

- Run server:

  ```bash
  make server
  ```

- Run test:

  ```bash
  make test
  ```
