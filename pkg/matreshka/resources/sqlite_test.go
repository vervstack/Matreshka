package resources

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SqliteConnectionString_Scenarios(t *testing.T) {
	cases := []struct {
		name string
		path string
		want string
	}{
		{"file path gets pragmas", "/app/data", "/app/data?" + sqlitePragmas},
		{"path with query gets pragmas appended", "file:/app/data?cache=shared", "file:/app/data?cache=shared&" + sqlitePragmas},
		{"in memory path stays untouched", ":memory:", ":memory:"},
		{"empty path stays untouched", "", ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sqlite := Sqlite{Path: tc.path}
			got := sqlite.ConnectionString()
			require.Equal(t, tc.want, got)
		})
	}
}
