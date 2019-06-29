package handler

import (
	"github.com/gin-gonic/gin"
	db "wikux-api/database"
	"net/http"
)

type sportTeamMemberList struct {
	UserId int64 `json:"userId" db:"user_id"`
	Fullname string `json:"fulllname" db:"fullname"`
	PositionName string `json:"positionName" db:"position_name"`
	SportTeamId int64 `json:"sportTeamId" db:"sport_team_id"`
}

func WikusamacupSportTeamMatchScoreList(c *gin.Context){
	var (
		s []db.WikusamacupSportTeamMatchScore
		err error
	)

	err = db.Wikufest.Select(&s, db.SqlWikusamacupSportTeamMatchScoreList)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": s,
		})
	}
}

func WikusamacupSportTeamMemberScoreCreate(c *gin.Context){
	var (
		s db.WikusamacupSportTeamScore
		err error
	)

	if err = c.BindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	r, err := db.Wikufest.Exec(db.SqlWikusamacupSportTeamMemberScoreCreate, s.SportTeamMatchId, s.SportTeamMemberId, s.Score)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"message": err.Error(),
		})
	} else {
		id, _  := r.LastInsertId()
		s.Id = id

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": s,
		})
	}
}

func WikusamacupSportTeamMemberCreate(c *gin.Context){
	var (
		m db.WikusamacupSportTeamMember
		err error
	)

	if err = c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	r, err := db.Wikufest.Exec(db.SqlWikusamacupSportTeamMemberCreate, m.PositionName, m.TeamMemberId, m.SportTeamId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"message": err.Error(),
		})
	} else {
		id, _  := r.LastInsertId()
		m.Id = id

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": m,
		})
	}
}

func WikusamacupSportTeamMemberList(c *gin.Context){
	var (
		m sportTeamMemberList
		ms []sportTeamMemberList
		err error
	)

	if err = c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	err = db.Wikufest.Select(&ms, db.SqlWikusamacupSportTeamMemberList, m.SportTeamId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": ms,
		})
	}
}

func WikusamacupSportTeamMatchCreate(c *gin.Context){
	var (
		m db.WikusamacupSportTeamMatch
		err error
	)

	if err = c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	r, err := db.Wikufest.Exec(db.SqlWikusamacupSportTeamMatchCreate, m.SportTeamHomeId, m.SportTeamAwayId, 
			m.StartDate, m.EndDate, m.IsActive)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"message": err.Error(),
		})
	} else {
		id, _  := r.LastInsertId()
		m.Id = id

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": m,
		})
	}
}

func WikusamacupSportTeamCreate(c *gin.Context){
	var (
		st db.WikusamacupSportTeam
		err error
	)

	if err = c.BindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})

		return
	}

	r, err := db.Wikufest.Exec(db.SqlWikusamacupSportTeamCreate, st.TeamName, 
					st.PersonInChargeId, st.SportId, st.IsActive)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		id, _  := r.LastInsertId()
		st.Id = id

		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": st,
		})
	}
}

func WikusamacupSportList(c *gin.Context){
	var (
		s []db.WikusamacupSport
		err error
	)

	if err = db.Wikufest.Select(&s, db.SqlWikusamacupSportList); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": nil,
			"data": s,
		})
	}

}