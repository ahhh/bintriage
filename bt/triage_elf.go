package bt

import (
	"debug/elf"
	"log"
)

// ElfBinTriage - Inject shellcode into an ELF binary
func ElfBinTriage(sourceFile string) error {
	elfFile, err := elf.Open(sourceFile)
	if err != nil {
		return err
	}
	for _, section := range elfFile.Sections {
		log.Printf("Section details: %+v", section)
	}

	symbols, err := elfFile.Symbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, symbol := range symbols {
		log.Printf("Symbol details: %+v", symbol)
	}

	dynsymbols, err := elfFile.DynamicSymbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, dynsymbol := range dynsymbols {
		log.Printf("Dynamic symbol details: %+v", dynsymbol)
	}

	libraries, err := elfFile.ImportedLibraries()
	if err != nil {
		log.Fatal(err)
	}
	for _, ilib := range libraries {
		log.Printf("Imported lib details: %+v", ilib)
	}

	impSymbs, err := elfFile.ImportedSymbols()
	if err != nil {
		log.Fatal(err)
	}
	for _, isymb := range impSymbs {
		log.Printf("Imported symbol details: %+v", isymb)
	}

	return nil
}
