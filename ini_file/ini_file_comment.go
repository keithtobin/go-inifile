package inifile

type IniFileComment struct {
	Key string
	Value string
	LineNumber int
	Type string
}

//This function will create a new struct and return it.
func NewIniFileComment(p string) (y *IniFileComment) {
	
	rv := new(IniFileComment)
	rv.Type = "IniFileComment"
	rv.Key = ""
	rv.Value = p
	rv.LineNumber = 0
	return rv
}

func (i IniFileComment) GetKey() (s string) {
	return i.Key
}

func (i IniFileComment) GetValue() (s string) {
	return i.Value
}

func (i IniFileComment) GetLineNumber() (s int) {
	return i.LineNumber
}

func (i IniFileComment) GetType() (s string) {
	return i.Type
}

func (i IniFileComment) GetString() (s string) {
	return i.Value
}

/*
func (i *IniFileComment) SetKey(k string) {
	i.Key = k
}
*/
func (i *IniFileComment) SetValue(v string) {
	i.Value = v
}
/*
func (i *IniFileComment) SetLineNumber(l int) {
	i.LineNumber = l
}
*/
