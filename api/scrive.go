package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"fluffy-enigma/config"
)

type ScriveClient struct {
	baseURL     string
	accessToken string
}

func NewScriveClient() *ScriveClient {
	return &ScriveClient{
		baseURL:     config.AppConfig.ScriveAPI.BaseURL,
		accessToken: config.AppConfig.ScriveAPI.AccessToken,
	}
}

type CreateDocumentRequest struct {
	Title       string   `json:"title"`
	Signatories []string `json:"signatories"`
	Document    string   `json:"document"` // Base64 encoded document
}

type CreateDocumentResponse struct {
	DocumentID string `json:"documentId"`
}

func (c *ScriveClient) CreateDocument(title string, signatories []string, documentBase64 string) (*CreateDocumentResponse, error) {
	reqBody := CreateDocumentRequest{
		Title:       title,
		Signatories: signatories,
		Document:    documentBase64,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/documents", c.baseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response CreateDocumentResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

func (c *ScriveClient) InitiateSigning(documentID string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/documents/%s/sign", c.baseURL, documentID), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
