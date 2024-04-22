package request

type UpdateTagRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=10" json:"name"`
}
