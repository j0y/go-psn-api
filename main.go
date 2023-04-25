package psn

import (
	"errors"
	"net/http"
)

const AuthBaseURL = "https://ca.account.sony.com/api/authz/v3/oauth"

var ErrNoCode = errors.New("there was a problem retrieving your PSN access code. Is your NPSSO code valid? To get a new NPSSO code, visit https://ca.account.sony.com/api/v1/ssocookie")

type PSN struct {
	http                  *http.Client
	accessToken           string
	refreshToken          string
	expiresIn             int32
	refreshTokenExpiresIn int32
}

// NewPsnAPI requires npsso for auth. It's available for any logged-in account at https://ca.account.sony.com/api/v1/ssocookie
func NewPsnAPI(npsso string) (*PSN, error) {
	psnAPI := &PSN{
		http: &http.Client{},
	}

	err := psnAPI.AuthWithNPSSO(npsso)
	if err != nil {
		return nil, err
	}
	return psnAPI, nil
}
