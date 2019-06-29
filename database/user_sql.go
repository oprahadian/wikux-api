package database

const (
	SqlUserLogin = `SELECT * FROM user 
							WHERE email=lower(?)
							AND password=sha(?) AND is_active=true LIMIT 1`

	SqlUserByEmail = `SELECT * FROM user 
						WHERE email=lower(?) LIMIT 1`

	SqlUpdatePassword = `UPDATE user SET password=sha(?), upd_date=CURRENT_TIMESTAMP WHERE id=?`

	SqlInsertUser = `INSERT INTO user(email, password, is_active, fullname, reg_date)
						VALUES(?, sha(?), ?, ?, current_timestamp)`

	SqlUserPermissionList = `select * from user_permission where user_id=?`

)