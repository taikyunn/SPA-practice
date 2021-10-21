package db

import (
	"fmt"
	"github.com/jinzhu/gorm"

	// エンティティ(データベースのテーブルの行に対応)
	entity "module/models/entity"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	DBNAME := "Shopping"
	CONNECT := USER + ":" + "@/" + DBNAME + "?parseTime=true"
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
	db.AutoMigrate(&entity.Product{})

	fmt.Println("db connected ", &db)

	return db
}

// 全件取得
func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db := gormConnect()

	db.Order("Id asc").Find(&products)

	defer db.Close()

	return products
}

// 1件取得
func FindProduct(productID int) []entity.Product {
	product := []entity.Product{}
	db := gormConnect()

	db.First(&product, productID)
	defer db.Close()

	return product
}

// 登録
func InsertProduct(registerProduct *entity.Product) {
	db := gormConnect()

	// insert
	db.Create(&registerProduct)
	defer db.Close()
}

// 更新処理
func UpdateStateProduct(productID int, productState int) {
	product := []entity.Product{}

	db := gormConnect()

	db.Model(&product).Where("ID = ?", productID).Update("status", productState)
	defer db.Close()
}

// 削除処理
func DeleteProduct(productID int) {
	product := []entity.Product{}

	db := gormConnect()

	db.Delete(&product, productID)
	defer db.Close()
}
