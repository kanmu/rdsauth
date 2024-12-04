# rdsauth

[![CI](https://github.com/winebarrel/rdsauth/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/rdsauth/actions/workflows/ci.yml)

Generates an auth token used to connect to a db with IAM credentials.

## Usage

```sh
$ MY_DB_HOST=database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com
$ export PGPASSWORD=$(rdsauth postgres://scott@$MY_DB_HOST)
$ psql -h $MY_DB_HOST -U scott
...
postgres=>
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
