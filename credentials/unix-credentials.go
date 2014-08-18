package libgocredentials

import(
	"fmt"
	"errors"
	"bytes"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//TODO
// file locking

// reads and modifys /etc/shadow format
// rewrites entire file on Add(), Modify() and Remove() calls
type Unix struct {
	filename	string
	handler		*os.File
	reader		*bufio.Reader
	nextIndex	uint64
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

func (unix *Unix) find(name string) (user UserInterface, e error) {
	if e = unix.open(); e != nil {
		return
	}

	unix.Reset()
	for {
		user, e = unix.Next()
		if e != nil {
			break
		} else if user.GetName() == name {
			break
		}
	}

	return
}

func (unix *Unix) Next() (user UserInterface, e error) {
	if e = unix.open(); e != nil {
		return
	}

	if unix.reader == nil {
		unix.reader = bufio.NewReader(unix.handler)
	}

	var line []byte
	line, e = unix.reader.ReadBytes('\n')
	if e != nil {
		return
	}

	user, e = unix.parseLine(line)

	user.setIndex(unix.nextIndex)
	unix.nextIndex += uint64(len(line))

	return
}

func (unix *Unix) Reset() {
	if unix.handler != nil {
		unix.handler.Seek(0, 0)
		unix.reader = bufio.NewReader(unix.handler)
		unix.nextIndex = 0
	}
}

func (unix *Unix) Print() {
	if e := unix.open(); e != nil {
		return
	}

	fmt.Println("Id -- User")

	reader := bufio.NewReader(unix.handler)
	for id := 1; ; id++ {
		line, _, e := reader.ReadLine()
		if e != nil {
			break
		}

		user, e := unix.parseLine(line)
		if e != nil {
			break
		}
		fmt.Printf("%v -- %s\n", id, user.GetName())
	}
}

func (unix *Unix) parseLine(line []byte) (user UserInterface, e error) {
	splitted := strings.Split(string(line), ":")
	if len(splitted) != 9 {
		e = errors.New("Unexpected file format!")
	}

	passworder, e := NewUnixPassworderParse(splitted[1])

	return &User{	name: splitted[0],
			passworder: passworder }, e
}

// writes to a temp file and then trys to move that temp file
func (unix *Unix) Add(user UserInterface) (e error) {
	if exists, _ := unix.Get(user.GetName()); exists {
		return UserExistsError
	}

	var index int64
	temp, e := ioutil.TempFile(path.Dir(unix.filename), path.Base(unix.filename))
	if e != nil {
		goto cleanup
	}
	defer temp.Close()

	e = unix.open()
	// old file exists, copy contents
	if e == nil {
		unix.handler.Seek(0, 0)
		index, e = io.Copy(temp, unix.handler)
		if e != nil {
			goto cleanup
		}
	}
	defer unix.Close()

	temp.Write(unix.format(user))

	e = os.Rename(temp.Name(), unix.filename)

cleanup:
	if e != nil {
		os.Remove(temp.Name())
	}

	user.setIndex(uint64(index))

	return
}

func (unix *Unix) Remove(user UserInterface) (e error) {
	temp, e := unix.copyTillUser(user)
	if e != nil {
		return e
	}
	defer unix.Close()

	return unix.skipOneCopyRest(temp)
}

func (unix *Unix) skipOneCopyRest(temp *os.File) (e error) {
	reader := bufio.NewReader(unix.handler)
	// advance pointer
	reader.ReadBytes('\n')

	var line []byte
	for {
		line, e = reader.ReadBytes('\n')
		if e != nil {
			break
		}

		_, e = temp.Write(line)
		if e != nil {
			goto cleanup
		}
	}

	if e == io.EOF {
		// EOF is no error
		e = nil
	}

	e = os.Rename(temp.Name(), unix.filename)

cleanup:
	if e != nil {
		temp.Close()
		os.Remove(temp.Name())
	}

	return
}

func (unix *Unix) Modify(user UserInterface) (e error) {
	if !user.HasChanged() {
		// this is not an error
		return
	}

	temp, e := unix.copyTillUser(user)
	if e != nil {
		return e
	}
	defer unix.Close()

	_, e = temp.Write(unix.format(user))
	if e != nil {
		return
	}

	return unix.skipOneCopyRest(temp)
}

func (unix *Unix) copyTillUser(user UserInterface) (temp *os.File, e error) {
	if user == nil {
		return nil, EmptyError
	}

	user, e = unix.find(user.GetName())
	if e != nil {
		return nil, UserDoesntExistError
	}

	temp, e = ioutil.TempFile(path.Dir(unix.filename), path.Base(unix.filename))
	if e != nil {
		goto cleanup
	}

	e = unix.open()
	if e != nil {
		return nil, FileNotExistsError
	}

	unix.handler.Seek(0, 0)
	_, e = io.CopyN(temp, unix.handler, int64(user.getIndex()))
	if e != nil {
		goto cleanup
	}

cleanup:
	if e != nil {
		os.Remove(temp.Name())
		return
	}

	return
}

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
		for i := 3; i <=7; i++ {
			buffer.WriteString(":")
		}
	}
	buffer.WriteString("\n")

	return buffer.Bytes()
}

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
