package libgocredentials

import(
	"fmt"
	sqldb "database/sql"
)

type Sql struct {
	config		*SqlConfig
	database	*sqldb.DB
	rowsNext	*sqldb.Rows
}

func NewSql(sqlConfig *SqlConfig) *Sql {
	sql := &Sql{ config: sqlConfig }
	sql.config.init()
	sql.database = sql.config.Database
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

func (sql *Sql) Add(user UserInterface) (e error) {
	query := "INSERT INTO "	+ sql.config.Table +
			" ( " +
			sql.config.columnUser + ", " +
			sql.config.columnHash

	if sql.config.columnHashType == "" {
		query += " ) " +
			" VALUES ( ?, ? )"

		if DEBUG { fmt.Println(query) }

		_, e = sql.database.Exec(query,
						user.GetName(),
						user.GetPassworder().format(),
					)
	} else {
		query += " , " +
			sql.config.columnHashType + ", " +
			sql.config.columnSalt + " " +
			" ) " +
			" VALUES ( ?, ?, ?, ? )"

		if DEBUG { fmt.Println(query) }

		_, e = sql.database.Exec(query,
						user.GetName(),
						user.GetPassworder().GetPasswordHash(),
						user.GetPassworder().GetHashType(),
						user.GetPassworder().GetSalt(),
						)
	}

	// try to create the tables
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

func (sql *Sql) Modify(user UserInterface) (e error) {
	if !user.HasChanged() {
	    return nil
	}

	query := "UPDATE " + sql.config.Table +
			" SET " + sql.config.columnHash + " = ? "

	if sql.config.columnHashType == "" {
		query += " WHERE " + sql.config.columnUser + " = ? "

		if DEBUG { fmt.Println(query) }

		_, e = sql.database.Exec(query,
						user.GetPassworder().format(),
						user.GetName(),
					    )
	} else {
		query +=", " + sql.config.columnHashType + " = ? " +
			sql.config.columnSalt + " = ? " +
			" WHERE " + sql.config.columnUser + " = ? "

		if DEBUG { fmt.Println(query) }

		_, e = sql.database.Exec(query,
						user.GetPassworder().GetPasswordHash(),
						user.GetPassworder().GetHashType(),
						user.GetPassworder().GetSalt(),
						user.GetName(),
						)
	}

	if DEBUG { fmt.Println(e) }
	if e != nil {  return e }

	return nil
}

func (sql *Sql) Remove(user UserInterface) (e error) {
	query := "DELETE FROM " + sql.config.Table +
			" WHERE " + sql.config.columnUser + " = ?"
			// " LIMIT 1" // not standard

	if DEBUG { fmt.Println(query) }

	if _, e = sql.database.Exec(query, user.GetName()); e != nil { return }

	return nil
}

func (sql *Sql) Next() (UserInterface, error) {
	if sql.rowsNext == nil {
		query := "SELECT " + sql.config.allColumns +
				" FROM " + sql.config.Table

		if DEBUG { fmt.Println(query) }

		rows, e := sql.database.Query(query)

		if DEBUG { fmt.Println(e) }
		if e != nil { return nil, e }

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

func (sql *Sql) Close() error {
	return sql.database.Close()
}

func (sql *Sql) find(name string) (UserInterface, error) {
	query := "SELECT " +
			sql.config.allColumns +
			" FROM " + sql.config.Table +
			" WHERE " + sql.config.columnUser + " = ?"

	if DEBUG { fmt.Println(query) }

	rows, e := sql.database.Query(query, name)
	defer rows.Close()

	if DEBUG { fmt.Println(e) }
	if e != nil { return nil, e }

	if !rows.Next() {
		return nil, nil
	}

	return sql.result2User(rows)
}

func (sql *Sql) result2User(rows *sqldb.Rows) (user UserInterface, e error) {
	var passworder PassworderInterface

	var name string
	var hash string

	if sql.config.columnHashType == "" {
		if e = rows.Scan(&name, &hash); e != nil { return nil, e }

		user = NewUser(name)
		passworder, e = NewPassworderParse(hash)
		user.setPassworder(passworder)
		if DEBUG { fmt.Println(hash, user.GetPassworder()) }
	} else {
		var hashType string
		var salt string
		if e = rows.Scan(&name, &hash, &hashType, &salt); e != nil {
			return nil, e
		}

		user = NewUser(name)
		passworder = NewPassworderParsed(hash, hashType, salt)
		user.setPassworder(passworder)
	}

	return
}

func (sql *Sql) create() (e error) {
	query := "CREATE TABLE IF NOT EXISTS " + sql.config.Table +
			"( " +
			sql.config.columnUser + " varchar(100) UNIQUE " +
			", " + sql.config.columnHash + " varchar(100) "
	if sql.config.columnHashType != "" {
		query += ", " + sql.config.columnHashType + " varchar(10) " +
			 ", " + sql.config.columnSalt + " varchar(50) "
	}

	query += ")"

	if DEBUG { fmt.Println(query) }

	if _, e := sql.database.Exec(query); e != nil { return e }

	return nil
}

func (sql *Sql) isNoSuchTableE(e error) bool {
	if e == nil {
		return false
	}

	message := e.Error()
	return len(message) > 13 && message[:14] == "no such table:"
}
