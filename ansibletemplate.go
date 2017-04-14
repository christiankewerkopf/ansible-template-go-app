package main

import "os"
import "fmt"

import flag "github.com/ogier/pflag"

var rolename string
var templatedir string
var installdir string

func init() {
	flag.StringVarP(&rolename, "rolename", "r", "fucknut", "help message for flagname")
	flag.StringVarP(&installdir, "installdir", "i", "./", "help message for flagname")
	flag.Parse()
}

func main() {
	templatedir = fmt.Sprintf("%s/%s/", installdir, rolename)

	CreateTemplateDir(templatedir)
	CreateTemplateDir(templatedir + "files")
	CreateTemplateDir(templatedir + "defaults")
	CreateTemplateDir(templatedir + "handlers")
	CreateTemplateDir(templatedir + "test")
	CreateTemplateDir(templatedir + "meta")
	CreateTemplateDir(templatedir + "tasks")
	CreateTemplateDir(templatedir + "templates")
	CreateTemplateDir(templatedir + "vars")
	CreateTemplateDir(templatedir + "test/integration")
	CreateTemplateDir(templatedir + "test/integration/" + rolename)
	CreateTemplateDir(templatedir + "test/integration/spec")

	os.Create(templatedir + "/defaults/main.yml")
	os.Create(templatedir + "/handlers/main.yml")
	os.Create(templatedir + "/meta/main.yml")
	os.Create(templatedir + "/tasks/main.yml")
	os.Create(templatedir + "/vars/main.yml")
	os.Create(templatedir + "/test/integration/spec/" + rolename + "_spec.rb")
	os.Create(templatedir + "/test/integration/" + rolename + "/default.yml")
}

//CreateTemplateDir creates a directory
func CreateTemplateDir(n string) {
	os.Mkdir(n, os.FileMode(511))
	fmt.Printf("created directory %s with filemode 511\n", n)
}
