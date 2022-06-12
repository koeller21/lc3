package main

import (
	"fmt"
	hw "lc3/lc3_hw"
	instruction "lc3/lc3_inst"
	utils "lc3/lc3_utils"
	"os"
)

func read_image(image_name string) bool {
	return true
}

func load_vm_images() {

	// provide at least one argument
	// which is a vm image to execute
	if !(len(os.Args) > 2) {
		fmt.Println("go run lc3.go [image-file1] ...")
	}

	// read all images provided
	for i := 0; i < len(os.Args); i++ {
		if !read_image(os.Args[i]) {
			fmt.Println("Failed to load image: %s", os.Args[i])
			os.Exit(1)
		}
	}

}

func cpu() {

	// exactly one condition flag should be set at any
	// given time, so just set the Z flag
	hw.Reg[hw.R_COND] = uint16(hw.FL_ZRO)

	hw.Reg[hw.R_PC] = hw.CPU_START

	var running bool = true

	for running {

		// FETCH -> DECODE -> EXECUTE

		// --- FETCH
		// get instruction at memory address pointed to by PC
		var instr uint16 = utils.Mem_read(hw.Reg[hw.R_PC])
		// increase PC
		hw.Reg[hw.R_PC] = hw.Reg[hw.R_PC] + 1

		// --- DECODE
		// get current operation opcode
		var op hw.Opcode = hw.Opcode(instr >> 12)

		// --- EXECUTE
		switch op {
		case hw.OP_ADD:
			instruction.Add(instr)
			break
		case hw.OP_AND:
			break
		case hw.OP_NOT:
			break
		case hw.OP_BR:
			break
		case hw.OP_JMP:
			break
		case hw.OP_JSR:
			break
		case hw.OP_LD:
			break
		case hw.OP_LDI:
			break
		case hw.OP_LDR:
			break
		case hw.OP_LEA:
			break
		case hw.OP_ST:
			break
		case hw.OP_STI:
			break
		case hw.OP_STR:
			break
		case hw.OP_TRAP:
			break
		case hw.OP_RES:
		case hw.OP_RTI:
		default: // bad opcode
			break
		}
	}

}

func main() {

	//cpu()
	//

	b := 0b0000100000000000
	s := fmt.Sprintf("%b", b)
	fmt.Println(s)
	instruction.Jsr(uint16(b))

}
