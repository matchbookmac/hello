package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type githubUser struct {
	Name  string `json:"name"`
	Login string `json:"login"`
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func query(user string) (githubUser, error) {
	resp, err := http.Get("https://api.github.com/users/" + user)

	if err != nil {
		return githubUser{}, err
	}

	defer resp.Body.Close()

	var dat githubUser

	if err := json.NewDecoder(resp.Body).Decode(&dat); err != nil {
		return dat, err
	}

	return dat, nil
}

func main() {
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		user := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := query(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":8080", nil)
}
