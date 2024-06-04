package api

//import "io"
//import "os"
//import "fmt"
import (
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Logic(c *gin.Context) Result
}

type DataList struct {
	Count int64       `json:"count"`
	List  interface{} `json:"list"`
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Logic(c *gin.Context, val IServer) {
	if err := c.ShouldBind(val); err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, val.Logic(c))
}

func LogicJson(c *gin.Context, val IServer) {
	c.JSON(200, val.Logic(c))
}

func SBJson(c *gin.Context, val IServer) {
	if err := c.ShouldBindJSON(val); err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, val.Logic(c))
}

func SBUri(c *gin.Context, val IServer) {
	if err := c.ShouldBindUri(val); err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, val.Logic(c))
}

func SBHeader(c *gin.Context, val IServer) {
	if err := c.ShouldBindHeader(val); err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, val.Logic(c))
}
