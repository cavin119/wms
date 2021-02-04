package resource

import (
	"7youo-wms/app/response"
	"github.com/gin-gonic/gin"
)

func GetMenus(c *gin.Context) {
	response.Ok(c)
}
