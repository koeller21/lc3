package lc3_utils

import (
	hw "lc3/lc3_hw"
)

func Mem_write() {

}

func Mem_read(address uint16) uint16 {
	return 12
}

// sign extension function
func Sext(bit_vector uint16, bit_count int16) uint16 {

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

func Update_flags(register uint16) {

	if hw.Reg[register] == 0 {
		hw.Reg[hw.R_COND] = uint16(hw.FL_ZRO)
	} else if (hw.Reg[register] >> 15) == 1 {
		hw.Reg[hw.R_COND] = uint16(hw.FL_NEG)
	} else {
		hw.Reg[hw.R_COND] = uint16(hw.FL_POS)
	}
}
