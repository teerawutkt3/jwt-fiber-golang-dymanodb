package auth

type ProductHandler struct {
	svc ProductService
}

func NewProductHandler(svc ProductService) ProductHandler {
	return ProductHandler{svc}
}

