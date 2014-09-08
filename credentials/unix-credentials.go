package libgocredentials

import(
	"bytes"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	//"time"
)

//TODO
// file locking

// reads and modifys /etc/shadow format
// rewrites entire file on Add(), Modify() and Remove() calls
// writes are done to a temp file which is then moved
type Unix struct {
	filename	string
	handler		*os.File
	reader		*bufio.Reader
	nextIndex	uint64
	// modificationTime time.Time
}

func NewUnix(newFilename string) *Unix {
	return &Unix{ filename: newFilename }
}

func (unix *Unix) IsAuthenticated(name, password string) bool {
	found, user := unix.Get(name)
	return found && user.GetPassworder().TestPassword(password)
}

func (unix *Unix) New(user, password string) (UserInterface, error) {
	return CreateUnixUser(user, password)
}

func (unix *Unix) Get(name string) (found bool, user UserInterface) {
	user, e := unix.find(name)
	return (e == nil && user != nil), user
}

func (unix *Unix) Next() (user UserInterface, e error) {
	user, e, _, _ = unix.next(false)
	return
}

func (unix *Unix) Reset() {
	if unix.handler != nil {
		unix.handler.Seek(0, 0)
		unix.reader = bufio.NewReader(unix.handler)
		unix.nextIndex = 0
	}
}

// writes to a temp file and then trys to move that temp file
func (unix *Unix) Add(user UserInterface) (e error) {
	if exists, _ := unix.Get(user.GetName()); exists { return UserExistsError }

	var index int64
	temp, e := ioutil.TempFile(path.Dir(unix.filename), path.Base(unix.filename))
	if e != nil { goto cleanup }
	defer temp.Close()

	e = unix.open()
	// old file exists, copy contents
	if e == nil {
		unix.handler.Seek(0, 0)
		index, e = io.Copy(temp, unix.handler)
		if e != nil { goto cleanup }
	}
	defer unix.Close()

	if _, e = temp.Write(unix.format(user)); e != nil { goto cleanup }

	user.setIndex(uint64(index))

cleanup:
	if e != nil { os.Remove(temp.Name()) }

	return unix.commit(temp)
}

func (unix *Unix) Remove(user UserInterface) (e error) {
	temp, e := unix.copyTillUser(user)
	if e != nil { goto cleanup }
	defer unix.Close()
	defer unix.Reset()

	if e = unix.copyRest(temp); e != nil { goto cleanup }

cleanup:
	if e != nil { os.Remove(temp.Name()) }

	return unix.commit(temp)
}

func (unix *Unix) Modify(user UserInterface) (e error) {
	// this is not an error
	if !user.HasChanged() { return }

	temp, e := unix.copyTillUser(user)
	if e != nil { return e }
	defer unix.Close()
	defer unix.Reset()

	_, e = temp.Write(unix.format(user))
	if e != nil { goto cleanup }

	if e = unix.copyRest(temp); e != nil { goto cleanup }

cleanup:
	if e != nil { os.Remove(temp.Name()) }


	return unix.commit(temp)
}

// TODO locking should be implemented here with modification time check
func (unix *Unix) commit(temp *os.File) (e error) {
	// this is not an error
	if unix.handler == nil { return	}

	/*
	check modification time
	if modificationTime, e := unix.handler.Stat().ModTime(); e != nil {
		goto cleanup
	}
	if modificationTime != unix.modificationTime { goto cleanup }
	*/

	e = os.Rename(temp.Name(), unix.handler.Name());

//cleanup:
	if e != nil {
		os.Remove(temp.Name())
		return TransactionAbortedError
	}

	return
}

// copys until a user is found
func (unix *Unix) copyTillUser(user UserInterface) (temp *os.File, e error) {
	if user == nil { return nil, EmptyError }

	temp, e = ioutil.TempFile(path.Dir(unix.filename), path.Base(unix.filename))
	if e != nil { return }

	return temp, unix.doCopy(temp, user)
}

// copys whats left
// io.copy not used, because it copied the whole file again
// i havn't got the time to debug that, probably some pointer issue
func (unix *Unix) copyRest(temp *os.File) (e error) {
	return unix.doCopy(temp, nil)
}

