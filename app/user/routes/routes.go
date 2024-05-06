package routes

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"goMysql/app/user/services/account"
)

func UserAccount(e *echo.Echo, db *sql.DB) {
	name := "user"
	e.POST("/"+name+"/singup", func(c echo.Context) error {
		return account.SingupController(c, db)
	})
	e.POST("/"+name+"/singin", func(c echo.Context) error {
		return account.SinginController(c, db)
	})
}
