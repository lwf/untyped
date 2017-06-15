package untyped

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Value struct {
	val interface{}
}

func (v *Value) MSI() (map[string]interface{}, error) {
	mii, ok := v.val.(map[string]interface{})
	if !ok {
		return nil, v.convErr("map")
	}
	return mii, nil
}

func (v *Value) String() (string, error) {
	s, ok := v.val.(string)
	if !ok {
		return "", v.convErr("string")
	}
	return s, nil
}

func (v *Value) Array() ([]interface{}, error) {
	a, ok := v.val.([]interface{})
	if !ok {
		return nil, v.convErr("array")
	}
	return a, nil
}

func (v *Value) Decode(out interface{}) error {
	return mapstructure.Decode(v.val, out)
}

func (v *Value) convErr(t string) error {
	return fmt.Errorf("could not convert '%v' (%T) to %s", v.val, v.val, t)
}
