# movies-api

`movies-api` is an API that returns information on movies, based on a text search. Currently it uses [omdbapi.com](https://omdbapi.com) to do so.

## Working with movies-api

`movies-api` requires go 1.13 to be built.

Here are some basic go commands you can use to work with movies-api.

* To build: `go build .`
* To run (after building): `./movies-api`
* To execute tests (in all packages): `go test -v ./...`

## Endpoints

Fetch all movies without any kind of sort

```
GET /movies?q=Title of movie to search
```

Fetch all movies sort by year of release and title

```
GET /movies-sorted?q=Title of movie to search
```

NOTE: The API by default listens on port 5432; this is set in the main function.
