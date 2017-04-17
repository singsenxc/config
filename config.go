package config

import (
	"github.com/pelletier/go-toml"
	"time"
)

var c_path = "config.toml"

var c *toml.TomlTree

func SetPath(p string) {
	c_path = p
}

func ReLoad() error {
	var err error
	c, err = toml.LoadFile(c_path)
	return err
}

func init() {
	ReLoad()
}

func String(key string, def ...string) string {
	v, ok := c.Get(key).(string)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

func Int64(key string, def ...int64) int64 {
	v, ok := c.Get(key).(int64)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

func Float64(key string, def ...float64) float64 {
	v, ok := c.Get(key).(float64)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

func Bool(key string, def ...bool) bool {
	v, ok := c.Get(key).(bool)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

func Time(key string, def ...time.Time) time.Time {
	v, ok := c.Get(key).(time.Time)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

func Array(key string) []interface{} {
	v, _ := c.Get(key).([]interface{})
	return v
}
