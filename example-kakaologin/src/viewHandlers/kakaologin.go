package viewHandlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"example-kakaologin/src/env"
	plog "example-kakaologin/src/logprint"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type requestTokenForm struct {
	grant_type   string
	client_id    string
	redirect_uri string
	code         string
}

type responseToken struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             uint32 `json:"expires_in"`
	Scope                 string `json:"scope"`
	RefreshTokenExpiresIn uint32 `json:"refresh_token_expires_in"`
}

func handlePanic(res http.ResponseWriter) {
	if exception := recover(); exception != nil {
		plog.PrintError("Panic!!")
		http.Error(res, "Exception", http.StatusInternalServerError)
	}
}

func getAuthCodeParams(req *http.Request) (string, error) {
	plog.PrintInfo("getAuthCodeParams Enter...\n")
	fmt.Println(req.URL)
	code, err := req.URL.Query()["code"]
	if !err || len(code[0]) < 1 {
		plog.PrintError("Cannot find code param")
		return "", errors.New("Cannot find code param")
	}
	// plog.PrintDebug("auth code : ", code[0])
	plog.PrintInfo("getAuthCodeParams Leave...\n")
	return code[0], nil
}
func requestToken(code string) (*responseToken, error) {
	plog.PrintInfo("requestToken Enter...\n")
	bodyBuffer := url.Values{
		"grant_type":   {"authorization_code"},
		"client_id":    {env.GetStringEnv("KAKAO_CLIENT_ID")},
		"redirect_uri": {env.GetStringEnv("KAKAO_REDIRECT_URL")},
		"code":         {code}}
	requestData, err := http.NewRequest("POST", "https://kauth.kakao.com/oauth/token", bytes.NewBufferString(bodyBuffer.Encode()))
	if err != nil {
		plog.PrintError("Failed to create request form\n")
		return nil, errors.New("Failed to create request form")
	}
	requestData.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	result, err := httpClient.Do(requestData)
	if err != nil {
		plog.PrintError("Failed to do http request\n")
		return nil, errors.New("Failed to do http request")
	}
	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	// plog.PrintInfo("response string : ", string(body))
	var tokenStructure responseToken
	err = json.Unmarshal(body, &tokenStructure)
	if err != nil {
		plog.PrintError("Failed to decode json body\n")
		return nil, errors.New("Failed to decode json body")
	}
	plog.PrintInfo("response body : %+v\n", tokenStructure)

	plog.PrintInfo("requestToken Leave...\n")
	return &tokenStructure, nil
}

func KakaoLogin(res http.ResponseWriter, req *http.Request, pageName string) {
	page, err := loadPage(pageName)
	if err != nil {
		fmt.Println("Not Found : ", pageName, ", redirect to login...")
		http.Redirect(res, req, "/login/kakaologin", http.StatusFound)
		return
	}
	page.ClientId = env.GetStringEnv("KAKAO_CLIENT_ID")
	page.RedirectURL = env.GetStringEnv("KAKAO_REDIRECT_URL")

	renderTemplate(res, "kakaologin", page)
	defer handlePanic(res)
}

func KakaoAuth(res http.ResponseWriter, req *http.Request, pageName string) {
	plog.PrintInfo("KakaoAuth Enter...\n")
	authCode, err := getAuthCodeParams(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	tokenStructure, err := requestToken(authCode)
	if err != nil {
		return
	}
	pageData := &Page{Name: "Kakao Login Result", ClientId: env.GetStringEnv("KAKAO_CLIENT_ID"), Token: tokenStructure.AccessToken}
	renderTemplate(res, "kakaoauth", pageData)
	plog.PrintInfo("KakaoAuth Leave...\n")
	defer handlePanic(res)
}
