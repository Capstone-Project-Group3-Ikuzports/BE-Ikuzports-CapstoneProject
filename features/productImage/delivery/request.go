package delivery

import "ikuzports/features/productImage"

type ProductImageRequest struct {
	ProductImage string `json:"product_image" form:"product_image"`
	ProductID    uint   `json:"product_id" form:"product_id"`
}

func toCore(productImageInput ProductImageRequest) productImage.ProductImageCore {
	productImageCoreData := productImage.ProductImageCore{
		URL:       productImageInput.ProductImage,
		ProductID: productImageInput.ProductID,
	}
	return productImageCoreData
}
