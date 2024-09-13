package httpclient

import (
    "net/http"
    "time"
    "log"
)

type APIClient struct {
    BaseURL string
    Client  *http.Client
}

// NewAPIClient crée une nouvelle instance de APIClient avec un timeout
func NewAPIClient(baseURL string) *APIClient {
    return &APIClient{
        BaseURL: baseURL,
        Client:  &http.Client{
            Timeout: 10 * time.Second, // Timeout de 10 secondes
        },
    }
}

// DoRequest exécute une requête HTTP et gère les erreurs
func (c *APIClient) DoRequest(req *http.Request) (*http.Response, error) {
    // Optionnel : Log des requêtes pour le débogage
    log.Printf("Making request to %s %s", req.Method, req.URL)

    resp, err := c.Client.Do(req)
    if err != nil {
        log.Printf("Error executing request: %v", err)
        return nil, err
    }

    return resp, nil
}
