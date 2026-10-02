package main

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	jsonBody, err := cfg.client.LocationAreaRequest(cfg.next)
	if err != nil {
		return err
	}
	cfg.prev = jsonBody.Previous
	cfg.next = jsonBody.Next
	for _, area := range jsonBody.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	jsonBody, err := cfg.client.LocationAreaRequest(cfg.prev)
	if err != nil {
		return err
	}
	cfg.prev = jsonBody.Previous
	cfg.next = jsonBody.Next
	for _, area := range jsonBody.Results {
		fmt.Println(area.Name)
	}
	return nil
}
