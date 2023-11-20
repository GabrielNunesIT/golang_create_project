package createproject

import (
	"flag"
	"log"
	"os"
	"slices"

	"github.com/GabrielNunesIT/golang_create_project/config"
)

// This is a func that will create folder and files for the project according to the project name

var projectNameFlag = flag.String("projectName", "project-template", "Use this flag to specify the project name.")
var goVersion = flag.String("goVersion", "1.21", "Use this flag to specify the go version to use.")

func CreateProject() {
	flag.Parse()

	log.Printf("\nCreating project template with project name: %s", *projectNameFlag)

	log.Printf("\nCreating project folders...")
	// Create the project folder
	createFolders(*projectNameFlag)

	log.Printf("\nCreating project files...")
	// Create the project files
	createFiles(*projectNameFlag)

	log.Printf("\nCreating project README...")
	createReadme(*projectNameFlag)

}

func createFolders(projectName string) (err error) {
	//Get current working directory
	var cmdPath string
	cmdPath, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Create the project folder
	err = os.MkdirAll(cmdPath+"/"+projectName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	folderNames := config.GetFolders()
	specialFolders := []string{"cmd", "internal", "pkg"}

	// Create the folders
	for _, folderName := range folderNames {
		if slices.Contains(specialFolders, folderName) {
			err = os.MkdirAll(cmdPath+"/"+projectName+"/"+folderName+"/"+projectName, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			err = os.MkdirAll(cmdPath+"/"+projectName+"/"+folderName, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return
}

func createFiles(projectName string) (err error) {
	//Get current working directory
	var cmdPath string
	cmdPath, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(cmdPath+"/"+projectName+"/internal/"+projectName+"/"+projectName+".go", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	file.WriteString("package " + projectName + "\n")
	file.Close()

	file, err = os.OpenFile(cmdPath+"/"+projectName+"/cmd/"+projectName+"/cmd.go", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	file.WriteString("package cmd_" + projectName + "\n")
	file.Close()

	file, err = os.OpenFile(cmdPath+"/"+projectName+"/"+"go.mod", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	file.WriteString("module " + projectName + "\n")
	file.WriteString("go " + *goVersion + "\n")
	file.Close()

	return err
}

func createReadme(projectName string) (err error) {
	//create the readme file

	var cmdPath string
	cmdPath, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(cmdPath+"/"+projectName+"/"+"README.md", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("This is a golang template project."))
	return err
}
