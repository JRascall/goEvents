# goEvents
Simple event/hook driven library 

## Installation

1. Download and install it:

```sh
go get -u github.com/JRascall/goEvents
```

2. Import it in your code:

```go
import "github.com/JRascall/goEvents"
```

## Examples

```go
package main

import "github.com/JRascall/goEvents"

func main() {

    eventEmitter := CreateEventEmitter()

    eventEmitter.On("example", func(args IEventArgs) 

    })

    eventEmitter.Call("example", CreateEventArgs(nil))
}
```
