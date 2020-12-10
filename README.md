# Overview

This is a sample project to demonstrate how to use Rod to setup an end-to-end testing project.

Use standard Go commands to run the project, such as run `go test` to execute all tests.

## Add a test case

To create a test, first, create a file with suffix "_test.go", such as the [calculator_test.go](calculator_test.go).
Then create a function with a capitalized name, such as

```go
func (t T) Add() {

}
```

If the function name is not capitalized, it will be treated as a helper function, not a test case.

## Filter tests

You can use regex to filter tests you want to run, such as

```bash
go test -run /Add
```

The above will only run the test `Add`.

## Debugging

Same as the tutorial here: [See what's under the hood](https://go-rod.github.io/#/get-started/README?id=see-what39s-under-the-hood)
