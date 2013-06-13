package z80

// MemoryAccessor is an interface to access memory addressed by the
// Z80. 
// It defines four read/write method for accessing memory, taking
// into account contention when needed. In systems where memory
// contention is not an issue ReadByte and WriteByte should simply
// call ReadByteInternal and WriteByteInternal.
type MemoryAccessor interface {
	// ReadByte reads a byte from address taking into account
	// contention.
	ReadByte(address uint16) byte

	// ReadByteInternal reads a byte from address without taking
	// into account contention.
	ReadByteInternal(address uint16) byte

	// WriteByte writes a byte at address taking into account
	// contention.
	WriteByte(address uint16, value byte)

	// WriteByteInternal writes a byte at address without taking
	// into account contention.
	WriteByteInternal(address uint16, value byte)

	// Follow contention methods. Leave unimplemented if you don't
	// care about memory contention.

	// ContendRead increments the Tstates counter by time as a
	// result of a memory read at the given address.
	ContendRead(address uint16, time int)

	ContendReadNoMreq(address uint16, time int)
	ContendReadNoMreq_loop(address uint16, time int, count uint)

	ContendWriteNoMreq(address uint16, time int)
	ContendWriteNoMreq_loop(address uint16, time int, count uint)

	Read(address uint16) byte
	Write(address uint16, value byte, protectROM bool)

	// Data returns the memory content.
	Data() []byte
}

// MemoryReader is a simple interface that defines only a ReadByte
// method. It's used mainly by the disassembler.
type MemoryReader interface {
	ReadByte(address uint16) byte
}
