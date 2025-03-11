# gogator

## Start Postgres server in the background
sudo service postgresql start

## Enter the Shell Environment for postgresql using psql
sudo -u postgres psql

## enter and view tables (gator is the name of the Database)
psql gator;
\dt;

## connect to database
\c gator;

## generate Go Code from SQLC scripts
sqlc generate

## migrating up / down
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" down