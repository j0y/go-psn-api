package psn

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestPSN(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fail()
	}
	psnAPI, err := NewPsnAPI(os.Getenv("NPSSO"))
	if err != nil {
		t.Fail()
	}

	search, err := psnAPI.MakeUniversalSearch("ikemenzi", []domain{SOCIAL_ALL_ACCOUNTS})
	if err != nil {
		t.Fail()
	}
	if len(search.DomainResponses) == 0 {
		t.Fail()
	}
	if len(search.DomainResponses[0].Results) == 0 {
		t.Fail()
	}

	userID := search.DomainResponses[0].Results[0].SocialMetadata.AccountID
	_, err = psnAPI.GetUserTitles(userID, 0, 7)
	if err != nil {
		t.Fail()
	}
}
