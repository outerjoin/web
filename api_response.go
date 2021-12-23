package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/outerjoin/do"
)

type ApiResponse struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data"`
	Errors  []do.ErrorPlus `json:"errors"`
}

func (a *ApiResponse) SetData(data interface{}) error {
	a.Data = data
	a.Success = true
	return nil
}

func (a *ApiResponse) AddError(errs ...error) error {
	if a.Errors == nil {
		a.Errors = []do.ErrorPlus{}
	}

	tmp := make([]do.ErrorPlus, len(errs))
	for i := 0; i < len(errs); i++ {
		tmp[i] = do.ErrorPlus{Message: errs[i].Error()}
	}

	a.Errors = append(a.Errors, tmp...)

	return nil
}

func (a *ApiResponse) AddErrorPlus(errs ...do.ErrorPlus) error {
	if a.Errors == nil {
		a.Errors = []do.ErrorPlus{}
	}

	a.Errors = append(a.Errors, errs...)

	return nil
}

func (a *ApiResponse) Scribe(c echo.Context) {
	c.JSON(http.StatusOK, a)
}
