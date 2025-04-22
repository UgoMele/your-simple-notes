package cli

import (
	"fmt"
	"log"
	"strings"
	"your-simple-notes/internal/filerw"
)

func AddNote(addArg, notesPath, noteExt string) error {
	category := "default"
	var name string
	if i := strings.IndexRune(addArg, ':'); i != -1 {
		argData := strings.SplitN(addArg, string(':'), 2)
		category = argData[0]
		name = argData[1]
	} else {
		name = addArg
	}

	filePath, err := filerw.Create(notesPath, category, name+noteExt)
	if err != nil {
		return err
	}
	err = filerw.RunInCli(filePath)
	return err
}

func Home(notesPath string) error {
	files, err := filerw.GetLastNotes(notesPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(`                                       
@@@ @@@        @@@@@@        @@@  @@@  
@@@ @@@       @@@@@@@        @@@@ @@@  
@@! !@@       !@@            @@!@!@@@  
!@! @!!       !@!            !@!!@!@!  
 !@!@!        !!@@!!         @!@ !!@!  
  @!!!         !!@!!!        !@!  !!!  
  !!:              !:!       !!:  !!!  
  :!:    :!:      !:!   :!:  :!:  !:!  
   ::    :::  :::: ::   :::   ::   ::  
   :     :::  :: : :    :::  ::    :   
	`)

	fmt.Println("latest edited files:")
	fmt.Println()

	lastCategory := ""

	for i, file := range files {
		if file.Name == "" {
			break
		}
		if file.Category != lastCategory {
			lastCategory = file.Category
			fmt.Printf("     %s\n", file.Category)
		}
		fmt.Printf("[%d]  |  %s\n", i, file.Name)
	}

	var input int
	fmt.Print("\n> ")
	_, err = fmt.Scanf("%d", &input)
	if err != nil {
		return err
	}

	for i, f := range files {
		if input == i {
			return filerw.RunInCli(f.Path)
		}
	}

	return nil
}
