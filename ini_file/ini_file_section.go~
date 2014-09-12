package inifile

type IniFileSection struct {
	Key string
	//Items []interface{}
	Items []Shaper
	LineNumber int
}

//This function will search the items in this section and return the matching item. 
func (s IniFileSection) GetItem(k string) (v Shaper, e error) {

	for _, m := range s.Items {
		if m.GetKey() == k {
			return m, nil
		}
	}

	return nil, nil
}

//This function will return all the KeyValuePairs for this section.
func (s IniFileSection) GetItems() (m []Shaper, e error){

	return s.Items, nil
}


func (s IniFileSection) GetString() (m string, e error){

	rv := "[" + s.Key + "]\n" 
	for _, m := range s.Items {
		rv = rv + m.GetString()+ "\n"	
	}
	return rv, nil
}

//This function will return this sections key.
//The key  is the name of the section.
//For example, it's name in the inifile will be like [Global].
func (i IniFileSection) GetKey() (s string) {
	return i.Key
}


