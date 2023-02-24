# manga
![](https://raw.githubusercontent.com/ritarock/manga/main/etc/manga.png)

## Usage
### install
```
$ make install
```

### manga update
This subcommand is used to update manga data.
This subcommand must be executed first.

```
$ manga update
```

### manga view
This subcommand is used to view data.

This command execute it will start the server.

```
$ manga view
```

You can specify a query path. Default is the month of execution.

ex) http://localhost:8080/manga?yyyy=2022&mm=1

You can also use GraphQL playground.

http://localhost:8080/playground

## develop
- Golang
- CLI
    - cobra
- ORM
    - ent
- GraphQL
    - gqlgen
