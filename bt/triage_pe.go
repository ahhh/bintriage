package bt

import (
	"debug/pe"
	"log"
)

// PeBinTriage - get more info on a PE binary
func PeBinTriage(sourceFile string) error {
	peFile, err := pe.Open(sourceFile)
	if err != nil {
		return err
	}
	for _, section := range peFile.Sections {
		log.Println(cyan.Sprintf("Section details: %+v", section))
	}

	for _, symbol := range peFile.Symbols {
		log.Println(blue.Sprintf("Symbol details: %+v", symbol))
	}

	libraries, err := peFile.ImportedLibraries()
	if err != nil {
		log.Fatal(err)
	}
	for _, ilib := range libraries {
		log.Println(cyan.Sprintf("Imported lib details: %+v", ilib))
	}

	impSymbs, err := peFile.ImportedSymbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, isymb := range impSymbs {
		log.Println(blue.Sprintf("Imported symbol details: %+v", isymb))
	}

	return nil
}
