package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/sabarmartua/FAQ/dto"
    "github.com/sabarmartua/FAQ/helper"
    "github.com/sabarmartua/FAQ/service"
)

type FAQController interface {
    All(ctx *gin.Context)
    FindByID(ctx *gin.Context)
    Insert(ctx *gin.Context)
    Update(ctx *gin.Context)
    Delete(ctx *gin.Context)
}

type faqController struct {
    faqService service.FAQService
}

func NewFAQController(faqService service.FAQService) FAQController {
    return &faqController{
        faqService: faqService,
    }
}

func (c *faqController) All(ctx *gin.Context) {
    faqs := c.faqService.All()
    ctx.JSON(http.StatusOK, faqs)
}

func (c *faqController) FindByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal mendapatkan ID", "Tidak ada parameter ID yang ditemukan", helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    faq := c.faqService.FindByID(id)
    if faq.ID == 0 {
        res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data dengan ID yang diberikan", helper.EmptyObj{})
        ctx.JSON(http.StatusNotFound, res)
        return
    }

    ctx.JSON(http.StatusOK, faq)
}

func (c *faqController) Insert(ctx *gin.Context) {
    var faqCreateDTO dto.NewFAQDTO
    errDTO := ctx.ShouldBind(&faqCreateDTO)
    if errDTO != nil {
        res := helper.BuildErrorResponse("Gagal memproses permintaan", errDTO.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    result, err := c.faqService.Insert(faqCreateDTO)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal menambahkan data", err.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }

    response := helper.BuildResponse(true, "OK!", result)
    ctx.JSON(http.StatusCreated, response)
}

func (c *faqController) Update(ctx *gin.Context) {
    var faqUpdateDTO dto.UpdateFAQDTO
    errDTO := ctx.ShouldBind(&faqUpdateDTO)
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

    faqUpdateDTO.ID = id
    result, err := c.faqService.Update(faqUpdateDTO)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal memperbarui data", err.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }

    response := helper.BuildResponse(true, "OK!", result)
    ctx.JSON(http.StatusOK, response)
}

func (c *faqController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal mendapatkan ID", "Tidak ada parameter ID yang ditemukan", helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    err = c.faqService.Delete(id)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal menghapus data", err.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }

    res := helper.BuildResponse(true, "Berhasil dihapus", helper.EmptyObj{})
    ctx.JSON(http.StatusOK, res)
}
