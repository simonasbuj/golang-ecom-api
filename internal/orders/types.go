package orders

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int64 `jsong:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}
