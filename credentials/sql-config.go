package libgocredentials

import(
	sqldb "database/sql"
)


type SqlConfig struct {
	base, table, columnUser, columnHash, columnHashType, columnSalt string
	fullTable, allColumns string
	user User
	database *sqldb.DB
}

func (config *SqlConfig) init() {
	if config.columnUser == "" {
		config.columnUser = "user"
	}

	if config.columnHash == "" {
		config.columnHash = "password"

		// set only if columnHash is empty
		// allows to use a different passworder
		if config.columnHashType == "" {
			config.columnHashType =  "hashType"
		}

		if config.columnSalt == "" {
			config.columnSalt = "salt"
		}
	}

	//config.fullTable = config.base + "." + config.table
	config.allColumns = config.columnUser + ", " +
				config.columnHash + ", "

	if config.columnHashType != "" {
		config.allColumns += config.columnHashType + ", "
	}

	if config.columnSalt != "" {
		config.allColumns += config.columnSalt
	}

}
