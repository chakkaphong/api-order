package utils

import (
	"api-order/internal/global/responses"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func GetPaginationQuery(c echo.Context) responses.Paging {
	page := c.Request().Header.Get("Page")
	limit := c.Request().Header.Get("Limit")

	pageNumber, _ := strconv.Atoi(page)
	limitNumber, _ := strconv.Atoi(limit)

	return responses.Paging{
		Page:  lo.Ternary(pageNumber != 0, pageNumber, 1),
		Limit: lo.Ternary(limitNumber != 0, limitNumber, 10),
	}
}
