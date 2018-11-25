# goHooks
Simple event/hook driven library 

## Installation

1. Download and install it:

```sh
go get -u github.com/JRascall/goHooks
```

2. Import it in your code:

```go
import "github.com/JRascall/goHooks"
```

## Examples

```go
package main

import "github.com/JRascall/goHooks"

func main() {

    hookSys := CreateHookSystem()

    hookSys.On("example", func(args IHookArgs) 
        data := args.Data()
        one := data["One"]
    })

    args := make(map[string]interface{})
    args["One"] = 1
    hookSys.Call("example", CreateHookArgs(args))
}
```
