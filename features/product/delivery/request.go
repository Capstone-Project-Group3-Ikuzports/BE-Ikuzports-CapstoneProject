package delivery

import "ikuzports/features/product"

type ProductRequest struct {
	Name           string `json:"name" form:"name"`
	Price          uint   `json:"price" form:"price"`
	Description    string `json:"description" form:"description"`
	ItemCategoryID uint   `json:"itemcategory_id" form:"itemcategory_id"`
}

func toCore(productInput ProductRequest) product.ProductCore {
	productCoreData := product.ProductCore{
		Name:           productInput.Name,
		Price:          productInput.Price,
		Description:    productInput.Description,
		ItemCategoryID: productInput.ItemCategoryID,
	}
	return productCoreData
}
