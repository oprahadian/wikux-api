package main

import (
	"github.com/gin-gonic/gin"
	"wikux-api/database"
	"wikux-api/handler"
	"wikux-api/config"
)

func main() {
	database.InitDB()
	
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "It works",
		})
	})

	func(){
		rg := r.Group("/User")
		rg.GET("/PermissionList", handler.UserPermissionList)
		rg.POST("/LoginCheck", handler.UserLoginCheck)
		rg.POST("/ForgotPassword", handler.UserForgotPassword)
		rg.POST("/Registration", handler.UserRegistration)
	}()

	func(){
		rg := r.Group("/Wikufest")
		rg.GET("/CourseSessionList", handler.Root)
		rg.GET("/CourseSessionEnrollmentList", handler.Root)
		rg.POST("/CourseSessionEnrollment", handler.Root)
	}()
	
	func(){
		rg := r.Group("/Wikusamacup")
		rg.GET("/SportList", handler.WikusamacupSportList)
		rg.GET("/SportTeamMatchScoreList", handler.WikusamacupSportTeamMatchScoreList)
		rg.GET("/SportTeamMemberList", handler.WikusamacupSportTeamMemberList)
		rg.GET("/SportTeamMatchScoreByMatchIdList", handler.WikusamacupSportTeamMatchScoreByMatchIdList)
		rg.POST("/SportCreate", handler.WikusamacupSportCreate)
		rg.POST("/SportTeamCreate", handler.WikusamacupSportTeamCreate)
		rg.POST("/SportTeamMatchCreate", handler.WikusamacupSportTeamMatchCreate)
		rg.POST("/SportTeamMemberCreate", handler.WikusamacupSportTeamMemberCreate)
		rg.POST("/SportTeamMemberScoreCreate", handler.WikusamacupSportTeamMemberScoreCreate)
	}()

	func(){
		rg := r.Group("/Wikusama")
		rg.GET("/WordpressPostList", handler.WikusamaWordpressPostList)
	}()

	func(){
		rg := r.Group("/Service")
		rg.POST	("/SendEmail", handler.ServiceSendEmail)
	}()
	r.Run("0.0.0.0:" + config.HttpPort)
}