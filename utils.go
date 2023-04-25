package psn

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (p *PSN) request(method string, url string, body io.Reader, value interface{}) error {
	req, err := http.NewRequest(
		method,
		url,
		body,
	)
	if err != nil {
		return errors.WithStack(fmt.Errorf("can't create new %s request %w: ", method, err))
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.accessToken))

	resp, err := p.http.Do(req)
	if err != nil {
		return errors.WithStack(fmt.Errorf("can't execute %s request %w: ", method, err))
	}

	defer func() {
		err = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return errors.WithStack(fmt.Errorf("bad request. Status: %d - %s", resp.StatusCode, resp.Status))
	}

	err = json.NewDecoder(resp.Body).Decode(&value)
	if err != nil {
		return errors.WithStack(fmt.Errorf("can't decode %s request %w: ", method, err))
	}

	return nil
}
