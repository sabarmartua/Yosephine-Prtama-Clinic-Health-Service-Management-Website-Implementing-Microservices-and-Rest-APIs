package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Artikel/dto"
	"github.com/sabarmartua/Artikel/helper"
	"github.com/sabarmartua/Artikel/service"
)

type ArtikelController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RelatedByCategory(ctx *gin.Context)
}

type artikelController struct {
	artikelService service.ArtikelService
}

func NewArtikelController(artikelService service.ArtikelService) ArtikelController {
	return &artikelController{
		artikelService: artikelService,
	}
}

func (c *artikelController) All(ctx *gin.Context) {
	artikels, err := c.artikelService.All()
	if err != nil {
		res := helper.BuildErrorResponse("Failed to fetch articles", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	ctx.JSON(http.StatusOK, artikels)
}

func (c *artikelController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	artikel, err := c.artikelService.FindByID(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to fetch article", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	if artikel.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, artikel)
}

func (c *artikelController) Insert(ctx *gin.Context) {
	var artikelDTO dto.NewArtikelDTO
	err := ctx.ShouldBindJSON(&artikelDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	artikel, err := c.artikelService.Create(artikelDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to create article", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Article created successfully", artikel)
	ctx.JSON(http.StatusCreated, res)
}

func (c *artikelController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var artikelDTO dto.UpdateArtikelDTO
	err = ctx.ShouldBindJSON(&artikelDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	artikel, err := c.artikelService.Update(id, artikelDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to update article", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Article updated successfully", artikel)
	ctx.JSON(http.StatusOK, res)
}

func (c *artikelController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.artikelService.Delete(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete article", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Article deleted successfully", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *artikelController) RelatedByCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	relatedArticles, err := c.artikelService.GetRelatedByCategory(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to fetch related articles", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	if len(relatedArticles) == 0 {
		res := helper.BuildErrorResponse("Data not found", "No related articles found for the given category ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, relatedArticles)
}
