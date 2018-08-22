package iotool

type Cp struct {
	From,
	To		string

	fromInfo,
	toInfo		*FileInfo
}

// get cached FileInfo from *Cp.From.
func (cp *Cp) FromInfo() (FileInfoInterface, error) {
	if cp.fromInfo == nil { var e error
		cp.fromInfo, e = NewFileInfo(cp.From); if e != nil { return nil, e }
	}; return cp.fromInfo, nil
}

// get cached FileInfo from *Cp.To.
func (cp *Cp) ToInfo() (FileInfoInterface, error) {
	if cp.toInfo == nil { var e error
		cp.toInfo, e = NewFileInfo(cp.To); if e != nil { return nil, e }
	}; return cp.toInfo, nil
}
