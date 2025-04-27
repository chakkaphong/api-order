package responses

func InternalServerError() (res Response) {
	return Response{
		Code:    "internal_server_error",
		Message: "Internal server error.",
		Data:    nil,
	}
}

func BadRequest(msg ...string) (res ResponseBadRequest) {
	m := "Bad request."
	if len(msg) > 0 {
		m = msg[0]
	}
	return ResponseBadRequest{
		Code:    "bad_request",
		Message: m,
		Data:    nil,
	}
}

func DataNotFound(code ...string) (res Response) {
	c := "data_not_found"
	if len(code) > 0 {
		c = code[0]
	}
	return Response{
		Code:    c,
		Message: "Data not found.",
		Data:    nil,
	}
}
