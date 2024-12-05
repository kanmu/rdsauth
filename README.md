# rdsauth

[![CI](https://github.com/kanmu/rdsauth/actions/workflows/ci.yml/badge.svg)](https://github.com/kanmu/rdsauth/actions/workflows/ci.yml)

rdsauth is a tool to generate an auth token used to connect to a db with IAM credentials.

## Download

https://github.com/kanmu/rdsauth/releases/latest

## Usage

```
Usage: rdsauth <url> [flags]

Arguments:
  <url>    Database URL

Flags:
  -h, --help       Show help.
  -e, --export     Output as environment variable.
      --version
```

### PostgreSQL

```sh
$ MY_DB_HOST=database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com
$ $(rdsauth -e postgres://scott@$MY_DB_HOST)
$ psql -h $MY_DB_HOST -U scott
...
postgres=>
```

### MySQL

```sh
$ MY_DB_HOST=database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com
$ $(rdsauth -e mysql://scott@$MY_DB_HOST)
$ psql -h $MY_DB_HOST -u scott --enable-cleartext-plugin
...
mysql>
```

### CNAME support

```sh
$ dig +short cname my-db.example.com
database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com

$ $(rdsauth -e postgres://scott@my-db.example.com)
$ psql -h my-db.example.com -U scott
...
postgres=>
```
