# URL Safe, Swear Free Base64 ID Generator

This is a simple tool to generate (youtube alphabet) base64 strings and ensure they do not have block listed words in them.
IDs are guaranteed to start with a letter.

To prevent infinite generation (if you have a vast block list) the tool will panic if it cannot generate a string after 100 attempts.

## Installation

```bash
go get github.com/tadatodo/safebase64
```

## Usage

```go

package main

import (
    "fmt"
    "github.com/tadatodo/safebase64"
)



func main() {
    // will also block all numeric variations like "m01st"
    s := safebase64.New([]string{"moist"})
    id := s.Generate(12)
    fmt.Println(id)
}
```
