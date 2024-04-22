package service

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagsRepository,
		Validate:       validate,
	}
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.CreateTagRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []response.TagsReponse {
	result := t.TagsRepository.FindAll()

	// var tags []response.TagsReponse
	tags := make([]response.TagsReponse, 0, len(result))

	for _, value := range result {
		tag := response.TagsReponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImpl) FindById(TagsId int) response.TagsReponse {
	tagData, err := t.TagsRepository.FindById(TagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsReponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.UpdateTagRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)

}
