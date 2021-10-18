package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	ID          int    `gorm:"primary_key;not null"`
	ProductName string `gorm:"type varchar(200);not null"`
	Memo        string `gorm:"type:varchar(400);not null"`
	Status      string `gorm:"type:char(2);not null"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	DBNAME := "Shopping"
	PROTOCOL := "tcp(localhost:3306)"
	CONNECT := USER + ":" + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション
	db.AutoMigrate(&Product{})

	fmt.Println("db connected ", &db)

	return db
}

// 商品登録
func insertProduct(registerProduct *Product) {
	db := gormConnect()

	db.Create(&registerProduct)
	defer db.Close()
}

// 全件取得
func getAllProduct() []Product {
	db := gormConnect()
	var products []Product

	db.Order("ID asc").Find(&products)
	defer db.Close()
	return products
}

func main() {

	// 構造体初期化

	var product = Product{
		ProductName: "テスト商品",
		Memo:        "テスト商品です",
		Status:      "01",
	}

	insertProduct(&product)

	resultProducts := getAllProduct()

	for i := range resultProducts {
		fmt.Printf("index: %d, 商品ID: %d, 商品名: %s, メモ: %s, ステータス: %s\n",
			i, resultProducts[i].ID, resultProducts[i].ProductName, resultProducts[i].Memo, resultProducts[i].Status)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})
	r.Run()
}
