package lc3_inst

import (
	hw "lc3/lc3_hw"
	utils "lc3/lc3_utils"
)

// Add - adds contents of <R1> or imm5 to <R0> and stores in <R2>
// ADD R2 R0 R1
// |15 14 13 12 . 11 10 9 . 8 7 6 . 5 . 4 3 . 2 1 0|
//     opcode		dr		 sr1   mode pad    sr2
// ADD R2 R0 2
// |15 14 13 12 . 11 10 9 . 8 7 6 . 5 . 4 3 2 1 0|
//     opcode		dr		 sr1   mode   imm5
func Add(instruction uint16) {

	// get destination register dr
	var dr uint16 = (instruction >> 9) & 0x7

	// get source register 1, sr1
	var sr1 uint16 = (instruction >> 6) & 0x7

	// get bit 5 flag indicating
	// register or immediate mode
	var imm_flag uint16 = (instruction >> 5) & 0x1

	// if imm_flag is 1, then we're
	// in immediate mode, therefore
	// we add the imm5 bitfield directly to sr1
	if imm_flag == 0x1 {

		// get 5-bit immediate value
		// by and-ing with 31 (10)
		var imm5 uint16 = utils.Sext(instruction&0x1F, 5)

		// add imm5 to sr1 and store in dr
		hw.Reg[dr] = hw.Reg[sr1] + imm5

	} else {
		// else, we add sr2 to sr1

		// get sr2
		var sr2 uint16 = instruction & 0x7

		// add sr2 to sr1 and assign to dr
		hw.Reg[dr] = hw.Reg[sr1] + hw.Reg[sr2]

	}

	utils.Update_flags(dr)

}

// load indirect - loads <mem_location> to register
// LDI R0
// |15 14 13 12 . 11 10 9 . 8 7 6 5 4 3 2 1 0|
//     opcode		dr			pc offset
func Ldi(instruction uint16) {

	// get destination register
	var dr uint16 = (instruction >> 9) & 0x7

	// get 9-bit pc offset
	var pc_offset uint16 = utils.Sext(instruction&0x1FF, 9)

	// add pc offset to current PC
	mem_location_value := utils.Mem_read(hw.Reg[hw.R_PC] + pc_offset)

	// get value at mem_location_value and assign to dr
	hw.Reg[dr] = utils.Mem_read(mem_location_value)

	utils.Update_flags(dr)

}

func And(instruction uint16) {

	// get destination register dr
	var dr uint16 = (instruction >> 9) & 0x7

	// get source register 1, sr1
	var sr1 uint16 = (instruction >> 6) & 0x7

	// get bit 5 flag indicating
	// register or immediate mode
	var imm_flag uint16 = (instruction >> 5) & 0x1

	// if imm_flag is 1, then we're
	// in immediate mode, therefore
	// we and the imm5 bitfield directly to sr1
	if imm_flag == 0x1 {

		// get 5-bit immediate value
		// by and-ing with 31 (10)
		var imm5 uint16 = utils.Sext(instruction&0x1F, 5)

		// add imm5 to sr1 and store in dr
		hw.Reg[dr] = hw.Reg[sr1] & imm5

	} else {
		// else, we add sr2 to sr1

		// get sr2
		var sr2 uint16 = instruction & 0x7

		// add sr2 to sr1 and assign to dr
		hw.Reg[dr] = hw.Reg[sr1] & hw.Reg[sr2]

	}

}

func Br(instruction uint16) {

	var cond_flag uint16 = (instruction >> 9) & 0x7

	if (cond_flag & hw.Reg[hw.R_COND]) == 1 {
		hw.Reg[hw.R_PC] = hw.Reg[hw.R_PC] + utils.Sext(instruction&0x1FF, 9)
	}

}

func Jmp(instruction uint16) {

	var base_r uint16 = (instruction >> 6) & 0x7
	hw.Reg[hw.R_PC] = hw.Reg[base_r]

}

func Jsr(instruction uint16) {

	// get PC
	var tmp_pc uint16 = hw.Reg[hw.R_PC]

	// get jsr flag
	var flag uint16 = (instruction >> 11) & 0x1

	// JSR
	if flag == 0x1 {

		hw.Reg[hw.R_PC] = tmp_pc + utils.Sext(instruction&0x7FF, 11)

	} else {
		// JSRR

		var base_r uint16 = (instruction >> 6) & 0x7

		hw.Reg[hw.R_PC] = hw.Reg[base_r]
	}

}

func Ld(instruction uint16) {

	// get direct register
	var dr uint16 = (instruction >> 9) & 0x7

	// get 9-bit PC offset
	var pc_offset uint16 = instruction & 0x1FF

	hw.Reg[dr] = utils.Mem_read(hw.Reg[hw.R_PC] + pc_offset)

	utils.Update_flags(dr)
}

func Ldr(instruction uint16) {

	// get direct register
	var dr uint16 = (instruction >> 9) & 0x7

	// get base register
	var base_r uint16 = (instruction >> 6) & 0x7

	// get 6-bit PC offset
	var pc_offset uint16 = instruction & 0x3F

	var offset_total uint16 = base_r + utils.Sext(pc_offset, 6)
	hw.Reg[dr] = utils.Mem_read(offset_total)

	utils.Update_flags(dr)
}

func Lea(instruction uint16) {

	// get direct register
	var dr uint16 = (instruction >> 9) & 0x7

	// get 9-bit PC offset
	var pc_offset uint16 = instruction & 0x1FF

	hw.Reg[dr] = hw.Reg[hw.R_PC] + utils.Sext(pc_offset, 9)

}

func Not(instruction uint16) {

	// get destination register
	var dr uint16 = (instruction >> 9) & 0x7

	// get source register
	var sr uint16 = (instruction >> 6) & 0x7

	hw.Reg[dr] = ^hw.Reg[sr]

	utils.Update_flags(dr)

}

func St(instruction uint16) {

	// get source register
	var sr uint16 = (instruction >> 9) & 0x7

	// get 9-bit PC offset
	var pc_offset uint16 = utils.Sext(instruction&0x1FF, 9)

	utils.Mem_write(hw.Reg[hw.R_PC]+pc_offset, hw.Reg[sr])

}

func Sti(instruction uint16) {

	// get source register
	var sr uint16 = (instruction >> 9) & 0x7

	// get 9-bit PC offset
	var pc_offset uint16 = utils.Sext(instruction&0x1FF, 9)

	utils.Mem_write(utils.Mem_read(hw.Reg[hw.R_PC]+pc_offset), hw.Reg[sr])
}

func Str(instruction uint16) {

	// get source register
	var sr uint16 = (instruction >> 9) & 0x7

	// get base register
	var base_r uint16 = (instruction >> 6) & 0x7

	// get 6-bit PC offset
	var pc_offset uint16 = utils.Sext(instruction&0x3F, 6)

	utils.Mem_write(hw.Reg[base_r]+pc_offset, hw.Reg[sr])

}
