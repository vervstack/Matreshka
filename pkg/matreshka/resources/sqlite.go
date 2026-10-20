package resources

import "strings"

const (
	SqliteResourceName = "sqlite"

	sqliteInMemoryPath = ":memory:"
	sqlitePragmas      = "_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
)

type Sqlite struct {
	Name `yaml:"resource_name" env:"-"`

	Path             string `yaml:"path"`
	MigrationsFolder `yaml:"migrations_folder,omitempty"`
}

func NewSqlite(n Name) Resource {
	return &Sqlite{
		Name:             n,
		Path:             "/app/data",
		MigrationsFolder: "./migrations",
	}
}

func (p *Sqlite) GetType() string {
	return SqliteResourceName
}

func (p *Sqlite) ConnectionString() string {
	if p.Path == "" || p.Path == sqliteInMemoryPath {
		return p.Path
	}

	separator := "?"
	if strings.Contains(p.Path, "?") {
		separator = "&"
	}

	return p.Path + separator + sqlitePragmas
}

func (p *Sqlite) SqlDialect() string {
	return "sqlite"
}
