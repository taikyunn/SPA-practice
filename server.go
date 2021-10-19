package main

import (
	// ロギングを行うパッケージ
	"log"
	// HTTPを扱うパッケージ
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	controller "module/controllers/controller"
)

func main() {
	// サーバーの起動
	serve()
}

func serve() {
	r := gin.Default()

	// ファイルのパス指定
	r.Static("/views", "./views")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	r.StaticFS("/shoppingapp", http.Dir("./views/static"))

	// 全ての商品情報のJSONを返す
	r.GET("/fetchAllProducts", controller.FetchAllProducts)

	// 1つの商品情報の状態のJSONを返す
	r.GET("/fetchProduct", controller.FindProduct)

	// 商品情報をDBへ登録する
	r.GET("/addProduct", controller.AddProduct)

	// 商品情報の状態を変更する
	r.POST("/changeStateProduct", controller.ChangeStateProduct)

	// 商品情報を削除する
	r.POST("/deleteProduct", controller.DeleteProduct)

	if err := r.Run(); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
