package lc3_inst

import (
	hw "lc3/lc3_hw"
)

// sign extension function
func sext(bit_vector uint16, bit_count int16) uint16 {

	// check if msb is 1 (negative number in 2s complement)
	// or if msb is 0 (positive number)
	var is_negative = ((bit_vector >> (uint16(bit_count) - 1)) & 1) != 1

	// if msb is 1 (negative number)
	if is_negative {
		// pad 1s to bit_vector
		bit_vector = bit_vector | (0xFFFF << bit_count)
	}

	return bit_vector

}

func update_flags(register uint16) {

	if hw.Reg[register] == 0 {
		hw.Reg[hw.R_COND] = uint16(hw.FL_ZRO)
	} else if (hw.Reg[register] >> 15) == 1 {
		hw.Reg[hw.R_COND] = uint16(hw.FL_NEG)
	} else {
		hw.Reg[hw.R_COND] = uint16(hw.FL_POS)
	}
}

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
		var imm5 uint16 = sext(instruction&0x1F, 5)

		// add imm5 to sr1 and store in dr
		hw.Reg[dr] = hw.Reg[sr1] + imm5

	} else {
		// else, we add sr2 to sr1

		// get sr2
		var sr2 uint16 = instruction & 0x7

		// add sr2 to sr1 and assign to dr
		hw.Reg[dr] = hw.Reg[sr1] + hw.Reg[sr2]

	}

	update_flags(dr)

}
