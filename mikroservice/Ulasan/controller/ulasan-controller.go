package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Ulasan/dto"
	"github.com/sabarmartua/Ulasan/model"
	"github.com/sabarmartua/Ulasan/helper"
	"github.com/sabarmartua/Ulasan/service"
)

type UlasanController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type ulasanController struct {
	UlasanService service.UlasanService
}

func NewUlasanController(ulasanService service.UlasanService) UlasanController {
	return &ulasanController{
		UlasanService: ulasanService,
	}
}

func (uc *ulasanController) Create(ctx *gin.Context) {
	var ulasanDTO dto.NewUlasanDTO
	errDTO := ctx.ShouldBind(&ulasanDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := uc.UlasanService.Create(ulasanDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (uc *ulasanController) Update(ctx *gin.Context) {
	var ulasanDTO dto.UpdateUlasanDTO
	errDTO := ctx.ShouldBind(&ulasanDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	idStr := ctx.Param("id")
	id, errID := strconv.ParseUint(idStr, 10, 64)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ulasanDTO.ID = id
	result := uc.UlasanService.Update(ulasanDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (uc *ulasanController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, errID := strconv.ParseUint(idStr, 10, 64)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ulasan := model.Ulasan{ID: id}
	uc.UlasanService.Delete(ulasan)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (uc *ulasanController) GetAll(ctx *gin.Context) {
	ulasans := uc.UlasanService.GetAll()
	response := helper.BuildResponse(true, "OK!", ulasans)
	ctx.JSON(http.StatusOK, response)
}
