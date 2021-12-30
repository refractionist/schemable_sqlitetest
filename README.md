# Schemable SQLITE test

Runs the schemable test suite against an sqlite3 database.

NOTE: Only works with go 1.18, since schemable uses generics.

```sh
$ git clone https://github.com/refractionist/schemable_sqlitetest
$ cd schemable_sqlitestest
$ go1.18beta1 test -v
```