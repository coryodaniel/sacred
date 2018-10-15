package confluence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type ConfluenceError struct{}

func (e ConfluenceError) Error() string {
	return fmt.Sprintf("Confluence error: %+v", e)
}

type ContentService struct {
	sling *sling.Sling
}

type Client struct {
	ContentService *ContentService
}

func NewClient(domain string, token string, httpClient *http.Client) *Client {
	return &Client{
		ContentService: NewContentService(domain, token, httpClient),
	}
}

func NewContentService(domain string, token string, httpClient *http.Client) *ContentService {
	baseURL := fmt.Sprintf("https://%s/wiki/rest/api/", domain)
	auth := fmt.Sprintf("Basic %s", token)
	svc := sling.New().Client(httpClient).Base(baseURL).
		Set("Authorization", auth).
		Set("Content-Type", "application/json")
	return &ContentService{sling: svc}
}

func (c *ContentService) Get(id string) (*Content, *http.Response, error) {
	content := new(Content)
	confluenceError := new(ConfluenceError)
	path := fmt.Sprintf("content/%s", id)

	resp, err := c.sling.New().Path(path).
		Receive(content, confluenceError)

	if err == nil {
		err = confluenceError
	}
	return content, resp, err
}

func (s *ContentService) Update(id string, contentBody ContentRequest) (*Content, *http.Response, error) {
	content := new(Content)
	confluenceError := new(ConfluenceError)
	path := fmt.Sprintf("content/%s", id)

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(contentBody)

	resp, err := s.sling.New().Put(path).
		Body(buf).
		Set("Content-Type", "application/json").
		Receive(content, confluenceError)

	if err == nil {
		err = confluenceError
	}
	return content, resp, err
}

// ContentRequestPayload creates a content request HTTP body
func ContentRequestPayload(spaceId string, version int, title string, html string) ContentRequest {
	return ContentRequest{
		Status: "current",
		Version: &Version{
			Number: version + 1,
		},
		Type: "page",
		Space: &Space{
			Key: spaceId,
		},
		Title: title,
		Body: &Body{
			StorageView: &View{
				Value:          html,
				Representation: "storage",
			},
		},
	}
}
