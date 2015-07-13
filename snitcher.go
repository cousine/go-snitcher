package snitcher

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Snitcher struct {
	Token string
}

func NewSnitch(token string) *Snitcher {
	return &Snitcher{
		Token: token,
	}
}

func (snitcher *Snitcher) Snitch() error {
	uri := fmt.Sprintf("https://nosnch.in/%v", snitcher.Token)
	_, err := http.Get(uri)

	return err
}

func (snitcher *Snitcher) SnitchWithMessage(message string) error {
	uri := fmt.Sprintf("https://nosnch.in/%v", snitcher.Token)
	form := url.Values{}
	form.Add("message", message)

	request, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	_, err = httpClient.Do(request)
	if err != nil {
		return err
	}

	return nil
}
