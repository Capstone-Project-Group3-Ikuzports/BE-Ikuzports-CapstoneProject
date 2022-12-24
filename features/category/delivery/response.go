package delivery

import "ikuzports/features/category"

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func fromCore(dataCore category.CategoryCore) CategoryResponse {
	return CategoryResponse{
		ID:   dataCore.ID,
		Name: dataCore.Name,
	}
}

func fromCoreList(dataCore []category.CategoryCore) []CategoryResponse {
	var dataResponse []CategoryResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
