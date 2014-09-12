package inifile

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

const (
	TEST_FILE = "/home/keith/dev/go/src/github.com/keithtobin/go-inifile/test/nova.conf.sample"
)

func TestNew(t *testing.T) {
	Convey("Create and return a inifile object", t, func() {
		file, err := New()
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestNewFromFile(t *testing.T) {
	Convey("Create and return a inifile object, object is also loaded with data from ini config file on disk", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestNumberOfSections(t *testing.T) {
	Convey("Get the number of sections", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		ns := file.NumberOfSections()
		So(ns, ShouldBeGreaterThan, 0)
		

	})
}

func TestGetSections(t *testing.T) {
	Convey("Get the sections from the inifile", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		sections := file.GetSections()
		So(sections, ShouldNotBeNil)
		So(len(sections), ShouldBeGreaterThan, 0)

		

	})
}

func TestGetSection(t *testing.T) {
	Convey("Get a section form the sections.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		nf, _ := file.GetSection("spice")
		So(nf, ShouldNotBeNil)
		

	})
}


func TestSectionExists(t *testing.T) {
	Convey("Check iof the section exists.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}

//Todo:
func TestSectionAdd(t *testing.T) {
	Convey("Add a section to the inifile.Sections.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		r := file.SectionAdd(IniFileSection{Key : "mysection"})
		So(r, ShouldBeNil)

		n := file.SectionExists("mysection")
		So(n, ShouldBeTrue)
		
	})
}

//Todo:
func TestSectionDelete(t *testing.T) {
	Convey("Delete a section from the sections.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		
		r := file.SectionDelete("spice")
		So(r, ShouldBeNil)

		n = file.SectionExists("spice")
		So(n, ShouldBeFalse)
		

	})
}

//Todo:
func TestReadFile(t *testing.T) {
	Convey("Read the content of a ini file on disk into inifile.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}

//Todo:
func TestGetString(t *testing.T) {
	Convey("Get the string content of the inifile.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}

//Todo:
func TestSaveTo(t *testing.T) {
	Convey("Save the inifile to a file on disk.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}

//Todo:
func TestSave(t *testing.T) {
	Convey("Save the inifile to a file on disk.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}

//Todo:
func TestChangeSectionKeyValue(t *testing.T) {
	Convey("Change a section key value.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}


//Todo:
func TestDeleteSectionKeyValue(t *testing.T) {
	Convey("Delete/remove a section key value pair.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}

//Todo:
func TestEnableDisableSectionKeyValue(t *testing.T) {
	Convey("Enable or disable a section key value.", t, func() {
		file, err := NewFromFile(TEST_FILE)
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		n := file.SectionExists("spice")
		So(n, ShouldBeTrue)
		

	})
}
