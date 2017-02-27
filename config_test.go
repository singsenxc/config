package config

import (
	"testing"
)

func Test_String(t *testing.T) {
	t.Log(String("key"))
	t.Log(String("foo"))
	t.Log(String("string"))
	t.Log(String("not_exist_key","default value"))
}

func Test_Int(t *testing.T) {
	t.Log(Int("int"))
	t.Log(Int32("int"))
	t.Log(Int64("int"))
	t.Log(Int("bogus_int"))
}

func Test_Float(t *testing.T){
	t.Log(String("float"))
	t.Log(Float("float"))
	t.Log(Float64("float"))
	t.Log(Float32("float"))
	t.Log(Float("foo"))
}

func Test_Bool(t *testing.T) {
	t.Log(Bool("bool"))
	t.Log(Bool("bool2"))
	t.Log(Bool("bool3"))
	t.Log(Bool("string"))
}