package system

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/sirupsen/logrus"
)

func AddAccesstokenToRedis(id int64, accessToken string, log *logrus.Entry) {

	UserToAccessToken := fmt.Sprintf(constants.UserToAccessTokenList, strconv.FormatInt(id, 10))

	AccessTokenToUser := fmt.Sprintf(constants.AccessTokenToUser, accessToken)

	fmt.Println(UserToAccessToken, AccessTokenToUser)

	redis := SocialContext.Redis

	count, err := redis.LLen(context.TODO(), UserToAccessToken).Result()
	if err != nil {
		log.Errorln(err)
	}

	if count > 0 {

		accessTokens, err := redis.LRange(context.TODO(), UserToAccessToken, 0, 0).Result()
		if err != nil {
			log.Errorln(err)
		}

		for _, accessToken := range accessTokens {
			count, err := redis.Exists(context.TODO(), accessToken).Result()
			if err != nil {
				log.Errorln(err)
			}

			if count == 0 {
				redis.LRem(context.TODO(), UserToAccessToken, 1, accessToken)
			}
		}
	}
	redis.LPush(context.TODO(), UserToAccessToken, AccessTokenToUser)

	go RefreshUserToAccessTokenList(id)
}

func RefreshUserToAccessTokenList(id int64) {

	redis := SocialContext.Redis

	UserToAccessToken := fmt.Sprintf(constants.UserToAccessTokenList, strconv.FormatInt(id, 10))

	redis.Expire(context.TODO(), UserToAccessToken, constants.AccessTokenExpiry)
}

func RefreshAccessToken(accessToken string) {
	redis := SocialContext.Redis

	AccessTokenToUser := fmt.Sprintf(constants.AccessTokenToUser, accessToken)

	redis.Expire(context.TODO(), AccessTokenToUser, constants.AccessTokenExpiry)
}

func RefreshUserDetails(id int64, accessToken string) {
	RefreshUserToAccessTokenList(id)

	RefreshAccessToken(accessToken)

}
