package bt

import (
	"debug/macho"
	"log"
)

// FatBinTriage - Get more info on a Mach-O binary
func FatBinTriage(sourceFile string) error {
	fatFile, err := macho.OpenFat(sourceFile)
	if err != nil {
		return err
	}
	for _, arch := range fatFile.Arches {
		log.Println(red.Sprintf("Arch details: %v", arch))
		machoFile := arch.File
		for _, section := range machoFile.Sections {
			log.Println(cyan.Sprintf("Section details: %+v", section))
		}

		for _, symbol := range machoFile.Symtab.Syms {
			log.Println(blue.Sprintf("Symbol details: %+v", symbol))
		}

		libraries, err := machoFile.ImportedLibraries()
		if err != nil {
			log.Fatal(red.Sprintf(err.Error()))
		}
		for _, ilib := range libraries {
			log.Println(cyan.Sprintf("Imported lib details: %+v", ilib))
		}

		impSymbs, err := machoFile.ImportedSymbols()
		if err != nil {
			log.Fatal(red.Sprintf(err.Error()))
		}
		for _, isymb := range impSymbs {
			log.Println(blue.Sprintf("Imported symbol details: %+v", isymb))
		}
	}

	return nil
}
