package responses

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Success(i interface{}) (res ResponseSuccess) {
	return ResponseSuccess{
		Code:    "success",
		Message: "Success",
		Data:    i,
	}
}

func SuccessPaging(i interface{}, p Paging, echoCtx ...echo.Context) (res ResponsePaging) {
	res = ResponsePaging{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i

	if len(echoCtx) > 0 {
		c := echoCtx[0]
		if c != nil {
			// set to header
			c.Response().Header().Set("X-Total-Count", fmt.Sprintf("%d", p.Total))
			c.Response().Header().Set("X-Page", fmt.Sprintf("%d", p.Page))
			c.Response().Header().Set("X-Limit", fmt.Sprintf("%d", p.Limit))
		}
	}

	return res
}
