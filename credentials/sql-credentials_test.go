package libgocredentials

import(
	"testing"
	sqldb "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func TestSql(t *testing.T) {
	var e error
	config := &SqlConfig{
				base: ":memory:",
				//base: "anewfile.sqlite",
				//base: "", // creates a temporary database
				table: "ran555dom",
			    }

	if config.database, e = sqldb.Open("sqlite3", config.base); e != nil {
		t.Errorf("error opening sqlite database: %s", e.Error())
	}

	sql := NewSql(config)
	defer sql.Close()

	//sql.Begin()

	genericCredentialsTest(t, sql)

	//sql.Commit()
}
