package service

import "github.com/arielpierot/iris-framework/model"

type InterfaceEstoqueService interface {
	List(int, int) ([]model.Produto, error)
	Fetch(string) (model.Produto, error)
	Create(model.Produto) (model.Produto, error)
	Update(string, model.Produto) (model.Produto, error)
	Delete(string) ([]model.Produto, error)
}
