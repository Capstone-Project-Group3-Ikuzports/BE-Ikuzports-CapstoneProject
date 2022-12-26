package delivery

import "ikuzports/features/productImage"

type ProductImageResponse struct {
	ID           uint   `json:"id"`
	ProductImage string `json:"product_image"`
	ProductID    uint   `json:"product_id"`
}

func fromCore(dataCore productImage.ProductImageCore) ProductImageResponse {
	return ProductImageResponse{
		ID:           dataCore.ID,
		ProductImage: dataCore.URL,
		ProductID:    dataCore.ProductID,
	}
}

func fromCoreList(dataCore []productImage.ProductImageCore) []ProductImageResponse {
	var dataResponse []ProductImageResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
