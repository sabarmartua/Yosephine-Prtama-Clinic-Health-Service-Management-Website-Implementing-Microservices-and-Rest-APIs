package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Antrian/dto"
	"github.com/sabarmartua/Antrian/model"
	"github.com/sabarmartua/Antrian/helper"
	"github.com/sabarmartua/Antrian/service"
)

type AntrianController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByUserID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type antrianController struct {
	AntrianService service.AntrianService
}

func NewAntrianController(antrianService service.AntrianService) AntrianController {
	return &antrianController{
		AntrianService: antrianService,
	}
}

func (ac *antrianController) Create(ctx *gin.Context) {
	var antrianDTO dto.NewAntrianDTO
	errDTO := ctx.ShouldBind(&antrianDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := ac.AntrianService.Create(antrianDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (ac *antrianController) GetAll(ctx *gin.Context) {
	antrians := ac.AntrianService.GetAll()
	response := helper.BuildResponse(true, "OK!", antrians)
	ctx.JSON(http.StatusOK, response)
}

func (ac *antrianController) GetByUserID(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint64)

	antrians := ac.AntrianService.GetByUserID(userID)
	response := helper.BuildResponse(true, "OK!", antrians)
	ctx.JSON(http.StatusOK, response)
}

func (ac *antrianController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	antrian := model.Antrian{ID: id}
	ac.AntrianService.Delete(antrian)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
