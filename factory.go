package untyped

type Factory func() interface{}

func MapFactory() interface{} {
	return make(map[string]interface{})
}
