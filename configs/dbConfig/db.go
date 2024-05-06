package dbConfig

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	username := "root"
	password := ""
	hostname := "localhost"
	port := "3306"
	dbname := "mysqlgo"

	// MySQL bağlantı dizisi oluştur
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	// MySQL veritabanına bağlan
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Bağlantıyı test et
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("MySQL veritabanına başarıyla bağlandı")

	//err = createTable.CreateTable(db, "createTable/sql/user_create.sql")
	//if err != nil {
	//	fmt.Println("Tablo oluşturma hatası:", err)
	//	return nil, err
	//}

	return db, nil
}