// actual copy function
// if tillUser is nil, it copies everything
func (unix *Unix) doCopy(temp *os.File, tillUser UserInterface) (e error) {
	var found bool

	e = unix.open()
	if e != nil { goto cleanup }

	if tillUser == nil {
		for {
			_, e, _, line := unix.next(true)
			if e != nil {  break }

			temp.Write(line)
		}
	} else {
		for {
			_, e, name, line := unix.next(true)
			if e != nil {
				break
			} else if name == tillUser.GetName() {
				found = true
				break
			}

			temp.Write(line)
		}

		if !found { e = UserDoesntExistError }
	}

cleanup:
	if e != nil { return os.Remove(temp.Name()) }

	return
}

// parses only the name from the format
func (unix *Unix) parseName(line []byte) (name string, e error) {
	splitted := strings.SplitN(string(line), ":", 2)

	// dont return yet, maybe the caller can debug better
	// even with the garbled format
	if len(splitted) != 2 { e = UnexpectedFileFormatError }

	name = splitted[0]

	return
}

// parse the whole line and return UserInterface
func (unix *Unix) parseLine(line []byte) (user UserInterface, e error) {
	splitted := strings.Split(string(line), ":")

	// dont return yet, maybe the caller can debug better
	// even with the garbled format
	if len(splitted) != 9 { e = UnexpectedFileFormatError }

	passworder, e := NewUnixPassworderParse(splitted[1])

	return &User{	name: splitted[0],
			passworder: passworder }, e
}

// format a user
func (unix *Unix) format(user UserInterface) (line []byte) {
	buffer := bytes.NewBuffer(line)
	buffer.WriteString(user.GetName())
	buffer.WriteString(":")
	buffer.Write(user.GetPassworder().format())
	if unixUser, ok := user.(*UnixUser); ok {
		buffer.WriteString(":")
		buffer.WriteString(unixUser.daysPasswordChanged)
		buffer.WriteString(":")
		buffer.WriteString(unixUser.daysPasswordCanBeChanged)
		buffer.WriteString(":")
		buffer.WriteString(unixUser.daysPasswordMustBeChanged)
		buffer.WriteString(":")
		buffer.WriteString(unixUser.daysPasswordChangeWarning)
		buffer.WriteString(":")
		buffer.WriteString(unixUser.daysExpiration)
		buffer.WriteString(":")
		buffer.WriteString(unixUser.daysExpiered)
		buffer.WriteString(":")
		buffer.WriteString(unixUser.reserved)
	} else {
		for i := 3; i <= 7; i++ {
			buffer.WriteString(":")
		}
	}

	buffer.WriteString("\n")

	return buffer.Bytes()
}

// read next user
// if onlyName == true the user will be parsed into
// UserInterface. name return is then nil!
//
// if onlyName == true the user won't be parsed. user return is then nil!
func (unix *Unix) next(onlyName bool) (user UserInterface, e error, name string,
					line []byte) {
	if e = unix.open(); e != nil { return }

	if unix.reader == nil {
		unix.reader = bufio.NewReader(unix.handler)
	}

	line, e = unix.reader.ReadBytes('\n')
	if e != nil { return }

	if onlyName {
		name, e = unix.parseName(line)
	} else {
		user, e = unix.parseLine(line)
		user.setIndex(unix.nextIndex)
		unix.nextIndex += uint64(len(line))
	}

	return
}

// find a user in file
func (unix *Unix) find(name string) (user UserInterface, e error) {
	if e = unix.open(); e != nil { return }

	unix.Reset()
	for {
		user, e = unix.Next()
		if e != nil || user.GetName() == name {
			break
		}
	}

	return
}

// open
func (unix *Unix) open() (e error) {
	if unix.handler == nil {
		var handler *os.File
		if handler, e = os.Open(unix.filename); e == nil {
			unix.handler = handler
		} else {
			return
		}
	}

	return
}

func (unix *Unix) Close() (e error) {
	if unix.handler != nil {
		e = unix.handler.Close()
		unix.handler = nil
	}

	return
}
