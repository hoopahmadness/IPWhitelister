package main

import (
	"fmt"
)

type IPCheckRequest struct {
	IP        string   `json:"ip"`
	Whitelist []string `json:"whitelist"`
	Lang      string   `json:"lang"`
}

func (req IPCheckRequest) validate() error {
	if req.IP == "" {
		return fmt.Errorf("Missing IP address in request")
	}
	return nil
}
