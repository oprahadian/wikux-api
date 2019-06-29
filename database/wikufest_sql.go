package database

const (
	SqlWikufestCourseSessionAll = `select t3.instructor_name,t1.*, t2.* from course_session t1 
									inner join course t2 on (t1.course_id = t2.id)
									left join (
										select t1.course_id,GROUP_CONCAT(t2.fullname SEPARATOR ', ') instructor_name 
											from course_instructor t1 
										inner join user t2 on (t1.instructor_id = t2.id) 
										group by t1.course_id
									) t3 on (t2.id = t3.course_id)
									where t2.is_active = true and t1.start_time <= CURRENT_TIMESTAMP
										and t1.end_time > CURRENT_TIMESTAMP ORDER BY t1.start_time ASC`

	SqlWikufestCourseSessionByUserId = `select t4.instructor_name, t2.*,t3.* from course_session_enrollment t1 
								inner join course_session t2 on (t1.course_session_id = t2.id)
								inner join course t3 on (t2.course_id = t3.id)
								left join (
									select t1.course_id, GROUP_CONCAT(t2.fullname SEPARATOR ', ') instructor_name 
										from course_instructor t1 
									inner join app_user t2 on (t1.instructor_id = t2.id) 
									group by t1.course_id
								) t4 on (t3.id=t4.course_id)
								where t1.user_id=? ORDER BY t2.start_time ASC`

	SqlWikufestAvailableQuota = `SELECT available_quota FROM course_session WHERE id=? LIMIT 1`

	SqlWikufestInsertCourseSessionEnrollment = `INSERT INTO course_session_enrollment VALUES (?, ?, current_timestamp)`

	SqlWikufestUpdateMinAvailableQuota = `UPDATE course_session SET available_quota=available_quota-1 WHERE id=? AND available_quota >=1`
)