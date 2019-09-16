package signatures

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/constants"
)

func GenerateSignature(data string) string {
	h := hmac.New(sha256.New, []byte(constants.IgSigKey))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func GenerateSignatureForPost(data string) string {
	return "ig_sig_key_version=" + constants.SigKeyVersion + "&signed_body=" + GenerateSignature(data) + url.QueryEscape(data)
}

func SignData(data map[string] /*string */ interface{}, exclude *[]string) map[string]interface{} {
	result := map[string]interface{}{}
	if exclude != nil {
		for _, item := range *(exclude) {
			if _, ok := data[item]; ok {
				result[item] = data[item]
			}
		}
		/*castedData := map[string]string{}
		for k, v := range data {
			switch v.(type) {
			case string:
				castedData[k] = v.(string)
			case float32, float64:
				castedData[k] = fmt.Sprintf("%f", v.(float32))
			case int, int8, int16, int32, int64:
				castedData[k] = fmt.Sprintf("%d", v.(int64))
			case bool:
				castedData[k] = strconv.FormatBool(v.(bool))
			}
		}
		*/
	}
	// ReorderByHashCodeIsNeededHere
	dataBytes, _ := json.Marshal(data)
	result["ig_sig_key_version"] = constants.SigKeyVersion
	result["signed_body"] = GenerateSignature(string(dataBytes)) + "." + string(dataBytes)
	// ReorderByHashCodeIsNeededHere
	return result
}

func GenerateDeviceId() string {
	nFormat := utils.NumberFormat(utils.MicroTime(), 7, "", "")
	megaRandomHash := fmt.Sprintf("%x", md5.Sum([]byte(nFormat)))
	return "android-" + megaRandomHash[16:]
}

func GenerateUUID(keepDashes bool) string {
	uuId := fmt.Sprintf("%04x%04x-%04x-%04x-%04x-%04x%04x%04x",
		utils.MtRand(0, 0xffff),
		utils.MtRand(0, 0xffff),
		utils.MtRand(0, 0xffff),
		utils.MtRand(0, 0x0fff)|0x4000,
		utils.MtRand(0, 0x3fff)|0x8000,
		utils.MtRand(0, 0xffff),
		utils.MtRand(0, 0xffff),
		utils.MtRand(0, 0xffff))
	if !keepDashes {
		uuId = strings.Replace(uuId, "-", "", -1)
	}
	return uuId
}

func HashCode(str string) (code int) {
	code = 0
	runes := []rune(str)
	for _, k := range runes {
		code = (-code + (code << 5) + int(k)) & 0xFFFFFFFF
	}
	if strconv.IntSize > 4 {
		if code > 0x7FFFFFFF {
			code -= 0x100000000
		} else if code < -0x80000000 {
			code += 0x100000000
		}
	}
	return
}

// Function not needed in Golang
func ReorderByHashCode(data map[string]string) (result map[string]string) {
	/*hashCodes := map[string]int{}
	newData := map[string]string{}
	hashCodesInts := []int{}
	for k := range data {
		h := HashCode(k)
		hashCodes[k] = h
		hashCodesInts = append(hashCodesInts, h)
	}
	sort.Ints(hashCodesInts)
	for _, k := range hashCodesInts {
		newData[]
	}
	pretty.Println(hashCodes)*/
	//sort.Sort()
	return
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile(`#^[a-f\d]{8}-(?:[a-f\d]{4}-){3}[a-f\d]{12}$#D`) // TODO: Fix Regexp
	return r.MatchString(uuid)
}
