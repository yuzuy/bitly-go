package bitly

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const baseURLV4 = "https://api-ssl.bitly.com/v4"

type Client struct {
	Token string
}

func New(token string) Client {
	return Client{
		Token: token,
	}
}

func (c Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	return req,nil
}

type ShortenConfig struct {
	Domain    string
	GroupGUID string
}

type shortenRequest struct {
	LongURL   string `json:"long_url"`
	Domain    string `json:"domain"`
	GroupGUID string `json:"group_guid"`
}

type ShortenResponse BitlinkBody

func (c Client) Shorten(url string, config ...ShortenConfig) (*ShortenResponse, error) {
	var domain, groupGUID string
	if config != nil {
		domain = config[0].Domain
		groupGUID = config[0].GroupGUID
	}
	reqBody := shortenRequest{
		LongURL:   url,
		Domain:    domain,
		GroupGUID: groupGUID,
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, err := c.newRequest(http.MethodPost, baseURLV4+"/shorten", bytes.NewReader(reqBodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, unmarshalError(decoder)
	}
	shortenResp := new(ShortenResponse)
	_ = decoder.Decode(shortenResp)

	return shortenResp, nil
}
