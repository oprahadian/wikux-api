package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	db "wikux-api/database"
	"crypto/rand"
    "encoding/base64"
)

func generateRandomString(s int) (string, error) {
    b, err := func (n int) ([]byte, error) {
		b := make([]byte, n)
		_, err := rand.Read(b)
		if err != nil {
			return nil, err
		}
	
		return b, nil
	}(s)
    return base64.URLEncoding.EncodeToString(b), err
}

func UserRegistration(c *gin.Context){
	var ul db.UserLogin
	var err error

	if err = c.BindJSON(&ul); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})
		
		return
	}

	r, err := db.Wikufest.Exec(db.SqlInsertUser,
				ul.Email, ul.Password, ul.IsActive, ul.Fullname)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		userId, _  := r.LastInsertId()
		ul.Id = userId
		ul.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": ul,
		})
	}
}

func UserForgotPassword(c *gin.Context) {
	var ul db.UserLogin
	var err error

	if err = c.BindJSON(&ul); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	u, err := func(ul db.UserLogin) (db.UserLogin, error){
		var (
			u db.UserLogin
			err error
		)

		err = db.Wikufest.Get(&u, db.SqlUserByEmail, ul.Email)
		if err != nil {
			return db.UserLogin{},  err
		}
		
		return u, nil
	}(ul)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		np, _ := generateRandomString(16)

		go func(np string, u db.UserLogin){
			_, err := db.Wikufest.Exec(db.SqlUpdatePassword, np, u.Id)
			if err != nil {
				log.Print(err)
			}
		}(np, u)
		
		u.Password = np

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": u,
		})
	}
}

func UserPermissionList(c *gin.Context) {
	var (
		p db.UserPermission
		pl []db.UserPermission
		err error
	)

	if err = c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	if err = db.Wikufest.Select(&pl, db.SqlUserPermissionList, p.UserId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": pl,
		})
	}
}

func UserLoginCheck(c *gin.Context) {
	var (
		ul db.UserLogin
		err error
	)

	if err = c.BindJSON(&ul); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	u, err := func(ul db.UserLogin) (db.UserLogin, error){
		var (
			u db.UserLogin
			err error
		)

		err = db.Wikufest.Get(&u, db.SqlUserLogin, ul.Email, ul.Password); 
		if err != nil {
			return db.UserLogin{},  err
		}
		
		return u, nil
	}(ul)
	
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": u,
		})
	}
}