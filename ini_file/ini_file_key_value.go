package inifile


type IniFileIKeyValue struct {
	Key string
	Value string
	LineNumber int
	Type string
}

//This function will create a new struct and return it.
func NewIniFileIKeyValue(k string, v string, l int) (y *IniFileIKeyValue) {
	
	rv := new(IniFileIKeyValue)
	rv.Type = "IniFileIKeyValue"
	rv.Key = k
	rv.Value = v
	rv.LineNumber = l
	return rv
}

func (i IniFileIKeyValue) GetKey() (s string) {
	return i.Key
}

func (i IniFileIKeyValue) GetValue() (s string) {
	return i.Value
}

func (i IniFileIKeyValue) GetLineNumber() (s int) {
	return i.LineNumber
}

func (i IniFileIKeyValue) GetType() (s string) {
	return i.Type
}

func (i IniFileIKeyValue) GetString() (s string) {
	return i.Key + "=" + i.Value
}

/*
func (i *IniFileIKeyValue) SetKey(k string) {
	i.Key = k
}
*/
func (i *IniFileIKeyValue) SetValue(v string) {
	i.Value = v
}
/*
func (i *IniFileIKeyValue) SetLineNumber(l int) {
	i.LineNumber = l
}
*/
