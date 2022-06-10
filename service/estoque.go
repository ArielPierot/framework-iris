package service

import (
	"errors"

	"github.com/arielpierot/iris-framework/model"
	"github.com/arielpierot/iris-framework/repository"
	"gorm.io/gorm"
)

type EstoqueService struct {
	repository repository.InterfaceEstoqueRepository
}

func NewServiceEstoque(db *gorm.DB) InterfaceEstoqueService {
	return EstoqueService{repository: repository.NewEstoqueRepository(db)}
}

func (s EstoqueService) List(page, pageSize int) (produtos []model.Produto, err error) {
	produtos, err = s.repository.List(page, pageSize)
	if len(produtos) == 0 {
		err = errors.New("Não há produtos cadastrados")
		return
	}

	return
}

func (s EstoqueService) Fetch(codigo string) (produto model.Produto, err error) {
	produto, err = s.repository.Fetch(codigo)
	return
}

func (s EstoqueService) Create(produtoDTO model.Produto) (produto model.Produto, err error) {
	if produtoDTO.PrecoInvalido() {
		err = errors.New("O preço por precisa ser inferior ao preço de.")
		return
	}
	if produtoDTO.EstoqueInvalido() {
		err = errors.New("O estoque de corte não pode ser maior que o estoque total.")
		return
	}
	produtoDTO.Estoque.EstoqueDisponivel = produtoDTO.Estoque.EstoqueTotal - produtoDTO.Estoque.EstoqueCorte
	produto, err = s.repository.Create(produtoDTO)
	return
}

func (s EstoqueService) Update(codigo string, produtoDTO model.Produto) (produto model.Produto, err error) {
	if produtoDTO.PrecoInvalido() {
		err = errors.New("O preço por precisa ser inferior ao preço de.")
		return
	}
	if produtoDTO.EstoqueInvalido() {
		err = errors.New("O estoque de corte não pode ser maior que o estoque total.")
		return
	}
	produto, err = s.repository.Update(codigo, produtoDTO)
	return
}

func (s EstoqueService) Delete(codigo string) (produtos []model.Produto, err error) {
	produtos = s.repository.Delete(codigo)
	if len(produtos) == 0 {
		err = errors.New("Nenhum registro foi deletado.")
	}
	return
}
