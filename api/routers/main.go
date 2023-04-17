package routers

// パッケージのインポート
import (
	// "github.com/Simo-C3/tyoukankakuC3-2023/api/routers"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	api := e.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			InitToDo(v1)
		}
	}
}
