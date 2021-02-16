# DnD

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/dnd.svg)](https://pkg.go.dev/github.com/brittonhayes/dnd)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/brittonhayes/dnd?color=blue&label=Latest%20Version&sort=semver)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/dnd)](https://goreportcard.com/report/github.com/brittonhayes/dnd)
![Test](https://github.com/brittonhayes/dnd/workflows/Test/badge.svg)
[![codecov](https://codecov.io/gh/brittonhayes/dnd/branch/main/graph/badge.svg?token=VN11FU4LBW)](https://codecov.io/gh/brittonhayes/dnd)

<img width="700" src="https://repository-images.githubusercontent.com/328138488/50a49700-522a-11eb-85be-7d8244592f86"></img>

> A Go Client for the Dungeons and Dragons 5e SRD REST API

## Installation

Install with the go get command

```go
go get github.com/brittonhayes/dnd
```

## Documentation

View the full docs on [pkg.go.dev](https://pkg.go.dev/github.com/brittonhayes/dnd)

View the API here https://www.dnd5eapi.co/

## Usage

Using the package is as easy as create client, pick the endpoint, and run the method. This applies across every data
type, so it is consistent across the board. Here's a simple example of how to fetch a rule from the DnD 5e ruleset.

```go
func main() {
    // Create a dnd client
    c := dnd.NewClient()
    
    // Fetch DnD rules about adventuring
    r, _ := c.Rules.Find("adventuring")
    
    // Print out the rule's name
    fmt.Println("Name", r.Name)
}
```

## Examples

For example uses of the package, check out the [example](_example) directory

## Development

If you'd like to contribute to DnD\, make sure you have mage installed: https://magefile.org

```shell
# Download dependencies and run tests
go run main.go download
go test ./...
```

---

> Social image by Ashley Mcnamara https://twitter.com/ashleymcnamara 💖
