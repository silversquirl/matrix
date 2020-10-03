// +build amd64,!nosse

// func ·Mul4(a, b mgl.Mat4) mgl.Mat4
TEXT ·Mul4(SB),$0-192
	//    LOAD A    //
	// Load rows
	MOVUPS a_0+0(FP), X0 
	MOVUPS a_4+16(FP), X1 
	MOVUPS a_8+32(FP), X2 
	MOVUPS a_12+48(FP), X3 

	//    TRANSPOSE    //
	// Mix rows
	VSHUFPS $0x44, X1, X0, X4
	VSHUFPS $0x44, X3, X2, X5
	VSHUFPS $0xee, X1, X0, X6
	VSHUFPS $0xee, X3, X2, X7

	// Create columns
	VSHUFPS $0x88, X5, X4, X0
	VSHUFPS $0xdd, X5, X4, X1
	VSHUFPS $0x88, X7, X6, X2
	VSHUFPS $0xdd, X7, X6, X3

	//    LOAD B    //
	MOVUPS b_0+64(FP), X4 
	MOVUPS b_4+80(FP), X5 
	MOVUPS b_8+96(FP), X6 
	MOVUPS b_12+112(FP), X7 

	//    MULTIPLY    //
#define CALC_ROW(row) \
	VDPPS $0xf1, X0, row, X8 \
	VDPPS $0xf1, X1, row, X9 \
	VDPPS $0xf1, X2, row, X10 \
	VDPPS $0xf1, X3, row, X11 \
	SHUFPS $0x00, X9, X8 \
	SHUFPS $0x00, X11, X10 \
	SHUFPS $0x88, X10, X8
	
	CALC_ROW(X4)
	MOVUPS X8, ret_0+128(FP)
	CALC_ROW(X5)
	MOVUPS X8, ret_4+144(FP)
	CALC_ROW(X6)
	MOVUPS X8, ret_8+160(FP)
	CALC_ROW(X7)
	MOVUPS X8, ret_12+176(FP)

	RET
