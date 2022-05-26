package system

import (
	"strconv"
	"time"

	"github.com/Masher828/MessengerBackend/common-packages/constants"
	"github.com/sirupsen/logrus"
)

func GetUTCTime() time.Time {
	return time.Now().UTC()
}

func GetOffsetAndLimit(incomingOffset, incomingLimit []string, defaultOffset, defaultLimit int64, log *logrus.Entry) (int64, int64) {

	offset := defaultOffset

	var err error = nil

	if len(incomingOffset) > 0 {
		offset, err = strconv.ParseInt(incomingOffset[0], 10, 64)
		if err != nil {
			log.Errorln(err)
			offset = constants.DefaultOffset
		}
	}

	limit := defaultLimit

	if len(incomingLimit) > 0 {
		limit, err = strconv.ParseInt(incomingLimit[0], 10, 64)
		if err != nil {
			log.Errorln(err)
			limit = constants.DefaultLimit
		}
	}
	return offset, limit
}
