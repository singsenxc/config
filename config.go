package config

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

var mConfig map[string]string

const DEFAULT_CONFIG = "default.config"

func init() {
	if mConfig == nil {
		mConfig = make(map[string]string)
		load(DEFAULT_CONFIG)
	}
}

func Reload(p string) {
	mConfig = make(map[string]string)
	load(p)
}

func load(p string) {
	f, e := os.Open(p)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var err error = nil
	var line []byte
	for err == nil {
		line, _, err = r.ReadLine() //逐行读取
		s := strings.TrimSpace(string(line))
		if len(s) == 0 { //跳过空行
			continue
		}
		if strings.HasPrefix(s, "//") {
			continue
		}
		i := strings.Index(s, "=")
		commentIndex := strings.Index(s, "//")
		if commentIndex < 0 {
			commentIndex = len(s)
		}
		if i > 0 {
			key := s[:i]
			v := s[i+1 : commentIndex]
			mConfig[strings.TrimSpace(key)] = strings.TrimSpace(v)
		}
	}
}

func String(key string, def ...interface{}) string {
	v := mConfig[key]
	if v == "" && len(def) > 0 {
		v, _ = def[0].(string)
	}
	return v
}

func Int(key string, def ...interface{}) int {
	s := String(key)
	var v int
	if s == "" && len(def) > 0 {
		v, _ = def[0].(int)
		return v
	}
	v, _ = strconv.Atoi(s)
	return v
}

func Int64(key string, def ...interface{}) int64 {
	s := String(key)
	var v int64
	if s == "" && len(def) > 0 {
		v, _ = def[0].(int64)
		return v
	}
	v, _ = strconv.ParseInt(s, 10, 64) //10进制 64位
	return v
}

func Int32(key string, def ...interface{}) int32 {
	s := String(key)
	var v int32
	if s == "" && len(def) > 0 {
		v, _ = def[0].(int32)
		return v
	}
	tmp, _ := strconv.ParseInt(s, 10, 32) //10进制 32位
	return int32(tmp)
}

func Int8(key string, def ...interface{}) int8 {
	s := String(key)
	var v int8
	if s == "" && len(def) > 0 {
		v, _ = def[0].(int8)
		return v
	}
	tmp, _ := strconv.ParseInt(s, 10, 8) //10进制 8位
	return int8(tmp)
}

func Float64(key string, def ...interface{}) float64 {
	s := String(key)
	var v float64
	if s == "" && len(def) > 0 {
		v, _ = def[0].(float64)
		return v
	}
	v, _ = strconv.ParseFloat(s, 64)
	return v
}
func Float32(key string, def ...interface{}) float32 {
	s := String(key)
	var v float32
	if s == "" && len(def) > 0 {
		v, _ = def[0].(float32)
		return v
	}
	tmp, _ := strconv.ParseFloat(s, 64)
	return float32(tmp)
}

func Float(key string, def ...interface{}) float64 {
	return Float64(key, def...)
}

func Bool(key string, def ...interface{}) bool {
	s := String(key)
	var v bool
	if s == "" && len(def) > 0 {
		v, _ = def[0].(bool)
		return v
	}
	v, _ = strconv.ParseBool(s)
	return v
}

func StringArray(key string) []string {
	s := String(key)
	var v []string
	json.Unmarshal([]byte(s), &v)
	return v
}
