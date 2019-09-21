package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "golang.org/x/image/webp"

	"github.com/rwcarlsen/goexif/tiff"

	"github.com/rwcarlsen/goexif/exif"
)

var DefaultTempPath = ""

const (
	BoundaryChars  = "-_1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	BoundaryLength = 30
)

var lastUploadId *string = nil

func SaveImage(path string, img image.Image) (err error) {
	var f *os.File
	f, err = os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	err = jpeg.Encode(f, img, nil)
	return
}

func CreateTempFile(outputDir string, namePrefix *string) (f *os.File, err error) {
	finalPrefix := "TEMP"
	if namePrefix != nil {
		finalPrefix = *namePrefix
	}
	finalPrefix = fmt.Sprintf("INSTA%s_", finalPrefix)
	f, err = ioutil.TempFile(outputDir, finalPrefix)
	if err != nil {
		return
	}
	return
}

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

func IsDirectory(filePath string) bool {
	if fi, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		switch mode := fi.Mode(); {
		case mode.IsDir():
			return true
		}
	}
	return true
}

func RemoveDirectory(filePath string) bool {
	if fi, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		switch mode := fi.Mode(); {
		case mode.IsDir():
			return true
		}
	}
	return true
}

func Realpath(fpath string) (string, error) {

	if len(fpath) == 0 {
		return "", os.ErrInvalid
	}

	if !filepath.IsAbs(fpath) {
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		fpath = filepath.Join(pwd, fpath)
	}

	path := []byte(fpath)
	nlinks := 0
	start := 1
	prev := 1
	for start < len(path) {
		c := nextComponent(path, start)
		cur := c[start:]

		switch {

		case len(cur) == 0:
			copy(path[start:], path[start+1:])
			path = path[0 : len(path)-1]

		case len(cur) == 1 && cur[0] == '.':
			if start+2 < len(path) {
				copy(path[start:], path[start+2:])
			}
			path = path[0 : len(path)-2]

		case len(cur) == 2 && cur[0] == '.' && cur[1] == '.':
			copy(path[prev:], path[start+2:])
			path = path[0 : len(path)+prev-(start+2)]
			prev = 1
			start = 1

		default:

			fi, err := os.Lstat(string(c))
			if err != nil {
				return "", err
			}
			if isSymlink(fi) {

				nlinks++
				if nlinks > 16 {
					return "", os.ErrInvalid
				}

				var link string
				link, err = os.Readlink(string(c))
				after := string(path[len(c):])

				// switch symlink component with its real path
				path = switchSymlinkCom(path, start, link, after)

				prev = 1
				start = 1
			} else {
				// Directories
				prev = start
				start = len(c) + 1
			}
		}
	}

	for len(path) > 1 && path[len(path)-1] == os.PathSeparator {
		path = path[0 : len(path)-1]
	}
	return string(path), nil

}

// switch a symbolic link component to its real path
func switchSymlinkCom(path []byte, start int, link, after string) []byte {

	if link[0] == os.PathSeparator {
		// Absolute links
		return []byte(filepath.Join(link, after))
	}

	// Relative links
	return []byte(filepath.Join(string(path[0:start]), link, after))
}

func isSymlink(fi os.FileInfo) bool {
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}

func nextComponent(path []byte, start int) []byte {
	v := bytes.IndexByte(path[start:], os.PathSeparator)
	if v < 0 {
		return path
	}
	return path[0 : start+v]
}

func GetImageDimension(imagePath string) (width int, height int, err error) {
	var file *os.File
	file, err = os.Open(imagePath)
	if err != nil {
		return
	}
	defer file.Close()
	var img image.Config
	img, _, err = image.DecodeConfig(file)
	if err != nil {
		return
	}
	width = img.Width
	height = img.Height
	return
}

func GetImageOrientation(imagePath string) (orientation string, err error) {
	var file *os.File
	file, err = os.Open(imagePath)
	if err != nil {
		return
	}
	defer file.Close()

	file.Seek(0, io.SeekStart)
	var x *exif.Exif
	x, err = exif.Decode(file)
	if err != nil {
		orientation = "1"
		return
	}
	if x != nil {
		var orient *tiff.Tag
		orient, err = x.Get(exif.Orientation)
		if err != nil {
			orientation = "1"
			return
		}
		if orient != nil {
			orientation = orient.String()
			return
		}
	}
	orientation = "1"
	return
}

func GetFileSize(filePath string) (size int64, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(filePath)
	if err != nil {
		return
	}
	size = fi.Size()
	return
}

func GuessImageFormat(filePath string) (format string, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	_, format, err = image.DecodeConfig(file)
	return
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
