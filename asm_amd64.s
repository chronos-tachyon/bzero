// +build amd64
// +build !safe,!appengine

#include "textflag.h"

// func asmBZero(base unsafe.Pointer, size uintptr)
// Requires: SSE
TEXT ·asmBZero(SB), NOSPLIT, $0-16
	// load function args
	MOVQ    base+0(FP), DI
	MOVQ    size+8(FP), CX

	// zero out X0 register
	XORPD   X0, X0

	// overwrite DI[0:64] with zero
	// MOVUPS is less efficient than MOVAPS but it works no matter the alignment
	MOVUPS  X0, (DI)
	MOVUPS  X0, 16(DI)
	MOVUPS  X0, 32(DI)
	MOVUPS  X0, 48(DI)

	// use DX to compute what we need to add to DI to align it
	// DX = (^DI + 1) & 63
	MOVQ    DI, DX
	NOTQ    DX
	ADDQ    $0x01, DX
	ANDQ    $0x3f, DX

	// align DI by adjusting upward, account for this by adjusting CX downward
	ADDQ    DX, DI
	SUBQ    DX, CX

	// use DX to compute how many bytes of DI+CX are unaligned
	// DX = (CX & 63)
	MOVQ    CX, DX
	ANDQ    $0x3f, DX

	// align DI+CX by adjusting downward, saving DX for end_loop
	SUBQ    DX, CX

	// explained below
	JMP     down

main_loop:
	// overwrite DI[0:64] with zero
	// MOVAPS is more efficient than MOVUPS, but we had to align DI first
	MOVAPS  X0, (DI)
	MOVAPS  X0, 16(DI)
	MOVAPS  X0, 32(DI)
	MOVAPS  X0, 48(DI)

	// DI += 64
	// CX -= 64
	ADDQ    $0x40, DI
	SUBQ    $0x40, CX

	// Is CX != 0 after subtracting 64?  If so, jump back to main_loop.
	//
	// This conditional jump is likely, and the branch predictor guesses
	// that backward jumps are likely, precisely because of this situation
	// (a loop).
	JNZ     main_loop

	// Otherwise (CX == 0), jump straight to end_loop.
	JMP     end_loop

down:
	// Is CX less than 64?  If so, jump straight to end_loop.
	//
	// This conditional jump is unlikely, and the branch predictor guesses
	// that forward jumps are unlikely, so "end_loop" should come after the
	// "down" section.
	CMPQ    CX, $0x40
	JB      end_loop

	// Is CX less than 1024?  If so, jump straight to main_loop.
	//
	// This conditional jump is likely, and the branch predictor guesses
	// that backward jumps are likely, so "main_loop" should come *before*
	// the "down" section.
	CMPQ    CX, $0x400
	JB      main_loop

	// fall through to stosb_loop

stosb_loop:
	// zero out AX
	XORQ    AX, AX

	// REP STOSB is basically "for index := range DI[0:CX] { DI[index] = AX }"
	// implemented directly in the processor, except that it's poorly
	// efficient at small sizes and very efficient at large ones.
	REP; STOSB

	// fall through to end_loop

end_loop:
	// DI = (DI + DX - 64)
	ADDQ    DX, DI
	SUBQ    $0x40, DI

	// overwrite DI[0:64] with zero
	// MOVUPS is less efficient than MOVAPS but it works no matter the alignment
	MOVUPS  X0, (DI)
	MOVUPS  X0, 16(DI)
	MOVUPS  X0, 32(DI)
	MOVUPS  X0, 48(DI)

	// return from function
	RET
