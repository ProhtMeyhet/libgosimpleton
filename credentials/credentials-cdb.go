package libgomessage

import(
	"errors"
	cdb "github.com/jbarham/go-cdb"
)

UnknownUserError := errors.New("User unknown!")
WrongPasswordError := errors.New("Password wrong!")

type Cdb struct {
	databaseFile string
	database *cdb.Cdb
}

func NewCdb(database string) *Cdb {
	return &Cdb{ databaseFile: database }
}

func (cdb *Cdb) Add(user, password string) {

}

func (cdb *Cdb) IsAuthenticated(user, password string) (bool, error) {
	if e := cdb.open(); e != nil {
		return false, e
	}

	if storedPasswordHash, e := cdb.database.Find(user); e != nil {
		return false, UnknownUserError
	}
}

func (cdb *Cdb) open() error {
	if cdb.database == nil {
		if database, e := cdb.Open(credentials.databaseFile); e == nil {
			cdb.database = database
		} else {
			return e
		}
	}

	return nil
}
