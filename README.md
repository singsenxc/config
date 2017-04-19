# config
go configuration base on toml  
dependency: [go-toml](https://github.com/pelletier/go-toml)  

Usage:
```go
package main

import (
 "github.com/x-croz/config"
)

func main() {
	config.SetPath("config.toml") //change config path, or use default path without set path and reload config.
	config.ReLoad()               //reload config
	config.String("foo")
	config.Int64("bar")           //only support int64
	config.Float64("pie")         //only support float64
	config.String("user.name")
}
```
config.toml:
```toml
foo = "foo"
bar = 123
pie = 3.14159
[user]
name = ""
```
