package handler

type Context interface {
	Param(string) string
	PostForm(string) interface{}
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
