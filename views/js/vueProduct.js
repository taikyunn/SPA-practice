new Vue({
  // 「el」プロパティーでVueの表示を反映する場所＝HTML要素のセレクター（id）を定義
  // つまりid=appの範囲に対して操作を適用するということ
  el: '#app',

  // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
  data: {
      // 商品情報
      products: [],
      // 品名
      productName: '',
      // メモ
      productMemo: '',
      // 商品情報の状態
      current: -1,
      // 商品情報の状態一覧
      options: [
          { value: -1, label: 'すべて' },
          { value:  0, label: '未購入' },
          { value:  1, label: '購入済' }
      ],
      // true：入力済・false：未入力
      isEntered: false
  },

  // 算出プロパティ
  computed: {
      // 商品情報の状態一覧を表示する
      labels() {
          // reduceメソッド：配列の要素を一つずつ取り出し、指定した処理を行なっていき最終的に一つの値を返す関数
          // a:アキュムレータb:現在値
          // b:optionsの中身をさしていそう。
          return this.options.reduce(function (a, b) {
              // Object.assignメソッド：全ての列挙可能な自身のプロパティの値を1つ以上コピー元オブジェクトからコピー先オブジェクトにコピーする
              return Object.assign(a, { [b.value]: b.label })
          }, {})
      },
      // 表示対象の商品情報を返却する
      computedProducts() {
        return this.products.filter(function (el) {
          var option = this.current < 0 ? true : this.current === el.state
          return option
        }, this)
      },
      // 入力チェック
      validate() {
          var isEnteredProductName = 0 < this.productName.length
          this.isEntered = isEnteredProductName
          return isEnteredProductName
      }
  },

  // インスタンス作成時の処理
  // インスタンスが作成された後き同期的に呼ばれる。マウンティングの前に実行される
  created: function() {
      this.doFetchAllProducts()
  },

  // メソッド定義
  methods: {
      // 全ての商品情報を取得する
      doFetchAllProducts() {
          axios.get('/fetchAllProducts')
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // 取得してきた値をresultProductsにおく
                  var resultProducts = response.data

                  // サーバから取得した商品情報をdataに設定する
                  this.products = resultProducts
              }
          })
      },
      // １つの商品情報を取得する
      doFetchProduct(product) {
          // 送られてきたproduct.idを元にDBデータを取得してくる
          // 今回はデータの更新をdoChangeProductStateで行なったので、更新したデータを再度取得し直して返す。
          axios.get('/fetchProduct', {
              params: {
                  productID: product.id
              }
          })
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // 取得してきたデータを含め諸々を一度resultProductという変数に定義
                  var resultProduct = response.data

                  // 選択された商品情報のインデックスを取得する
                  // インデックスは表のNoのこと。つまり表の何番に再描画するかの指定をする
                  var index = this.products.indexOf(product)

                  // spliceを使うとdataプロパティの配列の要素をリアクティブに変更できる
                  // resultProduct[0]はデータベースに登録された値を表示している。
                  // this.productsは2つの配列が用意されている。1つめは今登録した値を取得してきた時の値。二つ目は表全体の値を保持している

                  // spliceは既存の要素を取り除いたり置き換えたり、新しく追加したりすることができる。
                  // spliceは
                  // 第一引数：変更したい配列の位置の指定。
                  // 第二引数：消す要素数
                  // 第三引数：変更要素
                  console.log(this.products)
                  this.products.splice(index, 1, resultProduct[0])
              }
          })
      },
      // 商品情報を登録する
      doAddProduct() {
          // サーバへ送信するパラメータ
          const params = new URLSearchParams();
          params.append('productName', this.productName)
          params.append('productMemo', this.productMemo)

          axios.post('/addProduct', params)
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // 商品情報を取得する
                  this.doFetchAllProducts()

                  // 入力値を初期化する
                  this.initInputValue()
              }
          })
      },
      // 商品情報の状態を変更する
      doChangeProductState(product) {
          // サーバへ送信するパラメータ
          const params = new URLSearchParams();
          //product.idでパラメータ情報のidにアクセスできる
          params.append('productID', product.id)
          //product.stateでパラメータ情報のstateにアクセスできる
          params.append('productState', product.state)

          // postメソッドで送信。第一引数が送信先のURL。第二引数がパラメータ
          // ここで商品状態を変更しDBに登録する
          axios.post('/changeStateProduct', params)
          // レスポンスが返ってきた後の処理
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // productには状態変更に使用したデータ情報が入っている。
                  // 商品情報を取得する
                  this.doFetchProduct(product)
              }
          })
      },
      // 商品情報を削除する
      doDeleteProduct(product) {
          // サーバへ送信するパラメータ
          const params = new URLSearchParams();
          params.append('productID', product.id)

          axios.post('/deleteProduct', params)
          .then(response => {
              if (response.status != 200) {
                  throw new Error('レスポンスエラー')
              } else {
                  // 商品情報を取得する
                  this.doFetchAllProducts()
              }
          })
      },
      // 入力値を初期化する
      initInputValue() {
          this.current = -1
          this.productName = ''
          this.productMemo = ''
      }
  }
})
