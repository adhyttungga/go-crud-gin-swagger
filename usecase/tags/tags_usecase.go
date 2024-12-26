package tags_usecase

import (
	"github.com/adhyttungga/go-crud-gin-swagger/model/dto"
	"github.com/adhyttungga/go-crud-gin-swagger/model/entity"
	tags_repository "github.com/adhyttungga/go-crud-gin-swagger/repository/tags"
	"github.com/go-playground/validator/v10"
)

type TagsUsecase interface {
	Create(tags dto.ReqCreateTags)
	FindAll() []dto.ResTags
	FindById(tagsId string) dto.ResTags
}

type TagsUsecaseImpl struct {
	TagsRepository tags_repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsUsecaseImpl(tagsRepository tags_repository.TagsRepository, validate *validator.Validate) TagsUsecase {
	return &TagsUsecaseImpl{
		TagsRepository: tagsRepository,
		Validate:       validate,
	}
}

// Create implements TagsService
func (t *TagsUsecaseImpl) Create(tags dto.ReqCreateTags) {
	if err := t.Validate.Struct(tags); err != nil {
		panic(err)
	}

	tagsEntity := entity.Tags{
		Name: tags.Name,
	}

	t.TagsRepository.Save(tagsEntity)
}

// FindAll implements TagsService
func (t *TagsUsecaseImpl) FindAll() []dto.ResTags {
	result := t.TagsRepository.FindAll()

	var tags []dto.ResTags
	for _, value := range result {
		tag := dto.ResTags{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService
func (t *TagsUsecaseImpl) FindById(tagsId string) dto.ResTags {
	tagData, err := t.TagsRepository.FindById(tagsId)
	if err != nil {
		panic(err)
	}

	resTags := dto.ResTags{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return resTags
}
