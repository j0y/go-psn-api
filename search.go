package psn

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const SearchBaseURL = "https://m.np.playstation.com/api/search"

type domain string

const BLENDED domain = "Blended"
const CONCEPT_GAME domain = "ConceptGame"
const CONCEPT_GAME_ALL domain = "ConceptGameAll"
const CONCEPT_GAME_MOBILE_APP domain = "ConceptGameMobileApp"
const CONCEPT_GAME_PS_NOW domain = "ConceptGamePsNow"
const CONCEPT_GAME_EXCLUDING_PS_NOW domain = "ConceptGameExcludingPsNow"
const CONCEPT_GAME_ADDONS domain = "ConceptGameAddons"
const CONCEPT_APPLICATION_GAME domain = "ConceptApplicationGame"
const CONCEPT_APPLICATION_VIDEO domain = "ConceptApplicationVideo"
const CONCEPT_OTHER_NON_GAME domain = "ConceptOtherNonGame"
const VIDEO_CATALOG domain = "VideoCatalog"
const VIDEO_CATALOG_MOVIE domain = "VideoCatalogMovie"
const VIDEO_CATALOG_TV_SERIES domain = "VideoCatalogTVSeries"
const SOCIAL_ALL_ACCOUNTS domain = "SocialAllAccounts"
const SOCIAL_FRIENDS domain = "SocialFriends"
const SOCIAL_FRIENDS_OF_FRIENDS domain = "SocialFriendsOfFriends"
const SOCIAL_FRIENDS_AND_FRIENDS_OF_FRIENDS domain = "SocialFriendsAndFriendsOfFriends"
const SOCIAL_UNRELATED_AND_FRIENDS_OF_FRIENDS domain = "SocialUnrelatedAndFriendsOfFriends"
const SOCIAL_UNRELATED domain = "SocialUnrelated"

type domainStruct struct {
	Domain domain `json:"domain"`
}

type SearchRequest struct {
	SearchTerm     string         `json:"searchTerm"`
	DomainRequests []domainStruct `json:"domainRequests"`
}

type UniversalSearchResponse struct {
	DomainResponses []struct {
		Domain               string   `json:"domain"`
		DomainTitle          string   `json:"domainTitle"`
		DomainTitleMessageID string   `json:"domainTitleMessageId"`
		DomainTitleHighlight []string `json:"domainTitleHighlight"`
		ZeroState            bool     `json:"zeroState"`
		Next                 string   `json:"next"`
		TotalResultCount     int      `json:"totalResultCount"`
		Results              []struct {
			ID             string  `json:"id"`
			Type           string  `json:"type"`
			Score          float64 `json:"score"`
			SocialMetadata struct {
				AccountID            string `json:"accountId"`
				Country              string `json:"country"`
				Language             string `json:"language"`
				OnlineID             string `json:"onlineId"`
				IsPsPlus             bool   `json:"isPsPlus"`
				IsOfficiallyVerified bool   `json:"isOfficiallyVerified"`
				AvatarURL            string `json:"avatarUrl"`
				VerifiedUserName     string `json:"verifiedUserName"`
				Highlights           struct {
					OnlineID  []string `json:"onlineId"`
					FirstName []string `json:"firstName,omitempty"`
					LastName  []string `json:"lastName,omitempty"`
				} `json:"highlights"`
				FirstName     string `json:"firstName,omitempty"`
				LastName      string `json:"lastName,omitempty"`
				ProfilePicURL string `json:"profilePicUrl,omitempty"`
				MiddleName    string `json:"middleName,omitempty"`
			} `json:"socialMetadata"`
		} `json:"results"`
	} `json:"domainResponses"`
	ResponseStatus []struct {
		Status        string `json:"status"`
		StatusMessage string `json:"statusMessage"`
	} `json:"responseStatus"`
	FallbackQueried bool `json:"fallbackQueried"`
}

func (p *PSN) MakeUniversalSearch(searchTerm string, domains []domain) (UniversalSearchResponse, error) {
	urlPath := fmt.Sprintf("%s/v1/universalSearch", SearchBaseURL)

	search := SearchRequest{SearchTerm: searchTerm}
	for _, searchDomain := range domains {
		search.DomainRequests = append(search.DomainRequests, domainStruct{Domain: searchDomain})
	}
	data, err := json.Marshal(search)
	if err != nil {
		return UniversalSearchResponse{}, errors.WithStack(err)
	}

	var response UniversalSearchResponse
	err = p.request(http.MethodPost, urlPath, strings.NewReader(string(data)), &response)
	if err != nil {
		return UniversalSearchResponse{}, err
	}

	return response, nil
}
