package routes

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	admin "goMysql/app/admin/routes"
	user "goMysql/app/user/routes"
)

func InitRoutes(e *echo.Echo, db *sql.DB) {
	user.UserAccount(e, db)
	admin.AdminAccount(e, db)
}
