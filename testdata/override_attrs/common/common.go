package common

type Optional[T any] struct {
	Value T
	Set   bool
}

type Foo struct {
	A string `json:"a"`
}

type Bar struct {
	B int `json:"b"`
}
