package bt

// BinTriage - function to get some struct info on a binary
func BinTriage(sourceFile string) error {

	binType, err := BinaryMagic(sourceFile)
	var bintriage func(string) error
	switch binType {
	case ELF:
		bintriage = ElfBinTriage
	case MACHO:
		bintriage = MachoBinTriage
	case FAT:
		bintriage = FatBinTriage
	case PE:
		bintriage = PeBinTriage
	case ERROR:
		return err
	}
	return bintriage(sourceFile)
}
