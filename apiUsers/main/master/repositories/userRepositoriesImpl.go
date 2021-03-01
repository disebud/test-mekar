package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/disebud/api-users-test/main/master/constanta"
	"github.com/disebud/api-users-test/main/master/models"
)

type UserRepoImpl struct {
	db *sql.DB
}

func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db}
}

func (s UserRepoImpl) GetAllUsers() ([]*models.User, error) {

	dataUsers := []*models.User{}
	// query := fmt.Sprintf(`select user_id ,username,date_of_birth,no_ktp, j.job_name as job, e.education_name as last_education
	// from user us
	// join job j on j.job_id=us.job
	// join education e on e.id=us.last_education
	// order by 1;`)

	query := constanta.GETALLUSERS

	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		Users := models.User{}
		var err = data.Scan(&Users.UserID, &Users.UserName, &Users.DateOfBirth, &Users.NoKtp, &Users.Job, &Users.Education)

		if err != nil {
			return nil, err
		}

		dataUsers = append(dataUsers, &Users)
	}

	return dataUsers, nil
}

func (s UserRepoImpl) GetJob() ([]*models.Job, error) {

	dataJob := []*models.Job{}
	query := constanta.GETJOB

	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		Job := models.Job{}
		var err = data.Scan(&Job.ID, &Job.Name)

		if err != nil {
			return nil, err
		}

		dataJob = append(dataJob, &Job)
	}

	return dataJob, nil
}

func (s UserRepoImpl) GetEducation() ([]*models.Education, error) {

	dataEducatiion := []*models.Education{}
	query := constanta.GETEDUCATION

	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		Educatiion := models.Education{}
		var err = data.Scan(&Educatiion.ID, &Educatiion.Name)

		if err != nil {
			return nil, err
		}

		dataEducatiion = append(dataEducatiion, &Educatiion)
	}

	return dataEducatiion, nil
}

func (s *UserRepoImpl) CreateUser(User models.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	log.Println(`Data`, User)

	query := constanta.CREATEUSER
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(User.UserName, User.DateOfBirth, User.NoKtp, User.Job, User.Education)
	if err != nil {
		return tx.Rollback()
	}
	stmt.Close()
	return tx.Commit()
}

func (s UserRepoImpl) GetUserById(UserId string) (*models.User, error) {
	User := new(models.User)
	// query := `SELECT mr.id_User , mr.name_User,mr.floor_location,mr.price,mrs.name_status_User
	// FROM
	// 	m_User mr
	// 		JOIN
	// 	m_status_User mrs ON mr.status_User = mrs.id_status_User  WHERE mr.id_User = ?;`
	query := constanta.GETUSERBYID
	if err := s.db.QueryRow(query, UserId).Scan(&User.UserID, &User.UserName, &User.DateOfBirth, &User.NoKtp, &User.Job, &User.Education); err != nil {
		return nil, err
	}
	return User, nil
}

func (s UserRepoImpl) DeleteUserById(UserID string) (*models.User, error) {
	User := new(models.User)
	query := "DELETE FROM m_User WHERE id_User = ?"
	_, err := s.db.Query(query, UserID)
	if err != nil {
		return nil, err
	}
	return User, nil

}

func (s UserRepoImpl) UpdateUser(UserID string, User models.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`UPDATE user SET username = ?, date_of_birth = ?, no_ktp = ?, job = ? , education = ? WHERE user_id = ?;`)
	if err != nil {
		return err
	}
	fmt.Println("Data", User)

	_, err = stmt.Exec(User.UserName, User.DateOfBirth, User.NoKtp, User.Job, User.Education, User.UserID)

	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}
