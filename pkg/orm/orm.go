package orm

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	defaultProtocol = "http"
	defaultPort     = "5431"
	defaultHost     = "127.0.0.1"
)

// IClient ...
type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// OClient - used for ...
type OClient struct {
	baseURL string
	client  IClient
}

// NewORMClient ...
func NewORMClient() *OClient {
	api := &OClient{}
	envIP := os.Getenv("API_HOST")
	if envIP != "" {
		defaultHost = envIP
	}
	envPort := os.Getenv("API_PORT")
	if envPort != "" {
		defaultPort = envPort
	}

	api.baseURL = fmt.Sprintf("%s://%s:%s", defaultProtocol, defaultHost, defaultPort)

	api.client = &http.Client{}

	return api
}

// Get is a function that will generate a GET reqest
func (m *OClient) Get(
	uri string,
	headers http.Header,
) (string, *http.Response, error) {
	var payload io.Reader
	req, err := m.generateRequest(uri, http.MethodGet, payload, headers)
	if err != nil {
		return "", nil, err
	}
	return m.runRequest(req)
}

func (m *OClient) generateRequest(uri string, method string, body io.Reader,
	headers http.Header) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", m.baseURL, uri)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// For PUT method add Content-Type if doesn't exist
	// for media type test use "missing" value to don't send
	// Content-Type and test it's missing handling
	if method == "PUT" {
		values := headers.Values("Content-Type")
		if len(values) == 0 {
			headers.Add("Content-Type", "application/json")
		} else {
			for _, item := range values {
				if item == "missing" {
					headers.Del("Content-Type")
				}
			}
		}
	}
	req.Header = headers
	return req, nil
}

func (m *OClient) runRequest(req *http.Request) (string, *http.Response, error) {
	res, err := m.client.Do(req)
	if err != nil {
		return "", nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	strBody := string(body)
	return strBody, res, err
}

// Post is a function that will generate a Post reqest
func (m *OClient) Post(
	uri string,
	data string,
	headers http.Header,
) (string, *http.Response, error) {
	payload := strings.NewReader(data)
	req, err := m.generateRequest(uri, http.MethodPost, payload, headers)
	if err != nil {
		return "", nil, err
	}
	return m.runRequest(req)
}

// Put is a function that will generate a PUT reqest
func (m *OClient) Put(
	uri string,
	data string,
	headers http.Header,
) (string, *http.Response, error) {
	payload := strings.NewReader(data)
	req, err := m.generateRequest(uri, http.MethodPut, payload, headers)
	if err != nil {
		return "", nil, err
	}
	return m.runRequest(req)
}

// Delete is a function that will generate a DELETE reqest
func (m *OClient) Delete(uri string, headers http.Header) (string, *http.Response, error) {
	req, err := m.generateRequest(uri, http.MethodDelete, nil, headers)
	if err != nil {
		return "", nil, err
	}
	return m.runRequest(req)
}

// CreateReservation etc...
func CreateReservation(v interface{}) error {
	// do something
	if v != nil {
		return nil
	}
	return errors.New("a")
}
