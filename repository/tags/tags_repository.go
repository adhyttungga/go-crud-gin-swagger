package tags_repository

import (
	"errors"

	"github.com/adhyttungga/go-crud-gin-swagger/model/entity"
	"gorm.io/gorm"
)

type TagsRepository interface {
	Save(tags entity.Tags)
	FindById(tagsId string) (tags entity.Tags, err error)
	FindAll() []entity.Tags
}

type TagsRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagsRepository(DB *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{DB}
}

// Save implements TagsRepository
func (t *TagsRepositoryImpl) Save(tags entity.Tags) {
	if err := t.DB.Create(&tags).Error; err != nil {
		panic(err)
	}
}

// FindById implements TagsRepository
func (t *TagsRepositoryImpl) FindById(tagsId string) (tags entity.Tags, err error) {
	var tag entity.Tags
	result := t.DB.Where("id = ?", tagsId).Find(&tag)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// FindAll implements TagsRepository
func (t *TagsRepositoryImpl) FindAll() []entity.Tags {
	var tags []entity.Tags
	if err := t.DB.Find(&tags).Error; err != nil {
		panic(err)
	}

	return tags
}
