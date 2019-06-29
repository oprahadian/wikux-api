package database

const (
	SqlWikusamacupSportList = `select * from sport order by name`
	
	SqlWikusamacupSportTeamCreate = `INSERT INTO sport_team
			(team_name, person_in_charge_id, sport_id, reg_date, is_active)
			VALUES(?, ?, ?, current_timestamp, ?)`

	SqlWikusamacupSportTeamMatchCreate = `INSERT INTO sport_team_match
			(sport_team_home_id, sport_team_away_id, start_date, end_date, is_active, reg_date)
			VALUES(?, ?, ?, ?, ?, current_timestamp)`

	SqlWikusamacupSportTeamMemberCreate = `INSERT INTO sport_team_member
			(position_name, team_member_id, sport_team_id, reg_date)
			VALUES(?, ?, ?, current_timestamp)`

	SqlWikusamacupSportTeamMemberScoreCreate = `INSERT INTO sport_team_member_score
			(sport_team_match_id, sport_team_member_id, score, reg_date)
			VALUES(?, ?, ?, current_timestamp)`

	SqlWikusamacupSportTeamMatchScoreList =	`select 
			tm.id sport_team_match_id,
				sth.id sport_team_home_id, 
			 sth.team_name sport_team_home_name,
			 sta.id sport_team_away_id, 
			 sta.team_name sport_team_away_name,
			 IFNULL(sum(tmsh.score), 0) score_team_home,
			 IFNULL(sum(tmsa.score),0) score_team_away
			from sport_team_match tm
				left join sport_team sth on (tm.sport_team_home_id = sth.id)
				left join sport_team sta on (tm.sport_team_away_id = sta.id)
				left join sport_team_member stmh on (sth.id = stmh.sport_team_id)
				left join sport_team_member stma on (sta.id = stma.sport_team_id)
				 left join sport_team_member_score tmsh on (tmsh.sport_team_match_id = tm.id AND stmh.id = tmsh.sport_team_member_id)
				 left join sport_team_member_score tmsa on (tmsa.sport_team_match_id = tm.id AND stma.id = tmsa.sport_team_member_id)
				 group by tm.id,sth.id,sth.team_name,sta.id, sta.team_name`

)