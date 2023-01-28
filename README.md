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

This command execute it will start the server. (`http://localhost:8080/manga`)

```
$ manga view
```

You can specify a query path. Default is the date of execution.

ex) `http://localhost:8080/manga?yyyy=2022&mm=1`

## develop
- Golang
- CLI
    - cobra
- ORM
    - ent
