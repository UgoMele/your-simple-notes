package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"your-simple-notes/internal/cli"
)

const (
	notesDirName = "ysn-notes"
	noteExt      = ".md"
)

func main() {
	//get home path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	// home path + notes folder
	notesPath := filepath.Join(homeDir, notesDirName)
	//check if it exists
	if _, err = os.Stat(notesPath); err != nil {
		err = os.Mkdir(notesPath, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}
	
	//parsing arguments
	var noteArg string
	var helpArg bool

	flag.StringVar(&noteArg, "a", "", "specify file name: -a filename or -a category:filename")
	flag.BoolVar(&helpArg, "h", false, "HELP? :)")

	flag.Parse()

	if len(os.Args) == 1 {
		err := cli.Home(notesPath)
		if err != nil {
			log.Println(err)
		}
	} else {
		flag.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "a":
				err := cli.AddNote(f.Value.String(), notesPath, noteExt)
				if err != nil {
					log.Println(err)
				}
			case "h":
				flag.Usage()
			}
		})
	}
}
