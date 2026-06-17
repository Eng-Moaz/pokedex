package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
    Count    int     `json:"count"`
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func ReqAndUnmarshal(apiUrl string) (LocationAreaResponse, error){
	var locationArea LocationAreaResponse
	res, err := http.Get(apiUrl)
	if err != nil{
		return locationArea, err 
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return locationArea, fmt.Errorf("failed to retrieve %v", res.StatusCode)
	}
	
	if err != nil{
		return locationArea, err
	}

	err = json.Unmarshal(body, &locationArea)
	if err != nil{
		return locationArea, err
	}

	return locationArea, nil
}


