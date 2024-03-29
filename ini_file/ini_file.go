//
//
//
//
//
//
//
//
// Notes:
//		Disabled keys in the inifile are defined as keys that are commented out.
//		When you enable a key it will be uncommented in the inifile.
//		when you disable a key it will be commented out.
package inifile

import (
	"os"
	"bufio"
	"io"
	//"fmt"
	"strings"
	"io/ioutil"
	"errors"

)

type Shaper interface {
	GetKey() string
	GetValue() string
	GetLineNumber() int 
	GetType() string  
	GetString() string
	//SetKey(v string) 
	SetValue(v string) 
	//SetLineNumber(v int)  

}


//This represents a ini file and all it's data.
type IniFile struct {
	path string //This holds the path to the ini file
	Sections []IniFileSection //This holds the ini file sections
}

//This will create a new IniFile and return it.
func New() (m IniFile, err error) {
	m = IniFile{}
	return m, nil
}

//This will create a new IniFile read a file and return it.
func NewFromFile(s string) (m IniFile, err error) {
	m = IniFile{}
	_, err = m.ReadFile(s)
	return m, err
}

/* This looks like a duplicit, remove.
//This will create a new IniFile and read a file from disk and return the IniFile.
func NewAndReadFile(s string) (m IniFile, err error) {
	m = IniFile{}
	_, e := m.ReadFile(s)
	return m, e
}
*/

//This function will return the number of sections in the ini file.
func (i IniFile) NumberOfSections() (n int) {
	return len(i.Sections)
}

//This function will return a slice holding the sections in the ini file.
func (i IniFile) GetSections() (n []IniFileSection) {
	return i.Sections
}

//This function will search and return a section.
func (i IniFile) GetSection(n string) (s IniFileSection, e error) {

	if len(i.Sections) > 0 {

		for _, v := range i.Sections {
			if v.Key == n {
				return v, nil
			}
		} 
	}

	//Todo: Put no nextion found error here
	return IniFileSection{}, nil
}

//This runction will return true if section exists or false if not.
func (i *IniFile) SectionExists(s string) (t bool) {
	
	if len(i.Sections) > 0 {
		for _, m := range i.Sections {
 			if m.GetKey() == s {
				return true
			}
		}
	}

	return false
}

//This function will add a section to the inifile sectiosn.
func (i *IniFile) SectionAdd(s IniFileSection) (e error) {
	i.Sections = append(i.Sections, s)	
	return nil
}

//This function will delete a section to the inifile sectiosn.
func (i *IniFile) SectionDelete(s string) (e error) {

	found := false
	foundIndex := 0
	for m, v := range(i.Sections) {
		if v.Key == s {
				
			found = true
			foundIndex = m
			break
		}
	}
	
	if found == true {
		i.Sections = append(i.Sections[:foundIndex], i.Sections[foundIndex+1:]...)
	}	

	return nil
}

//This function will read a ini file from disk into a IniFile.
func (i *IniFile) ReadFile(s string) (m int, e error) {

	i.path = s

	//Create global section
	var sectionIndex = 0
	i.Sections = make ([]IniFileSection,1)
	i.Sections[sectionIndex] = IniFileSection{Key: "Global", LineNumber: 0}
 
	f, err := os.Open(s)
  	if err != nil {
    		return 0, err
  	}

 	defer f.Close()
	
	var lineNumber int
	//var sectionItemsCounter int
	r := bufio.NewReader(f)
  	for {
    		path, err := r.ReadString(10) // 0x0A separator = newline
    		if err == io.EOF {	
      			break
    		} else if err != nil {
      			return 0,err 
		}
		
		path = strings.TrimSpace(path)
		
		if len(path) == 0 {//Empty line
			//ni := IniFileEmpty{}
			i.Sections[sectionIndex].Items = append(i.Sections[sectionIndex].Items, NewIniFileEmpty(lineNumber))
			//i.Sections[sectionIndex].Items = append(i.Sections[sectionIndex].Items, NewIniFileEmpty(lineNumber))
			//i.Sections[sectionIndex].Items = append(i.Sections[sectionIndex].Items, new(IniFileEmpty))
			//i.Sections[sectionIndex].Items[sectionItemsCounter] = ni
			//sectionItemsCounter++	
		} else if path[0] == 35 {//Comment line
			//ni := IniFileComment{}
			i.Sections[sectionIndex].Items = append(i.Sections[sectionIndex].Items, NewIniFileComment(path))
			//i.Sections[sectionIndex].Items[sectionItemsCounter] = ni
			//sectionItemsCounter++	
		} else if path[0] == 91 {//Section line
			sectionIndex++
			i.Sections = append(i.Sections,IniFileSection{})
			v, _ := extractSection(path)
			i.Sections[sectionIndex] = IniFileSection{Key: v, LineNumber: lineNumber}
			i.Sections[sectionIndex].Items = make([]Shaper, 0)
			//sectionItemsCounter = 0

		} else {//Key value line	
			//ni := IniFileIKeyValue{}
			key, value, _ := extractKeyValue(path)
			//i.Sections[sectionIndex].Items = append(i.Sections[sectionIndex].Items, IniFileIKeyValue{Key : key, Value : value})
			i.Sections[sectionIndex].Items = append(i.Sections[sectionIndex].Items, NewIniFileIKeyValue(key, value, lineNumber))		
			//i.Sections[sectionIndex].Items[sectionItemsCounter] = ni
			//sectionItemsCounter++		
		}
		
		lineNumber++
		
	}

	return lineNumber, nil;
}

