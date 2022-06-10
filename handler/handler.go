package handler

import "github.com/kataras/iris/v12"

type Handler interface {
	Create(iris.Context)
	List(iris.Context)
	Fetch(iris.Context)
	Update(iris.Context)
	Delete(iris.Context)
}
