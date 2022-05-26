package system

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/Masher828/MessengerBackend/common-packages/log"
	"github.com/Masher828/MessengerBackend/common-packages/models"
	"github.com/sirupsen/logrus"
	"github.com/zenazn/goji/web"
)

type Controller struct {
}

type Application struct {
}

func (application *Application) Route(controller interface{}, route string, isPublic bool) interface{} {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				logrus.Errorln("panic is caught", err)
				response := make(map[string]interface{})
				response["message"] = InternalServerErr
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		if !isPublic && c.Env[constants.AuthFailed].(bool) {
			http.Redirect(w, r, "", http.StatusUnauthorized)
		} else {

			var logger *logrus.Entry
			if !isPublic {
				userContext := c.Env[constants.UserContext].(models.UserModelContext)

				logger = log.GetDefaultLogger(userContext.Id, r.RequestURI, r.Method)

			} else {

				logger = log.GetDefaultLogger(0, r.RequestURI, r.Method)

			}

			methodInterface := reflect.ValueOf(controller).MethodByName(route).Interface()

			method := methodInterface.(func(c web.C, w http.ResponseWriter, r *http.Request, logger *logrus.Entry) ([]byte, error))
			response, err := method(c, w, r, logger)
			if err != nil {
				// TODO ADD FILTER FOR SYSTEM GENERATED ERROR Or THROWN BY USER so not leak important info to frontend
				if err.Error() == "EOF" {
					err = InvalidPayloadData
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte(response))
			}

		}
	}
	return fn
}

func (application *Application) ApplyAuth(c *web.C, h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		} else {
			accessToken, err := getAccessToken(r)
			if err != nil {
				fmt.Println(err)
				c.Env[constants.AuthFailed] = true
			} else {

				userContext, err := GetUserContext(accessToken)
				if err != nil {
					fmt.Println(err)
					fmt.Println("check")
					c.Env[constants.AuthFailed] = true
				} else {
					c.Env[constants.AuthFailed] = false
					c.Env[constants.UserContext] = *userContext
				}

			}

		}

		h.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}

func getAccessToken(r *http.Request) (string, error) {

	var accessToken string
	tokenBearer := r.Header.Get("Authorization")

	var err error = nil

	if len(tokenBearer) != 0 {
		s := strings.Split(tokenBearer, " ")

		if len(s) != 2 || strings.ToLower(s[0]) != "bearer" || len(s[1]) == 0 {
			return accessToken, UnauthorizedErr
		}
		accessToken = s[1]
	} else {
		err = UnauthorizedErr
	}

	return accessToken, err

}

func GetUserContext(accessToken string) (*models.UserModelContext, error) {

	redisDb := SocialContext.Redis

	key := "accessToken:"
	result := redisDb.Get(context.TODO(), key+accessToken)

	if result.Err() != nil {
		fmt.Println(result.Err())
		return nil, AccessTokenDoesNotExistErr
	}

	data, err := result.Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var userContext models.UserModelContext

	err = json.Unmarshal([]byte(data), &userContext)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &userContext, nil

}
