# Golang-CRUD-PostgreSQL

Install dependencies
There are 3 packages we are going to use in this project:
Open the terminal inside the go-postgres project.

1. gorilla/mux router
Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

go get -u github.com/gorilla/mux

2. lib/pq driver
A pure Go postgres driver for Go’s database/sql package.

go get github.com/lib/pq

3. joho/godotenv
We are going to use godotenv package to read the .env file. The .env file is used to save the environment variables. Environment variables is used to keep the sensitive data safe. Learn more about how to use environment variables in golang.

go get github.com/joho/godotenv