//This function will extract a key/value pair from the incomming string.
func extractKeyValue(p string) (k string, v string, e error) {
	i := strings.Index(p,"=")
	if i != -1 {
		key := p[:i]
		value := p[i+1:len(p)]
		return key, value, nil
	}

	return "", "", nil
}

//This function will extract the section name from the section line.
func extractSection(s string) (name string, e error) {
	s = strings.TrimSpace(s)
	
	i := strings.Index(s,"[")
	if i < 0 {
		//ToDo: Error handle here
	}
	
	if i != 0 {
		//ToDo: Error handle here
	}

	m := strings.Index(s,"]")
	if m < 0 {
		//ToDo: Error handle here
	}	

	r := s[i+1:m]

	return strings.TrimSpace(r), nil
	
}

//This function will return the ini file as a string.
func (i IniFile) GetString() (s string) {

	rv := ""
	for _, m := range i.Sections {
		tv, _ := m.GetString()
		rv = rv + tv + "\n"
	}
	return rv
}

//This will save IniFile to any file on disk.
func (i IniFile) SaveTo(s string) (e error) {
	
	err := ioutil.WriteFile(s, []byte(i.GetString()), 0644)
    	if err != nil { return err }

	return nil
}

//This will save IniFile to the file on disk from witch the file was loaded.
func (i IniFile) Save() (e error) {
	
	err := ioutil.WriteFile(i.path, []byte(i.GetString()), 0644)
    	if err != nil { return err }

	return nil
}



//This function will search and find a section key/value pair and change the value.
//if the section dose not exist and can not be found a error will be returned.
//This function will only work on keys that are enabled.
func (i *IniFile) ChangeSectionKeyValue(s string, k string, v string) (e error) {

	if i.SectionExists(s) == false { return errors.New("Section '" + s + "'" + "dose not exist")}
	
	se, err := i.GetSection(s)
	if err != nil {return err}

	it, err := se.GetItem(k)
	if err != nil {return err}

	it.SetValue(v)

	return nil
}

//This function will delete/remove a Keyvalue from a section, by this we mean the key will be totaly delted.note present.
func (i *IniFile) DeleteSectionKeyValue(s string, k string) (e error) {

	if i.SectionExists(s) == false { return errors.New("Section '" + s + "'" + "dose not exist")}
	

	return nil
}

//This function will enable or disable a section keyvalue, enable/disable means the keyvalue will be commented out or uncommented. 
func (i *IniFile) EnableDisableSectionKeyValue(s string, k string, y bool) (e error) {

	if i.SectionExists(s) == false { return errors.New("Section '" + s + "'" + "dose not exist")}
	

	return nil
}

/*

//This function will clear the content of the current ini file.
func (i IniFile) Clear() (e error) {
	
	return nil
}

//This function will add a key value to a section
func (i IniFile) AddKeyValueToSection(k string, v string, s string) (e error) {
	
	return nil
}


*/ 




