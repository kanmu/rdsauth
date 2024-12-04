# rdsauth

[![CI](https://github.com/winebarrel/rdsauth/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/rdsauth/actions/workflows/ci.yml)

Generates an auth token used to connect to a db with IAM credentials.

## Download

https://github.com/winebarrel/rdsauth/releases/latest

## Usage

### PostgreSQL

```sh
$ MY_DB_HOST=database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com
$ export PGPASSWORD=$(rdsauth postgres://scott@$MY_DB_HOST)
$ psql -h $MY_DB_HOST -U scott
...
postgres=>
```

### MySQL

```sh
$ MY_DB_HOST=database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com
$ export MYSQL_PWD=$(rdsauth mysql://scott@$MY_DB_HOST)
$ psql -h $MY_DB_HOST -u scott --enable-cleartext-plugin
...
mysql>
```

### CNAME support

```sh
$ dig +short cname my-db.example.com
database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com

$ export PGPASSWORD=$(rdsauth postgres://scott@my-db.example.com)
$ psql -h my-db.example.com -U scott
...
postgres=>
```
