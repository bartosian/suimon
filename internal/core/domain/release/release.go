package release

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const releseApiUrl = "https://api.github.com/repos/MystenLabs/sui/releases?per_page=50"

// Release represents a GitHub release
type Release struct {
	TagName     string `json:"tag_name"`
	CommitHash  string `json:"target_commitish"`
	Name        string `json:"name"`
	Draft       bool   `json:"draft"`
	PreRelease  bool   `json:"prerelease"`
	PublishedAt string `json:"published_at"`
	CreatedAt   string `json:"created_at"`
	URL         string `json:"html_url"`
	Author      struct {
		Login string `json:"login"`
	} `json:"author"`
}

// getReleases fetches releases for a given repo and filters them by network name
func GetReleases(networkName string) ([]Release, error) {
	resp, err := http.Get(releseApiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var allReleases []Release
	err = json.Unmarshal(body, &allReleases)
	if err != nil {
		return nil, err
	}

	var filteredReleases []Release
	for _, release := range allReleases {
		if strings.HasPrefix(strings.ToLower(release.Name), strings.ToLower(networkName)) {
			filteredReleases = append(filteredReleases, release)
		}
	}

	return filteredReleases, nil
}
