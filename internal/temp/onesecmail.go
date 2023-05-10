package temp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Mail struct {
	ID      int
	Date    string
	From    string
	Subject string
}

type OneSecMail struct {
	domain string
	mail   string
	user   string
	mails  map[int]struct{}
	lock   *sync.Mutex
}

type Single struct {
	Body string
}

func (osm *OneSecMail) Fetch() (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://www.1secmail.com/api/v1/?action=genRandomMailbox")
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	var mail []string
	if err := json.NewDecoder(resp.Body).Decode(&mail); err != nil {
		return "", err
	}

	osm.mail = mail[0]

	splitted := strings.Split(osm.mail, "@")
	if len(splitted) == 2 {
		osm.user = splitted[0]
		osm.domain = splitted[1]
	} else {
		return "", errors.New("error with email format")
	}

	return osm.mail, nil
}

func (osm *OneSecMail) Check() ([]Content, error) {
	if osm.mails == nil {
		osm.mails = make(map[int]struct{})
	}
	if osm.lock == nil {
		osm.lock = &sync.Mutex{}
	}
	osm.lock.Lock()
	defer osm.lock.Unlock()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	ur := fmt.Sprintf("https://www.1secmail.com/api/v1/?action=getMessages&login=%s&domain=%s",
		osm.user,
		osm.domain)
	resp, err := client.Get(ur)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mails []Mail
	if err := json.NewDecoder(resp.Body).Decode(&mails); err != nil {
		return nil, err
	}

	var newMails []Content

	for _, mail := range mails {
		if _, ok := osm.mails[mail.ID]; !ok {
			newMails = append(newMails, Content{
				ID:      mail.ID,
				Subject: mail.Subject,
				From:    mail.From,
				Body:    "",
			})
			osm.mails[mail.ID] = struct{}{}
		}
	}

	return newMails, nil
}

func (osm *OneSecMail) Read(val, mail any) error {
	id := val.(int)
	m := mail.(*Content)
	ur := fmt.Sprintf("https://www.1secmail.com/api/v1/?action=readMessage&login=%s&domain=%s&id=%d",
		osm.user, osm.domain, id)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(ur)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var single Single
	if err := json.NewDecoder(resp.Body).Decode(&single); err != nil {
		return err
	}

	m.Body = single.Body
	return nil
}

func (osm *OneSecMail) FetchURL() string {
	return ""
}
