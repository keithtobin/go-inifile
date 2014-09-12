package main

import (
	"github.com/keithtobin/go-inifile/ini_file"
	"fmt"
	"os"

)


func main() {
	file, err := inifile.NewFromFile("/home/keith/dev/go/src/cloud-guy.net/ini_config/nova.conf.sample")
	if err != nil {
		fmt.Println("Error opening file..............")
		os.Exit(1)
	}

	fmt.Println(file.NumberOfSections()) 

	for _,i := range file.GetSections() {
		fmt.Println(i.Key) 
	}

	nf, _ := file.GetSection("spice")
	fmt.Println(nf)

	for _, m := range nf.Items {
		fmt.Println(m.GetType() + "(" + m.GetString() + ")")
	}

	fmt.Print(file.GetString())
	
	file.ChangeSectionKeyValue("spice", "agent_enabled", "false")

	file.Save()

}
