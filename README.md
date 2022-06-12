### A golang-based implementation of the LC3 Architecture

## HW Architecture

- RAM: 65,536 @ 16 bit
- CPU: non-pipelined, register-based von Neumann
- 8 General Purpose registers @ 16 bit

## ISA

-	OP_BR   // branch
-	OP_ADD  // add
-	OP_LD   // load
-	OP_ST   // store
-	OP_JSR  // jump register
-	OP_AND  // bitwise and
-	OP_LDR  // load register
-	OP_STR  // store register
-	OP_RTI  // not used
-	OP_NOT  // bitwise not
-	OP_LDI  // load indirect
-	OP_STI  // store indirect
-	OP_JMP  // jump
-	OP_RES  // reserved (unused)
-	OP_LEA  // load effective address
-	OP_TRAP // execute trap

