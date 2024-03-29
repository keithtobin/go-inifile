package inifile

type IniFileEmpty struct {
	Key string
	Value string
	LineNumber int
	Type string
	
}

//This function will create a new struct and return it.
func NewIniFileEmpty(LineNumber int) (y *IniFileEmpty) {
	
	rv := new(IniFileEmpty)
	rv.Type = "IniFileEmpty"
	rv.Key = ""
	rv.Value = ""
	rv.LineNumber = 0
	return rv
}

func (i IniFileEmpty) GetKey() (s string) {
	return i.Key
}

func (i IniFileEmpty) GetValue() (s string) {
	return i.Value
}

func (i IniFileEmpty) GetLineNumber() (s int) {
	return i.LineNumber
}

func (i IniFileEmpty) GetType() (s string) {
	return i.Type
}

func (i IniFileEmpty) GetString() (s string) {
	return i.Value
}
/*
func (i *IniFileEmpty) SetKey(k string) {
	i.Key = k
}
*/
func (i *IniFileEmpty) SetValue(v string) {
	i.Value = v
}
/*
func (i *IniFileEmpty) SetLineNumber(l int) {
	i.LineNumber = l
}
*/
