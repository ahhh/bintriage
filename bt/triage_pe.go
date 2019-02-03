package bt

import (
	"log"

	"github.com/Binject/debug/pe"
)

// PeBinTriage - get more info on a PE binary
func PeBinTriage(sourceFile string) error {
	peFile, err := pe.Open(sourceFile)
	if err != nil {
		return err
	}
	for _, section := range peFile.Sections {
		log.Printf("Section details: %+v", section)
	}

	for _, symbol := range peFile.Symbols {
		log.Printf("Symbol details: %+v", symbol)
	}

	libraries, err := peFile.ImportedLibraries()
	if err != nil {
		log.Fatal(err)
	}
	for _, ilib := range libraries {
		log.Printf("Imported lib details: %+v", ilib)
	}

	impSymbs, err := peFile.ImportedSymbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, isymb := range impSymbs {
		log.Printf("Imported symbol details: %+v", isymb)
	}

	return nil
}
