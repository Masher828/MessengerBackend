package repository

import (
	"database/sql"
	"fmt"

	"github.com/Masher828/MessengerBackend/authapp/models"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func GetNextHibernateSequence() int64 {
	query := "select nextval('hibernate_sequence')"
	db := system.SocialContext.PostgresDB
	var id int64
	db.QueryRow(query).Scan(&id)
	return id
}

func InsertUserToDB(user *models.UserModel, log *logrus.Entry) error {
	query := `INSERT INTO social_user (id, name, email, password, contact, country_code, country, is_locked,
		date_of_birth, last_updated, date_created) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	db := system.SocialContext.PostgresDB

	now := system.GetUTCTime()
	_, err := db.Exec(query, user.Id, user.FullName, user.Email, user.Password, user.Contact, user.CountryCode, false,
		user.Country, user.DateOfBirth, now, now)

	if err != nil {
		log.Errorln(err)
	}
	return err
}

func GetUserByEmail(email string, log *logrus.Entry) (*models.UserModel, error) {
	query := `SELECT id, name, email, password FROM social_user WHERE email = $1`

	db := system.SocialContext.PostgresDB

	var user models.UserModel

	err := db.QueryRow(query, email).Scan(&user.Id, &user.FullName, &user.Email, &user.Password)

	if err != nil {
		log.Errorln(err)
	}
	return &user, err
}

func GetAllUsers(log *logrus.Entry) (map[int64]string, error) {
	query := "SELECT name, id FROM social_user"

	db := system.SocialContext.PostgresDB

	rows, err := db.Query(query)
	if err != nil {
		log.Errorln(err)
		return map[int64]string{}, err
	}

	names := make(map[int64]string)
	for rows.Next() {
		var name string
		var id int64
		err = rows.Scan(&name, &id)
		if err != nil {
			log.Error(err)
			continue
		}

		names[id] = name
	}
	return names, nil
}

func GetUsersById(log *logrus.Entry, userIds interface{}) (map[int64]*models.UserModel, int, error) {

	db := system.SocialContext.PostgresDB

	query := "SELECT id, name, email, contact, country_code, country, date_of_birth, date_created, last_updated, last_login " +
		"from social_user where id = ANY($1)"

	rows, err := db.Query(query, pq.Array(userIds))
	if err != nil {
		log.Errorln(err)
		return nil, 0, err
	}
	count := 0
	var id int64

	result := make(map[int64]*models.UserModel)

	var name, email, contact, country_code, country sql.NullString

	var date_created, date_of_birth, last_updated, last_login sql.NullTime
	for rows.Next() {

		err = rows.Scan(&id, &name, &email, &contact, &country_code, &country, &date_of_birth, &date_created, &last_updated, &last_login)
		if err != nil {
			log.Errorln(err)
			continue
		}

		count += 1

		user := models.UserModel{
			Id:          id,
			FullName:    name.String,
			Email:       email.String,
			Contact:     contact.String,
			CountryCode: country_code.String,
			DateCreated: date_created.Time,
			LastUpdated: last_updated.Time,
			LastLogin:   last_login.Time,
		}

		result[id] = &user

	}

	return result, count, err

}

func CheckIfUsersExist(log *logrus.Entry, userIds []int64) (bool, error) {
	db := system.SocialContext.PostgresDB

	query := "SELECT count(*) from social_user where id = ANY($1)"

	var count int
	err := db.QueryRow(query, pq.Array(userIds)).Scan(&count)
	if err != nil {
		log.Errorln(err)
	}

	return count == len(userIds), err
}

func UpdateLastLoginTime(log *logrus.Entry, userId int64) error {
	db := system.SocialContext.PostgresDB

	query := "UPDATE social_user SET last_login = $1 WHERE id = $2"

	now := system.GetUTCTime()

	_, err := db.Exec(query, now, userId)
	if err != nil {
		log.Errorln(err)
	}

	return err
}

func ToggleUserlock(log *logrus.Entry, emailId string, lock bool) error {

	db := system.SocialContext.PostgresDB

	query := "UPDATE social_user SET is_locked = $1 WHERE email = $2"

	_, err := db.Query(query, lock, emailId)
	if err != nil {
		log.Errorln(err)
	}
	return err
}

func IsUserLocked(log *logrus.Entry, emailId string) (bool, error) {

	db := system.SocialContext.PostgresDB

	query := "SELECT is_locked FROM social_user WHERE email = $1"

	var isLocked sql.NullBool

	err := db.QueryRow(query, emailId).Scan(&isLocked)

	if err != nil {
		log.Errorln(err)
	}

	return isLocked.Bool, err

}

func GetUsersBySearchPattern(userid int64, pattern string, offset, limit int64, log *logrus.Entry) ([]*models.UserSearchDetails, error) {

	db := system.SocialContext.PostgresDB

	pattern = "'" + pattern + "%'"

	query := "SELECT id, name, email FROM social_user WHERE email like %s or name like %s OFFSET %d LIMIT %d ;"

	rows, err := db.Query(fmt.Sprintf(query, pattern, pattern, offset, limit))
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	var result []*models.UserSearchDetails

	for rows.Next() {
		var user models.UserSearchDetails
		var (
			id    int64
			name  string
			email string
		)

		err = rows.Scan(&id, &name, &email)

		if err != nil {
			log.Errorln(err)
			continue
		}

		user.Id = id
		user.FullName = name
		user.Email = email
		result = append(result, &user)

	}

	return result, nil

}
