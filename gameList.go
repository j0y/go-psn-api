package psn

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const GameListBaseURL = "https://m.np.playstation.com/api/gamelist"

type UserTitlesResponse struct {
	Titles []struct {
		TitleID           string `json:"titleId"`
		Name              string `json:"name"`
		LocalizedName     string `json:"localizedName"`
		ImageURL          string `json:"imageUrl"`
		LocalizedImageURL string `json:"localizedImageUrl"`
		Category          string `json:"category"`
		Service           string `json:"service"`
		PlayCount         int    `json:"playCount"`
		Concept           struct {
			ID       int      `json:"id"`
			TitleIds []string `json:"titleIds"`
			Name     string   `json:"name"`
			Media    struct {
				Audios []interface{} `json:"audios"`
				Videos []interface{} `json:"videos"`
				Images []struct {
					URL    string `json:"url"`
					Format string `json:"format"`
					Type   string `json:"type"`
				} `json:"images"`
			} `json:"media"`
			Genres        []string `json:"genres"`
			LocalizedName struct {
				DefaultLanguage string `json:"defaultLanguage"`
				Metadata        struct {
					ArAE   string `json:"ar-AE"`
					DaDK   string `json:"da-DK"`
					DeDE   string `json:"de-DE"`
					EnGB   string `json:"en-GB"`
					EnUS   string `json:"en-US"`
					Es419  string `json:"es-419"`
					EsES   string `json:"es-ES"`
					FiFI   string `json:"fi-FI"`
					FrCA   string `json:"fr-CA"`
					FrFR   string `json:"fr-FR"`
					ItIT   string `json:"it-IT"`
					JaJP   string `json:"ja-JP"`
					KoKR   string `json:"ko-KR"`
					NlNL   string `json:"nl-NL"`
					NoNO   string `json:"no-NO"`
					PlPL   string `json:"pl-PL"`
					PtBR   string `json:"pt-BR"`
					PtPT   string `json:"pt-PT"`
					RuRU   string `json:"ru-RU"`
					SvSE   string `json:"sv-SE"`
					TrTR   string `json:"tr-TR"`
					UkUA   string `json:"uk-UA"`
					ZhHans string `json:"zh-Hans,omitempty"`
					ZhHant string `json:"zh-Hant,omitempty"`
				} `json:"metadata"`
			} `json:"localizedName"`
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"concept"`
		Media struct {
			Audios []interface{} `json:"audios"`
			Videos []interface{} `json:"videos"`
			Images []struct {
				URL    string `json:"url"`
				Format string `json:"format"`
				Type   string `json:"type"`
			} `json:"images"`
		} `json:"media"`
		FirstPlayedDateTime time.Time `json:"firstPlayedDateTime"`
		LastPlayedDateTime  time.Time `json:"lastPlayedDateTime"`
		PlayDuration        string    `json:"playDuration"`
	} `json:"titles"`
	NextOffset     int `json:"nextOffset"`
	PreviousOffset int `json:"previousOffset"`
	TotalItemCount int `json:"totalItemCount"`
}

type UserTitleResponse struct {
	TitleID           string `json:"titleId"`
	Name              string `json:"name"`
	LocalizedName     string `json:"localizedName"`
	ImageURL          string `json:"imageUrl"`
	LocalizedImageURL string `json:"localizedImageUrl"`
	Category          string `json:"category"`
	Service           string `json:"service"`
	PlayCount         int    `json:"playCount"`
	Concept           struct {
		ID       int      `json:"id"`
		TitleIds []string `json:"titleIds"`
		Type     string   `json:"type"`
		Name     string   `json:"name"`
		Media    struct {
			Audios []interface{} `json:"audios"`
			Videos []interface{} `json:"videos"`
			Images []struct {
				URL    string `json:"url"`
				Format string `json:"format"`
				Type   string `json:"type"`
			} `json:"images"`
		} `json:"media"`
		Genres        []string `json:"genres"`
		LocalizedName struct {
			DefaultLanguage string `json:"defaultLanguage"`
			Metadata        struct {
				ArAE   string `json:"ar-AE"`
				DaDK   string `json:"da-DK"`
				DeDE   string `json:"de-DE"`
				EnGB   string `json:"en-GB"`
				EnUS   string `json:"en-US"`
				Es419  string `json:"es-419"`
				EsES   string `json:"es-ES"`
				FiFI   string `json:"fi-FI"`
				FrCA   string `json:"fr-CA"`
				FrFR   string `json:"fr-FR"`
				ItIT   string `json:"it-IT"`
				JaJP   string `json:"ja-JP"`
				KoKR   string `json:"ko-KR"`
				NlNL   string `json:"nl-NL"`
				NoNO   string `json:"no-NO"`
				PlPL   string `json:"pl-PL"`
				PtBR   string `json:"pt-BR"`
				PtPT   string `json:"pt-PT"`
				RuRU   string `json:"ru-RU"`
				SvSE   string `json:"sv-SE"`
				TrTR   string `json:"tr-TR"`
				UkUA   string `json:"uk-UA"`
				ZhHans string `json:"zh-Hans"`
				ZhHant string `json:"zh-Hant"`
			} `json:"metadata"`
		} `json:"localizedName"`
		LocalizedMedia struct {
			DefaultLanguage string `json:"defaultLanguage"`
			Metadata        struct {
				ArAE struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"ar-AE"`
				DaDK struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"da-DK"`
				DeDE struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"de-DE"`
				EnGB struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"en-GB"`
				EnUS struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"en-US"`
				Es419 struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"es-419"`
				EsES struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"es-ES"`
				FiFI struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"fi-FI"`
				FrCA struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"fr-CA"`
				FrFR struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"fr-FR"`
				ItIT struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"it-IT"`
				JaJP struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"ja-JP"`
				KoKR struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"ko-KR"`
				NlNL struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"nl-NL"`
				NoNO struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"no-NO"`
				PlPL struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"pl-PL"`
				PtBR struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"pt-BR"`
				PtPT struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"pt-PT"`
				RuRU struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"ru-RU"`
				SvSE struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"sv-SE"`
				TrTR struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"tr-TR"`
				UkUA struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"uk-UA"`
				ZhHans struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"zh-Hans"`
				ZhHant struct {
					Audios []interface{} `json:"audios"`
					Videos []interface{} `json:"videos"`
					Images []struct {
						URL    string `json:"url"`
						Format string `json:"format"`
						Type   string `json:"type"`
					} `json:"images"`
				} `json:"zh-Hant"`
			} `json:"metadata"`
		} `json:"localizedMedia"`
		Country  string `json:"country"`
		Language string `json:"language"`
	} `json:"concept"`
	Media struct {
		Audios []interface{} `json:"audios"`
		Videos []interface{} `json:"videos"`
		Images []struct {
			URL    string `json:"url"`
			Format string `json:"format"`
			Type   string `json:"type"`
		} `json:"images"`
	} `json:"media"`
	FirstPlayedDateTime time.Time `json:"firstPlayedDateTime"`
	LastPlayedDateTime  time.Time `json:"lastPlayedDateTime"`
	PlayDuration        string    `json:"playDuration"`
}

func (p *PSN) GetUserTitles(accountID string, offset int32, limit int32) (UserTitlesResponse, error) {
	if limit > 300 {
		return UserTitlesResponse{}, errors.New("limit must be less than or equal to 300")
	}
	urlPath := fmt.Sprintf("%s/v2/users/%s/titles?limit=%d&offset=%d", GameListBaseURL, accountID, limit, offset)

	var response UserTitlesResponse
	err := p.request(http.MethodGet, urlPath, nil, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// GetUserTitle returns user game info. titleID example: "CUSA07994_00"
func (p *PSN) GetUserTitle(accountID string, titleID string) (UserTitleResponse, error) {
	urlPath := fmt.Sprintf("%s/v2/users/%s/titles/%s", GameListBaseURL, accountID, titleID)

	var response UserTitleResponse
	err := p.request(http.MethodGet, urlPath, nil, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
