package tables

import (
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainrelease "github.com/bartosian/suimon/internal/core/domain/metrics"
)

var (
	ColumnsConfigRelease = ColumnsConfig{
		enums.ColumnNameIndex:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameReleaseName: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameDraft:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCommit:      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNamePublishedAt: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameAuthor:      NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNamePreRelease:  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameURL:         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameCreatedAt:   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}
	RowsRelease = RowsConfig{
		0: {
			enums.ColumnNameIndex,
			enums.ColumnNameReleaseName,
			enums.ColumnNameDraft,
			enums.ColumnNamePreRelease,
			enums.ColumnNameCommit,
			enums.ColumnNameURL,
			enums.ColumnNameAuthor,
			enums.ColumnNameCreatedAt,
			enums.ColumnNamePublishedAt,
		},
	}
)

// GetReleaseColumnValues returns a map of ColumnName keys to corresponding values for the specified release.
// The function retrieves information about the release from the provided metrics.Release object and formats it into a map of ColumnName keys and corresponding values.
// Returns a map of ColumnName keys to corresponding values.
func GetReleaseColumnValues(idx int, release *domainrelease.Release) ColumnValues {
	return ColumnValues{
		enums.ColumnNameIndex:       idx + 1,
		enums.ColumnNameReleaseName: release.Name,
		enums.ColumnNameTagName:     release.TagName,
		enums.ColumnNameCommit:      release.CommitHash,
		enums.ColumnNameAuthor:      release.Author.Login,
		enums.ColumnNamePublishedAt: release.PublishedAt,
		enums.ColumnNameDraft:       release.Draft,
		enums.ColumnNamePreRelease:  release.PreRelease,
		enums.ColumnNameURL:         release.URL,
		enums.ColumnNameCreatedAt:   release.CreatedAt,
	}
}
