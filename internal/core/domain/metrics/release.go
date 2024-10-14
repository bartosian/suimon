package metrics

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

const releseAPIURL = "https://api.github.com/repos/MystenLabs/sui/releases?per_page=50"

// Release represents a GitHub release.
type Release struct {
	TagName     string `json:"tag_name"`
	CommitHash  string `json:"target_commitish"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	CreatedAt   string `json:"created_at"`
	URL         string `json:"html_url"`
	Author      struct {
		Login string `json:"login"`
	} `json:"author"`
	Draft      bool `json:"draft"`
	PreRelease bool `json:"prerelease"`
}

// getReleases fetches releases for a given repo and filters them by network name.
func GetReleases(networkName string) ([]Release, error) {
	resp, err := http.Get(releseAPIURL)
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			slog.Error("failed to close response body", "error", closeErr)
		}
	}()

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
