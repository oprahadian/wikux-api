package database

type WikusamacupSport struct {
	Id int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	IsActive bool `json:"isActive" db:"is_active"`
}

type WikusamacupSportTeam struct {
	Id int64 `json:"id" db:"id"`
	TeamName string `json:"teamName" db:"team_name"`
	PersonInChargeId int64 `json:"personInChargeId" db:"person_in_charge_id"`
	SportId int64 `json:"sportId" db:"sport_id"`
	CreatedDate string `json:"createdDate" db:"created_date"`
	IsActive bool `json:"isActive" db:"is_active"`
}

type WikusamacupSportTeamMatch struct {
	Id int64 `json:"id" db:"id"`
	SportTeamHomeId  int64`json:"sportTeamHomeId" db:"sport_team_home_id"`
	SportTeamAwayId int64 `json:"sportTeamAwayId" db:"sport_team_away_id"`
	StartDate string `json:"startDate" db:"start_date"`
	EndDate string `json:"endDate" db:"end_date"`
	IsActive bool `json:"isActive" db:"is_active"`
	CreatedDate string `json:"createdDate" db:"created_date"`
}

type WikusamacupSportTeamMember struct {
	Id int64 `json:"id" db:"id"`
	PositionName string `json:"positionName" db:"position_name"`
	TeamMemberId  int64`json:"teamMemberId" db:"team_member_id"`
	SportTeamId  int64`json:"sportTeamId" db:"sport_team_id"`
	CreatedDate string `json:"createdDate" db:"created_date"`
}

type WikusamacupSportTeamScore struct {
	Id int64 `json:"id" db:"id"`
	SportTeamMatchId int64 `json:"sportTeamMatchId" db:"sport_team_match_id"`
	SportTeamMemberId int64 `json:"sportTeamMemberId" db:"sport_team_member_id"`
	Score int `json:"score" db:"score"`
	ScoreDate string `json:"scoreDate" db:"score_date"`
	CreatedDate string `json:"createdDate" db:"created_date"`
}

type WikusamacupSportTeamMatchScore struct {
	SportTeamMatchId int64 `json:"sportTeamMatchId" db:"sport_team_match_id"`
	SportTeamHomeId int64 `json:"sportTeamHomeId" db:"sport_team_home_id"`
	SportTeamHomeName string `json:"sportTeamHomeName" db:"sport_team_home_name"`
	SportTeamAwayId int64 `json:"sportTeamAwayId" db:"sport_team_away_id"`
	SportTeamAwayName string `json:"sportTeamAwayName" db:"sport_team_away_name"`
	ScoreTeamHome int  `json:"scoreTeamHome" db:"score_team_home"`
	ScoreTeamAway int  `json:"scoreTeamAway" db:"score_team_away"`
}

type WikusamacupSportTeamMatchScoreByMatchId struct {
	SportTeamMatchId int64 `json:"sportTeamMatchId" db:"sport_team_match_id"`
	Fullname string `json:"fullname" db:"fullname"`
	SportTeamId int64 `json:"sportTeamId" db:"sport_team_id"`
	SportTeamName string `json:"sportTeamName" db:"sport_team_name"`
	Score int `json:"score" db:"score"`
	ScoreDate string `json:"scoreDate" db:"score_date"`
	ScoreMinuteFromStart int `json:"scoreMinuteFromStart" db:"score_minute_from_start"`
}