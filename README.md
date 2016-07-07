# golog

this is simple and easy go logging library

## Install

```
go get github.com/wapj/golog
```

## Usage
import the golog and just use it. 
or set some writer

```
package main

import "github.com/wapj/golog"

func main() {

    // not printed below cause default logging level is warning
    golog.Debug("Will not Print")
    golog.Info("Also will not print")
    
    golog.Warn("Warning! warning! this is warning message!")
    golog.Error("Use it carefully for find error")
    golog.Fatal("This is emergency situation! You have to resolve this problem right now!") 
   
}
```

## License

Apache