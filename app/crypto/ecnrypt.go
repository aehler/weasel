package crypto

import (
	"crypto/sha1"
	"math/rand"
	"time"
	"fmt"
)

const salt = "o0d*0sfJFMxWea2kd#sel#fajBee"

func Encrypt(s, ls string) string {

	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%s%s%s", salt, s, ls))))

}

func GenSessionId(i uint, l_salt string) string {

	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%s%d%s%s%d", salt, i, l_salt, time.Now(), rand.Int()))))

}
