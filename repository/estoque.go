package repository

import (
	"errors"

	"github.com/arielpierot/iris-framework/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EstoqueRepository struct {
	db *gorm.DB
}

func NewEstoqueRepository(db *gorm.DB) InterfaceEstoqueRepository {
	return EstoqueRepository{db}
}

func (s EstoqueRepository) List(page, pageSize int) (produtos []model.Produto, err error) {
	s.db.Joins("Estoque").Scopes(Paginate(page, pageSize)).Find(&produtos)
	if len(produtos) == 0 {
		err = errors.New("Não há produtos cadastrados")
		return
	}
	return
}

func (s EstoqueRepository) Fetch(codigo string) (produto model.Produto, err error) {
	s.db.Joins("Estoque").Where("codigo = ?", codigo).First(&produto)
	if len(produto.Codigo) == 0 {
		err = errors.New("Não foi possível encontrar o produto pelo código.")
		return
	}
	return
}

func (s EstoqueRepository) Create(produtoDTO model.Produto) (produto model.Produto, err error) {
	result := s.db.Create(&produtoDTO)
	s.db.Joins("Estoque").First(&produto, produtoDTO.ID)
	err = result.Error
	return
}
func (s EstoqueRepository) Update(codigo string, produtoDTO model.Produto) (produto model.Produto, err error) {
	var estoque model.Estoque
	s.db.Where("codigo = ?", codigo).First(&produto)
	if len(produto.Codigo) == 0 {
		err = errors.New("Não foi possível encontrar o produto pelo código.")
		return
	}
	produtoDTO.Estoque.EstoqueDisponivel = produtoDTO.Estoque.EstoqueTotal - produtoDTO.Estoque.EstoqueCorte
	s.db.First(&estoque, produto.EstoqueID)
	s.db.Model(&produto).Updates(produtoDTO)
	s.db.Model(&estoque).Updates(produtoDTO.Estoque)
	s.db.Joins("Estoque").First(&produto, produto.ID)
	return
}

func (s EstoqueRepository) Delete(codigo string) (produtos []model.Produto) {
	s.db.Clauses(clause.Returning{}).Where("codigo = ?", codigo).Delete(&produtos)
	return
}
