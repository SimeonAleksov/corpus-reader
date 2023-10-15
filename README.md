## Corpus Reader


#### Running the HTTP Server

```shell
$ go run main.go runserver
INFO   [2023-10-15T22:53:00+02:00] [HTTP]: Configured logrus logger.
INFO   [2023-10-15T22:53:00+02:00] [HTTP]: Sucessfully configured http server.
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /healthcheck              --> nu/corpus-reader/infrastructure/router.webEngine.healthcheck.func1 (1 handlers)
[GIN-debug] POST   /counter                  --> nu/corpus-reader/infrastructure/router.webEngine.search.func1 (1 handlers)
INFO   [2023-10-15T22:53:00+02:00] [HTTP]: Starting HTTP Server                          port=8080
```

#### Running the CLI

```shell
$ go run main.go counter --dir ./corpus --word John
INFO   [2023-10-15T22:51:34+02:00] [COUNTER-COMMAND]: Searching for: configuration in ./corpus
INFO   [2023-10-15T22:51:34+02:00] [COUNTER-COMMAND]: count: 12
```

### Benchmarks

Benchmarking

```shell
$ go test -v ./... -bench=. -run=^# -benchtime=100x -benchmem
```
Random small strings and pattern search

```shell
goos: darwin
goarch: arm64
pkg: nu/corpus-reader/application/repository
BenchmarkSearch
BenchmarkSearch-10    	     100	       472.9 ns/op	     248 B/op	       8 allocs/op
```

Searching within the `corpus` directory

```shell
goos: darwin
goarch: arm64
pkg: nu/corpus-reader/application/services
BenchmarkSearchInDirectory
BenchmarkSearchInDirectory-10    	     100	  37824605 ns/op	37264451 B/op	     537 allocs/op
```


### Task Description

Please implement a program that scans through a provided repository and counts how many times a specific word appears in it. The challenge consists of two parts:

#### Part 1 - Command Line Interface (CLI)

Implement a CLI program in Go that accepts a path to a corpus repository and the word that we are interested in:

**Example of Execution:**
```shell
counter --dir ./corpus --word john
```

**Output:**
```
count: 58
```

#### Part 2 - HTTP Server

Expose the functionality from part 1 via an HTTP server.

**Example of Execution:**
```shell
curl -X POST -H "Content-Type: application/json" -d '{"directory":"corpus", "word":"john"}' http://localhost:8080/counter
```

**Output:**
```json
{ "count": 58 }
```

### Remarks

- The solution can be as simple or as complex as you want.
- Aim for GO idiomatic code.
- Prioritize code maintainability.
- Faster is better.
- Include tests, documentation, and benchmarks if necessary.
- Follow Git best practices.
- You are welcome to enhance the challenge in any manner you prefer.
