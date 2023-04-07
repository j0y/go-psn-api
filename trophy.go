package psn

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const TrophyBaseURL = "https://m.np.playstation.com/api/trophy"

type TrophyTitlesResponse struct {
	TrophyTitles []struct {
		NpServiceName       string `json:"npServiceName"`
		NpCommunicationID   string `json:"npCommunicationId"`
		TrophySetVersion    string `json:"trophySetVersion"`
		TrophyTitleName     string `json:"trophyTitleName"`
		TrophyTitleDetail   string `json:"trophyTitleDetail"`
		TrophyTitleIconURL  string `json:"trophyTitleIconUrl"`
		TrophyTitlePlatform string `json:"trophyTitlePlatform"`
		DefinedTrophies     struct {
			Bronze   int `json:"bronze"`
			Silver   int `json:"silver"`
			Gold     int `json:"gold"`
			Platinum int `json:"platinum"`
		} `json:"definedTrophies"`
		EarnedTrophies struct {
			Bronze   int `json:"bronze"`
			Silver   int `json:"silver"`
			Gold     int `json:"gold"`
			Platinum int `json:"platinum"`
		} `json:"earnedTrophies"`
		HiddenFlag          bool      `json:"hiddenFlag"`
		HasTrophyGroups     bool      `json:"hasTrophyGroups"`
		Progress            int       `json:"progress"`
		LastUpdatedDateTime time.Time `json:"lastUpdatedDateTime"`
	} `json:"trophyTitles"`
	TotalItemCount int `json:"totalItemCount"`
}

func (p *PSN) GetUserTrophyTitles(accountID string) (TrophyTitlesResponse, error) {
	urlPath := fmt.Sprintf("%s/v1/users/%s/trophyTitles", TrophyBaseURL, accountID)

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return TrophyTitlesResponse{}, errors.WithStack(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.accessToken))

	res, err := p.http.Do(req)
	if err != nil {
		return TrophyTitlesResponse{}, errors.WithStack(err)
	}
	defer res.Body.Close()

	var response TrophyTitlesResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, errors.WithStack(err)
	}
	return response, nil
}
