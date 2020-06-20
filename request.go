package main

type IPCheckRequest struct {
	IP        string   `json:"ip"`
	Whitelist []string `json:"whitelist"`
}
