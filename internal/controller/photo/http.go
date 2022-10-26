package photo

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/photo"
)

type PhotoController interface {
	PostPhoto(ctx *gin.Context)
	GetAllPhotos(ctx *gin.Context)
	PutPhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHttp struct {
	photoUsecase photo.PhotoUsecase
	ctx          context.Context
}

func NewUserHttp(ctx context.Context, photoUsecase photo.PhotoUsecase) PhotoController {
	return &photoHttp{
		photoUsecase: photoUsecase,
		ctx:          ctx,
	}
}

func (p photoHttp) PostPhoto(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	var bindPhoto domain.PhotoAdd

	err := ctx.ShouldBindJSON(&bindPhoto)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	addedPhoto, err := p.photoUsecase.AddPhoto(p.ctx, &userId, &domain.Photo{
		Title:    bindPhoto.Title,
		Caption:  bindPhoto.Caption,
		PhotoUrl: bindPhoto.PhotoUrl,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusCreated, addedPhoto)
}

func (p photoHttp) GetAllPhotos(ctx *gin.Context) {
	photos, err := p.photoUsecase.GetAllPhotos(p.ctx)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (p photoHttp) PutPhoto(ctx *gin.Context) {
	paramPhotoId := ctx.Param("photoId")
	conv, _ := strconv.ParseUint(paramPhotoId, 10, 64)
	photoId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	var bindPhoto domain.PhotoUpdateData

	err := ctx.ShouldBindJSON(&bindPhoto)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	updatedPhoto, err := p.photoUsecase.UpdatePhoto(p.ctx, &userId, &photoId, &domain.Photo{
		Title:    bindPhoto.Title,
		Caption:  bindPhoto.Caption,
		PhotoUrl: bindPhoto.PhotoUrl,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (p photoHttp) DeletePhoto(ctx *gin.Context) {
	paramPhotoId := ctx.Param("photoId")
	conv, _ := strconv.ParseUint(paramPhotoId, 10, 64)
	photoId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	_, err := p.photoUsecase.DeletePhoto(p.ctx, &userId, &photoId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}