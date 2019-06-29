package handler

import (
	"github.com/gin-gonic/gin"
	db "wikux-api/database"
	"net/http"
)

type WordressPostListForm struct {
	TermId int `json:"termId"`
}

func WikusamaWordpressPostList(c *gin.Context){
	var p []db.WikuPosts
	var f WordressPostListForm
	var err error

	if err = c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	err = db.Wordpress.Select(&p, db.SqlWordpressPosts, f.TermId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": p,
		})
	}
}