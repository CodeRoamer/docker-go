package api

import (
	"reflect"
	"net/url"
	"strings"
	"strconv"
	"encoding/json"
)


// Struct to JSON
func ParseStruct2Json(opts interface {}) (string, error) {
	//if struct is nil
	if opts != nil {
		if buf, err := json.Marshal(opts); err == nil {
			return string(buf), nil
		} else {
			return "", err
		}
	}

	return "", nil
}



// Struct to Multi-value Map
func ParseStruct2MultiMap(opts interface {}) map[string][]string {

	var ret_map = url.Values(make(map[string][]string))

	if opts == nil {
		return ret_map
	}

	// type of opts
	value := reflect.ValueOf(opts)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return ret_map
	}

	// iterate the field of struct
	for i := 0; i < value.NumField(); i++ {
		// type of the i-th field
		field := value.Type().Field(i)
		if field.PkgPath != "" {
			continue
		}

		// field name (if has json name)
		key := field.Tag.Get("qs")
		if key == "" {
			key = strings.ToLower(field.Name)
		} else if key == "-" {
			continue
		}

		v := value.Field(i)
		switch v.Kind() {
		case reflect.Bool:
			if v.Bool() {
				ret_map.Add(key, "1")
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v.Int() > 0 {
				ret_map.Add(key, strconv.FormatInt(v.Int(), 10))
			}
		case reflect.Float32, reflect.Float64:
			if v.Float() > 0 {
				ret_map.Add(key, strconv.FormatFloat(v.Float(), 'f', -1, 64))
			}
		case reflect.String:
			if v.String() != "" {
				ret_map.Add(key, v.String())
			}
		case reflect.Ptr:
			if !v.IsNil() {
				if b, err := json.Marshal(v.Interface()); err == nil {
					ret_map.Add(key, string(b))
				}
			}
		}
	}

	return ret_map
}


// Struct to query string
func ParseStruct2QueryString(opts interface {}) string {

	return url.Values(ParseStruct2MultiMap(opts)).Encode()

}

