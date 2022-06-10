package handler

import (
	"strconv"

	"github.com/arielpierot/iris-framework/model"
	"github.com/arielpierot/iris-framework/service"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type HandlerEstoque struct {
	estoqueService service.InterfaceEstoqueService
}

func NewHandlerEstoque(db *gorm.DB) Handler {
	return HandlerEstoque{estoqueService: service.NewServiceEstoque(db)}
}

func (h HandlerEstoque) List(ctx iris.Context) {
	page, err := strconv.Atoi(ctx.URLParam("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(ctx.URLParam("page_size"))
	if err != nil {
		pageSize = 10
	}

	produtos, err := h.estoqueService.List(page, pageSize)
	if err != nil {
		_ = ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(produtos)
}

func (h HandlerEstoque) Fetch(ctx iris.Context) {
	codigo := ctx.Params().Get("codigo")

	produto, err := h.estoqueService.Fetch(codigo)
	if err != nil {
		_ = ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(produto)
}

func (h HandlerEstoque) Create(ctx iris.Context) {
	var produtoDTO model.Produto
	err := ctx.ReadJSON(&produtoDTO)
	if err != nil {
		_ = ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Não foi possível receber o payload.").DetailErr(err))
		return
	}

	produto, err := h.estoqueService.Create(produtoDTO)
	if err != nil {
		_ = ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, _ = ctx.JSON(produto)
}

func (h HandlerEstoque) Update(ctx iris.Context) {
	var (
		produtoDTO model.Produto
		produto    model.Produto
	)

	err := ctx.ReadJSON(&produtoDTO)
	if err != nil {
		_ = ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Não foi possível receber o payload.").DetailErr(err))
		return
	}

	codigo := ctx.Params().Get("codigo")

	produto, err = h.estoqueService.Update(codigo, produtoDTO)
	if err != nil {
		_ = ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(produto)
}

func (h HandlerEstoque) Delete(ctx iris.Context) {
	codigo := ctx.Params().Get("codigo")

	produtos, err := h.estoqueService.Delete(codigo)
	if err != nil {
		_ = ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().DetailErr(err))
		return
	}
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(produtos)
}
