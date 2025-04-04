package order

type OrderCreateRequest struct {
	ProductName string `json:"productname"`
	Amount      string `json:"amount"`
}

type OrderCreateResponse struct {
	OrderId string `json:"orderid"`
}

type OrderGetByIdRequest struct {
	OrderId string `json:"orderid"`
}

type OrderGetByIdResponse struct {
	ProductName string `json:"productname"`
	Amount      string `json:"amount"`
}

type OrderGetAllOrdersByUserRequest struct {
}

type OrderGetAllOrdersByUserResponse struct {
	Orders []Order `json:"orders"`
}
