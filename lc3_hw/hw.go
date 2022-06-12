package lc3_hw

// maximum addressable memory
const MEMORY_MAX int = 1 << 16

// memory
var Memory [MEMORY_MAX]uint16

// register type
type Register uint16

// registers
const (
	R_R0 Register = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC    // program counter
	R_COND  // condition flags
	R_COUNT // register count
)

// register storage
var Reg [R_COUNT]uint16

// opcode type
type Opcode uint16

// opcodes
const (
	OP_BR   Opcode = iota // branch
	OP_ADD                // add
	OP_LD                 // load
	OP_ST                 // store
	OP_JSR                // jump register
	OP_AND                // bitwise and
	OP_LDR                // load register
	OP_STR                // store register
	OP_RTI                // unused
	OP_NOT                // bitwise not
	OP_LDI                // load indirect
	OP_STI                // store indirect
	OP_JMP                // jump
	OP_RES                // reserved (unused)
	OP_LEA                // load effective address
	OP_TRAP               // execute trap
)

// condition flag type
type ConditionFlag uint16

// condition flags
const (
	FL_POS ConditionFlag = 1 << 0
	FL_ZRO ConditionFlag = 1 << 1
	FL_NEG ConditionFlag = 1 << 2
)

// set the CPU starting position to
// memory location 0x3000 as the default
const CPU_START = 0x3000
