package delivery

import (
	itemcategory "ikuzports/features/itemCategory"
)

type ItemCategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func fromCore(dataCore itemcategory.ItemCategoryCore) ItemCategoryResponse {
	return ItemCategoryResponse{
		ID:   dataCore.ID,
		Name: dataCore.Name,
	}
}

func fromCoreList(dataCore []itemcategory.ItemCategoryCore) []ItemCategoryResponse {
	var dataResponse []ItemCategoryResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
