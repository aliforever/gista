package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"github.com/kr/pretty"
)

type ChallengeResponse struct {
	Config      *ChallengeResponseConfig    `json:"config"`
	EntryData   *ChallengeResponseEntryData `json:"entry_data"`
	RollOutHash string                      `json:"rollout_hash"`
}

type ChallengeResponseConfig struct {
	CsrfToken string `json:"csrf_token"`
}

type ChallengeResponseEntryData struct {
	Challenge *[]ChallengeResponseEntryDataChallenge `json:"Challenge"`
}

type ChallengeResponseEntryDataChallenge struct {
	ChallengeType string                                         `json:"challengeType"`
	Fields        *ChallengeResponseEntryDataChallengeFields     `json:"fields"`
	Navigation    *ChallengeResponseEntryDataChallengeNavigation `json:"navigation"`
	Type          string                                         `json:"type"`
}

type ChallengeResponseEntryDataChallengeFields struct {
	Choice string `json:"choice"`
	Email  string `json:"email"`
}

type ChallengeResponseEntryDataChallengeNavigation struct {
	Forward string `json:"forward"`
	Replay  string `json:"replay"`
	Dismiss string `json:"dismiss"`
}

type PostChallengeChoiceResponse struct {
	Location      string                                `json:"location"`
	ChallengeType string                                `json:"challengeType"`
	ExtraData     *PostChallengeChoiceResponseExtraData `json:"extraData"`
	Type          string                                `json:"type"`
	Status        string                                `json:"status"`
}

type PostChallengeChoiceResponseExtraData struct {
	TypeName string                                         `json:"__typename"`
	Content  *[]PostChallengeChoiceResponseExtraDataContent `json:"content"`
}

type PostChallengeChoiceResponseExtraDataContent struct {
	TypeName string `json:"__typename"`
	Title    string `json:"title"`
	Text     string `json:"text"`
}

type SolveChallengeResponse struct {
	Message   string                           `json:"message"`
	Challenge *SolveChallengeResponseChallenge `json:"challenge"`
	Location  string                           `json:"location"`
	Type      string                           `json:"type"`
	Status    string                           `json:"status"`
}

type SolveChallengeResponseChallenge struct {
	ChallengeType string   `json:"challengeType"`
	Errors        []string `json:"errors"`
}

type ChallengeSolver struct {
	client *http.Client
	appId  string
}

func NewChallengeSolver() *ChallengeSolver {
	j, _ := cookiejar.New(nil)
	return &ChallengeSolver{client: &http.Client{Jar: j}, appId: "936619743392459"}
}

func (ncs *ChallengeSolver) GetChallengeByUrl(address string) (response *ChallengeResponse, err error) {
	var res *http.Response
	req, _ := http.NewRequest("GET", address, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763")
	res, err = ncs.client.Do(req)
	if err != nil {
		return
	}
	pretty.Println(res.Header)
	pretty.Println(res.Cookies())
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	r := regexp.MustCompile(`<script type="text/javascript">window\._sharedData = ({[\s\S]+});</script>`)
	m := r.FindAllStringSubmatch(string(b), 1)

	response = &ChallengeResponse{}
	err = json.Unmarshal([]byte(m[0][1]), response)
	return
}

func (ncs *ChallengeSolver) GetSolveChallengeByEmail(address, csrf, ajaxId string) (response *PostChallengeChoiceResponse, err error) {
	values := url.Values{}
	values.Add("choice", "1")
	p, _ := http.NewRequest("POST", address, strings.NewReader(values.Encode()))
	p.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	p.Header.Set("X-CSRFToken", csrf)
	p.Header.Set("X-Instagram-AJAX", ajaxId)
	p.Header.Set("X-IG-App-ID", ncs.appId)
	p.Header.Set("X-Requested-With", "XMLHttpRequest")
	p.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var res *http.Response
	res, err = ncs.client.Do(p)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	response = &PostChallengeChoiceResponse{}
	err = json.Unmarshal(b, response)
	return
}

func (ncs *ChallengeSolver) SolveChallenge(address, csrf, ajaxId, securityCode string) (response *SolveChallengeResponse, err error) {
	values := url.Values{}
	values.Add("security_code", securityCode)
	p, _ := http.NewRequest("POST", address, strings.NewReader(values.Encode()))
	p.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	p.Header.Set("X-CSRFToken", csrf)
	p.Header.Set("X-Instagram-AJAX", ajaxId)
	p.Header.Set("X-IG-App-ID", ncs.appId)
	p.Header.Set("X-Requested-With", "XMLHttpRequest")
	p.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var res *http.Response
	res, err = ncs.client.Do(p)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	response = &SolveChallengeResponse{}
	err = json.Unmarshal(b, response)
	return
}
