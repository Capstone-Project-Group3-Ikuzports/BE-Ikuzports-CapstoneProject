package helper

func FailedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Failed",
		"message": msg,
	}
}

func SuccessResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": msg,
	}
}

func SuccessWithDataResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": msg,
		"data":    data,
	}
}
