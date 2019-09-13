package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-errors/errors"

	"github.com/aliforever/gista/signatures"
)

var DefaultTempPath = ""

const (
	BoundaryChars  = "-_1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	BoundaryLength = 30
)

var lastUploadId *string = nil

func GenerateUploadId(useNano bool) (result string) {
	if !useNano {
		for true {
			result := NumberFormat(MicroTime()*1000, 0, "", "")
			if lastUploadId != nil && result == *lastUploadId {
				time.Sleep(time.Microsecond * 1000)
			} else {
				lastUploadId = &result
			}
		}
	} else {
		result = NumberFormat(MicroTime()-float64(getLastMonday()), 6, "", "")
		result += fmt.Sprintf("%d", MtRand(1, 999))
	}
	return
}

func HashCode(data string) int {
	result := 0
	for _, v := range data {
		//$result = (-$result + ($result << 5) + ord($string[$i])) & 0xFFFFFFFF;
		result = (-result + result<<5) + int(v)&0xFFFFFFFF
	}
	if result > 0x7FFFFFFF {
		result -= 0x100000000
	} else if result < -0x80000000 {
		result += 0x100000000
	}
	return result
}

func ReorderByHashCode(data map[string]interface{}) {
	hashCodes := map[string]interface{}{}
	for k := range data {
		hashCodes[k] = HashCode(k)
	}

}

func MicroTime() float64 {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	micSeconds := float64(now.Nanosecond()) / 1000000000
	return float64(now.Unix()) + micSeconds
}

func MtRand(min, max int64) int64 {

	return rand.Int63n(max-min+1) + min
}

func getLastMonday() int64 {
	now := time.Now()
	loc, _ := time.LoadLocation("GMT")
	now = time.Date(now.Year(), now.Month(), now.Day(), 24, 00, 00, 00, loc)
	start := 0
	nowMinus7 := now.Add(-(time.Hour * 24 * 7))
	for start < 7 {
		if nowMinus7.Weekday() == time.Monday {
			break
		}
		nowMinus7 = nowMinus7.Add(time.Hour * 24)
		start++
	}
	return nowMinus7.Unix()
}

func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

func CreateFolder(folderPath string) error {
	return os.MkdirAll(folderPath, os.ModePerm)
}

func FileOrFolderExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileGetContents(path string) (content string, err error) {
	data, err := ioutil.ReadFile(path)
	return string(data), err
}

func GenerateUserBreadCrumb(size int) string {
	key := "iN4$aGr0m"
	date := MicroTime() * 1000
	term := (rand.Intn(2)+2)*1000 + size*(rand.Intn(6)+15)*100
	textChangeEventCount := math.Round(float64(size / (rand.Intn(2) + 2)))
	if textChangeEventCount == 0 {
		textChangeEventCount = 1
	}
	data := fmt.Sprintf("%d %d %f %f", size, term, textChangeEventCount, date)
	hmc := hmac.New(sha256.New, []byte(key))
	hmc.Write([]byte(data))
	d := fmt.Sprintf("%s\n%s\n", hex.EncodeToString(hmc.Sum(nil)), base64.StdEncoding.EncodeToString([]byte(data)))
	return base64.StdEncoding.EncodeToString([]byte(d))
}

func ThrowIfInvalidRankToken(rankToken string) (err error) {
	if !signatures.IsValidUUID(rankToken) {
		err = errors.New(rankToken + " is not a valid tank token")
	}
	return
}
