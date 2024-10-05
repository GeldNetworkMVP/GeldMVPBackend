package commons

import (
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
)

func ByteStingToInteger(byteValue string) (int64, error) {
	strVal := []byte(byteValue)
	encodedString := hex.EncodeToString(strVal)
	intValue, err := strconv.ParseInt(encodedString, 16, 64)
	if err != nil {
		logrus.Printf("Conversion failed: %s\n", err)
		return 0, errors.New("Conversion failed: %s\n" + err.Error())
	} else {
		return intValue, nil
	}
}
