// パッケージ名
package main

// パッケージのインポート
import (
	"log"
	"time"

	// "test/routers"

	"github.com/Simo-C3/tyoukankakuC3-2023/api/config"
	"github.com/Simo-C3/tyoukankakuC3-2023/api/db"
	"github.com/Simo-C3/tyoukankakuC3-2023/api/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// OpenAPI
	_ "github.com/Simo-C3/tyoukankakuC3-2023/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title   超感覚C3-2023 API
// @version  1.0
// @license.name Simo-C3
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config.LoadConfig()

	e := gin.Default()

	// cors設定
	e.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
				"http://localhost:3000",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
				"POST",
				"GET",
				"PUT",
				"DELETE",
				"PATCH",
				"OPTIONS",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
				"Content-Type",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
}))


	routers.Router(e)

	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	db.Init()
	db.AutoMigration()

	if err := e.Run("0.0.0.0:8000"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}

	db.Close()
}
