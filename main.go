package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"goMysql/configs/dbConfig"
	"goMysql/configs/routes"
)

func main() {
	db, err := dbConfig.InitDB()
	if err != nil {
		fmt.Println("MySQL bağlantısı başarısız:", err)
		return
	}
	e := echo.New()

	routes.InitRoutes(e, db)
	fmt.Println("Rotalar oluşturuldu")

	err = e.Start(":8080")
	if err != nil {
		fmt.Println("Sunucu başlatma hatası:", err)
	}
}
