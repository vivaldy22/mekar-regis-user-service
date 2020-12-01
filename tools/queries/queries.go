package queries

const (
	GET_ALL_USERS  = `SELECT * FROM tb_user WHERE user_status = 1`
	GET_USER_BY_ID = `SELECT * FROM tb_user WHERE user_id = ? AND user_status = 1`
	CREATE_USER    = `INSERT INTO tb_user VALUES (?, ?, ?, ?, ?, ?, 1)`
	UPDATE_USER    = `UPDATE tb_user 
						SET user_name = ?,
							user_bday = ?,
							user_ktp = ?,
							user_job = ?,
							user_edu = ?
						WHERE user_id = ?
						AND user_status = 1`
	DELETE_USER = `UPDATE tb_user
					SET user_status = 0
					WHERE user_id = ?`

	GET_ALL_JOBS = `SELECT * FROM tb_job`
	GET_ALL_EDU  = `SELECT * FROM tb_edu`

	GET_PASSWORD = `SELECT admin_pass FROM tb_admin WHERE admin_uname = ?`
)
