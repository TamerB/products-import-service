# products-import-service
A service to import products, countries, and stocks from csv files

## Overview

This service provides an API to import products, countries, and stocks from CSV files:
```
/readyz                                                         GET
/healthz                                                        GET

/import                                                         POST
```

## Developer builds

#### Generating Server using Swagger

+ This project uses ***Swagger version: v0.26.0***.
+ go-swagger configuration commands are stored in `/Makefile`.
+ To check or install ***go-swagger***, run `make check_install`.
+ To generate server with ***go-swagger*** run `make generate_server`.


#### Builds

From the project's root directory:

```
make generate_server
go get -u all
go mod tidy

make build
```

## Running

```bash
#!/bin/sh

export DB_DRIVER=postgres
export DB_SOURCE=postgresql://<username>:<password>@localhost:5432/<db-name>?sslmode=<enable-or-disable>

make run
```

### Arguments

#### `-v`
Verbose mode, configures the service to output debug level log entries.

### Environment Variables
#### `DB_DRIVER`
Database driver.
#### `DB_SOURCE`
Database source.
#### `DB_CONTAINER_NAME`
Database Container name. This is used by Makefile to Run DB Docker container, create DB, and drop DB.
### `DB_NAME`
Database name


## Curl

After running the application, please go to `http://<host>:<port>/docs` for example CURL requests.

Example: `http://localhost:8080/docs`

### GET Curl requests on terminal:

```
curl -i http://<host>:<port>/readyz
curl -i http://<host>:<port>/healthz
```

#### Examples:
```
curl -i http://localhost:9090/readyz
curl -i http://localhost:9090/healthz
```

#### CSV file columns:
```
country,sku,name,stock_change
```

## Test

To run tests, from the project's root directory, run `make test` in terminal.