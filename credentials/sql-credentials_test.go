package libgocredentials

import(
	"testing"
	sqldb "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func TestSql(t *testing.T) {
	var e error
	config := &SqlConfig{
				Base: ":memory:",
				//Base: "anewfile.sqlite",
				//Base: "", // creates a temporary database
				Table: "ran555dom",
			    }

	if config.Database, e = sqldb.Open("sqlite3", config.Base); e != nil {
		t.Errorf("error opening sqlite database: %s", e.Error())
	}

	sql := NewSql(config)
	defer sql.Close()

	//sql.Begin()

	genericCredentialsTest(t, sql)

	//sql.Commit()
}
