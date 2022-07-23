package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getting(c *gin.Context) {
	raw, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	//st := Sample{message: "xxfff", data: "test"}
	//c.JSON(http.StatusOK, gin.H{"(test": "a", "test2": "b"})
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, string(raw))
}
