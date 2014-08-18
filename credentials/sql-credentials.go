package libgocredentials

import(
	"fmt"
	sqldb "database/sql"
	 _ "github.com/mattn/go-sqlite3"
)

//TODO allow for not using columns hashType and salt
type Sql struct {
	config		*SqlConfig
	database	*sqldb.DB
	rowsNext	*sqldb.Rows
}

func NewSql(sqlConfig *SqlConfig) *Sql {
	sql := &Sql{ config: sqlConfig }
	sql.config.init()
	sql.database = sql.config.database
	return sql
}

func (sql *Sql) SetTransactionIsolationLevel(to string) {
	if IsTransactionIsolationLevel(to) {
		sql.database.Exec("SET TRANSACTION ISOLATION LEVEL " + to)
	}
}

func (sql *Sql) Begin() {
	sql.database.Exec("BEGIN TRANSACTION")
}

func (sql *Sql) Commit() {
	sql.database.Exec("COMMIT")
}

func (sql *Sql) Rollback() {
	sql.database.Exec("ROLLBACK")
}

func (sql *Sql) IsAuthenticated(name, password string) bool {
	found, user := sql.Get(name)
	return found && user.GetPassworder().TestPassword(password)
}

func (sql *Sql) Get(name string) (bool, UserInterface) {
	user, e := sql.find(name)
	return (e == nil && user != nil), user
}

func (sql *Sql) New(user, password string) (UserInterface, error) {
	return CreateUser(user, password)
}

func (sql *Sql) Add(user UserInterface) error {
	query := "INSERT INTO "	+ sql.config.table +
			" ( " +
			sql.config.columnUser + ", " +
			sql.config.columnHash + ", " +
			sql.config.columnHashType + ", " +
			sql.config.columnSalt + " " +
			" ) " +
			" VALUES ( ?, ?, ?, ? )"

	if DEBUG { fmt.Println(query) }

	_, e := sql.database.Exec(query,
					user.GetName(),
					user.GetPassworder().GetPasswordHash(),
					user.GetPassworder().GetHashType(),
					user.GetPassworder().GetSalt(),
					)
	if e != nil {
		if DEBUG { fmt.Println(e.Error()) }
		if sql.isNoSuchTableE(e) {
			if e := sql.create(); e != nil {
				if DEBUG { fmt.Println(e.Error()) }
				return e
			} else {
				return sql.Add(user)
			}
			return e
		}
		return e
	}

	return nil
}

func (sql *Sql) Modify(user UserInterface) error {
	if !user.HasChanged() {
	    return nil
	}

	query := "UPDATE "+ sql.config.table +
			" SET " + sql.config.columnHash + " = ?, " +
			sql.config.columnHashType + " = ?, " +
			sql.config.columnSalt + " = ? " +
			" WHERE " + sql.config.columnUser + " = ? "

	if DEBUG { fmt.Println(query) }

	_, e := sql.database.Exec(query,
					user.GetPassworder().GetPasswordHash(),
					user.GetPassworder().GetHashType(),
					user.GetPassworder().GetSalt(),
					user.GetName(),
					)
	if e != nil {
		if DEBUG { fmt.Println(e.Error()) }
		return e
	}

	return nil
}

func (sql *Sql) Remove(user UserInterface) error {
	query := "DELETE FROM " + sql.config.table +
			" WHERE " + sql.config.columnUser + " = ?"
			// " LIMIT 1" // not standard

	if DEBUG { fmt.Println(query) }

	if _, e := sql.database.Exec(query, user.GetName()); e != nil {
		if DEBUG { fmt.Println(e.Error()) }
		return e
	}

	return nil
}

func (sql *Sql) Next() (UserInterface, error) {
	if sql.rowsNext == nil {
		query := "SELECT " + sql.config.allColumns +
				" FROM " + sql.config.table

		if DEBUG { fmt.Println(query) }

		rows, e := sql.database.Query(query)
		if e != nil {
			if DEBUG { fmt.Println(e.Error()) }
			return nil, e
		}

		sql.rowsNext = rows
	}

	if !sql.rowsNext.Next() {
	    sql.Reset()
	    return nil, EOFError
	}

	return sql.result2User(sql.rowsNext)
}

func (sql *Sql) Reset() {
	if sql.rowsNext != nil {
		sql.rowsNext.Close()
		sql.rowsNext = nil
	}
}

func (sql *Sql) Close() {
	sql.database.Close()
}

func (sql *Sql) find(name string) (UserInterface, error) {
	query := "SELECT " +
			//"COUNT( " + config.columnUser + " ) AS __countcredentials8888__, " +
			sql.config.allColumns +
			" FROM " + sql.config.table +
			" WHERE " + sql.config.columnUser + " = ?"

	if DEBUG { fmt.Println(query) }

	rows, e := sql.database.Query(query, name)
	if e != nil {
	    if DEBUG { fmt.Println(e.Error()) }
	    return nil, e
	}

	if !rows.Next() {
		return nil, nil
	}

	user, e := sql.result2User(rows)
	rows.Close()

	return user, e
}

func (sql *Sql) result2User(rows *sqldb.Rows) (UserInterface, error) {
	var name string
	var hash string
	var hashType string
	var salt string

	if e := rows.Scan(&name, &hash, &hashType, &salt); e != nil {
		if DEBUG { fmt.Println(e.Error()) }
		return nil, e
	}

	user := NewUser(name)
	user.passworder = NewPassworderParsed(hash, hashType, salt)

	return user, nil
}

func (sql *Sql) create() (e error) {
	query := "CREATE TABLE IF NOT EXISTS " + sql.config.table +
			"( " +
			sql.config.columnUser + " varchar(100) UNIQUE, " +
			sql.config.columnHash + " varchar(100), " +
			sql.config.columnHashType + " varchar(10), " +
			sql.config.columnSalt + " varchar(50) " +
			")"

	if DEBUG { fmt.Println(query) }

	if _, e := sql.database.Exec(query); e != nil {
		if DEBUG { fmt.Println(e.Error()) }
		return e
	}

	return nil
}

func (sql *Sql) isNoSuchTableE(e error) bool {
	if e == nil {
	    return false
	}

	message := e.Error()
	return len(message) > 13 && message[:14] == "no such table:"
}
