// 首页

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 返回首页
func Index(c *gin.Context) {
	data := gin.H{
		"PageTitle": "首页",
	}
	c.HTML(http.StatusOK, "stock.html", data)
	return
}
