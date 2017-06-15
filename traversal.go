package untyped

import "errors"

var Nil = errors.New("unexpected nil returned")

type traversal struct {
	f []TraversalFunc
}

func Traversal(f ...TraversalFunc) *traversal {
	return &traversal{f}
}

func (t *traversal) Get(m interface{}) (*Value, error) {
	return t.access(m)
}

func (t *traversal) Set(m interface{}, s setter, val interface{}) error {
	v, err := t.access(m)
	if err != nil {
		return err
	}
	if _, err := s.set(v, val); err != nil {
		return err
	}
	return nil
}

func (t *traversal) Branch(f ...TraversalFunc) *traversal {
	fs := make([]TraversalFunc, 0, 0)
	fs = append(fs, t.f...)
	fs = append(fs, f...)
	return &traversal{fs}
}

func (t *traversal) access(m interface{}) (*Value, error) {
	var err error
	p := &Value{m}
	for _, fn := range t.f {
		p, err = fn(p)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

type TraversalFunc func(*Value) (*Value, error)

func Get(g getter) TraversalFunc {
	return func(v *Value) (*Value, error) {
		v, err := g.get(v)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, Nil
		}
		return v, nil
	}
}

func GetOrDefault(g getter, def interface{}) TraversalFunc {
	return func(v *Value) (*Value, error) {
		v, err := g.get(v)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return &Value{def}, nil
		}
		return v, nil
	}
}

func GetOrCreate(g getsetter, factory Factory) TraversalFunc {
	return func(v *Value) (*Value, error) {
		val, err := g.get(v)
		if err != nil {
			return nil, err
		}
		if val.val == nil {
			val, err = g.set(v, factory())
			if err != nil {
				return nil, err
			}
		}
		return val, nil
	}
}
