package gin_util

import (
	"everdale-wiki/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RespStatusOK(ctx *gin.Context, args ...interface{}) {
	var data interface{}
	switch len(args) {
	case 1:
		data = args[0]
	default:
		logger.Warning.Json(map[string]interface{}{
			"flag": "RespStatusOK",
			"err":  "len(args) != 1",
		})
	}

	currentTime := time.Now().Unix()

	logger.Info.Json(map[string]interface{}{
		"flag":         "RespStatusOK",
		"current_time": currentTime,
		"data":         data,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code":        0,
		"message":     "success",
		"currentTime": currentTime,
		"data":        data,
	})
}
