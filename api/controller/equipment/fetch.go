package equipment

import (
	"encoding/json"
	"net/http"
	"prototype/constant"
)

func FetchImage(equipment string) (string, error) {
	// Get Publick API from pexel
	url := "https://api.pexels.com/v1/search?query=" + equipment
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Authorization to access api pexel with your secret api pexel
	req.Header.Set("Authorization", constant.ImageApiKey())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Decode to json from url images
	var response struct {
		Photos []struct {
			Src struct {
				Original string `json:"original"`
			} `json:"src"`
		} `json:"photos"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	// If image is not found return, and if dound get the first image from api list
	if len(response.Photos) == 0 {
		return "", nil
	}
	return response.Photos[0].Src.Original, nil
}
