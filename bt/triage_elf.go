package bt

import (
	"log"

	"github.com/Binject/debug/elf"
)

// ElfBinTriage - Inject shellcode into an ELF binary
func ElfBinTriage(sourceFile string) error {
	elfFile, err := elf.Open(sourceFile)
	if err != nil {
		return err

	}

	// Class
	log.Println(blue.Printf("Elf Class: %+v", elfFile.Class))
	// Data
	log.Println(cyan.Printf("Data: %+v", elfFile.Data))
	// Version
	log.Println(blue.Printf("Elf Version: %+v", elfFile.Version))
	//OSABI
	log.Println(cyan.Printf("OSABI: %+v", elfFile.OSABI))
	//ABIVersion
	log.Println(blue.Printf("Elf ABIVersion: %+v", elfFile.ABIVersion))

	for _, section := range elfFile.Sections {
		log.Println(cyan.Sprintf("Section details: %+v", section))
	}

	symbols, err := elfFile.Symbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, symbol := range symbols {
		log.Println(blue.Sprintf("Symbol details: %+v", symbol))
	}

	dynsymbols, err := elfFile.DynamicSymbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, dynsymbol := range dynsymbols {
		log.Println(cyan.Sprintf("Dynamic symbol details: %+v", dynsymbol))
	}

	libraries, err := elfFile.ImportedLibraries()
	if err != nil {
		log.Fatal(err)
	}
	for _, ilib := range libraries {
		log.Println(blue.Sprintf("Imported lib details: %+v", ilib))
	}

	impSymbs, err := elfFile.ImportedSymbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, isymb := range impSymbs {
		log.Println(cyan.Sprintf("Imported symbol details: %+v", isymb))
	}

	return nil
}
