package controller

import (
	strconv "strconv"

	"github.com/gin-gonic/gin"

	// エンティティ（データベースのテーブルの行に対応）
	entity "module/models/entity"
	// entity "../../models/entity"

	// DBアクセスモジュール
	db "module/models/db"
)

// 商品の購入状態の定義(0:未購入 1:購入済)
const (
	NotPurchased = 0
	Purchased    = 1
)

// 全ての商品情報を取得
func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()
	c.JSON(200, resultProducts)
}

// 指定したIDの商品情報を取得
func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")
	productID, _ := strconv.Atoi(productIDStr)

	resultProduct := db.FindProduct(productID)

	c.JSON(200, resultProduct)
}

// 商品登録
func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	var product = entity.Product{
		ProductName: productName,
		Memo:        productMemo,
		Status:      NotPurchased,
	}

	db.InsertProduct(&product)
}

// 商品情報の状態の変更
func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	// 商品状態が未購入の場合
	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	db.UpdateStateProduct(productID, changeState)
}

// 商品削除
func DeleteProduct(c *gin.Context) {
	productIdStr := c.PostForm("productID")
	productID, _ := strconv.Atoi(productIdStr)

	db.DeleteProduct(productID)
}
