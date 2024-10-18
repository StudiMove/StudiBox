package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL string
	Client  *http.Client
}

// NewAPIClient crée une nouvelle instance de APIClient avec un timeout
func NewAPIClient(baseURL string) *APIClient {
	// S'assurer que l'URL de base contient le schéma HTTP/HTTPS
	if baseURL == "" || (len(baseURL) > 4 && baseURL[:4] != "http") {
		log.Fatalf("API_BASE_URL non valide : %s", baseURL)
	}

	return &APIClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// DoRequest exécute une requête HTTP et gère les erreurs
func (c *APIClient) DoRequest(req *http.Request) (*http.Response, error) {
	log.Printf("Making request to %s %s", req.Method, req.URL)

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Printf("Error executing request: %v", err)
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Received non-2xx status code: %d - %s", resp.StatusCode, string(body))
		return resp, fmt.Errorf("non-2xx status code: %d", resp.StatusCode)
	}

	return resp, nil
}

// PostJSON envoie une requête POST avec un corps JSON
func (c *APIClient) PostJSON(endpoint string, body interface{}, headers map[string]string) (*http.Response, error) {
	url := c.BaseURL + endpoint

	// Convertir le corps en JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	// Ajouter les en-têtes personnalisés
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.DoRequest(req)
}

// ParseJSONResponse lit et parse la réponse JSON dans une structure donnée
func ParseJSONResponse(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Remplacement de ioutil.ReadAll par io.ReadAll
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return nil
}
