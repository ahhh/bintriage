package bt

import (
	"debug/macho"
	"log"
)

// MachoBinTriage - Get more info on a Mach-O binary
func MachoBinTriage(sourceFile string) error {
	machoFile, err := macho.Open(sourceFile)
	if err != nil {
		return err
	}
	for _, section := range machoFile.Sections {
		log.Printf("Section details: %+v", section)
	}

	for _, symbol := range machoFile.Symtab.Syms {
		log.Printf("Symbol details: %+v", symbol)
	}

	libraries, err := machoFile.ImportedLibraries()
	if err != nil {
		log.Fatal(err)
	}
	for _, ilib := range libraries {
		log.Printf("Imported lib details: %+v", ilib)
	}

	impSymbs, err := machoFile.ImportedSymbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, isymb := range impSymbs {
		log.Printf("Imported symbol details: %+v", isymb)
	}

	return nil
}
