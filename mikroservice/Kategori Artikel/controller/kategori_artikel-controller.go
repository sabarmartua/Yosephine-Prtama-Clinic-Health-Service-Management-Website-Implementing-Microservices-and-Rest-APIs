package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/sabarmartua/Kategori-Artikel/dto"
    "github.com/sabarmartua/Kategori-Artikel/helper"
    "github.com/sabarmartua/Kategori-Artikel/service"
)

// KategoriArtikelController adalah kontrak tentang apa yang dapat dilakukan oleh controller ini
type KategoriArtikelController interface {
    All(ctx *gin.Context)
    FindByID(ctx *gin.Context)
    Insert(ctx *gin.Context)
    Update(ctx *gin.Context)
    Delete(ctx *gin.Context)
}

type kategoriArtikelController struct {
    KategoriArtikelService service.KategoriArtikelService
}

// NewKategoriArtikelController membuat instance baru dari KategoriArtikelController
func NewKategoriArtikelController(KategoriArtikelService service.KategoriArtikelService) KategoriArtikelController {
    return &kategoriArtikelController{
        KategoriArtikelService: KategoriArtikelService,
    }
}

func (c *kategoriArtikelController) All(ctx *gin.Context) {
    kategoriArtikels := c.KategoriArtikelService.All()
    ctx.JSON(http.StatusOK, kategoriArtikels)
}

func (c *kategoriArtikelController) FindByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal mendapatkan ID", "Tidak ada parameter ID yang ditemukan", helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    kategoriArtikel := c.KategoriArtikelService.FindByID(id)
    if kategoriArtikel.ID == 0 {
        res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data dengan ID yang diberikan", helper.EmptyObj{})
        ctx.JSON(http.StatusNotFound, res)
        return
    }

    ctx.JSON(http.StatusOK, kategoriArtikel)
}

func (c *kategoriArtikelController) Insert(ctx *gin.Context) {
    var kategoriArtikelCreateDTO dto.NewKategoriArtikelDTO
    errDTO := ctx.ShouldBind(&kategoriArtikelCreateDTO)
    if errDTO != nil {
        res := helper.BuildErrorResponse("Gagal memproses permintaan", errDTO.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }
    result := c.KategoriArtikelService.Insert(kategoriArtikelCreateDTO)
    response := helper.BuildResponse(true, "OK!", result)
    ctx.JSON(http.StatusCreated, response)
}

func (c *kategoriArtikelController) Update(ctx *gin.Context) {
    var kategoriArtikelUpdateDTO dto.UpdateKategoriArtikelDTO
    errDTO := ctx.ShouldBind(&kategoriArtikelUpdateDTO)
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
    kategoriArtikelUpdateDTO.ID = id // Convert id to uint
    result := c.KategoriArtikelService.Update(kategoriArtikelUpdateDTO)
    response := helper.BuildResponse(true, "OK!", result)
    ctx.JSON(http.StatusOK, response)
}

func (c *kategoriArtikelController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal mendapatkan ID", "Tidak ada parameter ID yang ditemukan", helper.EmptyObj{})
        ctx.JSON(http.StatusBadRequest, res)
        return
    }
    err = c.KategoriArtikelService.Delete(id)
    if err != nil {
        res := helper.BuildErrorResponse("Gagal menghapus data", err.Error(), helper.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }
    res := helper.BuildResponse(true, "Berhasil dihapus", helper.EmptyObj{})
    ctx.JSON(http.StatusOK, res)
}
