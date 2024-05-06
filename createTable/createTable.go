package createTable

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func CreateTable(db *sql.DB, sqlFilePath string) error {
	// SQL dosyasını oku
	sqlBytes, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		return err
	}

	// SQL komutlarını bir stringe dönüştür
	sqlQuery := string(sqlBytes)

	// SQL komutlarını MySQL veritabanında çalıştır
	_, err = db.Exec(sqlQuery)
	if err != nil {
		return err
	}

	fmt.Println("Tablo başarıyla oluşturuldu")
	return nil
}
