package config 

import (
	"errors"
	"fmt"
	"reflect"
	
	"github.com/spf13/viper"
)

const tagPrefix = "viper"

func populateConfig(c *Config) (*Config, error) {
	err := recursivelySet(reflect.ValueOf(c), "")
	if err != nil {
		return nil, err 
	}
	return c, nil 
}

func recursivelySet(v reflect.Value, prefix string) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("WTF")
	}
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("FML")
	}
	vType := reflect.TypeOf(v.Interface())
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := vType.Field(i)
		tag := prefix + getTag(t)
		switch f.Kind() {
		case reflect.Struct:
			if err := recursivelySet(f.Addr(), tag+"."); err != nil {
				return err 
			}
		case reflect.Int:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			configVal := int64(viper.GetInt(tag))
			f.SetInt(configVal)
		case reflect.String:
			f.SetString(viper.GetString(tag))
		case reflect.Bool:
			f.SetBool(viper.GetBool(tag))
		default:
			return fmt.Errorf("unexpected type detection ~ aborting: %s", f.Kind())
		}
	}
	return nil
}

func getTag(f reflect.StructField) string {
	t := f.Tag
	if t != "" {
		for _, prefix := range []string{tagPrefix, "mapstructure", "yaml"} {
			if v := t.Get(prefix); v != "" {
				return v
			}
		}
	}
	return f.Name
}