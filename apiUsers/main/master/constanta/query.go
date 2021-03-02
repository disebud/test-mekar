package constanta

const (

	// // GETALLUSERS menampilkan semua data user
	// GETALLUSERS = `SELECT user_id ,username,date_of_birth,no_ktp, j.job_name AS job, e.education_name AS last_education
	// FROM user us
	// JOIN job j ON j.job_id=us.job
	// JOIN education e ON e.id=us.education
	// ORDER BY 1;`

	// GETALLUSERS menampilkan semua data user dengan format yang sudah ada
	GETALLUSERS = `SELECT * FROM user ;`

	// GETJOB menampilkan semua data job
	GETJOB = `SELECT * FROM job ;`

	// GETEDUCATION menampilkan semua data education
	GETEDUCATION = `SELECT * FROM education ;`

	// CREATEUSER membuat data user baru
	CREATEUSER = `INSERT INTO user (username,date_of_birth,no_ktp,job,education) VALUES (?, ?, ?, ?, ?);`

	//GETUSERBYID mengambil data detail satu user dengan Id
	GETUSERBYID = ` SELECT user_id ,username,date_of_birth,no_ktp, j.job_name AS job, e.education_name AS last_education
	FROM user us
	JOIN job j ON j.job_id=us.job
	JOIN education e ON e.id=us.education
	WHERE user_id = ?;`

	//UPDATEUSER merubah data user
	UPDATEUSER = `UPDATE user SET username = ?, date_of_birth = ?, no_ktp = ?, job = ? , education = ? WHERE user_id = ?;`

	GETALLRESERVATION = `SELECT mr.id_room , mr.name_room,mr.floor_location,mr.price,mrs.id_status_room,mrs.name_status_room
							  FROM m_room mr JOIN m_status_room mrs ON mr.status_room = mrs.id_status_room  ORDER BY mr.id_room ASC;`
	CREATERESERVATION = `INSERT INTO m_room VALUES (?, ?, ?, ?, ?);`
)
