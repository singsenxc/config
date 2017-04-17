package config

import (
	"github.com/pelletier/go-toml"
	"time"
)

var c_path = "config.toml"

var c *toml.TomlTree

func init() {
	ReLoad()
}

// SetPath set custom config path
// default is "config.toml".
func SetPath(p string) {
	c_path = p
}

// Reload to reload config file if it's changed.
func ReLoad() error {
	var err error
	c, err = toml.LoadFile(c_path)
	return err
}

// String get a string value
// the type must be string(surrounded by double quotes), or the value will be nil.
//   key: the key in config file, it will be "tablename.key" in table
//   def: is an optional default value
func String(key string, def ...string) string {
	v, ok := c.Get(key).(string)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

// Int64 get an int64 value
// the type int in toml is 64 bit
//   key: the key in config file, it will be "tablename.key" in table
//   def: is an optional default value
func Int64(key string, def ...int64) int64 {
	v, ok := c.Get(key).(int64)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

// Float64 get a float64 value
// the type float in toml is 64 bit
//   key: the key in config file, it will be "tablename.key" in table
//   def: is an optional default value
func Float64(key string, def ...float64) float64 {
	v, ok := c.Get(key).(float64)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

// Bool get a bool value
// the type must be toml's bool in config file, just "true" or "false" (case sensitive)
//   key: the key in config file, it will be "tablename.key" in table
//   def: is an optional default value
func Bool(key string, def ...bool) bool {
	v, ok := c.Get(key).(bool)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

// Time get a time value
// the type must be toml's date in config file, base on "ISO 8601"
// e.g: 1979-05-27T07:32:00Z (it's not a string)
//   key: the key in config file, it will be "tablename.key" in table
//   def: is an optional default value
func Time(key string, def ...time.Time) time.Time {
	v, ok := c.Get(key).(time.Time)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

// Array get an array value
//   key: the key in config file, it will be "tablename.key" in table
//   def: is an optional default value
func Array(key string) []interface{} {
	v, _ := c.Get(key).([]interface{})
	return v
}
