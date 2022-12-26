package delivery

import "ikuzports/features/product"

type ProductResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Price            uint   `json:"price"`
	UserName         string `json:"user_name"`
	Description      string `json:"description"`
	ItemCategoryName string `json:"itemcategory_name"`
}

func fromCore(dataCore product.ProductCore) ProductResponse {
	return ProductResponse{
		ID:               dataCore.ID,
		Name:             dataCore.Name,
		Price:            dataCore.Price,
		UserName:         dataCore.User.Name,
		Description:      dataCore.Description,
		ItemCategoryName: dataCore.ItemCategory.Name,
	}
}

func fromCoreList(dataCore []product.ProductCore) []ProductResponse {
	var dataResponse []ProductResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
