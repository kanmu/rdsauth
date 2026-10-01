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
  -h, --help               Show help.
      --profile=STRING     AWS credentials profile name.
      --sso-role=STRING    Override sso_role_name for every profile, e.g.
                           ReadOnlyAccess.
      --device-auth        Sign in to IAM Identity Center with the OAuth 2.0
                           device authorization grant (no browser needed).
  -e, --export             Output as environment variable.
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

If the URL has no user, `PGUSER` is used.

```sh
$ export PGUSER=scott
$ $(rdsauth -e postgres://$MY_DB_HOST)
```

### MySQL

```sh
$ MY_DB_HOST=database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com
$ $(rdsauth -e mysql://scott@$MY_DB_HOST)
$ mysql -h $MY_DB_HOST -u scott --enable-cleartext-plugin
...
mysql>
```

### Device authorization grant

On a machine with no browser (e.g. over ssh), `--device-auth` signs in to IAM
Identity Center with the [OAuth 2.0 device authorization grant](https://github.com/winebarrel/awsdag)
instead of resolving credentials the usual way. The Identity Center settings
come from the profile in `~/.aws/config`.

```sh
$ $(rdsauth -e --device-auth --profile dev postgres://scott@$MY_DB_HOST)
Open the following URL in a browser and confirm the code:

  https://device.sso.us-east-1.amazonaws.com/?user_code=ABCD-EFGH
  ABCD-EFGH

$ psql -h $MY_DB_HOST -U scott
```

If the profile does not specify `sso_account_id` or `sso_role_name`, the
account and role are chosen interactively. `--sso-role` overrides the role
here too.

### CNAME support

```sh
$ dig +short cname my-db.example.com
database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com

$ $(rdsauth -e postgres://scott@my-db.example.com)
$ psql -h my-db.example.com -U scott
...
postgres=>
```
