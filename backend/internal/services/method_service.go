package services

func VerifyGetRequest(method string) bool {
	return !(method == "GET")
}

func VerifyPostRequest(method string) bool {
	return !(method == "POST")
}

func VerifyDeleteRequest(method string) bool {
	return !(method == "DELETE")
}
