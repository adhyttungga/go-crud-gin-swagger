package tags_delivery

import (
	"net/http"

	"github.com/adhyttungga/go-crud-gin-swagger/model/dto"
	tags_usecase "github.com/adhyttungga/go-crud-gin-swagger/usecase/tags"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TagsDelivery interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type TagsDeliveryImpl struct {
	TagsUsecase tags_usecase.TagsUsecase
}

func NewTagsDeliveryImpl(tagsUsecase tags_usecase.TagsUsecase) TagsDelivery {
	return &TagsDeliveryImpl{
		TagsUsecase: tagsUsecase,
	}
}

// CreateTags				godoc
// @Summary					Create tags
// @Description			Save tags data in Database.
// @Param						tags body dto.ReqCreateTags true "Create tags"
// @Produce					application/json
// @Tags						tags
// @Success					200 {object} dto.Response{data=dto.ResTags}
// @Router					/tags [post]
func (t *TagsDeliveryImpl) Create(ctx *gin.Context) {
	log.Info().Msg("Create Tags")
	reqCreateTags := dto.ReqCreateTags{}
	if err := ctx.ShouldBindJSON(&reqCreateTags); err != nil {
		panic(err)
	}

	t.TagsUsecase.Create(reqCreateTags)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

// FindAllTags 			godoc
// @Summary					Get All tags.
// @Description			Return list of tags.
// @Tags						tags
// @Success					200 {object} dto.Response{data=[]dto.ResTags}
// @Router					/tags [get]
func (t *TagsDeliveryImpl) FindAll(ctx *gin.Context) {
	log.Info().Msg("Find All Tags")
	resTgs := t.TagsUsecase.FindAll()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   resTgs,
	})
}

// FindByIdTags 		godoc
// @Summary					Get Single tags by id.
// @Param						tagId path string true "get tags by id"
// @Description			Returns tags whose tagId value matches id.
// @Produce					application/json
// @Tags						tags
// @Success					200 {object} dto.Response{data=dto.ResTags}
// @Router					/tags/{tagId} [get]
func (t *TagsDeliveryImpl) FindById(ctx *gin.Context) {
	log.Info().Msg("Find by Id Tags")
	tagId := ctx.Param("tagId")
	resTag := t.TagsUsecase.FindById(tagId)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   resTag,
	})
}
