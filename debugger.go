package z80

type DebugInstruction struct {
	Address uint16
	Mnemonic string
}

// Disassemble disassembles a Z80 instruction at the given address.
// It returns the string containing the mnemonic form of the
// instruction plus arguments, the address of the next instruction,
// and the eventual shift opcode.
func Disassemble(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	opcode := memory.ReadByte(address)
	return OpcodesDisMap[shift+int(opcode)](memory, address, shift)
}

// DisassemleN disassembles n memory instructions starting from address.
func DisassembleN(memory MemoryReader, address uint16, n int) []DebugInstruction {
	var (
		res string
		shift int
	)
	result := make([]DebugInstruction, n)
	i := 0
 	for i < n {
		instruction := DebugInstruction{Address: address}
		res, address, shift = Disassemble(memory, address, shift)
		instruction.Mnemonic = res
		if shift == 0 {
			result[i] = instruction
			i++
		}
	}
	return result
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
