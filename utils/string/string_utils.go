package string

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// convert uuid to string , remove "-"
	uuidString := strings.ReplaceAll(newUUID.String(), "-", "")
	// 10clitokenijkasdmfasikdjfpomgasdfgl,masdl;gmsdfpgk
	return strconv.Itoa(userId) + "clitoken" + uuidString
}
