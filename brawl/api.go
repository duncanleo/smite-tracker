package brawl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.brawlstars.com/v1"
)

// Client represents a Brawl Stars API client
type Client struct {
	BaseURL string
	Token   string
}

// NewClient create a new client
func NewClient(token string) Client {
	return Client{BaseURL: baseURL, Token: token}
}

// GetPlayerData get details of a player
func (c Client) GetPlayerData(playerTag string) (Player, error) {
	var player Player
	var url = fmt.Sprintf("%s/players/%s", c.BaseURL, url.QueryEscape(playerTag))
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return player, err
	}
	if resp.StatusCode != http.StatusOK {
		return player, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&player)
	return player, err
}
