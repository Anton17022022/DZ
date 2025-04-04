package product

// TODO тут бы надо сделать общий инстанс из дескрипшина и ид

type ProductCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductCreateResponse struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ProductUpdateRequest struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ProductUpdateResponse struct {
	ProductCreateResponse
}

type ProductDeleteRequest struct {
}

type ProductDeleteResponse struct {
	Message string `json:"message"`
}

type ProductGetAllRequest struct {
}

type ProductGetAllResponse struct {
	Products []Product `json:"products"`
}
