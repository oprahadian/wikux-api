package database

const (
	SqlWikuappsGetUser = `SELECT * FROM wikuapps_user 
							WHERE email=LOWER(?) OR id=? LIMIT 1`
	
	SqlWikuappsGetSession = `SELECT * FROM wikuapps_session 
						WHERE id=? OR session_id=? LIMIT 1`

	SqlWikuappsLogin = `SELECT * FROM wikuapps_user 
								WHERE email=LOWER(?) AND password=sha1(?) LIMIT 1`
	
	SqlWikuappsInsertUserMinimal = `INSERT INTO wikuapps_user (email, create_time) 
										VALUES (LOWER(?), current_timestamp)`
	
	SqlWikuappsInsertSession = `INSERT INTO wikuapps_session (wikuapps_user_id, session_id, create_time) 
									VALUES (?, UUID(), current_timestamp)`

	SqlWikuappsUpdatePassword = `UPDATE wikuapps_user SET password=sha1(?),
									update_time=CURRENT_TIMESTAMP WHERE id=? OR email=LOWER(?) LIMIT 1`

	SqlWikuappsUpdateUser = `UPDATE wikuapps_user
							SET fullname=?, phone=?, generation=?, graduate_year=?, 
							domicile=?, occupation=?, update_time=CURRENT_TIMESTAMP
							WHERE id=? OR email=lower(?) LIMIT 1 `


)