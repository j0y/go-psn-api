package psn

import (
	"fmt"
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

	var response TrophyTitlesResponse
	err := p.request(http.MethodGet, urlPath, nil, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
