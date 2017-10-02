package main

import (
	"encoding/json"
	"net/http"
)

type GitHub struct{}

type Repo struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
}

func (g GitHub) GetRepos() ([]Repo, error) {
	resp, err := http.Get("https://api.github.com/users/shavenking/repos")
	if err != nil {
		return []Repo{}, err
	}

	defer resp.Body.Close()

	repos := make([]Repo, 0)
	json.NewDecoder(resp.Body).Decode(&repos)

	return repos, nil
}
