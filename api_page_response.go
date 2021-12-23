package web

import (
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/outerjoin/do"
	"github.com/rightjoin/fig"
)

type Paging struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	Current int `json:"current"`
	Total   int `json:"total"`
	Chunk   int `json:"chunk"`
}

type ApiPageResponse struct {
	ApiResponse
	Paging `json:"paging"`
}

func (a *ApiPageResponse) Scribe(c echo.Context) {
	c.JSON(http.StatusOK, a)
}

func (a *ApiPageResponse) SetData(d interface{}, current, total int) error {
	a.Data = d
	a.Success = true

	a.Current = current
	a.Total = total
	a.Pages = int(math.Ceil(float64(total) / float64(a.Chunk)))

	return nil
}

func NewApiPageResponse(page, chunk int) ApiPageResponse {

	if page <= 0 {
		page = 1
	}

	max := fig.IntOr(25, "pagination.chunk")
	if chunk < 1 || chunk > max {
		chunk = max
	}

	return ApiPageResponse{
		Paging: Paging{
			Page:    page,
			Pages:   0,
			Total:   0,
			Current: 0,
			Chunk:   chunk,
		},
	}
}

func NewApiPageResponseFromRequest(c echo.Context) ApiPageResponse {

	page := do.ParseIntOr(c.QueryParam(":page"), 0)
	chunk := do.ParseIntOr(c.QueryParam(":chunk"), 0)

	return NewApiPageResponse(page, chunk)
}
