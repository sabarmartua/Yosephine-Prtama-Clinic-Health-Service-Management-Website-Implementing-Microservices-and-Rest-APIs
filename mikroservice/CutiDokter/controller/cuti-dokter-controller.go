package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/CutiDokter/dto"
	"github.com/sabarmartua/CutiDokter/helper"
	"github.com/sabarmartua/CutiDokter/service"
)

type CutiDokterController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type cutiDokterController struct {
	cutiDokterService service.CutiDokterService
}

func NewCutiDokterController(cutiDokterService service.CutiDokterService) CutiDokterController {
	return &cutiDokterController{
		cutiDokterService: cutiDokterService,
	}
}

func (c *cutiDokterController) All(ctx *gin.Context) {
	cutiDokters := c.cutiDokterService.All()
	ctx.JSON(http.StatusOK, cutiDokters)
}

func (c *cutiDokterController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Gagal mendapatkan ID", "Tidak ada parameter ID yang ditemukan", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	cutiDokter := c.cutiDokterService.FindByID(id)
	if cutiDokter.ID == 0 {
		res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data dengan ID yang diberikan", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, cutiDokter)
}

func (c *cutiDokterController) Insert(ctx *gin.Context) {
	var cutiDokterCreateDTO dto.NewCutiDokterDTO
	errDTO := ctx.ShouldBind(&cutiDokterCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.cutiDokterService.Insert(cutiDokterCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Gagal menambahkan data", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *cutiDokterController) Update(ctx *gin.Context) {
	var cutiDokterUpdateDTO dto.UpdateCutiDokterDTO
	errDTO := ctx.ShouldBind(&cutiDokterUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	idStr := ctx.Param("id")
	id, errID := strconv.ParseUint(idStr, 10, 64)
	if errID != nil {
		res := helper.BuildErrorResponse("Gagal mendapatkan ID", "Tidak ada parameter ID yang ditemukan", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	cutiDokterUpdateDTO.ID = id
	result, err := c.cutiDokterService.Update(cutiDokterUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Gagal memperbarui data", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *cutiDokterController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Gagal mendapatkan ID", "Tidak ada parameter ID yang ditemukan", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.cutiDokterService.Delete(id)
	if err != nil {
		res := helper.BuildErrorResponse("Gagal menghapus data", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Berhasil dihapus", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
