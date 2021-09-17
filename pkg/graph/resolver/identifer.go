package resolver

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func toGlobalID(tblName string, sequentialID int) string {
	src := []byte(fmt.Sprintf("%s:%d", tblName, sequentialID))

	return base64.StdEncoding.EncodeToString(src)
}

func fromGlobalID(globalId string) (string, int) {
	bytes, err := base64.StdEncoding.DecodeString(globalId)

	if err != nil {
		// TODO: loggerにする
		fmt.Println("Can not decode string")
	}

	vals := strings.Split(string(bytes), ":")
	tblName := vals[0]
	sequentialID, _ := strconv.Atoi(vals[1])

	return tblName, sequentialID
}
