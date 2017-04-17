package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log(String("string"))
	t.Log(Int64("int"))
	t.Log(Float64("float"))
	t.Log(Bool("bool"))
	t.Log(Array("array1"))
	t.Log(Array("array2"))

	//toml table
	t.Log(String("foo.a"))
	t.Log(Int64("foo.b"))
	t.Log(Float64("foo.c"))
	t.Log(Array("foo.d"))
}
