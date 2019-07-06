package database

const (
	SqlWikusamacupSportList = `select * from sport order by name`
	
	SqlWikusamacupSportTeamCreate = `INSERT INTO sport_team
			(team_name, person_in_charge_id, sport_id, created_date, is_active)
			VALUES(?, ?, ?, current_timestamp, ?)`

	SqlWikusamacupSportTeamMatchCreate = `INSERT INTO sport_team_match
			(sport_team_home_id, sport_team_away_id, start_date, end_date, is_active, created_date)
			VALUES(?, ?, ?, ?, ?, current_timestamp)`

	SqlWikusamacupSportTeamMemberCreate = `INSERT INTO sport_team_member
			(position_name, team_member_id, sport_team_id, created_date)
			VALUES(?, ?, ?, current_timestamp)`

	SqlWikusamacupSportTeamMemberList = `select u.id user_id, u.fullname,tm.position_name, tm.sport_team_id FROM sport_team_member tm
										left join user u on (tm.team_member_id = u.id)
											where sport_team_id=?`

	SqlWikusamacupSportTeamMemberScoreCreate = `INSERT INTO sport_team_member_score
			(sport_team_match_id, sport_team_member_id, score, score_date, created_date)
			VALUES(?, ?, ?, ?,current_timestamp)`

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

	SqlWikusamacupSportCreate = `INSERT INTO wikufest_db.sport
									(name, is_active) VALUES(?, ?)`


	SqlWikusamacupSportTeamMatchScoreByMatchIdList = `select t5.id sport_team_match_id,t3.fullname,t4.id sport_team_id,
	t4.team_name sport_team_name,t1.score,IFNULL(t1.score_date, t1.created_date) score_date,  
	TIMESTAMPDIFF(MINUTE, t5.start_date,IFNULL(t1.score_date, t1.created_date)) score_minute_from_start  from sport_team_member_score t1 
		inner join sport_team_member t2 on (t1.sport_team_member_id = t2.id)
		inner join user t3 on (t2.team_member_id = t3.id)
		inner join sport_team t4 on (t2.sport_team_id = t4.id)
		inner join sport_team_match t5 on (t1.sport_team_match_id = t5.id)
			where t5.id = ?
			order by t4.id asc, 6 asc
		`

)