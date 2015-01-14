# go-locale

Simple package to determine best locale based on a proper `Accept-Language`
header.

# Install

```
go get github.com/swhite24/go-locale/locale
```

# Usage

```golang
package main

import(
    "github.com/swhite24/go-locale/locale"
)

func main() {
    // Pass a proper accept-language header
    ls := locale.Read("en-US,en;q=0.8")

    // Ask which is best
    l := ls.Best()
    fmt.Println(l)

    // Prints:
    // en_US
}
```

# TODO

Add support for determining best based on a list of supported locale
definitions.
