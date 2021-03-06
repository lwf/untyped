package untyped

import "fmt"

type Getter interface {
	get(*Value) (*Value, error)
}

type Setter interface {
	set(*Value, interface{}) (*Value, error)
}

type Getsetter interface {
	Getter
	Setter
}

type mapGetSetter struct {
	key string
}

func (mg *mapGetSetter) get(v *Value) (*Value, error) {
	m, err := v.MSI()
	if err != nil {
		return nil, err
	}
	return &Value{m[mg.key]}, nil
}

func (mg *mapGetSetter) set(v *Value, val interface{}) (*Value, error) {
	m, err := v.MSI()
	if err != nil {
		return nil, err
	}
	m[mg.key] = val
	return &Value{m[mg.key]}, nil
}

func MapKey(key string) Getsetter {
	return &mapGetSetter{key}
}

type arrayGetSetter struct {
	idx int
}

func (a *arrayGetSetter) get(v *Value) (*Value, error) {
	m, err := v.Array()
	if err != nil {
		return nil, err
	}
	if len(m) <= a.idx {
		return nil, fmt.Errorf("array out of bounds (length %d)", len(m))
	}
	return &Value{m[a.idx]}, nil
}

func (a *arrayGetSetter) set(v *Value, val interface{}) (*Value, error) {
	arr, err := v.Array()
	if err != nil {
		return nil, err
	}
	arr[a.idx] = val
	return &Value{arr[a.idx]}, nil
}

func ArrayIndex(i int) Getsetter {
	return &arrayGetSetter{i}
}
