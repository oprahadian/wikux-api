package handler


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"math/rand"
	"fmt"
	"time"
	"wikux-api/common"
	db "wikux-api/database"
)


type WikuappsUserDetailForm struct {
	SessionID string `json:"sessionID"`
	Email string `json:"email"`
}

type WikuappsUserLoginForm struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func getUser(userId int64, email string)(db.WikuappsUser, error){
	var (
		u db.WikuappsUser
		err error
	)

	err = db.Wikufest.Get(&u, db.SqlWikuappsGetUser, email, userId)
	if err != nil {
		return db.WikuappsUser{},  err
	}

	u.Password = ""

	return u, nil
}

func resetPassword(userId int64, email string)(error){
	var err error
	rand.Seed(time.Now().UnixNano())
	newPassword := fmt.Sprintf("%v", rand.Intn(999999 - 100000) + 100000)

	_, err = db.Wikufest.Exec(db.SqlWikuappsUpdatePassword, newPassword, userId, email)
	if err != nil {
		return err
	}

	go common.CommonSendEmailMailgun(common.CommonSingleSendEmail{
		To: email,
		Subject: fmt.Sprintf("%v Your new WikuApps password", newPassword),
		Body: fmt.Sprintf("%v Your new WikuApps password", newPassword),
		ReplyTo: "publikasi@wikusama.org",
	})

	return nil
}

func WikuappsUpdateUser(c *gin.Context){
	var f db.WikuappsUser
	var err error

	if err = c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	u, err := getUser(f.Id, f.Email)
	
	_, err = db.Wikufest.Exec(db.SqlWikuappsUpdateUser, f.Fullname, 
		f.Phone,
		f.Generation,
		f.GraduateYear,
		f.Domicile,
		f.Occupation,
		f.Id,
		f.Email)

	u, err = getUser(f.Id, f.Email)
	
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

func WikuappsForgetPassword(c *gin.Context){
	var f WikuappsUserLoginForm
	var err error

	if err = c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	err = resetPassword(0, f.Email)

	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
		})
	}
}

func WikuappsLogin(c *gin.Context){
	var f WikuappsUserLoginForm
	var err error

	if err = c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	u, s, err := func(f WikuappsUserLoginForm) (db.WikuappsUser, db.WikuappsSession, error){
		var (
			u db.WikuappsUser
			s db.WikuappsSession
			err error
		)

		err = db.Wikufest.Get(&u, db.SqlWikuappsLogin, f.Email, f.Password) 
		if err != nil {
			return db.WikuappsUser{}, db.WikuappsSession{}, err
		}

		r, err := db.Wikufest.Exec(db.SqlWikuappsInsertSession, u.Id)
		sessionId, _  := r.LastInsertId()
		
		err = db.Wikufest.Get(&s, db.SqlWikuappsGetSession, sessionId, "") 
		if err != nil {
			return db.WikuappsUser{}, db.WikuappsSession{}, err
		}
		
		return u, s, nil
	}(f)
	
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		u.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": gin.H{
				"user": u,
				"session": s,
			},
		})
	}
}


func WikuappsSessionDetail(c *gin.Context){
	var f WikuappsUserDetailForm
	var s db.WikuappsSession
	var u db.WikuappsUser
	var err error

	if err = c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}
	
	err = db.Wikufest.Get(&s, db.SqlWikuappsGetSession, 0, f.SessionID) 
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	u, err = getUser(s.WikuappsUserId, "")
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": gin.H{
				"user": u,
				"session": s,
			},
		})
	}
}

func WikuappsCheckUser(c *gin.Context){
	var f WikuappsUserDetailForm
	var err error

	if err = c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	u, err := func(f WikuappsUserDetailForm) (db.WikuappsUser, error){
		var (
			u db.WikuappsUser
			err error
		)

		err = db.Wikufest.Get(&u, db.SqlWikuappsGetUser, f.Email, 0) 

		if err != nil {
			r, err := db.Wikufest.Exec(db.SqlWikuappsInsertUserMinimal, f.Email)
			userId, _  := r.LastInsertId()

			err = db.Wikufest.Get(&u, db.SqlWikuappsGetUser, "", userId)
			if err != nil {
				return db.WikuappsUser{},  err
			}
		}
		
		err = resetPassword(u.Id, u.Email)
		if err != nil {
			return db.WikuappsUser{},  err
		}
		
		return u, nil
	}(f)
	
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		u.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": u,
		})
	}
}

func WikuappsUserDetail(c *gin.Context){
	var f WikuappsUserDetailForm
	var err error

	if err = c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	u, err := func(f WikuappsUserDetailForm) (db.WikuappsUser, error){
		var (
			u db.WikuappsUser
			err error
		)
		
		err = db.Wikufest.Get(&u, db.SqlWikuappsGetUser, f.Email, 0)
		if err != nil {
			return db.WikuappsUser{},  err
		}
		
		return u, nil
	}(f)
	
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		u.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": u,
		})
	}
}