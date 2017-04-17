# config
go configuration base on toml  
dependency: [go-toml](https://github.com/pelletier/go-toml)  

Usage:
```go
package main

import (
 "github.com/singsenxc/config"
)

func main() {
	config.SetPath("config.toml")//change config path
	config.ReLoad() //reload config
	config.String("foo")
	config.Int64("bar")//just support int64
	config.String("user.name")
}
```
config.toml:
```toml
foo = "foo"
bar = 123

[user]
name = ""
```