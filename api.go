package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Data map[int]Link
var lastID int

func main() {
	createDBConnection()
	lastID = 0
	//Data = make(map[int]Link)
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupRoutes(r *gin.Engine) {
	r.POST("/srl", createHandler)
	r.GET("/srl/:id", redirectHandler)
}

type Link struct {
	ID        int    `json:"user_id"`
	LongLink  string `json:"long_link" binding:"required"`
	ShortLink string `json:"short_link"`
}

// POST
func SaveLongLink(c *gin.Context) {
	reqBody := Link{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": "invalid request body",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	lastID++
	reqBody.ID = lastID
	reqBody.ShortLink = "http://localhost:8080/srl/" + fmt.Sprint(lastID)
	Data[lastID] = reqBody
	c.JSON(http.StatusOK, reqBody)
	c.Writer.Header().Set("Content-Type", "application/json")
	return
}

// GET

func GetLongLink(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		res := gin.H{
			"error": "invalid request body",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		res := gin.H{
			"error": "invalid request body",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	if _, ok := Data[idInt]; !ok {
		res := gin.H{
			"error": "link not found",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	c.Redirect(http.StatusMovedPermanently, Data[idInt].LongLink)
	return
}
