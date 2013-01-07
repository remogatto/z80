package z80

// Disassemble disassambles Z80 instruction at address.
// It returns the string containing the mnemonic form of the
// instruction plus arguments, the address of the next instruction,
// and the eventual shift offset.
func Disassemble(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	opcode := memory.ReadByte(address)
	return OpcodesDisMap[shift+int(opcode)](memory, address, shift)
}


// PreviousInstruction returns the address of the instruction prior to
// the one at the given address.
func PreviousInstruction(memory MemoryReader, address uint16) uint16 {
	var currAddress uint16
	shift := 0
	for startingPoint := address - 20; startingPoint != address; startingPoint++ {
		addr := startingPoint
		for addr < address {
			_, currAddress, shift = Disassemble(memory, addr, shift)
			if (currAddress == address) {
				return addr
			}
			addr = currAddress
		}
	}
	return 0
}
