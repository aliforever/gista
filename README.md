# gista
Golang Instagram Private API

The library is ported from PHP (https://github.com/mgp25/Instagram-API), The mother of all private instagram  
libraries. 

## Installation
```sh
go get github.com/aliforever/gista
```
```go
package main
import (
	"fmt"
	"github.com/aliforever/gista"
)

func main() {
    ig,err := gista.New(nil)
    if err != nil {
    	fmt.Println(err)
    	return
    }
    username, password := "",""
    err = ig.Login(username, password, false)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```


Contribute to the library by completing the missing methods, simply copy them from the PHP API (with a tiny change).