package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"

	"io/ioutil"

	"github.com/go-errors/errors"
	"github.com/moul/http2curl"
	"github.com/spf13/viper"
)

type SuperAgent struct {
	Client *http.Client
	URL    string
	Dry    bool
}

func NewSuperAgent(rawurl string) *SuperAgent {
	return &SuperAgent{
		URL:    rawurl,
		Client: http.DefaultClient,
		Dry:    viper.GetBool("dry"),
	}
}

func (s *SuperAgent) doDry(req *http.Request) error {
	if s.Dry {
		command, err := http2curl.GetCurlCommand(req)
		if err != nil {
			return errors.New(err)
		}

		return errors.Errorf("Because you are using the dry option, the request will not be executed. You can execute this command using: \n\n%s", command)
	}
	return nil
}

func (s *SuperAgent) Delete() error {
	req, err := http.NewRequest("DELETE", s.URL, nil)
	if err != nil {
		return errors.New(err)
	}

	if err := s.doDry(req); err != nil {
		return err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return errors.New(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.Errorf("Expected status code %d, got %d.\n%s\n", http.StatusNoContent, resp.StatusCode, body)
	}

	return nil
}

func (s *SuperAgent) Get(o interface{}) error {
	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return errors.New(err)
	} else if o == nil {
		return errors.New("Can not pass nil")
	}

	if err := s.doDry(req); err != nil {
		return err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return errors.New(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.Errorf("Expected status code %d, got %d.\n%s\n", http.StatusOK, resp.StatusCode, body)
	} else if err := json.NewDecoder(resp.Body).Decode(o); err != nil {
		return errors.New(err)
	}

	return nil
}

func (s *SuperAgent) Create(o interface{}) error {
	return s.send("POST", o, o)
}

func (s *SuperAgent) POST(in, out interface{}) error {
	return s.send("POST", in, out)
}

func (s *SuperAgent) Update(o interface{}) error {
	return s.send("PUT", o, o)
}

func (s *SuperAgent) send(method string, in interface{}, out interface{}) error {
	if s.Client == nil {
		s.Client = http.DefaultClient
	}

	data, err := json.Marshal(in)
	if err != nil {
		return errors.New(err)
	}

	req, err := http.NewRequest(method, s.URL, bytes.NewReader(data))
	if err != nil {
		return errors.New(err)
	}

	req.Header.Set("Content-Type", "application/json")
	if err := s.doDry(req); err != nil {
		return err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return errors.New(err)
	}

	expectedStatus := http.StatusOK
	if method == "POST" {
		expectedStatus = http.StatusCreated
	}
	if resp.StatusCode != expectedStatus {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.Errorf("Expected status code %d, got %d.\n%s", expectedStatus, resp.StatusCode, body)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.Errorf("%s: %s", err, body)
	}

	return nil
}
