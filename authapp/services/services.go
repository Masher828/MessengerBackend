package services

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/Masher828/MessengerBackend/authapp/models"
	"github.com/Masher828/MessengerBackend/authapp/repository"
	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/Masher828/MessengerBackend/common-packages/system"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(user *models.UserModel, log *logrus.Entry) error {

	if _, err := user.IsValid(); err != nil {
		log.Errorln(err)
		return err
	}

	userDetails, err := repository.GetUserByEmail(user.Email, log)
	if err == nil && userDetails.Id != 0 {
		return system.EmailAlreadyExists
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), int(time.Now().Month()))
	if err != nil {
		log.Errorln(err)
		return err
	}

	user.Id = repository.GetNextHibernateSequence()
	user.Password = string(passwordBytes)

	err = repository.InsertUserToDB(user, log)
	if err != nil {
		log.Errorln(err)
	}

	return err
}

func UserSignIn(user *models.UserLoginModel, log *logrus.Entry) (*models.UserDetails, error) {
	if _, err := user.IsValid(); err != nil {
		log.Error(err)
		return nil, system.InvalidCredentialsErr
	}

	userDetails, err := repository.GetUserByEmail(user.Email, log)
	if err != nil || userDetails.Id == 0 {
		if err != nil {
			log.Errorln(err)
		}
		return nil, system.InvalidCredentialsErr
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))

	if err != nil {
		log.Errorln(err)
		fmt.Println("check")
		go IncrementIncorrectPasswordCountToRedis(user.Email, log)
		return nil, system.InvalidCredentialsErr
	}

	isLocked, err := repository.IsUserLocked(log, user.Email)
	fmt.Println(isLocked)
	if err != nil {
		log.Errorln(err)
	}

	if isLocked {
		return nil, system.AccountBlockedError
	}

	accessToken := uuid.NewString()
	redisDb := system.SocialContext.Redis

	userContext := userDetails.CreateUserContext()
	userContext.AccessToken = accessToken
	data, _ := json.Marshal(userContext)

	AccessTokenToUser := fmt.Sprintf(constants.AccessTokenToUser, accessToken)

	redisDb.Set(context.TODO(), AccessTokenToUser, data, constants.AccessTokenExpiry)

	go repository.UpdateLastLoginTime(log, userContext.Id)

	go system.AddAccesstokenToRedis(userContext.Id, userContext.AccessToken, log)

	//remove it later
	go ClearIncorrectPasswordCountFromRedis(log, user.Email)

	return userContext, nil
}

func IncrementIncorrectPasswordCountToRedis(emailId string, log *logrus.Entry) {

	redisDb := system.SocialContext.Redis

	key := fmt.Sprintf(constants.IncorrectPasswordCountEmail, emailId)
	data, err := redisDb.Get(context.TODO(), key).Result()
	if err != nil {
		log.Errorln(err)
	}

	count, err := strconv.Atoi(data)
	if err != nil {
		log.Errorln(err)
	}

	if count < constants.MaxIncorrectPasswordAllowed {
		redisDb.Set(context.TODO(), key, count+1, constants.IncorrectPasswordCountEmailExpiry)
	} else if count == constants.MaxIncorrectPasswordAllowed {
		err := repository.ToggleUserlock(log, emailId, true)
		if err != nil {
			log.Errorln(err)
		}
	}

}

func ClearIncorrectPasswordCountFromRedis(log *logrus.Entry, email string) {

	redisDb := system.SocialContext.Redis

	redisDb.Del(context.TODO(), fmt.Sprintf(constants.IncorrectPasswordCountEmail, email))

}

func GetAllUsers(log *logrus.Entry) (map[int64]string, error) {
	return repository.GetAllUsers(log)
}

func GetUsersBySearchPattern(userId int64, pattern string, offset, limit int64, log *logrus.Entry) ([]*models.UserSearchDetails, error) {

	return repository.GetUsersBySearchPattern(userId, regexp.QuoteMeta(pattern), offset, limit, log)
}
