package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/sabarmartua/Obat/dto"
    "github.com/sabarmartua/Obat/helper"
    "github.com/sabarmartua/Obat/service"
)

type ObatController interface {
    All(ctx *gin.Context)
    FindByID(ctx *gin.Context)
    Insert(ctx *gin.Context)
    Update(ctx *gin.Context)
    Delete(ctx *gin.Context)
}

type obatController struct {
    obatService service.ObatService
}

func NewObatController(obatService service.ObatService) ObatController {
    return &obatController{
        obatService: obatService,
    }
}

func (c *obatController) All(ctx *gin.Context) {
    obats := c.obatService.GetAll()
    ctx.JSON(http.StatusOK, obats)
}

func (c *obatController) FindByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        res := helper.BuildErrorResponse("Failed to get ID", "No ID parameter found", helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    obat := c.obatService.GetByID(id)
    if obat.ID == 0 {
        res := helper.BuildErrorResponse("Data not found", "No data found with the given ID", helper.EmptyObj{})
        ctx.JSON(http.StatusNotFound, res)
        return
    }

    ctx.JSON(http.StatusOK, obat)
}

func (c *obatController) Insert(ctx *gin.Context) {
    var obatCreateDTO dto.NewObatDTO
    errDTO := ctx.ShouldBind(&obatCreateDTO)
    if errDTO != nil {
        res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    result, err := c.obatService.Insert(obatCreateDTO)
    if err != nil {
        res := helper.BuildErrorResponse("Failed to add data", err.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }

    response := helper.BuildResponse(true, "OK!", result)
    ctx.JSON(http.StatusCreated, response)
}

func (c *obatController) Update(ctx *gin.Context) {
    var obatUpdateDTO dto.UpdateObatDTO
    errDTO := ctx.ShouldBind(&obatUpdateDTO)
    if errDTO != nil {
        res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    idStr := ctx.Param("id")
    id, errID := strconv.ParseUint(idStr, 10, 64)
    if errID != nil {
        res := helper.BuildErrorResponse("Failed to get ID", "No ID parameter found", helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    obatUpdateDTO.ID = id
    result, err := c.obatService.Update(obatUpdateDTO)
    if err != nil {
        res := helper.BuildErrorResponse("Failed to update data", err.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }

    response := helper.BuildResponse(true, "OK!", result)
    ctx.JSON(http.StatusOK, response)
}

func (c *obatController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        res := helper.BuildErrorResponse("Failed to get ID", "No ID parameter found", helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    err = c.obatService.Delete(id)
    if err != nil {
        res := helper.BuildErrorResponse("Failed to delete data", err.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }

    res := helper.BuildResponse(true, "Successfully deleted", helper.EmptyObj{})
    ctx.JSON(http.StatusOK, res)
}
