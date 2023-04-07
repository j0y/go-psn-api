package psn

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
)

type AuthTokensResponse struct {
	/** Used to retrieve data from the PSN API. */
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	/** Used to retrieve a new access token when it expires. */
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	/** When the access token will expire. */
	ExpiresIn int32 `json:"expires_in"`
	/** When the refresh token will expire. */
	RefreshTokenExpiresIn int32 `json:"refresh_token_expires_in"`
}

func (p *PSN) exchangeNpssoForCode(npsso string) (string, error) {
	values := url.Values{}
	values.Add("access_type", "offline")
	values.Add("client_id", "09515159-7237-4370-9b40-3806e67c0891")
	values.Add("redirect_uri", "com.scee.psxandroid.scecompcall://redirect")
	values.Add("response_type", "code")
	values.Add("scope", "psn:mobile.v2.core psn:clientapp")
	query := values.Encode()
	urlPath := fmt.Sprintf("%s/authorize?%s", AuthBaseURL, query)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { // don't follow redirects
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)

	if err != nil {
		return "", errors.WithStack(err)
	}
	req.Header.Add("Cookie", fmt.Sprintf("npsso=%s", npsso))

	res, err := client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer res.Body.Close()

	location := res.Header.Get("location")
	if !strings.Contains(location, "?code=") {
		return "", errors.WithStack(ErrNoCode)
	}

	params := strings.Split(location, "redirect/?")
	if len(params) < 2 {
		return "", errors.WithStack(ErrNoCode)
	}

	parseQuery, err := url.ParseQuery(params[1])
	if err != nil {
		return "", err
	}

	if !parseQuery.Has("code") {
		return "", errors.WithStack(ErrNoCode)
	}

	return parseQuery.Get("code"), nil
}

func (p *PSN) exchangeCodeForAccessToken(accessCode string) (AuthTokensResponse, error) {
	urlPath := fmt.Sprintf("%s/token", AuthBaseURL)

	values := url.Values{}
	values.Add("code", accessCode)
	values.Add("redirect_uri", "com.scee.psxandroid.scecompcall://redirect")
	values.Add("grant_type", "authorization_code")
	values.Add("token_format", "jwt")

	req, err := http.NewRequest(http.MethodPost, urlPath, strings.NewReader(values.Encode()))
	if err != nil {
		return AuthTokensResponse{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic MDk1MTUxNTktNzIzNy00MzcwLTliNDAtMzgwNmU2N2MwODkxOnVjUGprYTV0bnRCMktxc1A=")

	res, err := p.http.Do(req)
	if err != nil {
		return AuthTokensResponse{}, errors.WithStack(err)
	}
	defer res.Body.Close()

	var response AuthTokensResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, errors.WithStack(err)
	}
	return response, nil
}

func (p *PSN) AuthWithNPSSO(npsso string) error {
	code, err := p.exchangeNpssoForCode(npsso)
	if err != nil {
		return err
	}

	tokens, err := p.exchangeCodeForAccessToken(code)
	if err != nil {
		return err
	}

	p.accessToken = tokens.AccessToken
	p.expiresIn = tokens.ExpiresIn
	p.refreshToken = tokens.RefreshToken
	p.refreshTokenExpiresIn = tokens.RefreshTokenExpiresIn

	return nil
}
