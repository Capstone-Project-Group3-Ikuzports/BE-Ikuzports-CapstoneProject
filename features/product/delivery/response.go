package delivery

import "ikuzports/features/product"

type ProductResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Price            uint   `json:"price"`
	UserName         string `json:"user_name"`
	Description      string `json:"description"`
	ItemCategoryName string `json:"itemcategory_name"`
	City             string `json:"city"`
}

type ProductResponseImage struct {
	ID               uint           `json:"id"`
	Name             string         `json:"name"`
	Price            uint           `json:"price"`
	UserName         string         `json:"user_name"`
	Description      string         `json:"description"`
	ItemCategoryName string         `json:"itemcategory_name"`
	City             string         `json:"city"`
	ProductImage     []ProductImage `json:"product_image"`
}

type ProductImage struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

func fromCore(dataCore product.ProductCore) ProductResponse {
	return ProductResponse{
		ID:               dataCore.ID,
		Name:             dataCore.Name,
		Price:            dataCore.Price,
		UserName:         dataCore.User.Name,
		Description:      dataCore.Description,
		ItemCategoryName: dataCore.ItemCategory.Name,
		City:             dataCore.City,
	}
}

func fromCoreList(dataCore []product.ProductCore) []ProductResponse {
	var dataResponse []ProductResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreImage(data product.ProductImage) ProductImage {
	return ProductImage{
		ID:  data.ID,
		Url: data.URL,
	}
}

func fromImageCore(dataCore product.ProductCore) ProductResponseImage {
	return ProductResponseImage{
		ID:               dataCore.ID,
		Name:             dataCore.Name,
		Price:            dataCore.Price,
		UserName:         dataCore.User.Name,
		Description:      dataCore.Description,
		ItemCategoryName: dataCore.ItemCategory.Name,
		City:             dataCore.City,
		ProductImage:     fromCoreListImage(dataCore.ProductImage),
	}
}

func fromCoreListImage(dataCore []product.ProductImage) []ProductImage {
	var dataResponse []ProductImage
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreImage(v))
	}
	return dataResponse
}
