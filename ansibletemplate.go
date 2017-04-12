package main

import "os"
import "fmt"
import "path/filepath"

var rolename string = "testrole1"
var text string = "Creating Ansible-Test-kitchen template for"
var installdir string = "/Users/christiankewerkopf/cinq/git/go/src/ansibletemplate/"
var kitchenyml string = installdir + rolename + "/.kitchen.yml"
var templatedir string = installdir + rolename + "/"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
{
os.Mkdir(templatedir, os.FileMode(511))
os.Mkdir(templatedir + "files", os.FileMode(511))
os.Mkdir(templatedir + "defaults", os.FileMode(511))
os.Mkdir(templatedir + "handlers", os.FileMode(511))
os.Mkdir(templatedir + "test", os.FileMode(511))
os.Mkdir(templatedir + "meta", os.FileMode(511))
os.Mkdir(templatedir + "tasks", os.FileMode(511))
os.Mkdir(templatedir + "templates", os.FileMode(511))
os.Mkdir(templatedir + "vars", os.FileMode(511))
os.Mkdir(templatedir + "test/integration", os.FileMode(511))
os.Mkdir(templatedir + "test/integration/" + rolename, os.FileMode(511))
os.Mkdir(templatedir + "test/integration/spec", os.FileMode(511))
os.Mkdir(templatedir, os.FileMode(511))
}

{
os.Create(templatedir + "/defaults/main.yml")
os.Create(templatedir + "/handlers/main.yml")
os.Create(templatedir + "/meta/main.yml")
os.Create(templatedir + "/tasks/main.yml")
os.Create(templatedir + "/vars/main.yml")
os.Create(templatedir + "/test/integration/spec/rolename_spec.rb")
os.Create(templatedir + "/test/integration/" + rolename +"/default.yml")
}

{
err :=  os.Link(installdir + ".kitchen.yml", kitchenyml)

      if err != nil {
           fmt.Println(err)
           return
}
}

{
fmt.Println(text + " " + rolename)
}

{
files, _ := filepath.Glob(installdir + rolename + "/*")
fmt.Println("created",files)
}

}
