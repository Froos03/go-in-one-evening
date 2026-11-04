package projectpulse

import (
	"encoding/json"
	"fmt"
	"os"
)

type Site struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	ExpectedStatus int    `json:"expected_status"`
	TimeoutMS      int    `json:"timeout_ms"`
}

func LoadSites(name string) ([]Site, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("load sites: reading file: %w", err)
	}

	sites := []Site{}

	if err := json.Unmarshal(file, &sites); err != nil {
		return nil, fmt.Errorf("load sites: unmarshaling data: %w", err)
	}

	for i := range sites {
		if sites[i].Name == "" {
			return nil, fmt.Errorf("load sites: name required")
		}
		if sites[i].URL == "" {
			return nil, fmt.Errorf("load sites: url required for sites[%d]: %s", i, sites[i].Name)
		}

		if sites[i].ExpectedStatus == 0 {
			sites[i].ExpectedStatus = 200
		}
	}

	return sites, nil
}
