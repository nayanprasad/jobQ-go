
# migration

## tool : goose

install
```
brew install goose
```

set environment key
```
GOOSE_MIGRATION_DIR=./internal/storage/postgres/migrations
```

Create a new SQL migration.

```
goose -s create add_some_column sql
```
Created new file: 00001_add_some_column.sql


set environment key
```
GOOSE_DRIVER=DRIVER
GOOSE_DBSTRING=DBSTRING
```

Apply all available migrations.
```
goose up
```


## postgres driver
go get github.com/jackc/pgx/v5 

