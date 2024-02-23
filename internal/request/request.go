package request

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) Get(path string) (string, error) {
	client := restyClient(c.token)
	resp, err := client.R().Get(path)
	if err != nil {
		return "", err
	}
	logrus.Debugf("GET %s", path)
	logrus.Debugf("Response: %s", resp.String())
	return resp.String(), nil
}

func (c *Client) Post(path string, body interface{}) (string, error) {
	client := restyClient(c.token)
	resp, err := client.R().SetBody(body).Post(path)
	if err != nil {
		return "", err
	}
	logrus.Debugf("POST %s", path)
	logrus.Debugf("Body: %s", body)
	logrus.Debugf("Response: %s", resp.String())
	return resp.String(), nil
}

func restyClient(token string) *resty.Client {
	client := resty.New()
	client.SetBaseURL("https://www.kookapp.cn/api/v3")
	client.SetAuthScheme("Bot")
	client.SetAuthToken(token)
	return client
}
