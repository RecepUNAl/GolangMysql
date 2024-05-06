package account

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"goMysql/helper"
	"goMysql/types"
	"net/http"
)

func SingupController(c echo.Context, db *sql.DB) error {
	post := &types.Singup{}

	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := helper.ValidateStruct(post)
	if err != nil {
		res := types.ErrorMessage{Status: "error", Message: "Hatalı Veri Formatı", Data: result}
		return c.JSON(http.StatusBadRequest, res)
	}

	err = SingupModel(post, db)
	if err != nil {
		res := types.ErrorMessage{Status: "error", Message: "Kayıt İşlemi Başarısız", Data: err}
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, post)
}

func SinginController(c echo.Context, db *sql.DB) error {
	post := &types.Singin{}

	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := helper.ValidateStruct(post)
	if err != nil {
		res := types.ErrorMessage{Status: "error", Message: "Hatalı Veri Formatı", Data: result}
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, post)
}
