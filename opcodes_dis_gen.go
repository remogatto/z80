/*

Copyright (c) 2012 Andrea Fazzi

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.
8	 
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*/

//
// Automatically generated file -- DO NOT EDIT
//

package z80

import (
	"fmt"
)

func initOpcodesDis() {
	// BEGIN of non shifted opcodes
	/* NOP */
	OpcodesDisMap[0x00] = dis_instr__NOP
	/* LD BC,nnnn */
	OpcodesDisMap[0x01] = dis_instr__LD_BC_NNNN
	/* LD (BC),A */
	OpcodesDisMap[0x02] = dis_instr__LD_iBC_A
	/* INC BC */
	OpcodesDisMap[0x03] = dis_instr__INC_BC
	/* INC B */
	OpcodesDisMap[0x04] = dis_instr__INC_B
	/* DEC B */
	OpcodesDisMap[0x05] = dis_instr__DEC_B
	/* LD B,nn */
	OpcodesDisMap[0x06] = dis_instr__LD_B_NN
	/* RLCA */
	OpcodesDisMap[0x07] = dis_instr__RLCA
	/* EX AF,AF' */
	OpcodesDisMap[0x08] = dis_instr__EX_AF_AF
	/* ADD HL,BC */
	OpcodesDisMap[0x09] = dis_instr__ADD_HL_BC
	/* LD A,(BC) */
	OpcodesDisMap[0x0a] = dis_instr__LD_A_iBC
	/* DEC BC */
	OpcodesDisMap[0x0b] = dis_instr__DEC_BC
	/* INC C */
	OpcodesDisMap[0x0c] = dis_instr__INC_C
	/* DEC C */
	OpcodesDisMap[0x0d] = dis_instr__DEC_C
	/* LD C,nn */
	OpcodesDisMap[0x0e] = dis_instr__LD_C_NN
	/* RRCA */
	OpcodesDisMap[0x0f] = dis_instr__RRCA
	/* DJNZ offset */
	OpcodesDisMap[0x10] = dis_instr__DJNZ_OFFSET
	/* LD DE,nnnn */
	OpcodesDisMap[0x11] = dis_instr__LD_DE_NNNN
	/* LD (DE),A */
	OpcodesDisMap[0x12] = dis_instr__LD_iDE_A
	/* INC DE */
	OpcodesDisMap[0x13] = dis_instr__INC_DE
	/* INC D */
	OpcodesDisMap[0x14] = dis_instr__INC_D
	/* DEC D */
	OpcodesDisMap[0x15] = dis_instr__DEC_D
	/* LD D,nn */
	OpcodesDisMap[0x16] = dis_instr__LD_D_NN
	/* RLA */
	OpcodesDisMap[0x17] = dis_instr__RLA
	/* JR offset */
	OpcodesDisMap[0x18] = dis_instr__JR_OFFSET
	/* ADD HL,DE */
	OpcodesDisMap[0x19] = dis_instr__ADD_HL_DE
	/* LD A,(DE) */
	OpcodesDisMap[0x1a] = dis_instr__LD_A_iDE
	/* DEC DE */
	OpcodesDisMap[0x1b] = dis_instr__DEC_DE
	/* INC E */
	OpcodesDisMap[0x1c] = dis_instr__INC_E
	/* DEC E */
	OpcodesDisMap[0x1d] = dis_instr__DEC_E
	/* LD E,nn */
	OpcodesDisMap[0x1e] = dis_instr__LD_E_NN
	/* RRA */
	OpcodesDisMap[0x1f] = dis_instr__RRA
	/* JR NZ,offset */
	OpcodesDisMap[0x20] = dis_instr__JR_NZ_OFFSET
	/* LD HL,nnnn */
	OpcodesDisMap[0x21] = dis_instr__LD_HL_NNNN
	/* LD (nnnn),HL */
	OpcodesDisMap[0x22] = dis_instr__LD_iNNNN_HL
	/* INC HL */
	OpcodesDisMap[0x23] = dis_instr__INC_HL
	/* INC H */
	OpcodesDisMap[0x24] = dis_instr__INC_H
	/* DEC H */
	OpcodesDisMap[0x25] = dis_instr__DEC_H
	/* LD H,nn */
	OpcodesDisMap[0x26] = dis_instr__LD_H_NN
	/* DAA */
	OpcodesDisMap[0x27] = dis_instr__DAA
	/* JR Z,offset */
	OpcodesDisMap[0x28] = dis_instr__JR_Z_OFFSET
	/* ADD HL,HL */
	OpcodesDisMap[0x29] = dis_instr__ADD_HL_HL
	/* LD HL,(nnnn) */
	OpcodesDisMap[0x2a] = dis_instr__LD_HL_iNNNN
	/* DEC HL */
	OpcodesDisMap[0x2b] = dis_instr__DEC_HL
	/* INC L */
	OpcodesDisMap[0x2c] = dis_instr__INC_L
	/* DEC L */
	OpcodesDisMap[0x2d] = dis_instr__DEC_L
	/* LD L,nn */
	OpcodesDisMap[0x2e] = dis_instr__LD_L_NN
	/* CPL */
	OpcodesDisMap[0x2f] = dis_instr__CPL
	/* JR NC,offset */
	OpcodesDisMap[0x30] = dis_instr__JR_NC_OFFSET
	/* LD SP,nnnn */
	OpcodesDisMap[0x31] = dis_instr__LD_SP_NNNN
	/* LD (nnnn),A */
	OpcodesDisMap[0x32] = dis_instr__LD_iNNNN_A
	/* INC SP */
	OpcodesDisMap[0x33] = dis_instr__INC_SP
	/* INC (HL) */
	OpcodesDisMap[0x34] = dis_instr__INC_iHL
	/* DEC (HL) */
	OpcodesDisMap[0x35] = dis_instr__DEC_iHL
	/* LD (HL),nn */
	OpcodesDisMap[0x36] = dis_instr__LD_iHL_NN
	/* SCF */
	OpcodesDisMap[0x37] = dis_instr__SCF
	/* JR C,offset */
	OpcodesDisMap[0x38] = dis_instr__JR_C_OFFSET
	/* ADD HL,SP */
	OpcodesDisMap[0x39] = dis_instr__ADD_HL_SP
	/* LD A,(nnnn) */
	OpcodesDisMap[0x3a] = dis_instr__LD_A_iNNNN
	/* DEC SP */
	OpcodesDisMap[0x3b] = dis_instr__DEC_SP
	/* INC A */
	OpcodesDisMap[0x3c] = dis_instr__INC_A
	/* DEC A */
	OpcodesDisMap[0x3d] = dis_instr__DEC_A
	/* LD A,nn */
	OpcodesDisMap[0x3e] = dis_instr__LD_A_NN
	/* CCF */
	OpcodesDisMap[0x3f] = dis_instr__CCF
	/* LD B,B */
	OpcodesDisMap[0x40] = dis_instr__LD_B_B
	/* LD B,C */
	OpcodesDisMap[0x41] = dis_instr__LD_B_C
	/* LD B,D */
	OpcodesDisMap[0x42] = dis_instr__LD_B_D
	/* LD B,E */
	OpcodesDisMap[0x43] = dis_instr__LD_B_E
	/* LD B,H */
	OpcodesDisMap[0x44] = dis_instr__LD_B_H
	/* LD B,L */
	OpcodesDisMap[0x45] = dis_instr__LD_B_L
	/* LD B,(HL) */
	OpcodesDisMap[0x46] = dis_instr__LD_B_iHL
	/* LD B,A */
	OpcodesDisMap[0x47] = dis_instr__LD_B_A
	/* LD C,B */
	OpcodesDisMap[0x48] = dis_instr__LD_C_B
	/* LD C,C */
	OpcodesDisMap[0x49] = dis_instr__LD_C_C
	/* LD C,D */
	OpcodesDisMap[0x4a] = dis_instr__LD_C_D
	/* LD C,E */
	OpcodesDisMap[0x4b] = dis_instr__LD_C_E
	/* LD C,H */
	OpcodesDisMap[0x4c] = dis_instr__LD_C_H
	/* LD C,L */
	OpcodesDisMap[0x4d] = dis_instr__LD_C_L
	/* LD C,(HL) */
	OpcodesDisMap[0x4e] = dis_instr__LD_C_iHL
	/* LD C,A */
	OpcodesDisMap[0x4f] = dis_instr__LD_C_A
	/* LD D,B */
	OpcodesDisMap[0x50] = dis_instr__LD_D_B
	/* LD D,C */
	OpcodesDisMap[0x51] = dis_instr__LD_D_C
	/* LD D,D */
	OpcodesDisMap[0x52] = dis_instr__LD_D_D
	/* LD D,E */
	OpcodesDisMap[0x53] = dis_instr__LD_D_E
	/* LD D,H */
	OpcodesDisMap[0x54] = dis_instr__LD_D_H
	/* LD D,L */
	OpcodesDisMap[0x55] = dis_instr__LD_D_L
	/* LD D,(HL) */
	OpcodesDisMap[0x56] = dis_instr__LD_D_iHL
	/* LD D,A */
	OpcodesDisMap[0x57] = dis_instr__LD_D_A
	/* LD E,B */
	OpcodesDisMap[0x58] = dis_instr__LD_E_B
	/* LD E,C */
	OpcodesDisMap[0x59] = dis_instr__LD_E_C
	/* LD E,D */
	OpcodesDisMap[0x5a] = dis_instr__LD_E_D
	/* LD E,E */
	OpcodesDisMap[0x5b] = dis_instr__LD_E_E
	/* LD E,H */
	OpcodesDisMap[0x5c] = dis_instr__LD_E_H
	/* LD E,L */
	OpcodesDisMap[0x5d] = dis_instr__LD_E_L
	/* LD E,(HL) */
	OpcodesDisMap[0x5e] = dis_instr__LD_E_iHL
	/* LD E,A */
	OpcodesDisMap[0x5f] = dis_instr__LD_E_A
	/* LD H,B */
	OpcodesDisMap[0x60] = dis_instr__LD_H_B
	/* LD H,C */
	OpcodesDisMap[0x61] = dis_instr__LD_H_C
	/* LD H,D */
	OpcodesDisMap[0x62] = dis_instr__LD_H_D
	/* LD H,E */
	OpcodesDisMap[0x63] = dis_instr__LD_H_E
	/* LD H,H */
	OpcodesDisMap[0x64] = dis_instr__LD_H_H
	/* LD H,L */
	OpcodesDisMap[0x65] = dis_instr__LD_H_L
	/* LD H,(HL) */
	OpcodesDisMap[0x66] = dis_instr__LD_H_iHL
	/* LD H,A */
	OpcodesDisMap[0x67] = dis_instr__LD_H_A
	/* LD L,B */
	OpcodesDisMap[0x68] = dis_instr__LD_L_B
	/* LD L,C */
	OpcodesDisMap[0x69] = dis_instr__LD_L_C
	/* LD L,D */
	OpcodesDisMap[0x6a] = dis_instr__LD_L_D
	/* LD L,E */
	OpcodesDisMap[0x6b] = dis_instr__LD_L_E
	/* LD L,H */
	OpcodesDisMap[0x6c] = dis_instr__LD_L_H
	/* LD L,L */
	OpcodesDisMap[0x6d] = dis_instr__LD_L_L
	/* LD L,(HL) */
	OpcodesDisMap[0x6e] = dis_instr__LD_L_iHL
	/* LD L,A */
	OpcodesDisMap[0x6f] = dis_instr__LD_L_A
	/* LD (HL),B */
	OpcodesDisMap[0x70] = dis_instr__LD_iHL_B
	/* LD (HL),C */
	OpcodesDisMap[0x71] = dis_instr__LD_iHL_C
	/* LD (HL),D */
	OpcodesDisMap[0x72] = dis_instr__LD_iHL_D
	/* LD (HL),E */
	OpcodesDisMap[0x73] = dis_instr__LD_iHL_E
	/* LD (HL),H */
	OpcodesDisMap[0x74] = dis_instr__LD_iHL_H
	/* LD (HL),L */
	OpcodesDisMap[0x75] = dis_instr__LD_iHL_L
	/* HALT */
	OpcodesDisMap[0x76] = dis_instr__HALT
	/* LD (HL),A */
	OpcodesDisMap[0x77] = dis_instr__LD_iHL_A
	/* LD A,B */
	OpcodesDisMap[0x78] = dis_instr__LD_A_B
	/* LD A,C */
	OpcodesDisMap[0x79] = dis_instr__LD_A_C
	/* LD A,D */
	OpcodesDisMap[0x7a] = dis_instr__LD_A_D
	/* LD A,E */
	OpcodesDisMap[0x7b] = dis_instr__LD_A_E
	/* LD A,H */
	OpcodesDisMap[0x7c] = dis_instr__LD_A_H
	/* LD A,L */
	OpcodesDisMap[0x7d] = dis_instr__LD_A_L
	/* LD A,(HL) */
	OpcodesDisMap[0x7e] = dis_instr__LD_A_iHL
	/* LD A,A */
	OpcodesDisMap[0x7f] = dis_instr__LD_A_A
	/* ADD A,B */
	OpcodesDisMap[0x80] = dis_instr__ADD_A_B
	/* ADD A,C */
	OpcodesDisMap[0x81] = dis_instr__ADD_A_C
	/* ADD A,D */
	OpcodesDisMap[0x82] = dis_instr__ADD_A_D
	/* ADD A,E */
	OpcodesDisMap[0x83] = dis_instr__ADD_A_E
	/* ADD A,H */
	OpcodesDisMap[0x84] = dis_instr__ADD_A_H
	/* ADD A,L */
	OpcodesDisMap[0x85] = dis_instr__ADD_A_L
	/* ADD A,(HL) */
	OpcodesDisMap[0x86] = dis_instr__ADD_A_iHL
	/* ADD A,A */
	OpcodesDisMap[0x87] = dis_instr__ADD_A_A
	/* ADC A,B */
	OpcodesDisMap[0x88] = dis_instr__ADC_A_B
	/* ADC A,C */
	OpcodesDisMap[0x89] = dis_instr__ADC_A_C
	/* ADC A,D */
	OpcodesDisMap[0x8a] = dis_instr__ADC_A_D
	/* ADC A,E */
	OpcodesDisMap[0x8b] = dis_instr__ADC_A_E
	/* ADC A,H */
	OpcodesDisMap[0x8c] = dis_instr__ADC_A_H
	/* ADC A,L */
	OpcodesDisMap[0x8d] = dis_instr__ADC_A_L
	/* ADC A,(HL) */
	OpcodesDisMap[0x8e] = dis_instr__ADC_A_iHL
	/* ADC A,A */
	OpcodesDisMap[0x8f] = dis_instr__ADC_A_A
	/* SUB A,B */
	OpcodesDisMap[0x90] = dis_instr__SUB_A_B
	/* SUB A,C */
	OpcodesDisMap[0x91] = dis_instr__SUB_A_C
	/* SUB A,D */
	OpcodesDisMap[0x92] = dis_instr__SUB_A_D
	/* SUB A,E */
	OpcodesDisMap[0x93] = dis_instr__SUB_A_E
	/* SUB A,H */
	OpcodesDisMap[0x94] = dis_instr__SUB_A_H
	/* SUB A,L */
	OpcodesDisMap[0x95] = dis_instr__SUB_A_L
	/* SUB A,(HL) */
	OpcodesDisMap[0x96] = dis_instr__SUB_A_iHL
	/* SUB A,A */
	OpcodesDisMap[0x97] = dis_instr__SUB_A_A
	/* SBC A,B */
	OpcodesDisMap[0x98] = dis_instr__SBC_A_B
	/* SBC A,C */
	OpcodesDisMap[0x99] = dis_instr__SBC_A_C
	/* SBC A,D */
	OpcodesDisMap[0x9a] = dis_instr__SBC_A_D
	/* SBC A,E */
	OpcodesDisMap[0x9b] = dis_instr__SBC_A_E
	/* SBC A,H */
	OpcodesDisMap[0x9c] = dis_instr__SBC_A_H
	/* SBC A,L */
	OpcodesDisMap[0x9d] = dis_instr__SBC_A_L
	/* SBC A,(HL) */
	OpcodesDisMap[0x9e] = dis_instr__SBC_A_iHL
	/* SBC A,A */
	OpcodesDisMap[0x9f] = dis_instr__SBC_A_A
	/* AND A,B */
	OpcodesDisMap[0xa0] = dis_instr__AND_A_B
	/* AND A,C */
	OpcodesDisMap[0xa1] = dis_instr__AND_A_C
	/* AND A,D */
	OpcodesDisMap[0xa2] = dis_instr__AND_A_D
	/* AND A,E */
	OpcodesDisMap[0xa3] = dis_instr__AND_A_E
	/* AND A,H */
	OpcodesDisMap[0xa4] = dis_instr__AND_A_H
	/* AND A,L */
	OpcodesDisMap[0xa5] = dis_instr__AND_A_L
	/* AND A,(HL) */
	OpcodesDisMap[0xa6] = dis_instr__AND_A_iHL
	/* AND A,A */
	OpcodesDisMap[0xa7] = dis_instr__AND_A_A
	/* XOR A,B */
	OpcodesDisMap[0xa8] = dis_instr__XOR_A_B
	/* XOR A,C */
	OpcodesDisMap[0xa9] = dis_instr__XOR_A_C
	/* XOR A,D */
	OpcodesDisMap[0xaa] = dis_instr__XOR_A_D
	/* XOR A,E */
	OpcodesDisMap[0xab] = dis_instr__XOR_A_E
	/* XOR A,H */
	OpcodesDisMap[0xac] = dis_instr__XOR_A_H
	/* XOR A,L */
	OpcodesDisMap[0xad] = dis_instr__XOR_A_L
	/* XOR A,(HL) */
	OpcodesDisMap[0xae] = dis_instr__XOR_A_iHL
	/* XOR A,A */
	OpcodesDisMap[0xaf] = dis_instr__XOR_A_A
	/* OR A,B */
	OpcodesDisMap[0xb0] = dis_instr__OR_A_B
	/* OR A,C */
	OpcodesDisMap[0xb1] = dis_instr__OR_A_C
	/* OR A,D */
	OpcodesDisMap[0xb2] = dis_instr__OR_A_D
	/* OR A,E */
	OpcodesDisMap[0xb3] = dis_instr__OR_A_E
	/* OR A,H */
	OpcodesDisMap[0xb4] = dis_instr__OR_A_H
	/* OR A,L */
	OpcodesDisMap[0xb5] = dis_instr__OR_A_L
	/* OR A,(HL) */
	OpcodesDisMap[0xb6] = dis_instr__OR_A_iHL
	/* OR A,A */
	OpcodesDisMap[0xb7] = dis_instr__OR_A_A
	/* CP B */
	OpcodesDisMap[0xb8] = dis_instr__CP_B
	/* CP C */
	OpcodesDisMap[0xb9] = dis_instr__CP_C
	/* CP D */
	OpcodesDisMap[0xba] = dis_instr__CP_D
	/* CP E */
	OpcodesDisMap[0xbb] = dis_instr__CP_E
	/* CP H */
	OpcodesDisMap[0xbc] = dis_instr__CP_H
	/* CP L */
	OpcodesDisMap[0xbd] = dis_instr__CP_L
	/* CP (HL) */
	OpcodesDisMap[0xbe] = dis_instr__CP_iHL
	/* CP A */
	OpcodesDisMap[0xbf] = dis_instr__CP_A
	/* RET NZ */
	OpcodesDisMap[0xc0] = dis_instr__RET_NZ
	/* POP BC */
	OpcodesDisMap[0xc1] = dis_instr__POP_BC
	/* JP NZ,nnnn */
	OpcodesDisMap[0xc2] = dis_instr__JP_NZ_NNNN
	/* JP nnnn */
	OpcodesDisMap[0xc3] = dis_instr__JP_NNNN
	/* CALL NZ,nnnn */
	OpcodesDisMap[0xc4] = dis_instr__CALL_NZ_NNNN
	/* PUSH BC */
	OpcodesDisMap[0xc5] = dis_instr__PUSH_BC
	/* ADD A,nn */
	OpcodesDisMap[0xc6] = dis_instr__ADD_A_NN
	/* RST 00 */
	OpcodesDisMap[0xc7] = dis_instr__RST_00
	/* RET Z */
	OpcodesDisMap[0xc8] = dis_instr__RET_Z
	/* RET */
	OpcodesDisMap[0xc9] = dis_instr__RET
	/* JP Z,nnnn */
	OpcodesDisMap[0xca] = dis_instr__JP_Z_NNNN
	/* shift CB */
	OpcodesDisMap[0xcb] = dis_instr__SHIFT_CB
	/* CALL Z,nnnn */
	OpcodesDisMap[0xcc] = dis_instr__CALL_Z_NNNN
	/* CALL nnnn */
	OpcodesDisMap[0xcd] = dis_instr__CALL_NNNN
	/* ADC A,nn */
	OpcodesDisMap[0xce] = dis_instr__ADC_A_NN
	/* RST 8 */
	OpcodesDisMap[0xcf] = dis_instr__RST_8
	/* RET NC */
	OpcodesDisMap[0xd0] = dis_instr__RET_NC
	/* POP DE */
	OpcodesDisMap[0xd1] = dis_instr__POP_DE
	/* JP NC,nnnn */
	OpcodesDisMap[0xd2] = dis_instr__JP_NC_NNNN
	/* OUT (nn),A */
	OpcodesDisMap[0xd3] = dis_instr__OUT_iNN_A
	/* CALL NC,nnnn */
	OpcodesDisMap[0xd4] = dis_instr__CALL_NC_NNNN
	/* PUSH DE */
	OpcodesDisMap[0xd5] = dis_instr__PUSH_DE
	/* SUB nn */
	OpcodesDisMap[0xd6] = dis_instr__SUB_NN
	/* RST 10 */
	OpcodesDisMap[0xd7] = dis_instr__RST_10
	/* RET C */
	OpcodesDisMap[0xd8] = dis_instr__RET_C
	/* EXX */
	OpcodesDisMap[0xd9] = dis_instr__EXX
	/* JP C,nnnn */
	OpcodesDisMap[0xda] = dis_instr__JP_C_NNNN
	/* IN A,(nn) */
	OpcodesDisMap[0xdb] = dis_instr__IN_A_iNN
	/* CALL C,nnnn */
	OpcodesDisMap[0xdc] = dis_instr__CALL_C_NNNN
	/* shift DD */
	OpcodesDisMap[0xdd] = dis_instr__SHIFT_DD
	/* SBC A,nn */
	OpcodesDisMap[0xde] = dis_instr__SBC_A_NN
	/* RST 18 */
	OpcodesDisMap[0xdf] = dis_instr__RST_18
	/* RET PO */
	OpcodesDisMap[0xe0] = dis_instr__RET_PO
	/* POP HL */
	OpcodesDisMap[0xe1] = dis_instr__POP_HL
	/* JP PO,nnnn */
	OpcodesDisMap[0xe2] = dis_instr__JP_PO_NNNN
	/* EX (SP),HL */
	OpcodesDisMap[0xe3] = dis_instr__EX_iSP_HL
	/* CALL PO,nnnn */
	OpcodesDisMap[0xe4] = dis_instr__CALL_PO_NNNN
	/* PUSH HL */
	OpcodesDisMap[0xe5] = dis_instr__PUSH_HL
	/* AND nn */
	OpcodesDisMap[0xe6] = dis_instr__AND_NN
	/* RST 20 */
	OpcodesDisMap[0xe7] = dis_instr__RST_20
	/* RET PE */
	OpcodesDisMap[0xe8] = dis_instr__RET_PE
	/* JP HL */
	OpcodesDisMap[0xe9] = dis_instr__JP_HL
	/* JP PE,nnnn */
	OpcodesDisMap[0xea] = dis_instr__JP_PE_NNNN
	/* EX DE,HL */
	OpcodesDisMap[0xeb] = dis_instr__EX_DE_HL
	/* CALL PE,nnnn */
	OpcodesDisMap[0xec] = dis_instr__CALL_PE_NNNN
	/* shift ED */
	OpcodesDisMap[0xed] = dis_instr__SHIFT_ED
	/* XOR A,nn */
	OpcodesDisMap[0xee] = dis_instr__XOR_A_NN
	/* RST 28 */
	OpcodesDisMap[0xef] = dis_instr__RST_28
	/* RET P */
	OpcodesDisMap[0xf0] = dis_instr__RET_P
	/* POP AF */
	OpcodesDisMap[0xf1] = dis_instr__POP_AF
	/* JP P,nnnn */
	OpcodesDisMap[0xf2] = dis_instr__JP_P_NNNN
	/* DI */
	OpcodesDisMap[0xf3] = dis_instr__DI
	/* CALL P,nnnn */
	OpcodesDisMap[0xf4] = dis_instr__CALL_P_NNNN
	/* PUSH AF */
	OpcodesDisMap[0xf5] = dis_instr__PUSH_AF
	/* OR nn */
	OpcodesDisMap[0xf6] = dis_instr__OR_NN
	/* RST 30 */
	OpcodesDisMap[0xf7] = dis_instr__RST_30
	/* RET M */
	OpcodesDisMap[0xf8] = dis_instr__RET_M
	/* LD SP,HL */
	OpcodesDisMap[0xf9] = dis_instr__LD_SP_HL
	/* JP M,nnnn */
	OpcodesDisMap[0xfa] = dis_instr__JP_M_NNNN
	/* EI */
	OpcodesDisMap[0xfb] = dis_instr__EI
	/* CALL M,nnnn */
	OpcodesDisMap[0xfc] = dis_instr__CALL_M_NNNN
	/* shift FD */
	OpcodesDisMap[0xfd] = dis_instr__SHIFT_FD
	/* CP nn */
	OpcodesDisMap[0xfe] = dis_instr__CP_NN
	/* RST 38 */
	OpcodesDisMap[0xff] = dis_instr__RST_38

	// END of non shifted opcodes

	// BEGIN of 0xcb shifted opcodes
	/* RLC B */
	OpcodesDisMap[SHIFT_0xCB+0x00] = dis_instrCB__RLC_B
	/* RLC C */
	OpcodesDisMap[SHIFT_0xCB+0x01] = dis_instrCB__RLC_C
	/* RLC D */
	OpcodesDisMap[SHIFT_0xCB+0x02] = dis_instrCB__RLC_D
	/* RLC E */
	OpcodesDisMap[SHIFT_0xCB+0x03] = dis_instrCB__RLC_E
	/* RLC H */
	OpcodesDisMap[SHIFT_0xCB+0x04] = dis_instrCB__RLC_H
	/* RLC L */
	OpcodesDisMap[SHIFT_0xCB+0x05] = dis_instrCB__RLC_L
	/* RLC (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x06] = dis_instrCB__RLC_iHL
	/* RLC A */
	OpcodesDisMap[SHIFT_0xCB+0x07] = dis_instrCB__RLC_A
	/* RRC B */
	OpcodesDisMap[SHIFT_0xCB+0x08] = dis_instrCB__RRC_B
	/* RRC C */
	OpcodesDisMap[SHIFT_0xCB+0x09] = dis_instrCB__RRC_C
	/* RRC D */
	OpcodesDisMap[SHIFT_0xCB+0x0a] = dis_instrCB__RRC_D
	/* RRC E */
	OpcodesDisMap[SHIFT_0xCB+0x0b] = dis_instrCB__RRC_E
	/* RRC H */
	OpcodesDisMap[SHIFT_0xCB+0x0c] = dis_instrCB__RRC_H
	/* RRC L */
	OpcodesDisMap[SHIFT_0xCB+0x0d] = dis_instrCB__RRC_L
	/* RRC (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x0e] = dis_instrCB__RRC_iHL
	/* RRC A */
	OpcodesDisMap[SHIFT_0xCB+0x0f] = dis_instrCB__RRC_A
	/* RL B */
	OpcodesDisMap[SHIFT_0xCB+0x10] = dis_instrCB__RL_B
	/* RL C */
	OpcodesDisMap[SHIFT_0xCB+0x11] = dis_instrCB__RL_C
	/* RL D */
	OpcodesDisMap[SHIFT_0xCB+0x12] = dis_instrCB__RL_D
	/* RL E */
	OpcodesDisMap[SHIFT_0xCB+0x13] = dis_instrCB__RL_E
	/* RL H */
	OpcodesDisMap[SHIFT_0xCB+0x14] = dis_instrCB__RL_H
	/* RL L */
	OpcodesDisMap[SHIFT_0xCB+0x15] = dis_instrCB__RL_L
	/* RL (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x16] = dis_instrCB__RL_iHL
	/* RL A */
	OpcodesDisMap[SHIFT_0xCB+0x17] = dis_instrCB__RL_A
	/* RR B */
	OpcodesDisMap[SHIFT_0xCB+0x18] = dis_instrCB__RR_B
	/* RR C */
	OpcodesDisMap[SHIFT_0xCB+0x19] = dis_instrCB__RR_C
	/* RR D */
	OpcodesDisMap[SHIFT_0xCB+0x1a] = dis_instrCB__RR_D
	/* RR E */
	OpcodesDisMap[SHIFT_0xCB+0x1b] = dis_instrCB__RR_E
	/* RR H */
	OpcodesDisMap[SHIFT_0xCB+0x1c] = dis_instrCB__RR_H
	/* RR L */
	OpcodesDisMap[SHIFT_0xCB+0x1d] = dis_instrCB__RR_L
	/* RR (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x1e] = dis_instrCB__RR_iHL
	/* RR A */
	OpcodesDisMap[SHIFT_0xCB+0x1f] = dis_instrCB__RR_A
	/* SLA B */
	OpcodesDisMap[SHIFT_0xCB+0x20] = dis_instrCB__SLA_B
	/* SLA C */
	OpcodesDisMap[SHIFT_0xCB+0x21] = dis_instrCB__SLA_C
	/* SLA D */
	OpcodesDisMap[SHIFT_0xCB+0x22] = dis_instrCB__SLA_D
	/* SLA E */
	OpcodesDisMap[SHIFT_0xCB+0x23] = dis_instrCB__SLA_E
	/* SLA H */
	OpcodesDisMap[SHIFT_0xCB+0x24] = dis_instrCB__SLA_H
	/* SLA L */
	OpcodesDisMap[SHIFT_0xCB+0x25] = dis_instrCB__SLA_L
	/* SLA (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x26] = dis_instrCB__SLA_iHL
	/* SLA A */
	OpcodesDisMap[SHIFT_0xCB+0x27] = dis_instrCB__SLA_A
	/* SRA B */
	OpcodesDisMap[SHIFT_0xCB+0x28] = dis_instrCB__SRA_B
	/* SRA C */
	OpcodesDisMap[SHIFT_0xCB+0x29] = dis_instrCB__SRA_C
	/* SRA D */
	OpcodesDisMap[SHIFT_0xCB+0x2a] = dis_instrCB__SRA_D
	/* SRA E */
	OpcodesDisMap[SHIFT_0xCB+0x2b] = dis_instrCB__SRA_E
	/* SRA H */
	OpcodesDisMap[SHIFT_0xCB+0x2c] = dis_instrCB__SRA_H
	/* SRA L */
	OpcodesDisMap[SHIFT_0xCB+0x2d] = dis_instrCB__SRA_L
	/* SRA (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x2e] = dis_instrCB__SRA_iHL
	/* SRA A */
	OpcodesDisMap[SHIFT_0xCB+0x2f] = dis_instrCB__SRA_A
	/* SLL B */
	OpcodesDisMap[SHIFT_0xCB+0x30] = dis_instrCB__SLL_B
	/* SLL C */
	OpcodesDisMap[SHIFT_0xCB+0x31] = dis_instrCB__SLL_C
	/* SLL D */
	OpcodesDisMap[SHIFT_0xCB+0x32] = dis_instrCB__SLL_D
	/* SLL E */
	OpcodesDisMap[SHIFT_0xCB+0x33] = dis_instrCB__SLL_E
	/* SLL H */
	OpcodesDisMap[SHIFT_0xCB+0x34] = dis_instrCB__SLL_H
	/* SLL L */
	OpcodesDisMap[SHIFT_0xCB+0x35] = dis_instrCB__SLL_L
	/* SLL (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x36] = dis_instrCB__SLL_iHL
	/* SLL A */
	OpcodesDisMap[SHIFT_0xCB+0x37] = dis_instrCB__SLL_A
	/* SRL B */
	OpcodesDisMap[SHIFT_0xCB+0x38] = dis_instrCB__SRL_B
	/* SRL C */
	OpcodesDisMap[SHIFT_0xCB+0x39] = dis_instrCB__SRL_C
	/* SRL D */
	OpcodesDisMap[SHIFT_0xCB+0x3a] = dis_instrCB__SRL_D
	/* SRL E */
	OpcodesDisMap[SHIFT_0xCB+0x3b] = dis_instrCB__SRL_E
	/* SRL H */
	OpcodesDisMap[SHIFT_0xCB+0x3c] = dis_instrCB__SRL_H
	/* SRL L */
	OpcodesDisMap[SHIFT_0xCB+0x3d] = dis_instrCB__SRL_L
	/* SRL (HL) */
	OpcodesDisMap[SHIFT_0xCB+0x3e] = dis_instrCB__SRL_iHL
	/* SRL A */
	OpcodesDisMap[SHIFT_0xCB+0x3f] = dis_instrCB__SRL_A
	/* BIT 0,B */
	OpcodesDisMap[SHIFT_0xCB+0x40] = dis_instrCB__BIT_0_B
	/* BIT 0,C */
	OpcodesDisMap[SHIFT_0xCB+0x41] = dis_instrCB__BIT_0_C
	/* BIT 0,D */
	OpcodesDisMap[SHIFT_0xCB+0x42] = dis_instrCB__BIT_0_D
	/* BIT 0,E */
	OpcodesDisMap[SHIFT_0xCB+0x43] = dis_instrCB__BIT_0_E
	/* BIT 0,H */
	OpcodesDisMap[SHIFT_0xCB+0x44] = dis_instrCB__BIT_0_H
	/* BIT 0,L */
	OpcodesDisMap[SHIFT_0xCB+0x45] = dis_instrCB__BIT_0_L
	/* BIT 0,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x46] = dis_instrCB__BIT_0_iHL
	/* BIT 0,A */
	OpcodesDisMap[SHIFT_0xCB+0x47] = dis_instrCB__BIT_0_A
	/* BIT 1,B */
	OpcodesDisMap[SHIFT_0xCB+0x48] = dis_instrCB__BIT_1_B
	/* BIT 1,C */
	OpcodesDisMap[SHIFT_0xCB+0x49] = dis_instrCB__BIT_1_C
	/* BIT 1,D */
	OpcodesDisMap[SHIFT_0xCB+0x4a] = dis_instrCB__BIT_1_D
	/* BIT 1,E */
	OpcodesDisMap[SHIFT_0xCB+0x4b] = dis_instrCB__BIT_1_E
	/* BIT 1,H */
	OpcodesDisMap[SHIFT_0xCB+0x4c] = dis_instrCB__BIT_1_H
	/* BIT 1,L */
	OpcodesDisMap[SHIFT_0xCB+0x4d] = dis_instrCB__BIT_1_L
	/* BIT 1,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x4e] = dis_instrCB__BIT_1_iHL
	/* BIT 1,A */
	OpcodesDisMap[SHIFT_0xCB+0x4f] = dis_instrCB__BIT_1_A
	/* BIT 2,B */
	OpcodesDisMap[SHIFT_0xCB+0x50] = dis_instrCB__BIT_2_B
	/* BIT 2,C */
	OpcodesDisMap[SHIFT_0xCB+0x51] = dis_instrCB__BIT_2_C
	/* BIT 2,D */
	OpcodesDisMap[SHIFT_0xCB+0x52] = dis_instrCB__BIT_2_D
	/* BIT 2,E */
	OpcodesDisMap[SHIFT_0xCB+0x53] = dis_instrCB__BIT_2_E
	/* BIT 2,H */
	OpcodesDisMap[SHIFT_0xCB+0x54] = dis_instrCB__BIT_2_H
	/* BIT 2,L */
	OpcodesDisMap[SHIFT_0xCB+0x55] = dis_instrCB__BIT_2_L
	/* BIT 2,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x56] = dis_instrCB__BIT_2_iHL
	/* BIT 2,A */
	OpcodesDisMap[SHIFT_0xCB+0x57] = dis_instrCB__BIT_2_A
	/* BIT 3,B */
	OpcodesDisMap[SHIFT_0xCB+0x58] = dis_instrCB__BIT_3_B
	/* BIT 3,C */
	OpcodesDisMap[SHIFT_0xCB+0x59] = dis_instrCB__BIT_3_C
	/* BIT 3,D */
	OpcodesDisMap[SHIFT_0xCB+0x5a] = dis_instrCB__BIT_3_D
	/* BIT 3,E */
	OpcodesDisMap[SHIFT_0xCB+0x5b] = dis_instrCB__BIT_3_E
	/* BIT 3,H */
	OpcodesDisMap[SHIFT_0xCB+0x5c] = dis_instrCB__BIT_3_H
	/* BIT 3,L */
	OpcodesDisMap[SHIFT_0xCB+0x5d] = dis_instrCB__BIT_3_L
	/* BIT 3,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x5e] = dis_instrCB__BIT_3_iHL
	/* BIT 3,A */
	OpcodesDisMap[SHIFT_0xCB+0x5f] = dis_instrCB__BIT_3_A
	/* BIT 4,B */
	OpcodesDisMap[SHIFT_0xCB+0x60] = dis_instrCB__BIT_4_B
	/* BIT 4,C */
	OpcodesDisMap[SHIFT_0xCB+0x61] = dis_instrCB__BIT_4_C
	/* BIT 4,D */
	OpcodesDisMap[SHIFT_0xCB+0x62] = dis_instrCB__BIT_4_D
	/* BIT 4,E */
	OpcodesDisMap[SHIFT_0xCB+0x63] = dis_instrCB__BIT_4_E
	/* BIT 4,H */
	OpcodesDisMap[SHIFT_0xCB+0x64] = dis_instrCB__BIT_4_H
	/* BIT 4,L */
	OpcodesDisMap[SHIFT_0xCB+0x65] = dis_instrCB__BIT_4_L
	/* BIT 4,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x66] = dis_instrCB__BIT_4_iHL
	/* BIT 4,A */
	OpcodesDisMap[SHIFT_0xCB+0x67] = dis_instrCB__BIT_4_A
	/* BIT 5,B */
	OpcodesDisMap[SHIFT_0xCB+0x68] = dis_instrCB__BIT_5_B
	/* BIT 5,C */
	OpcodesDisMap[SHIFT_0xCB+0x69] = dis_instrCB__BIT_5_C
	/* BIT 5,D */
	OpcodesDisMap[SHIFT_0xCB+0x6a] = dis_instrCB__BIT_5_D
	/* BIT 5,E */
	OpcodesDisMap[SHIFT_0xCB+0x6b] = dis_instrCB__BIT_5_E
	/* BIT 5,H */
	OpcodesDisMap[SHIFT_0xCB+0x6c] = dis_instrCB__BIT_5_H
	/* BIT 5,L */
	OpcodesDisMap[SHIFT_0xCB+0x6d] = dis_instrCB__BIT_5_L
	/* BIT 5,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x6e] = dis_instrCB__BIT_5_iHL
	/* BIT 5,A */
	OpcodesDisMap[SHIFT_0xCB+0x6f] = dis_instrCB__BIT_5_A
	/* BIT 6,B */
	OpcodesDisMap[SHIFT_0xCB+0x70] = dis_instrCB__BIT_6_B
	/* BIT 6,C */
	OpcodesDisMap[SHIFT_0xCB+0x71] = dis_instrCB__BIT_6_C
	/* BIT 6,D */
	OpcodesDisMap[SHIFT_0xCB+0x72] = dis_instrCB__BIT_6_D
	/* BIT 6,E */
	OpcodesDisMap[SHIFT_0xCB+0x73] = dis_instrCB__BIT_6_E
	/* BIT 6,H */
	OpcodesDisMap[SHIFT_0xCB+0x74] = dis_instrCB__BIT_6_H
	/* BIT 6,L */
	OpcodesDisMap[SHIFT_0xCB+0x75] = dis_instrCB__BIT_6_L
	/* BIT 6,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x76] = dis_instrCB__BIT_6_iHL
	/* BIT 6,A */
	OpcodesDisMap[SHIFT_0xCB+0x77] = dis_instrCB__BIT_6_A
	/* BIT 7,B */
	OpcodesDisMap[SHIFT_0xCB+0x78] = dis_instrCB__BIT_7_B
	/* BIT 7,C */
	OpcodesDisMap[SHIFT_0xCB+0x79] = dis_instrCB__BIT_7_C
	/* BIT 7,D */
	OpcodesDisMap[SHIFT_0xCB+0x7a] = dis_instrCB__BIT_7_D
	/* BIT 7,E */
	OpcodesDisMap[SHIFT_0xCB+0x7b] = dis_instrCB__BIT_7_E
	/* BIT 7,H */
	OpcodesDisMap[SHIFT_0xCB+0x7c] = dis_instrCB__BIT_7_H
	/* BIT 7,L */
	OpcodesDisMap[SHIFT_0xCB+0x7d] = dis_instrCB__BIT_7_L
	/* BIT 7,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x7e] = dis_instrCB__BIT_7_iHL
	/* BIT 7,A */
	OpcodesDisMap[SHIFT_0xCB+0x7f] = dis_instrCB__BIT_7_A
	/* RES 0,B */
	OpcodesDisMap[SHIFT_0xCB+0x80] = dis_instrCB__RES_0_B
	/* RES 0,C */
	OpcodesDisMap[SHIFT_0xCB+0x81] = dis_instrCB__RES_0_C
	/* RES 0,D */
	OpcodesDisMap[SHIFT_0xCB+0x82] = dis_instrCB__RES_0_D
	/* RES 0,E */
	OpcodesDisMap[SHIFT_0xCB+0x83] = dis_instrCB__RES_0_E
	/* RES 0,H */
	OpcodesDisMap[SHIFT_0xCB+0x84] = dis_instrCB__RES_0_H
	/* RES 0,L */
	OpcodesDisMap[SHIFT_0xCB+0x85] = dis_instrCB__RES_0_L
	/* RES 0,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x86] = dis_instrCB__RES_0_iHL
	/* RES 0,A */
	OpcodesDisMap[SHIFT_0xCB+0x87] = dis_instrCB__RES_0_A
	/* RES 1,B */
	OpcodesDisMap[SHIFT_0xCB+0x88] = dis_instrCB__RES_1_B
	/* RES 1,C */
	OpcodesDisMap[SHIFT_0xCB+0x89] = dis_instrCB__RES_1_C
	/* RES 1,D */
	OpcodesDisMap[SHIFT_0xCB+0x8a] = dis_instrCB__RES_1_D
	/* RES 1,E */
	OpcodesDisMap[SHIFT_0xCB+0x8b] = dis_instrCB__RES_1_E
	/* RES 1,H */
	OpcodesDisMap[SHIFT_0xCB+0x8c] = dis_instrCB__RES_1_H
	/* RES 1,L */
	OpcodesDisMap[SHIFT_0xCB+0x8d] = dis_instrCB__RES_1_L
	/* RES 1,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x8e] = dis_instrCB__RES_1_iHL
	/* RES 1,A */
	OpcodesDisMap[SHIFT_0xCB+0x8f] = dis_instrCB__RES_1_A
	/* RES 2,B */
	OpcodesDisMap[SHIFT_0xCB+0x90] = dis_instrCB__RES_2_B
	/* RES 2,C */
	OpcodesDisMap[SHIFT_0xCB+0x91] = dis_instrCB__RES_2_C
	/* RES 2,D */
	OpcodesDisMap[SHIFT_0xCB+0x92] = dis_instrCB__RES_2_D
	/* RES 2,E */
	OpcodesDisMap[SHIFT_0xCB+0x93] = dis_instrCB__RES_2_E
	/* RES 2,H */
	OpcodesDisMap[SHIFT_0xCB+0x94] = dis_instrCB__RES_2_H
	/* RES 2,L */
	OpcodesDisMap[SHIFT_0xCB+0x95] = dis_instrCB__RES_2_L
	/* RES 2,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x96] = dis_instrCB__RES_2_iHL
	/* RES 2,A */
	OpcodesDisMap[SHIFT_0xCB+0x97] = dis_instrCB__RES_2_A
	/* RES 3,B */
	OpcodesDisMap[SHIFT_0xCB+0x98] = dis_instrCB__RES_3_B
	/* RES 3,C */
	OpcodesDisMap[SHIFT_0xCB+0x99] = dis_instrCB__RES_3_C
	/* RES 3,D */
	OpcodesDisMap[SHIFT_0xCB+0x9a] = dis_instrCB__RES_3_D
	/* RES 3,E */
	OpcodesDisMap[SHIFT_0xCB+0x9b] = dis_instrCB__RES_3_E
	/* RES 3,H */
	OpcodesDisMap[SHIFT_0xCB+0x9c] = dis_instrCB__RES_3_H
	/* RES 3,L */
	OpcodesDisMap[SHIFT_0xCB+0x9d] = dis_instrCB__RES_3_L
	/* RES 3,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0x9e] = dis_instrCB__RES_3_iHL
	/* RES 3,A */
	OpcodesDisMap[SHIFT_0xCB+0x9f] = dis_instrCB__RES_3_A
	/* RES 4,B */
	OpcodesDisMap[SHIFT_0xCB+0xa0] = dis_instrCB__RES_4_B
	/* RES 4,C */
	OpcodesDisMap[SHIFT_0xCB+0xa1] = dis_instrCB__RES_4_C
	/* RES 4,D */
	OpcodesDisMap[SHIFT_0xCB+0xa2] = dis_instrCB__RES_4_D
	/* RES 4,E */
	OpcodesDisMap[SHIFT_0xCB+0xa3] = dis_instrCB__RES_4_E
	/* RES 4,H */
	OpcodesDisMap[SHIFT_0xCB+0xa4] = dis_instrCB__RES_4_H
	/* RES 4,L */
	OpcodesDisMap[SHIFT_0xCB+0xa5] = dis_instrCB__RES_4_L
	/* RES 4,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xa6] = dis_instrCB__RES_4_iHL
	/* RES 4,A */
	OpcodesDisMap[SHIFT_0xCB+0xa7] = dis_instrCB__RES_4_A
	/* RES 5,B */
	OpcodesDisMap[SHIFT_0xCB+0xa8] = dis_instrCB__RES_5_B
	/* RES 5,C */
	OpcodesDisMap[SHIFT_0xCB+0xa9] = dis_instrCB__RES_5_C
	/* RES 5,D */
	OpcodesDisMap[SHIFT_0xCB+0xaa] = dis_instrCB__RES_5_D
	/* RES 5,E */
	OpcodesDisMap[SHIFT_0xCB+0xab] = dis_instrCB__RES_5_E
	/* RES 5,H */
	OpcodesDisMap[SHIFT_0xCB+0xac] = dis_instrCB__RES_5_H
	/* RES 5,L */
	OpcodesDisMap[SHIFT_0xCB+0xad] = dis_instrCB__RES_5_L
	/* RES 5,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xae] = dis_instrCB__RES_5_iHL
	/* RES 5,A */
	OpcodesDisMap[SHIFT_0xCB+0xaf] = dis_instrCB__RES_5_A
	/* RES 6,B */
	OpcodesDisMap[SHIFT_0xCB+0xb0] = dis_instrCB__RES_6_B
	/* RES 6,C */
	OpcodesDisMap[SHIFT_0xCB+0xb1] = dis_instrCB__RES_6_C
	/* RES 6,D */
	OpcodesDisMap[SHIFT_0xCB+0xb2] = dis_instrCB__RES_6_D
	/* RES 6,E */
	OpcodesDisMap[SHIFT_0xCB+0xb3] = dis_instrCB__RES_6_E
	/* RES 6,H */
	OpcodesDisMap[SHIFT_0xCB+0xb4] = dis_instrCB__RES_6_H
	/* RES 6,L */
	OpcodesDisMap[SHIFT_0xCB+0xb5] = dis_instrCB__RES_6_L
	/* RES 6,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xb6] = dis_instrCB__RES_6_iHL
	/* RES 6,A */
	OpcodesDisMap[SHIFT_0xCB+0xb7] = dis_instrCB__RES_6_A
	/* RES 7,B */
	OpcodesDisMap[SHIFT_0xCB+0xb8] = dis_instrCB__RES_7_B
	/* RES 7,C */
	OpcodesDisMap[SHIFT_0xCB+0xb9] = dis_instrCB__RES_7_C
	/* RES 7,D */
	OpcodesDisMap[SHIFT_0xCB+0xba] = dis_instrCB__RES_7_D
	/* RES 7,E */
	OpcodesDisMap[SHIFT_0xCB+0xbb] = dis_instrCB__RES_7_E
	/* RES 7,H */
	OpcodesDisMap[SHIFT_0xCB+0xbc] = dis_instrCB__RES_7_H
	/* RES 7,L */
	OpcodesDisMap[SHIFT_0xCB+0xbd] = dis_instrCB__RES_7_L
	/* RES 7,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xbe] = dis_instrCB__RES_7_iHL
	/* RES 7,A */
	OpcodesDisMap[SHIFT_0xCB+0xbf] = dis_instrCB__RES_7_A
	/* SET 0,B */
	OpcodesDisMap[SHIFT_0xCB+0xc0] = dis_instrCB__SET_0_B
	/* SET 0,C */
	OpcodesDisMap[SHIFT_0xCB+0xc1] = dis_instrCB__SET_0_C
	/* SET 0,D */
	OpcodesDisMap[SHIFT_0xCB+0xc2] = dis_instrCB__SET_0_D
	/* SET 0,E */
	OpcodesDisMap[SHIFT_0xCB+0xc3] = dis_instrCB__SET_0_E
	/* SET 0,H */
	OpcodesDisMap[SHIFT_0xCB+0xc4] = dis_instrCB__SET_0_H
	/* SET 0,L */
	OpcodesDisMap[SHIFT_0xCB+0xc5] = dis_instrCB__SET_0_L
	/* SET 0,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xc6] = dis_instrCB__SET_0_iHL
	/* SET 0,A */
	OpcodesDisMap[SHIFT_0xCB+0xc7] = dis_instrCB__SET_0_A
	/* SET 1,B */
	OpcodesDisMap[SHIFT_0xCB+0xc8] = dis_instrCB__SET_1_B
	/* SET 1,C */
	OpcodesDisMap[SHIFT_0xCB+0xc9] = dis_instrCB__SET_1_C
	/* SET 1,D */
	OpcodesDisMap[SHIFT_0xCB+0xca] = dis_instrCB__SET_1_D
	/* SET 1,E */
	OpcodesDisMap[SHIFT_0xCB+0xcb] = dis_instrCB__SET_1_E
	/* SET 1,H */
	OpcodesDisMap[SHIFT_0xCB+0xcc] = dis_instrCB__SET_1_H
	/* SET 1,L */
	OpcodesDisMap[SHIFT_0xCB+0xcd] = dis_instrCB__SET_1_L
	/* SET 1,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xce] = dis_instrCB__SET_1_iHL
	/* SET 1,A */
	OpcodesDisMap[SHIFT_0xCB+0xcf] = dis_instrCB__SET_1_A
	/* SET 2,B */
	OpcodesDisMap[SHIFT_0xCB+0xd0] = dis_instrCB__SET_2_B
	/* SET 2,C */
	OpcodesDisMap[SHIFT_0xCB+0xd1] = dis_instrCB__SET_2_C
	/* SET 2,D */
	OpcodesDisMap[SHIFT_0xCB+0xd2] = dis_instrCB__SET_2_D
	/* SET 2,E */
	OpcodesDisMap[SHIFT_0xCB+0xd3] = dis_instrCB__SET_2_E
	/* SET 2,H */
	OpcodesDisMap[SHIFT_0xCB+0xd4] = dis_instrCB__SET_2_H
	/* SET 2,L */
	OpcodesDisMap[SHIFT_0xCB+0xd5] = dis_instrCB__SET_2_L
	/* SET 2,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xd6] = dis_instrCB__SET_2_iHL
	/* SET 2,A */
	OpcodesDisMap[SHIFT_0xCB+0xd7] = dis_instrCB__SET_2_A
	/* SET 3,B */
	OpcodesDisMap[SHIFT_0xCB+0xd8] = dis_instrCB__SET_3_B
	/* SET 3,C */
	OpcodesDisMap[SHIFT_0xCB+0xd9] = dis_instrCB__SET_3_C
	/* SET 3,D */
	OpcodesDisMap[SHIFT_0xCB+0xda] = dis_instrCB__SET_3_D
	/* SET 3,E */
	OpcodesDisMap[SHIFT_0xCB+0xdb] = dis_instrCB__SET_3_E
	/* SET 3,H */
	OpcodesDisMap[SHIFT_0xCB+0xdc] = dis_instrCB__SET_3_H
	/* SET 3,L */
	OpcodesDisMap[SHIFT_0xCB+0xdd] = dis_instrCB__SET_3_L
	/* SET 3,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xde] = dis_instrCB__SET_3_iHL
	/* SET 3,A */
	OpcodesDisMap[SHIFT_0xCB+0xdf] = dis_instrCB__SET_3_A
	/* SET 4,B */
	OpcodesDisMap[SHIFT_0xCB+0xe0] = dis_instrCB__SET_4_B
	/* SET 4,C */
	OpcodesDisMap[SHIFT_0xCB+0xe1] = dis_instrCB__SET_4_C
	/* SET 4,D */
	OpcodesDisMap[SHIFT_0xCB+0xe2] = dis_instrCB__SET_4_D
	/* SET 4,E */
	OpcodesDisMap[SHIFT_0xCB+0xe3] = dis_instrCB__SET_4_E
	/* SET 4,H */
	OpcodesDisMap[SHIFT_0xCB+0xe4] = dis_instrCB__SET_4_H
	/* SET 4,L */
	OpcodesDisMap[SHIFT_0xCB+0xe5] = dis_instrCB__SET_4_L
	/* SET 4,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xe6] = dis_instrCB__SET_4_iHL
	/* SET 4,A */
	OpcodesDisMap[SHIFT_0xCB+0xe7] = dis_instrCB__SET_4_A
	/* SET 5,B */
	OpcodesDisMap[SHIFT_0xCB+0xe8] = dis_instrCB__SET_5_B
	/* SET 5,C */
	OpcodesDisMap[SHIFT_0xCB+0xe9] = dis_instrCB__SET_5_C
	/* SET 5,D */
	OpcodesDisMap[SHIFT_0xCB+0xea] = dis_instrCB__SET_5_D
	/* SET 5,E */
	OpcodesDisMap[SHIFT_0xCB+0xeb] = dis_instrCB__SET_5_E
	/* SET 5,H */
	OpcodesDisMap[SHIFT_0xCB+0xec] = dis_instrCB__SET_5_H
	/* SET 5,L */
	OpcodesDisMap[SHIFT_0xCB+0xed] = dis_instrCB__SET_5_L
	/* SET 5,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xee] = dis_instrCB__SET_5_iHL
	/* SET 5,A */
	OpcodesDisMap[SHIFT_0xCB+0xef] = dis_instrCB__SET_5_A
	/* SET 6,B */
	OpcodesDisMap[SHIFT_0xCB+0xf0] = dis_instrCB__SET_6_B
	/* SET 6,C */
	OpcodesDisMap[SHIFT_0xCB+0xf1] = dis_instrCB__SET_6_C
	/* SET 6,D */
	OpcodesDisMap[SHIFT_0xCB+0xf2] = dis_instrCB__SET_6_D
	/* SET 6,E */
	OpcodesDisMap[SHIFT_0xCB+0xf3] = dis_instrCB__SET_6_E
	/* SET 6,H */
	OpcodesDisMap[SHIFT_0xCB+0xf4] = dis_instrCB__SET_6_H
	/* SET 6,L */
	OpcodesDisMap[SHIFT_0xCB+0xf5] = dis_instrCB__SET_6_L
	/* SET 6,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xf6] = dis_instrCB__SET_6_iHL
	/* SET 6,A */
	OpcodesDisMap[SHIFT_0xCB+0xf7] = dis_instrCB__SET_6_A
	/* SET 7,B */
	OpcodesDisMap[SHIFT_0xCB+0xf8] = dis_instrCB__SET_7_B
	/* SET 7,C */
	OpcodesDisMap[SHIFT_0xCB+0xf9] = dis_instrCB__SET_7_C
	/* SET 7,D */
	OpcodesDisMap[SHIFT_0xCB+0xfa] = dis_instrCB__SET_7_D
	/* SET 7,E */
	OpcodesDisMap[SHIFT_0xCB+0xfb] = dis_instrCB__SET_7_E
	/* SET 7,H */
	OpcodesDisMap[SHIFT_0xCB+0xfc] = dis_instrCB__SET_7_H
	/* SET 7,L */
	OpcodesDisMap[SHIFT_0xCB+0xfd] = dis_instrCB__SET_7_L
	/* SET 7,(HL) */
	OpcodesDisMap[SHIFT_0xCB+0xfe] = dis_instrCB__SET_7_iHL
	/* SET 7,A */
	OpcodesDisMap[SHIFT_0xCB+0xff] = dis_instrCB__SET_7_A

	// END of 0xcb shifted opcodes

	// BEGIN of 0xed shifted opcodes
	/* IN B,(C) */
	OpcodesDisMap[SHIFT_0xED+0x40] = dis_instrED__IN_B_iC
	/* OUT (C),B */
	OpcodesDisMap[SHIFT_0xED+0x41] = dis_instrED__OUT_iC_B
	/* SBC HL,BC */
	OpcodesDisMap[SHIFT_0xED+0x42] = dis_instrED__SBC_HL_BC
	/* LD (nnnn),BC */
	OpcodesDisMap[SHIFT_0xED+0x43] = dis_instrED__LD_iNNNN_BC
	/* NEG */
	OpcodesDisMap[SHIFT_0xED+0x7c] = dis_instrED__NEG
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xED+0x44] = OpcodesDisMap[SHIFT_0xED+0x7c]
	OpcodesDisMap[SHIFT_0xED+0x4c] = OpcodesDisMap[SHIFT_0xED+0x7c]
	OpcodesDisMap[SHIFT_0xED+0x54] = OpcodesDisMap[SHIFT_0xED+0x7c]
	OpcodesDisMap[SHIFT_0xED+0x5c] = OpcodesDisMap[SHIFT_0xED+0x7c]
	OpcodesDisMap[SHIFT_0xED+0x64] = OpcodesDisMap[SHIFT_0xED+0x7c]
	OpcodesDisMap[SHIFT_0xED+0x6c] = OpcodesDisMap[SHIFT_0xED+0x7c]
	OpcodesDisMap[SHIFT_0xED+0x74] = OpcodesDisMap[SHIFT_0xED+0x7c]
	/* RETN */
	OpcodesDisMap[SHIFT_0xED+0x7d] = dis_instrED__RETN
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xED+0x45] = OpcodesDisMap[SHIFT_0xED+0x7d]
	OpcodesDisMap[SHIFT_0xED+0x4d] = OpcodesDisMap[SHIFT_0xED+0x7d]
	OpcodesDisMap[SHIFT_0xED+0x55] = OpcodesDisMap[SHIFT_0xED+0x7d]
	OpcodesDisMap[SHIFT_0xED+0x5d] = OpcodesDisMap[SHIFT_0xED+0x7d]
	OpcodesDisMap[SHIFT_0xED+0x65] = OpcodesDisMap[SHIFT_0xED+0x7d]
	OpcodesDisMap[SHIFT_0xED+0x6d] = OpcodesDisMap[SHIFT_0xED+0x7d]
	OpcodesDisMap[SHIFT_0xED+0x75] = OpcodesDisMap[SHIFT_0xED+0x7d]
	/* IM 0 */
	OpcodesDisMap[SHIFT_0xED+0x6e] = dis_instrED__IM_0
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xED+0x46] = OpcodesDisMap[SHIFT_0xED+0x6e]
	OpcodesDisMap[SHIFT_0xED+0x4e] = OpcodesDisMap[SHIFT_0xED+0x6e]
	OpcodesDisMap[SHIFT_0xED+0x66] = OpcodesDisMap[SHIFT_0xED+0x6e]
	/* LD I,A */
	OpcodesDisMap[SHIFT_0xED+0x47] = dis_instrED__LD_I_A
	/* IN C,(C) */
	OpcodesDisMap[SHIFT_0xED+0x48] = dis_instrED__IN_C_iC
	/* OUT (C),C */
	OpcodesDisMap[SHIFT_0xED+0x49] = dis_instrED__OUT_iC_C
	/* ADC HL,BC */
	OpcodesDisMap[SHIFT_0xED+0x4a] = dis_instrED__ADC_HL_BC
	/* LD BC,(nnnn) */
	OpcodesDisMap[SHIFT_0xED+0x4b] = dis_instrED__LD_BC_iNNNN
	/* LD R,A */
	OpcodesDisMap[SHIFT_0xED+0x4f] = dis_instrED__LD_R_A
	/* IN D,(C) */
	OpcodesDisMap[SHIFT_0xED+0x50] = dis_instrED__IN_D_iC
	/* OUT (C),D */
	OpcodesDisMap[SHIFT_0xED+0x51] = dis_instrED__OUT_iC_D
	/* SBC HL,DE */
	OpcodesDisMap[SHIFT_0xED+0x52] = dis_instrED__SBC_HL_DE
	/* LD (nnnn),DE */
	OpcodesDisMap[SHIFT_0xED+0x53] = dis_instrED__LD_iNNNN_DE
	/* IM 1 */
	OpcodesDisMap[SHIFT_0xED+0x76] = dis_instrED__IM_1
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xED+0x56] = OpcodesDisMap[SHIFT_0xED+0x76]
	/* LD A,I */
	OpcodesDisMap[SHIFT_0xED+0x57] = dis_instrED__LD_A_I
	/* IN E,(C) */
	OpcodesDisMap[SHIFT_0xED+0x58] = dis_instrED__IN_E_iC
	/* OUT (C),E */
	OpcodesDisMap[SHIFT_0xED+0x59] = dis_instrED__OUT_iC_E
	/* ADC HL,DE */
	OpcodesDisMap[SHIFT_0xED+0x5a] = dis_instrED__ADC_HL_DE
	/* LD DE,(nnnn) */
	OpcodesDisMap[SHIFT_0xED+0x5b] = dis_instrED__LD_DE_iNNNN
	/* IM 2 */
	OpcodesDisMap[SHIFT_0xED+0x7e] = dis_instrED__IM_2
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xED+0x5e] = OpcodesDisMap[SHIFT_0xED+0x7e]
	/* LD A,R */
	OpcodesDisMap[SHIFT_0xED+0x5f] = dis_instrED__LD_A_R
	/* IN H,(C) */
	OpcodesDisMap[SHIFT_0xED+0x60] = dis_instrED__IN_H_iC
	/* OUT (C),H */
	OpcodesDisMap[SHIFT_0xED+0x61] = dis_instrED__OUT_iC_H
	/* SBC HL,HL */
	OpcodesDisMap[SHIFT_0xED+0x62] = dis_instrED__SBC_HL_HL
	/* LD (nnnn),HL */
	OpcodesDisMap[SHIFT_0xED+0x63] = dis_instrED__LD_iNNNN_HL
	/* RRD */
	OpcodesDisMap[SHIFT_0xED+0x67] = dis_instrED__RRD
	/* IN L,(C) */
	OpcodesDisMap[SHIFT_0xED+0x68] = dis_instrED__IN_L_iC
	/* OUT (C),L */
	OpcodesDisMap[SHIFT_0xED+0x69] = dis_instrED__OUT_iC_L
	/* ADC HL,HL */
	OpcodesDisMap[SHIFT_0xED+0x6a] = dis_instrED__ADC_HL_HL
	/* LD HL,(nnnn) */
	OpcodesDisMap[SHIFT_0xED+0x6b] = dis_instrED__LD_HL_iNNNN
	/* RLD */
	OpcodesDisMap[SHIFT_0xED+0x6f] = dis_instrED__RLD
	/* IN F,(C) */
	OpcodesDisMap[SHIFT_0xED+0x70] = dis_instrED__IN_F_iC
	/* OUT (C),0 */
	OpcodesDisMap[SHIFT_0xED+0x71] = dis_instrED__OUT_iC_0
	/* SBC HL,SP */
	OpcodesDisMap[SHIFT_0xED+0x72] = dis_instrED__SBC_HL_SP
	/* LD (nnnn),SP */
	OpcodesDisMap[SHIFT_0xED+0x73] = dis_instrED__LD_iNNNN_SP
	/* IN A,(C) */
	OpcodesDisMap[SHIFT_0xED+0x78] = dis_instrED__IN_A_iC
	/* OUT (C),A */
	OpcodesDisMap[SHIFT_0xED+0x79] = dis_instrED__OUT_iC_A
	/* ADC HL,SP */
	OpcodesDisMap[SHIFT_0xED+0x7a] = dis_instrED__ADC_HL_SP
	/* LD SP,(nnnn) */
	OpcodesDisMap[SHIFT_0xED+0x7b] = dis_instrED__LD_SP_iNNNN
	/* LDI */
	OpcodesDisMap[SHIFT_0xED+0xa0] = dis_instrED__LDI
	/* CPI */
	OpcodesDisMap[SHIFT_0xED+0xa1] = dis_instrED__CPI
	/* INI */
	OpcodesDisMap[SHIFT_0xED+0xa2] = dis_instrED__INI
	/* OUTI */
	OpcodesDisMap[SHIFT_0xED+0xa3] = dis_instrED__OUTI
	/* LDD */
	OpcodesDisMap[SHIFT_0xED+0xa8] = dis_instrED__LDD
	/* CPD */
	OpcodesDisMap[SHIFT_0xED+0xa9] = dis_instrED__CPD
	/* IND */
	OpcodesDisMap[SHIFT_0xED+0xaa] = dis_instrED__IND
	/* OUTD */
	OpcodesDisMap[SHIFT_0xED+0xab] = dis_instrED__OUTD
	/* LDIR */
	OpcodesDisMap[SHIFT_0xED+0xb0] = dis_instrED__LDIR
	/* CPIR */
	OpcodesDisMap[SHIFT_0xED+0xb1] = dis_instrED__CPIR
	/* INIR */
	OpcodesDisMap[SHIFT_0xED+0xb2] = dis_instrED__INIR
	/* OTIR */
	OpcodesDisMap[SHIFT_0xED+0xb3] = dis_instrED__OTIR
	/* LDDR */
	OpcodesDisMap[SHIFT_0xED+0xb8] = dis_instrED__LDDR
	/* CPDR */
	OpcodesDisMap[SHIFT_0xED+0xb9] = dis_instrED__CPDR
	/* INDR */
	OpcodesDisMap[SHIFT_0xED+0xba] = dis_instrED__INDR
	/* OTDR */
	OpcodesDisMap[SHIFT_0xED+0xbb] = dis_instrED__OTDR
	/* slttrap */
	OpcodesDisMap[SHIFT_0xED+0xfb] = dis_instrED__SLTTRAP

	// END of 0xed shifted opcodes

	// BEGIN of 0xdd shifted opcodes
	/* ADD REGISTER,BC */
	OpcodesDisMap[SHIFT_0xDD+0x09] = dis_instrDD__ADD_REG_BC
	/* ADD REGISTER,DE */
	OpcodesDisMap[SHIFT_0xDD+0x19] = dis_instrDD__ADD_REG_DE
	/* LD REGISTER,nnnn */
	OpcodesDisMap[SHIFT_0xDD+0x21] = dis_instrDD__LD_REG_NNNN
	/* LD (nnnn),REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0x22] = dis_instrDD__LD_iNNNN_REG
	/* INC REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0x23] = dis_instrDD__INC_REG
	/* INC REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x24] = dis_instrDD__INC_REGH
	/* DEC REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x25] = dis_instrDD__DEC_REGH
	/* LD REGISTERH,nn */
	OpcodesDisMap[SHIFT_0xDD+0x26] = dis_instrDD__LD_REGH_NN
	/* ADD REGISTER,REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0x29] = dis_instrDD__ADD_REG_REG
	/* LD REGISTER,(nnnn) */
	OpcodesDisMap[SHIFT_0xDD+0x2a] = dis_instrDD__LD_REG_iNNNN
	/* DEC REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0x2b] = dis_instrDD__DEC_REG
	/* INC REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x2c] = dis_instrDD__INC_REGL
	/* DEC REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x2d] = dis_instrDD__DEC_REGL
	/* LD REGISTERL,nn */
	OpcodesDisMap[SHIFT_0xDD+0x2e] = dis_instrDD__LD_REGL_NN
	/* INC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x34] = dis_instrDD__INC_iREGpDD
	/* DEC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x35] = dis_instrDD__DEC_iREGpDD
	/* LD (REGISTER+dd),nn */
	OpcodesDisMap[SHIFT_0xDD+0x36] = dis_instrDD__LD_iREGpDD_NN
	/* ADD REGISTER,SP */
	OpcodesDisMap[SHIFT_0xDD+0x39] = dis_instrDD__ADD_REG_SP
	/* LD B,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x44] = dis_instrDD__LD_B_REGH
	/* LD B,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x45] = dis_instrDD__LD_B_REGL
	/* LD B,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x46] = dis_instrDD__LD_B_iREGpDD
	/* LD C,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x4c] = dis_instrDD__LD_C_REGH
	/* LD C,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x4d] = dis_instrDD__LD_C_REGL
	/* LD C,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x4e] = dis_instrDD__LD_C_iREGpDD
	/* LD D,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x54] = dis_instrDD__LD_D_REGH
	/* LD D,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x55] = dis_instrDD__LD_D_REGL
	/* LD D,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x56] = dis_instrDD__LD_D_iREGpDD
	/* LD E,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x5c] = dis_instrDD__LD_E_REGH
	/* LD E,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x5d] = dis_instrDD__LD_E_REGL
	/* LD E,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x5e] = dis_instrDD__LD_E_iREGpDD
	/* LD REGISTERH,B */
	OpcodesDisMap[SHIFT_0xDD+0x60] = dis_instrDD__LD_REGH_B
	/* LD REGISTERH,C */
	OpcodesDisMap[SHIFT_0xDD+0x61] = dis_instrDD__LD_REGH_C
	/* LD REGISTERH,D */
	OpcodesDisMap[SHIFT_0xDD+0x62] = dis_instrDD__LD_REGH_D
	/* LD REGISTERH,E */
	OpcodesDisMap[SHIFT_0xDD+0x63] = dis_instrDD__LD_REGH_E
	/* LD REGISTERH,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x64] = dis_instrDD__LD_REGH_REGH
	/* LD REGISTERH,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x65] = dis_instrDD__LD_REGH_REGL
	/* LD H,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x66] = dis_instrDD__LD_H_iREGpDD
	/* LD REGISTERH,A */
	OpcodesDisMap[SHIFT_0xDD+0x67] = dis_instrDD__LD_REGH_A
	/* LD REGISTERL,B */
	OpcodesDisMap[SHIFT_0xDD+0x68] = dis_instrDD__LD_REGL_B
	/* LD REGISTERL,C */
	OpcodesDisMap[SHIFT_0xDD+0x69] = dis_instrDD__LD_REGL_C
	/* LD REGISTERL,D */
	OpcodesDisMap[SHIFT_0xDD+0x6a] = dis_instrDD__LD_REGL_D
	/* LD REGISTERL,E */
	OpcodesDisMap[SHIFT_0xDD+0x6b] = dis_instrDD__LD_REGL_E
	/* LD REGISTERL,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x6c] = dis_instrDD__LD_REGL_REGH
	/* LD REGISTERL,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x6d] = dis_instrDD__LD_REGL_REGL
	/* LD L,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x6e] = dis_instrDD__LD_L_iREGpDD
	/* LD REGISTERL,A */
	OpcodesDisMap[SHIFT_0xDD+0x6f] = dis_instrDD__LD_REGL_A
	/* LD (REGISTER+dd),B */
	OpcodesDisMap[SHIFT_0xDD+0x70] = dis_instrDD__LD_iREGpDD_B
	/* LD (REGISTER+dd),C */
	OpcodesDisMap[SHIFT_0xDD+0x71] = dis_instrDD__LD_iREGpDD_C
	/* LD (REGISTER+dd),D */
	OpcodesDisMap[SHIFT_0xDD+0x72] = dis_instrDD__LD_iREGpDD_D
	/* LD (REGISTER+dd),E */
	OpcodesDisMap[SHIFT_0xDD+0x73] = dis_instrDD__LD_iREGpDD_E
	/* LD (REGISTER+dd),H */
	OpcodesDisMap[SHIFT_0xDD+0x74] = dis_instrDD__LD_iREGpDD_H
	/* LD (REGISTER+dd),L */
	OpcodesDisMap[SHIFT_0xDD+0x75] = dis_instrDD__LD_iREGpDD_L
	/* LD (REGISTER+dd),A */
	OpcodesDisMap[SHIFT_0xDD+0x77] = dis_instrDD__LD_iREGpDD_A
	/* LD A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x7c] = dis_instrDD__LD_A_REGH
	/* LD A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x7d] = dis_instrDD__LD_A_REGL
	/* LD A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x7e] = dis_instrDD__LD_A_iREGpDD
	/* ADD A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x84] = dis_instrDD__ADD_A_REGH
	/* ADD A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x85] = dis_instrDD__ADD_A_REGL
	/* ADD A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x86] = dis_instrDD__ADD_A_iREGpDD
	/* ADC A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x8c] = dis_instrDD__ADC_A_REGH
	/* ADC A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x8d] = dis_instrDD__ADC_A_REGL
	/* ADC A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x8e] = dis_instrDD__ADC_A_iREGpDD
	/* SUB A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x94] = dis_instrDD__SUB_A_REGH
	/* SUB A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x95] = dis_instrDD__SUB_A_REGL
	/* SUB A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x96] = dis_instrDD__SUB_A_iREGpDD
	/* SBC A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0x9c] = dis_instrDD__SBC_A_REGH
	/* SBC A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0x9d] = dis_instrDD__SBC_A_REGL
	/* SBC A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0x9e] = dis_instrDD__SBC_A_iREGpDD
	/* AND A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0xa4] = dis_instrDD__AND_A_REGH
	/* AND A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0xa5] = dis_instrDD__AND_A_REGL
	/* AND A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0xa6] = dis_instrDD__AND_A_iREGpDD
	/* XOR A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0xac] = dis_instrDD__XOR_A_REGH
	/* XOR A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0xad] = dis_instrDD__XOR_A_REGL
	/* XOR A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0xae] = dis_instrDD__XOR_A_iREGpDD
	/* OR A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0xb4] = dis_instrDD__OR_A_REGH
	/* OR A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0xb5] = dis_instrDD__OR_A_REGL
	/* OR A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0xb6] = dis_instrDD__OR_A_iREGpDD
	/* CP A,REGISTERH */
	OpcodesDisMap[SHIFT_0xDD+0xbc] = dis_instrDD__CP_A_REGH
	/* CP A,REGISTERL */
	OpcodesDisMap[SHIFT_0xDD+0xbd] = dis_instrDD__CP_A_REGL
	/* CP A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDD+0xbe] = dis_instrDD__CP_A_iREGpDD
	/* shift DDFDCB */
	OpcodesDisMap[SHIFT_0xDD+0xcb] = dis_instrDD__SHIFT_DDFDCB
	/* POP REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0xe1] = dis_instrDD__POP_REG
	/* EX (SP),REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0xe3] = dis_instrDD__EX_iSP_REG
	/* PUSH REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0xe5] = dis_instrDD__PUSH_REG
	/* JP REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0xe9] = dis_instrDD__JP_REG
	/* LD SP,REGISTER */
	OpcodesDisMap[SHIFT_0xDD+0xf9] = dis_instrDD__LD_SP_REG

	// END of 0xdd shifted opcodes

	// BEGIN of 0xfd shifted opcodes
	/* ADD REGISTER,BC */
	OpcodesDisMap[SHIFT_0xFD+0x09] = dis_instrFD__ADD_REG_BC
	/* ADD REGISTER,DE */
	OpcodesDisMap[SHIFT_0xFD+0x19] = dis_instrFD__ADD_REG_DE
	/* LD REGISTER,nnnn */
	OpcodesDisMap[SHIFT_0xFD+0x21] = dis_instrFD__LD_REG_NNNN
	/* LD (nnnn),REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0x22] = dis_instrFD__LD_iNNNN_REG
	/* INC REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0x23] = dis_instrFD__INC_REG
	/* INC REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x24] = dis_instrFD__INC_REGH
	/* DEC REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x25] = dis_instrFD__DEC_REGH
	/* LD REGISTERH,nn */
	OpcodesDisMap[SHIFT_0xFD+0x26] = dis_instrFD__LD_REGH_NN
	/* ADD REGISTER,REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0x29] = dis_instrFD__ADD_REG_REG
	/* LD REGISTER,(nnnn) */
	OpcodesDisMap[SHIFT_0xFD+0x2a] = dis_instrFD__LD_REG_iNNNN
	/* DEC REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0x2b] = dis_instrFD__DEC_REG
	/* INC REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x2c] = dis_instrFD__INC_REGL
	/* DEC REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x2d] = dis_instrFD__DEC_REGL
	/* LD REGISTERL,nn */
	OpcodesDisMap[SHIFT_0xFD+0x2e] = dis_instrFD__LD_REGL_NN
	/* INC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x34] = dis_instrFD__INC_iREGpDD
	/* DEC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x35] = dis_instrFD__DEC_iREGpDD
	/* LD (REGISTER+dd),nn */
	OpcodesDisMap[SHIFT_0xFD+0x36] = dis_instrFD__LD_iREGpDD_NN
	/* ADD REGISTER,SP */
	OpcodesDisMap[SHIFT_0xFD+0x39] = dis_instrFD__ADD_REG_SP
	/* LD B,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x44] = dis_instrFD__LD_B_REGH
	/* LD B,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x45] = dis_instrFD__LD_B_REGL
	/* LD B,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x46] = dis_instrFD__LD_B_iREGpDD
	/* LD C,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x4c] = dis_instrFD__LD_C_REGH
	/* LD C,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x4d] = dis_instrFD__LD_C_REGL
	/* LD C,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x4e] = dis_instrFD__LD_C_iREGpDD
	/* LD D,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x54] = dis_instrFD__LD_D_REGH
	/* LD D,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x55] = dis_instrFD__LD_D_REGL
	/* LD D,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x56] = dis_instrFD__LD_D_iREGpDD
	/* LD E,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x5c] = dis_instrFD__LD_E_REGH
	/* LD E,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x5d] = dis_instrFD__LD_E_REGL
	/* LD E,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x5e] = dis_instrFD__LD_E_iREGpDD
	/* LD REGISTERH,B */
	OpcodesDisMap[SHIFT_0xFD+0x60] = dis_instrFD__LD_REGH_B
	/* LD REGISTERH,C */
	OpcodesDisMap[SHIFT_0xFD+0x61] = dis_instrFD__LD_REGH_C
	/* LD REGISTERH,D */
	OpcodesDisMap[SHIFT_0xFD+0x62] = dis_instrFD__LD_REGH_D
	/* LD REGISTERH,E */
	OpcodesDisMap[SHIFT_0xFD+0x63] = dis_instrFD__LD_REGH_E
	/* LD REGISTERH,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x64] = dis_instrFD__LD_REGH_REGH
	/* LD REGISTERH,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x65] = dis_instrFD__LD_REGH_REGL
	/* LD H,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x66] = dis_instrFD__LD_H_iREGpDD
	/* LD REGISTERH,A */
	OpcodesDisMap[SHIFT_0xFD+0x67] = dis_instrFD__LD_REGH_A
	/* LD REGISTERL,B */
	OpcodesDisMap[SHIFT_0xFD+0x68] = dis_instrFD__LD_REGL_B
	/* LD REGISTERL,C */
	OpcodesDisMap[SHIFT_0xFD+0x69] = dis_instrFD__LD_REGL_C
	/* LD REGISTERL,D */
	OpcodesDisMap[SHIFT_0xFD+0x6a] = dis_instrFD__LD_REGL_D
	/* LD REGISTERL,E */
	OpcodesDisMap[SHIFT_0xFD+0x6b] = dis_instrFD__LD_REGL_E
	/* LD REGISTERL,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x6c] = dis_instrFD__LD_REGL_REGH
	/* LD REGISTERL,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x6d] = dis_instrFD__LD_REGL_REGL
	/* LD L,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x6e] = dis_instrFD__LD_L_iREGpDD
	/* LD REGISTERL,A */
	OpcodesDisMap[SHIFT_0xFD+0x6f] = dis_instrFD__LD_REGL_A
	/* LD (REGISTER+dd),B */
	OpcodesDisMap[SHIFT_0xFD+0x70] = dis_instrFD__LD_iREGpDD_B
	/* LD (REGISTER+dd),C */
	OpcodesDisMap[SHIFT_0xFD+0x71] = dis_instrFD__LD_iREGpDD_C
	/* LD (REGISTER+dd),D */
	OpcodesDisMap[SHIFT_0xFD+0x72] = dis_instrFD__LD_iREGpDD_D
	/* LD (REGISTER+dd),E */
	OpcodesDisMap[SHIFT_0xFD+0x73] = dis_instrFD__LD_iREGpDD_E
	/* LD (REGISTER+dd),H */
	OpcodesDisMap[SHIFT_0xFD+0x74] = dis_instrFD__LD_iREGpDD_H
	/* LD (REGISTER+dd),L */
	OpcodesDisMap[SHIFT_0xFD+0x75] = dis_instrFD__LD_iREGpDD_L
	/* LD (REGISTER+dd),A */
	OpcodesDisMap[SHIFT_0xFD+0x77] = dis_instrFD__LD_iREGpDD_A
	/* LD A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x7c] = dis_instrFD__LD_A_REGH
	/* LD A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x7d] = dis_instrFD__LD_A_REGL
	/* LD A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x7e] = dis_instrFD__LD_A_iREGpDD
	/* ADD A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x84] = dis_instrFD__ADD_A_REGH
	/* ADD A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x85] = dis_instrFD__ADD_A_REGL
	/* ADD A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x86] = dis_instrFD__ADD_A_iREGpDD
	/* ADC A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x8c] = dis_instrFD__ADC_A_REGH
	/* ADC A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x8d] = dis_instrFD__ADC_A_REGL
	/* ADC A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x8e] = dis_instrFD__ADC_A_iREGpDD
	/* SUB A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x94] = dis_instrFD__SUB_A_REGH
	/* SUB A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x95] = dis_instrFD__SUB_A_REGL
	/* SUB A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x96] = dis_instrFD__SUB_A_iREGpDD
	/* SBC A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0x9c] = dis_instrFD__SBC_A_REGH
	/* SBC A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0x9d] = dis_instrFD__SBC_A_REGL
	/* SBC A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0x9e] = dis_instrFD__SBC_A_iREGpDD
	/* AND A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0xa4] = dis_instrFD__AND_A_REGH
	/* AND A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0xa5] = dis_instrFD__AND_A_REGL
	/* AND A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0xa6] = dis_instrFD__AND_A_iREGpDD
	/* XOR A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0xac] = dis_instrFD__XOR_A_REGH
	/* XOR A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0xad] = dis_instrFD__XOR_A_REGL
	/* XOR A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0xae] = dis_instrFD__XOR_A_iREGpDD
	/* OR A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0xb4] = dis_instrFD__OR_A_REGH
	/* OR A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0xb5] = dis_instrFD__OR_A_REGL
	/* OR A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0xb6] = dis_instrFD__OR_A_iREGpDD
	/* CP A,REGISTERH */
	OpcodesDisMap[SHIFT_0xFD+0xbc] = dis_instrFD__CP_A_REGH
	/* CP A,REGISTERL */
	OpcodesDisMap[SHIFT_0xFD+0xbd] = dis_instrFD__CP_A_REGL
	/* CP A,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xFD+0xbe] = dis_instrFD__CP_A_iREGpDD
	/* shift DDFDCB */
	OpcodesDisMap[SHIFT_0xFD+0xcb] = dis_instrFD__SHIFT_DDFDCB
	/* POP REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0xe1] = dis_instrFD__POP_REG
	/* EX (SP),REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0xe3] = dis_instrFD__EX_iSP_REG
	/* PUSH REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0xe5] = dis_instrFD__PUSH_REG
	/* JP REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0xe9] = dis_instrFD__JP_REG
	/* LD SP,REGISTER */
	OpcodesDisMap[SHIFT_0xFD+0xf9] = dis_instrFD__LD_SP_REG

	// END of 0xfd shifted opcodes

	// BEGIN of 0xddfdcb shifted opcodes
	/* LD B,RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x00] = dis_instrDDCB__LD_B_RLC_iREGpDD
	/* LD C,RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x01] = dis_instrDDCB__LD_C_RLC_iREGpDD
	/* LD D,RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x02] = dis_instrDDCB__LD_D_RLC_iREGpDD
	/* LD E,RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x03] = dis_instrDDCB__LD_E_RLC_iREGpDD
	/* LD H,RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x04] = dis_instrDDCB__LD_H_RLC_iREGpDD
	/* LD L,RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x05] = dis_instrDDCB__LD_L_RLC_iREGpDD
	/* RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x06] = dis_instrDDCB__RLC_iREGpDD
	/* LD A,RLC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x07] = dis_instrDDCB__LD_A_RLC_iREGpDD
	/* LD B,RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x08] = dis_instrDDCB__LD_B_RRC_iREGpDD
	/* LD C,RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x09] = dis_instrDDCB__LD_C_RRC_iREGpDD
	/* LD D,RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x0a] = dis_instrDDCB__LD_D_RRC_iREGpDD
	/* LD E,RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x0b] = dis_instrDDCB__LD_E_RRC_iREGpDD
	/* LD H,RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x0c] = dis_instrDDCB__LD_H_RRC_iREGpDD
	/* LD L,RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x0d] = dis_instrDDCB__LD_L_RRC_iREGpDD
	/* RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x0e] = dis_instrDDCB__RRC_iREGpDD
	/* LD A,RRC (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x0f] = dis_instrDDCB__LD_A_RRC_iREGpDD
	/* LD B,RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x10] = dis_instrDDCB__LD_B_RL_iREGpDD
	/* LD C,RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x11] = dis_instrDDCB__LD_C_RL_iREGpDD
	/* LD D,RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x12] = dis_instrDDCB__LD_D_RL_iREGpDD
	/* LD E,RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x13] = dis_instrDDCB__LD_E_RL_iREGpDD
	/* LD H,RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x14] = dis_instrDDCB__LD_H_RL_iREGpDD
	/* LD L,RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x15] = dis_instrDDCB__LD_L_RL_iREGpDD
	/* RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x16] = dis_instrDDCB__RL_iREGpDD
	/* LD A,RL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x17] = dis_instrDDCB__LD_A_RL_iREGpDD
	/* LD B,RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x18] = dis_instrDDCB__LD_B_RR_iREGpDD
	/* LD C,RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x19] = dis_instrDDCB__LD_C_RR_iREGpDD
	/* LD D,RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x1a] = dis_instrDDCB__LD_D_RR_iREGpDD
	/* LD E,RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x1b] = dis_instrDDCB__LD_E_RR_iREGpDD
	/* LD H,RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x1c] = dis_instrDDCB__LD_H_RR_iREGpDD
	/* LD L,RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x1d] = dis_instrDDCB__LD_L_RR_iREGpDD
	/* RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x1e] = dis_instrDDCB__RR_iREGpDD
	/* LD A,RR (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x1f] = dis_instrDDCB__LD_A_RR_iREGpDD
	/* LD B,SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x20] = dis_instrDDCB__LD_B_SLA_iREGpDD
	/* LD C,SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x21] = dis_instrDDCB__LD_C_SLA_iREGpDD
	/* LD D,SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x22] = dis_instrDDCB__LD_D_SLA_iREGpDD
	/* LD E,SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x23] = dis_instrDDCB__LD_E_SLA_iREGpDD
	/* LD H,SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x24] = dis_instrDDCB__LD_H_SLA_iREGpDD
	/* LD L,SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x25] = dis_instrDDCB__LD_L_SLA_iREGpDD
	/* SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x26] = dis_instrDDCB__SLA_iREGpDD
	/* LD A,SLA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x27] = dis_instrDDCB__LD_A_SLA_iREGpDD
	/* LD B,SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x28] = dis_instrDDCB__LD_B_SRA_iREGpDD
	/* LD C,SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x29] = dis_instrDDCB__LD_C_SRA_iREGpDD
	/* LD D,SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x2a] = dis_instrDDCB__LD_D_SRA_iREGpDD
	/* LD E,SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x2b] = dis_instrDDCB__LD_E_SRA_iREGpDD
	/* LD H,SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x2c] = dis_instrDDCB__LD_H_SRA_iREGpDD
	/* LD L,SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x2d] = dis_instrDDCB__LD_L_SRA_iREGpDD
	/* SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x2e] = dis_instrDDCB__SRA_iREGpDD
	/* LD A,SRA (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x2f] = dis_instrDDCB__LD_A_SRA_iREGpDD
	/* LD B,SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x30] = dis_instrDDCB__LD_B_SLL_iREGpDD
	/* LD C,SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x31] = dis_instrDDCB__LD_C_SLL_iREGpDD
	/* LD D,SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x32] = dis_instrDDCB__LD_D_SLL_iREGpDD
	/* LD E,SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x33] = dis_instrDDCB__LD_E_SLL_iREGpDD
	/* LD H,SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x34] = dis_instrDDCB__LD_H_SLL_iREGpDD
	/* LD L,SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x35] = dis_instrDDCB__LD_L_SLL_iREGpDD
	/* SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x36] = dis_instrDDCB__SLL_iREGpDD
	/* LD A,SLL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x37] = dis_instrDDCB__LD_A_SLL_iREGpDD
	/* LD B,SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x38] = dis_instrDDCB__LD_B_SRL_iREGpDD
	/* LD C,SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x39] = dis_instrDDCB__LD_C_SRL_iREGpDD
	/* LD D,SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x3a] = dis_instrDDCB__LD_D_SRL_iREGpDD
	/* LD E,SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x3b] = dis_instrDDCB__LD_E_SRL_iREGpDD
	/* LD H,SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x3c] = dis_instrDDCB__LD_H_SRL_iREGpDD
	/* LD L,SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x3d] = dis_instrDDCB__LD_L_SRL_iREGpDD
	/* SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x3e] = dis_instrDDCB__SRL_iREGpDD
	/* LD A,SRL (REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x3f] = dis_instrDDCB__LD_A_SRL_iREGpDD
	/* BIT 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x47] = dis_instrDDCB__BIT_0_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x40] = OpcodesDisMap[SHIFT_0xDDCB+0x47]
	OpcodesDisMap[SHIFT_0xDDCB+0x41] = OpcodesDisMap[SHIFT_0xDDCB+0x47]
	OpcodesDisMap[SHIFT_0xDDCB+0x42] = OpcodesDisMap[SHIFT_0xDDCB+0x47]
	OpcodesDisMap[SHIFT_0xDDCB+0x43] = OpcodesDisMap[SHIFT_0xDDCB+0x47]
	OpcodesDisMap[SHIFT_0xDDCB+0x44] = OpcodesDisMap[SHIFT_0xDDCB+0x47]
	OpcodesDisMap[SHIFT_0xDDCB+0x45] = OpcodesDisMap[SHIFT_0xDDCB+0x47]
	OpcodesDisMap[SHIFT_0xDDCB+0x46] = OpcodesDisMap[SHIFT_0xDDCB+0x47]
	/* BIT 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x4f] = dis_instrDDCB__BIT_1_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x48] = OpcodesDisMap[SHIFT_0xDDCB+0x4f]
	OpcodesDisMap[SHIFT_0xDDCB+0x49] = OpcodesDisMap[SHIFT_0xDDCB+0x4f]
	OpcodesDisMap[SHIFT_0xDDCB+0x4a] = OpcodesDisMap[SHIFT_0xDDCB+0x4f]
	OpcodesDisMap[SHIFT_0xDDCB+0x4b] = OpcodesDisMap[SHIFT_0xDDCB+0x4f]
	OpcodesDisMap[SHIFT_0xDDCB+0x4c] = OpcodesDisMap[SHIFT_0xDDCB+0x4f]
	OpcodesDisMap[SHIFT_0xDDCB+0x4d] = OpcodesDisMap[SHIFT_0xDDCB+0x4f]
	OpcodesDisMap[SHIFT_0xDDCB+0x4e] = OpcodesDisMap[SHIFT_0xDDCB+0x4f]
	/* BIT 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x57] = dis_instrDDCB__BIT_2_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x50] = OpcodesDisMap[SHIFT_0xDDCB+0x57]
	OpcodesDisMap[SHIFT_0xDDCB+0x51] = OpcodesDisMap[SHIFT_0xDDCB+0x57]
	OpcodesDisMap[SHIFT_0xDDCB+0x52] = OpcodesDisMap[SHIFT_0xDDCB+0x57]
	OpcodesDisMap[SHIFT_0xDDCB+0x53] = OpcodesDisMap[SHIFT_0xDDCB+0x57]
	OpcodesDisMap[SHIFT_0xDDCB+0x54] = OpcodesDisMap[SHIFT_0xDDCB+0x57]
	OpcodesDisMap[SHIFT_0xDDCB+0x55] = OpcodesDisMap[SHIFT_0xDDCB+0x57]
	OpcodesDisMap[SHIFT_0xDDCB+0x56] = OpcodesDisMap[SHIFT_0xDDCB+0x57]
	/* BIT 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x5f] = dis_instrDDCB__BIT_3_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x58] = OpcodesDisMap[SHIFT_0xDDCB+0x5f]
	OpcodesDisMap[SHIFT_0xDDCB+0x59] = OpcodesDisMap[SHIFT_0xDDCB+0x5f]
	OpcodesDisMap[SHIFT_0xDDCB+0x5a] = OpcodesDisMap[SHIFT_0xDDCB+0x5f]
	OpcodesDisMap[SHIFT_0xDDCB+0x5b] = OpcodesDisMap[SHIFT_0xDDCB+0x5f]
	OpcodesDisMap[SHIFT_0xDDCB+0x5c] = OpcodesDisMap[SHIFT_0xDDCB+0x5f]
	OpcodesDisMap[SHIFT_0xDDCB+0x5d] = OpcodesDisMap[SHIFT_0xDDCB+0x5f]
	OpcodesDisMap[SHIFT_0xDDCB+0x5e] = OpcodesDisMap[SHIFT_0xDDCB+0x5f]
	/* BIT 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x67] = dis_instrDDCB__BIT_4_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x60] = OpcodesDisMap[SHIFT_0xDDCB+0x67]
	OpcodesDisMap[SHIFT_0xDDCB+0x61] = OpcodesDisMap[SHIFT_0xDDCB+0x67]
	OpcodesDisMap[SHIFT_0xDDCB+0x62] = OpcodesDisMap[SHIFT_0xDDCB+0x67]
	OpcodesDisMap[SHIFT_0xDDCB+0x63] = OpcodesDisMap[SHIFT_0xDDCB+0x67]
	OpcodesDisMap[SHIFT_0xDDCB+0x64] = OpcodesDisMap[SHIFT_0xDDCB+0x67]
	OpcodesDisMap[SHIFT_0xDDCB+0x65] = OpcodesDisMap[SHIFT_0xDDCB+0x67]
	OpcodesDisMap[SHIFT_0xDDCB+0x66] = OpcodesDisMap[SHIFT_0xDDCB+0x67]
	/* BIT 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x6f] = dis_instrDDCB__BIT_5_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x68] = OpcodesDisMap[SHIFT_0xDDCB+0x6f]
	OpcodesDisMap[SHIFT_0xDDCB+0x69] = OpcodesDisMap[SHIFT_0xDDCB+0x6f]
	OpcodesDisMap[SHIFT_0xDDCB+0x6a] = OpcodesDisMap[SHIFT_0xDDCB+0x6f]
	OpcodesDisMap[SHIFT_0xDDCB+0x6b] = OpcodesDisMap[SHIFT_0xDDCB+0x6f]
	OpcodesDisMap[SHIFT_0xDDCB+0x6c] = OpcodesDisMap[SHIFT_0xDDCB+0x6f]
	OpcodesDisMap[SHIFT_0xDDCB+0x6d] = OpcodesDisMap[SHIFT_0xDDCB+0x6f]
	OpcodesDisMap[SHIFT_0xDDCB+0x6e] = OpcodesDisMap[SHIFT_0xDDCB+0x6f]
	/* BIT 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x77] = dis_instrDDCB__BIT_6_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x70] = OpcodesDisMap[SHIFT_0xDDCB+0x77]
	OpcodesDisMap[SHIFT_0xDDCB+0x71] = OpcodesDisMap[SHIFT_0xDDCB+0x77]
	OpcodesDisMap[SHIFT_0xDDCB+0x72] = OpcodesDisMap[SHIFT_0xDDCB+0x77]
	OpcodesDisMap[SHIFT_0xDDCB+0x73] = OpcodesDisMap[SHIFT_0xDDCB+0x77]
	OpcodesDisMap[SHIFT_0xDDCB+0x74] = OpcodesDisMap[SHIFT_0xDDCB+0x77]
	OpcodesDisMap[SHIFT_0xDDCB+0x75] = OpcodesDisMap[SHIFT_0xDDCB+0x77]
	OpcodesDisMap[SHIFT_0xDDCB+0x76] = OpcodesDisMap[SHIFT_0xDDCB+0x77]
	/* BIT 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x7f] = dis_instrDDCB__BIT_7_iREGpDD
	// Fallthrough cases
	OpcodesDisMap[SHIFT_0xDDCB+0x78] = OpcodesDisMap[SHIFT_0xDDCB+0x7f]
	OpcodesDisMap[SHIFT_0xDDCB+0x79] = OpcodesDisMap[SHIFT_0xDDCB+0x7f]
	OpcodesDisMap[SHIFT_0xDDCB+0x7a] = OpcodesDisMap[SHIFT_0xDDCB+0x7f]
	OpcodesDisMap[SHIFT_0xDDCB+0x7b] = OpcodesDisMap[SHIFT_0xDDCB+0x7f]
	OpcodesDisMap[SHIFT_0xDDCB+0x7c] = OpcodesDisMap[SHIFT_0xDDCB+0x7f]
	OpcodesDisMap[SHIFT_0xDDCB+0x7d] = OpcodesDisMap[SHIFT_0xDDCB+0x7f]
	OpcodesDisMap[SHIFT_0xDDCB+0x7e] = OpcodesDisMap[SHIFT_0xDDCB+0x7f]
	/* LD B,RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x80] = dis_instrDDCB__LD_B_RES_0_iREGpDD
	/* LD C,RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x81] = dis_instrDDCB__LD_C_RES_0_iREGpDD
	/* LD D,RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x82] = dis_instrDDCB__LD_D_RES_0_iREGpDD
	/* LD E,RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x83] = dis_instrDDCB__LD_E_RES_0_iREGpDD
	/* LD H,RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x84] = dis_instrDDCB__LD_H_RES_0_iREGpDD
	/* LD L,RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x85] = dis_instrDDCB__LD_L_RES_0_iREGpDD
	/* RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x86] = dis_instrDDCB__RES_0_iREGpDD
	/* LD A,RES 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x87] = dis_instrDDCB__LD_A_RES_0_iREGpDD
	/* LD B,RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x88] = dis_instrDDCB__LD_B_RES_1_iREGpDD
	/* LD C,RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x89] = dis_instrDDCB__LD_C_RES_1_iREGpDD
	/* LD D,RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x8a] = dis_instrDDCB__LD_D_RES_1_iREGpDD
	/* LD E,RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x8b] = dis_instrDDCB__LD_E_RES_1_iREGpDD
	/* LD H,RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x8c] = dis_instrDDCB__LD_H_RES_1_iREGpDD
	/* LD L,RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x8d] = dis_instrDDCB__LD_L_RES_1_iREGpDD
	/* RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x8e] = dis_instrDDCB__RES_1_iREGpDD
	/* LD A,RES 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x8f] = dis_instrDDCB__LD_A_RES_1_iREGpDD
	/* LD B,RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x90] = dis_instrDDCB__LD_B_RES_2_iREGpDD
	/* LD C,RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x91] = dis_instrDDCB__LD_C_RES_2_iREGpDD
	/* LD D,RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x92] = dis_instrDDCB__LD_D_RES_2_iREGpDD
	/* LD E,RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x93] = dis_instrDDCB__LD_E_RES_2_iREGpDD
	/* LD H,RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x94] = dis_instrDDCB__LD_H_RES_2_iREGpDD
	/* LD L,RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x95] = dis_instrDDCB__LD_L_RES_2_iREGpDD
	/* RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x96] = dis_instrDDCB__RES_2_iREGpDD
	/* LD A,RES 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x97] = dis_instrDDCB__LD_A_RES_2_iREGpDD
	/* LD B,RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x98] = dis_instrDDCB__LD_B_RES_3_iREGpDD
	/* LD C,RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x99] = dis_instrDDCB__LD_C_RES_3_iREGpDD
	/* LD D,RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x9a] = dis_instrDDCB__LD_D_RES_3_iREGpDD
	/* LD E,RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x9b] = dis_instrDDCB__LD_E_RES_3_iREGpDD
	/* LD H,RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x9c] = dis_instrDDCB__LD_H_RES_3_iREGpDD
	/* LD L,RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x9d] = dis_instrDDCB__LD_L_RES_3_iREGpDD
	/* RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x9e] = dis_instrDDCB__RES_3_iREGpDD
	/* LD A,RES 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0x9f] = dis_instrDDCB__LD_A_RES_3_iREGpDD
	/* LD B,RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa0] = dis_instrDDCB__LD_B_RES_4_iREGpDD
	/* LD C,RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa1] = dis_instrDDCB__LD_C_RES_4_iREGpDD
	/* LD D,RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa2] = dis_instrDDCB__LD_D_RES_4_iREGpDD
	/* LD E,RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa3] = dis_instrDDCB__LD_E_RES_4_iREGpDD
	/* LD H,RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa4] = dis_instrDDCB__LD_H_RES_4_iREGpDD
	/* LD L,RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa5] = dis_instrDDCB__LD_L_RES_4_iREGpDD
	/* RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa6] = dis_instrDDCB__RES_4_iREGpDD
	/* LD A,RES 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa7] = dis_instrDDCB__LD_A_RES_4_iREGpDD
	/* LD B,RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa8] = dis_instrDDCB__LD_B_RES_5_iREGpDD
	/* LD C,RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xa9] = dis_instrDDCB__LD_C_RES_5_iREGpDD
	/* LD D,RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xaa] = dis_instrDDCB__LD_D_RES_5_iREGpDD
	/* LD E,RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xab] = dis_instrDDCB__LD_E_RES_5_iREGpDD
	/* LD H,RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xac] = dis_instrDDCB__LD_H_RES_5_iREGpDD
	/* LD L,RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xad] = dis_instrDDCB__LD_L_RES_5_iREGpDD
	/* RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xae] = dis_instrDDCB__RES_5_iREGpDD
	/* LD A,RES 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xaf] = dis_instrDDCB__LD_A_RES_5_iREGpDD
	/* LD B,RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb0] = dis_instrDDCB__LD_B_RES_6_iREGpDD
	/* LD C,RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb1] = dis_instrDDCB__LD_C_RES_6_iREGpDD
	/* LD D,RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb2] = dis_instrDDCB__LD_D_RES_6_iREGpDD
	/* LD E,RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb3] = dis_instrDDCB__LD_E_RES_6_iREGpDD
	/* LD H,RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb4] = dis_instrDDCB__LD_H_RES_6_iREGpDD
	/* LD L,RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb5] = dis_instrDDCB__LD_L_RES_6_iREGpDD
	/* RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb6] = dis_instrDDCB__RES_6_iREGpDD
	/* LD A,RES 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb7] = dis_instrDDCB__LD_A_RES_6_iREGpDD
	/* LD B,RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb8] = dis_instrDDCB__LD_B_RES_7_iREGpDD
	/* LD C,RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xb9] = dis_instrDDCB__LD_C_RES_7_iREGpDD
	/* LD D,RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xba] = dis_instrDDCB__LD_D_RES_7_iREGpDD
	/* LD E,RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xbb] = dis_instrDDCB__LD_E_RES_7_iREGpDD
	/* LD H,RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xbc] = dis_instrDDCB__LD_H_RES_7_iREGpDD
	/* LD L,RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xbd] = dis_instrDDCB__LD_L_RES_7_iREGpDD
	/* RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xbe] = dis_instrDDCB__RES_7_iREGpDD
	/* LD A,RES 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xbf] = dis_instrDDCB__LD_A_RES_7_iREGpDD
	/* LD B,SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc0] = dis_instrDDCB__LD_B_SET_0_iREGpDD
	/* LD C,SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc1] = dis_instrDDCB__LD_C_SET_0_iREGpDD
	/* LD D,SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc2] = dis_instrDDCB__LD_D_SET_0_iREGpDD
	/* LD E,SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc3] = dis_instrDDCB__LD_E_SET_0_iREGpDD
	/* LD H,SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc4] = dis_instrDDCB__LD_H_SET_0_iREGpDD
	/* LD L,SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc5] = dis_instrDDCB__LD_L_SET_0_iREGpDD
	/* SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc6] = dis_instrDDCB__SET_0_iREGpDD
	/* LD A,SET 0,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc7] = dis_instrDDCB__LD_A_SET_0_iREGpDD
	/* LD B,SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc8] = dis_instrDDCB__LD_B_SET_1_iREGpDD
	/* LD C,SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xc9] = dis_instrDDCB__LD_C_SET_1_iREGpDD
	/* LD D,SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xca] = dis_instrDDCB__LD_D_SET_1_iREGpDD
	/* LD E,SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xcb] = dis_instrDDCB__LD_E_SET_1_iREGpDD
	/* LD H,SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xcc] = dis_instrDDCB__LD_H_SET_1_iREGpDD
	/* LD L,SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xcd] = dis_instrDDCB__LD_L_SET_1_iREGpDD
	/* SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xce] = dis_instrDDCB__SET_1_iREGpDD
	/* LD A,SET 1,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xcf] = dis_instrDDCB__LD_A_SET_1_iREGpDD
	/* LD B,SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd0] = dis_instrDDCB__LD_B_SET_2_iREGpDD
	/* LD C,SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd1] = dis_instrDDCB__LD_C_SET_2_iREGpDD
	/* LD D,SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd2] = dis_instrDDCB__LD_D_SET_2_iREGpDD
	/* LD E,SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd3] = dis_instrDDCB__LD_E_SET_2_iREGpDD
	/* LD H,SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd4] = dis_instrDDCB__LD_H_SET_2_iREGpDD
	/* LD L,SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd5] = dis_instrDDCB__LD_L_SET_2_iREGpDD
	/* SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd6] = dis_instrDDCB__SET_2_iREGpDD
	/* LD A,SET 2,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd7] = dis_instrDDCB__LD_A_SET_2_iREGpDD
	/* LD B,SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd8] = dis_instrDDCB__LD_B_SET_3_iREGpDD
	/* LD C,SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xd9] = dis_instrDDCB__LD_C_SET_3_iREGpDD
	/* LD D,SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xda] = dis_instrDDCB__LD_D_SET_3_iREGpDD
	/* LD E,SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xdb] = dis_instrDDCB__LD_E_SET_3_iREGpDD
	/* LD H,SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xdc] = dis_instrDDCB__LD_H_SET_3_iREGpDD
	/* LD L,SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xdd] = dis_instrDDCB__LD_L_SET_3_iREGpDD
	/* SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xde] = dis_instrDDCB__SET_3_iREGpDD
	/* LD A,SET 3,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xdf] = dis_instrDDCB__LD_A_SET_3_iREGpDD
	/* LD B,SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe0] = dis_instrDDCB__LD_B_SET_4_iREGpDD
	/* LD C,SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe1] = dis_instrDDCB__LD_C_SET_4_iREGpDD
	/* LD D,SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe2] = dis_instrDDCB__LD_D_SET_4_iREGpDD
	/* LD E,SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe3] = dis_instrDDCB__LD_E_SET_4_iREGpDD
	/* LD H,SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe4] = dis_instrDDCB__LD_H_SET_4_iREGpDD
	/* LD L,SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe5] = dis_instrDDCB__LD_L_SET_4_iREGpDD
	/* SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe6] = dis_instrDDCB__SET_4_iREGpDD
	/* LD A,SET 4,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe7] = dis_instrDDCB__LD_A_SET_4_iREGpDD
	/* LD B,SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe8] = dis_instrDDCB__LD_B_SET_5_iREGpDD
	/* LD C,SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xe9] = dis_instrDDCB__LD_C_SET_5_iREGpDD
	/* LD D,SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xea] = dis_instrDDCB__LD_D_SET_5_iREGpDD
	/* LD E,SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xeb] = dis_instrDDCB__LD_E_SET_5_iREGpDD
	/* LD H,SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xec] = dis_instrDDCB__LD_H_SET_5_iREGpDD
	/* LD L,SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xed] = dis_instrDDCB__LD_L_SET_5_iREGpDD
	/* SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xee] = dis_instrDDCB__SET_5_iREGpDD
	/* LD A,SET 5,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xef] = dis_instrDDCB__LD_A_SET_5_iREGpDD
	/* LD B,SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf0] = dis_instrDDCB__LD_B_SET_6_iREGpDD
	/* LD C,SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf1] = dis_instrDDCB__LD_C_SET_6_iREGpDD
	/* LD D,SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf2] = dis_instrDDCB__LD_D_SET_6_iREGpDD
	/* LD E,SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf3] = dis_instrDDCB__LD_E_SET_6_iREGpDD
	/* LD H,SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf4] = dis_instrDDCB__LD_H_SET_6_iREGpDD
	/* LD L,SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf5] = dis_instrDDCB__LD_L_SET_6_iREGpDD
	/* SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf6] = dis_instrDDCB__SET_6_iREGpDD
	/* LD A,SET 6,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf7] = dis_instrDDCB__LD_A_SET_6_iREGpDD
	/* LD B,SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf8] = dis_instrDDCB__LD_B_SET_7_iREGpDD
	/* LD C,SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xf9] = dis_instrDDCB__LD_C_SET_7_iREGpDD
	/* LD D,SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xfa] = dis_instrDDCB__LD_D_SET_7_iREGpDD
	/* LD E,SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xfb] = dis_instrDDCB__LD_E_SET_7_iREGpDD
	/* LD H,SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xfc] = dis_instrDDCB__LD_H_SET_7_iREGpDD
	/* LD L,SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xfd] = dis_instrDDCB__LD_L_SET_7_iREGpDD
	/* SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xfe] = dis_instrDDCB__SET_7_iREGpDD
	/* LD A,SET 7,(REGISTER+dd) */
	OpcodesDisMap[SHIFT_0xDDCB+0xff] = dis_instrDDCB__LD_A_SET_7_iREGpDD

	// END of 0xddfdcb shifted opcodes
}

/* NOP */
func dis_instr__NOP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "NOP "
	shift = 0
	return result, address + 1, shift
}

/* LD BC,nnnn */
func dis_instr__LD_BC_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "BC" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* LD (BC),A */
func dis_instr__LD_iBC_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(BC)" + "," + "A"
	return result, address + 1, shift
}

/* INC BC */
func dis_instr__INC_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "BC"
	return result, address + 1, shift
}

/* INC B */
func dis_instr__INC_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* DEC B */
func dis_instr__DEC_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* LD B,nn */
func dis_instr__LD_B_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "B" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RLCA */
func dis_instr__RLCA(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLCA "
	shift = 0
	return result, address + 1, shift
}

/* EX AF,AF' */
func dis_instr__EX_AF_AF(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "EX "
	shift = 0
	result += "AF" + "," + "AF'"
	return result, address + 1, shift
}

/* ADD HL,BC */
func dis_instr__ADD_HL_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "HL" + "," + "BC"
	return result, address + 1, shift
}

/* LD A,(BC) */
func dis_instr__LD_A_iBC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "(BC)"
	return result, address + 1, shift
}

/* DEC BC */
func dis_instr__DEC_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "BC"
	return result, address + 1, shift
}

/* INC C */
func dis_instr__INC_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* DEC C */
func dis_instr__DEC_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* LD C,nn */
func dis_instr__LD_C_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "C" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RRCA */
func dis_instr__RRCA(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRCA "
	shift = 0
	return result, address + 1, shift
}

/* DJNZ offset */
func dis_instr__DJNZ_OFFSET(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DJNZ "
	shift = 0
	address++
	offset := memory.ReadByte(address)
	result += fmt.Sprintf("0x%x", offset)
	return result, address + 1, shift
}

/* LD DE,nnnn */
func dis_instr__LD_DE_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "DE" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* LD (DE),A */
func dis_instr__LD_iDE_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(DE)" + "," + "A"
	return result, address + 1, shift
}

/* INC DE */
func dis_instr__INC_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "DE"
	return result, address + 1, shift
}

/* INC D */
func dis_instr__INC_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* DEC D */
func dis_instr__DEC_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* LD D,nn */
func dis_instr__LD_D_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "D" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RLA */
func dis_instr__RLA(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLA "
	shift = 0
	return result, address + 1, shift
}

/* JR offset */
func dis_instr__JR_OFFSET(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JR "
	shift = 0
	address++
	offset := memory.ReadByte(address)
	result += fmt.Sprintf("0x%x", offset)
	return result, address + 1, shift
}

/* ADD HL,DE */
func dis_instr__ADD_HL_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "HL" + "," + "DE"
	return result, address + 1, shift
}

/* LD A,(DE) */
func dis_instr__LD_A_iDE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "(DE)"
	return result, address + 1, shift
}

/* DEC DE */
func dis_instr__DEC_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "DE"
	return result, address + 1, shift
}

/* INC E */
func dis_instr__INC_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* DEC E */
func dis_instr__DEC_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* LD E,nn */
func dis_instr__LD_E_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "E" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RRA */
func dis_instr__RRA(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRA "
	shift = 0
	return result, address + 1, shift
}

/* JR NZ,offset */
func dis_instr__JR_NZ_OFFSET(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JR "
	shift = 0
	address++
	offset := memory.ReadByte(address)
	result += "NZ" + "," + fmt.Sprintf("0x%x", offset)
	return result, address + 1, shift
}

/* LD HL,nnnn */
func dis_instr__LD_HL_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "HL" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* LD (nnnn),HL */
func dis_instr__LD_iNNNN_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "HL"
	return result, address + 1, shift
}

/* INC HL */
func dis_instr__INC_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "HL"
	return result, address + 1, shift
}

/* INC H */
func dis_instr__INC_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* DEC H */
func dis_instr__DEC_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* LD H,nn */
func dis_instr__LD_H_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "H" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* DAA */
func dis_instr__DAA(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DAA "
	shift = 0
	return result, address + 1, shift
}

/* JR Z,offset */
func dis_instr__JR_Z_OFFSET(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JR "
	shift = 0
	address++
	offset := memory.ReadByte(address)
	result += "Z" + "," + fmt.Sprintf("0x%x", offset)
	return result, address + 1, shift
}

/* ADD HL,HL */
func dis_instr__ADD_HL_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "HL" + "," + "HL"
	return result, address + 1, shift
}

/* LD HL,(nnnn) */
func dis_instr__LD_HL_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "HL" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* DEC HL */
func dis_instr__DEC_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "HL"
	return result, address + 1, shift
}

/* INC L */
func dis_instr__INC_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* DEC L */
func dis_instr__DEC_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* LD L,nn */
func dis_instr__LD_L_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "L" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* CPL */
func dis_instr__CPL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CPL "
	shift = 0
	return result, address + 1, shift
}

/* JR NC,offset */
func dis_instr__JR_NC_OFFSET(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JR "
	shift = 0
	address++
	offset := memory.ReadByte(address)
	result += "NC" + "," + fmt.Sprintf("0x%x", offset)
	return result, address + 1, shift
}

/* LD SP,nnnn */
func dis_instr__LD_SP_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "SP" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* LD (nnnn),A */
func dis_instr__LD_iNNNN_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "A"
	return result, address + 1, shift
}

/* INC SP */
func dis_instr__INC_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "SP"
	return result, address + 1, shift
}

/* INC (HL) */
func dis_instr__INC_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* DEC (HL) */
func dis_instr__DEC_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* LD (HL),nn */
func dis_instr__LD_iHL_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "(HL)" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* SCF */
func dis_instr__SCF(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SCF "
	shift = 0
	return result, address + 1, shift
}

/* JR C,offset */
func dis_instr__JR_C_OFFSET(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JR "
	shift = 0
	address++
	offset := memory.ReadByte(address)
	result += "C" + "," + fmt.Sprintf("0x%x", offset)
	return result, address + 1, shift
}

/* ADD HL,SP */
func dis_instr__ADD_HL_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "HL" + "," + "SP"
	return result, address + 1, shift
}

/* LD A,(nnnn) */
func dis_instr__LD_A_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "A" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* DEC SP */
func dis_instr__DEC_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "SP"
	return result, address + 1, shift
}

/* INC A */
func dis_instr__INC_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* DEC A */
func dis_instr__DEC_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* LD A,nn */
func dis_instr__LD_A_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* CCF */
func dis_instr__CCF(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CCF "
	shift = 0
	return result, address + 1, shift
}

/* LD B,B */
func dis_instr__LD_B_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "B"
	return result, address + 1, shift
}

/* LD B,C */
func dis_instr__LD_B_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "C"
	return result, address + 1, shift
}

/* LD B,D */
func dis_instr__LD_B_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "D"
	return result, address + 1, shift
}

/* LD B,E */
func dis_instr__LD_B_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "E"
	return result, address + 1, shift
}

/* LD B,H */
func dis_instr__LD_B_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "H"
	return result, address + 1, shift
}

/* LD B,L */
func dis_instr__LD_B_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "L"
	return result, address + 1, shift
}

/* LD B,(HL) */
func dis_instr__LD_B_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "(HL)"
	return result, address + 1, shift
}

/* LD B,A */
func dis_instr__LD_B_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "A"
	return result, address + 1, shift
}

/* LD C,B */
func dis_instr__LD_C_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "B"
	return result, address + 1, shift
}

/* LD C,C */
func dis_instr__LD_C_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "C"
	return result, address + 1, shift
}

/* LD C,D */
func dis_instr__LD_C_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "D"
	return result, address + 1, shift
}

/* LD C,E */
func dis_instr__LD_C_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "E"
	return result, address + 1, shift
}

/* LD C,H */
func dis_instr__LD_C_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "H"
	return result, address + 1, shift
}

/* LD C,L */
func dis_instr__LD_C_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "L"
	return result, address + 1, shift
}

/* LD C,(HL) */
func dis_instr__LD_C_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "(HL)"
	return result, address + 1, shift
}

/* LD C,A */
func dis_instr__LD_C_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "A"
	return result, address + 1, shift
}

/* LD D,B */
func dis_instr__LD_D_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "B"
	return result, address + 1, shift
}

/* LD D,C */
func dis_instr__LD_D_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "C"
	return result, address + 1, shift
}

/* LD D,D */
func dis_instr__LD_D_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "D"
	return result, address + 1, shift
}

/* LD D,E */
func dis_instr__LD_D_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "E"
	return result, address + 1, shift
}

/* LD D,H */
func dis_instr__LD_D_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "H"
	return result, address + 1, shift
}

/* LD D,L */
func dis_instr__LD_D_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "L"
	return result, address + 1, shift
}

/* LD D,(HL) */
func dis_instr__LD_D_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "(HL)"
	return result, address + 1, shift
}

/* LD D,A */
func dis_instr__LD_D_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "A"
	return result, address + 1, shift
}

/* LD E,B */
func dis_instr__LD_E_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "B"
	return result, address + 1, shift
}

/* LD E,C */
func dis_instr__LD_E_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "C"
	return result, address + 1, shift
}

/* LD E,D */
func dis_instr__LD_E_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "D"
	return result, address + 1, shift
}

/* LD E,E */
func dis_instr__LD_E_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "E"
	return result, address + 1, shift
}

/* LD E,H */
func dis_instr__LD_E_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "H"
	return result, address + 1, shift
}

/* LD E,L */
func dis_instr__LD_E_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "L"
	return result, address + 1, shift
}

/* LD E,(HL) */
func dis_instr__LD_E_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "(HL)"
	return result, address + 1, shift
}

/* LD E,A */
func dis_instr__LD_E_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "A"
	return result, address + 1, shift
}

/* LD H,B */
func dis_instr__LD_H_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "B"
	return result, address + 1, shift
}

/* LD H,C */
func dis_instr__LD_H_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "C"
	return result, address + 1, shift
}

/* LD H,D */
func dis_instr__LD_H_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "D"
	return result, address + 1, shift
}

/* LD H,E */
func dis_instr__LD_H_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "E"
	return result, address + 1, shift
}

/* LD H,H */
func dis_instr__LD_H_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "H"
	return result, address + 1, shift
}

/* LD H,L */
func dis_instr__LD_H_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "L"
	return result, address + 1, shift
}

/* LD H,(HL) */
func dis_instr__LD_H_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "(HL)"
	return result, address + 1, shift
}

/* LD H,A */
func dis_instr__LD_H_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "A"
	return result, address + 1, shift
}

/* LD L,B */
func dis_instr__LD_L_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "B"
	return result, address + 1, shift
}

/* LD L,C */
func dis_instr__LD_L_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "C"
	return result, address + 1, shift
}

/* LD L,D */
func dis_instr__LD_L_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "D"
	return result, address + 1, shift
}

/* LD L,E */
func dis_instr__LD_L_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "E"
	return result, address + 1, shift
}

/* LD L,H */
func dis_instr__LD_L_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "H"
	return result, address + 1, shift
}

/* LD L,L */
func dis_instr__LD_L_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "L"
	return result, address + 1, shift
}

/* LD L,(HL) */
func dis_instr__LD_L_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "(HL)"
	return result, address + 1, shift
}

/* LD L,A */
func dis_instr__LD_L_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "A"
	return result, address + 1, shift
}

/* LD (HL),B */
func dis_instr__LD_iHL_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(HL)" + "," + "B"
	return result, address + 1, shift
}

/* LD (HL),C */
func dis_instr__LD_iHL_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(HL)" + "," + "C"
	return result, address + 1, shift
}

/* LD (HL),D */
func dis_instr__LD_iHL_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(HL)" + "," + "D"
	return result, address + 1, shift
}

/* LD (HL),E */
func dis_instr__LD_iHL_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(HL)" + "," + "E"
	return result, address + 1, shift
}

/* LD (HL),H */
func dis_instr__LD_iHL_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(HL)" + "," + "H"
	return result, address + 1, shift
}

/* LD (HL),L */
func dis_instr__LD_iHL_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(HL)" + "," + "L"
	return result, address + 1, shift
}

/* HALT */
func dis_instr__HALT(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "HALT "
	shift = 0
	return result, address + 1, shift
}

/* LD (HL),A */
func dis_instr__LD_iHL_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "(HL)" + "," + "A"
	return result, address + 1, shift
}

/* LD A,B */
func dis_instr__LD_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* LD A,C */
func dis_instr__LD_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* LD A,D */
func dis_instr__LD_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* LD A,E */
func dis_instr__LD_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* LD A,H */
func dis_instr__LD_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* LD A,L */
func dis_instr__LD_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* LD A,(HL) */
func dis_instr__LD_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* LD A,A */
func dis_instr__LD_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* ADD A,B */
func dis_instr__ADD_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* ADD A,C */
func dis_instr__ADD_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* ADD A,D */
func dis_instr__ADD_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* ADD A,E */
func dis_instr__ADD_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* ADD A,H */
func dis_instr__ADD_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* ADD A,L */
func dis_instr__ADD_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* ADD A,(HL) */
func dis_instr__ADD_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* ADD A,A */
func dis_instr__ADD_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* ADC A,B */
func dis_instr__ADC_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* ADC A,C */
func dis_instr__ADC_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* ADC A,D */
func dis_instr__ADC_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* ADC A,E */
func dis_instr__ADC_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* ADC A,H */
func dis_instr__ADC_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* ADC A,L */
func dis_instr__ADC_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* ADC A,(HL) */
func dis_instr__ADC_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* ADC A,A */
func dis_instr__ADC_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* SUB A,B */
func dis_instr__SUB_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* SUB A,C */
func dis_instr__SUB_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* SUB A,D */
func dis_instr__SUB_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* SUB A,E */
func dis_instr__SUB_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* SUB A,H */
func dis_instr__SUB_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* SUB A,L */
func dis_instr__SUB_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* SUB A,(HL) */
func dis_instr__SUB_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* SUB A,A */
func dis_instr__SUB_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* SBC A,B */
func dis_instr__SBC_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* SBC A,C */
func dis_instr__SBC_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* SBC A,D */
func dis_instr__SBC_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* SBC A,E */
func dis_instr__SBC_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* SBC A,H */
func dis_instr__SBC_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* SBC A,L */
func dis_instr__SBC_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* SBC A,(HL) */
func dis_instr__SBC_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* SBC A,A */
func dis_instr__SBC_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* AND A,B */
func dis_instr__AND_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* AND A,C */
func dis_instr__AND_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* AND A,D */
func dis_instr__AND_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* AND A,E */
func dis_instr__AND_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* AND A,H */
func dis_instr__AND_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* AND A,L */
func dis_instr__AND_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* AND A,(HL) */
func dis_instr__AND_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* AND A,A */
func dis_instr__AND_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* XOR A,B */
func dis_instr__XOR_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* XOR A,C */
func dis_instr__XOR_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* XOR A,D */
func dis_instr__XOR_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* XOR A,E */
func dis_instr__XOR_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* XOR A,H */
func dis_instr__XOR_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* XOR A,L */
func dis_instr__XOR_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* XOR A,(HL) */
func dis_instr__XOR_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* XOR A,A */
func dis_instr__XOR_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* OR A,B */
func dis_instr__OR_A_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "B"
	return result, address + 1, shift
}

/* OR A,C */
func dis_instr__OR_A_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "C"
	return result, address + 1, shift
}

/* OR A,D */
func dis_instr__OR_A_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "D"
	return result, address + 1, shift
}

/* OR A,E */
func dis_instr__OR_A_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "E"
	return result, address + 1, shift
}

/* OR A,H */
func dis_instr__OR_A_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "H"
	return result, address + 1, shift
}

/* OR A,L */
func dis_instr__OR_A_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "L"
	return result, address + 1, shift
}

/* OR A,(HL) */
func dis_instr__OR_A_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "(HL)"
	return result, address + 1, shift
}

/* OR A,A */
func dis_instr__OR_A_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "A"
	return result, address + 1, shift
}

/* CP B */
func dis_instr__CP_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* CP C */
func dis_instr__CP_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* CP D */
func dis_instr__CP_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* CP E */
func dis_instr__CP_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* CP H */
func dis_instr__CP_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* CP L */
func dis_instr__CP_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* CP (HL) */
func dis_instr__CP_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* CP A */
func dis_instr__CP_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* RET NZ */
func dis_instr__RET_NZ(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "NZ"
	return result, address + 1, shift
}

/* POP BC */
func dis_instr__POP_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "POP "
	shift = 0
	result += "BC"
	return result, address + 1, shift
}

/* JP NZ,nnnn */
func dis_instr__JP_NZ_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "NZ" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* JP nnnn */
func dis_instr__JP_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* CALL NZ,nnnn */
func dis_instr__CALL_NZ_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "NZ" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* PUSH BC */
func dis_instr__PUSH_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "PUSH "
	shift = 0
	result += "BC"
	return result, address + 1, shift
}

/* ADD A,nn */
func dis_instr__ADD_A_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 00 */
func dis_instr__RST_00(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "00"
	return result, address + 1, shift
}

/* RET Z */
func dis_instr__RET_Z(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "Z"
	return result, address + 1, shift
}

/* RET */
func dis_instr__RET(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	return result, address + 1, shift
}

/* JP Z,nnnn */
func dis_instr__JP_Z_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "Z" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* shift CB */
func dis_instr__SHIFT_CB(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "shift "
	shift = SHIFT_0xCB
	result += "CB"
	return result, address + 1, shift
}

/* CALL Z,nnnn */
func dis_instr__CALL_Z_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "Z" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* CALL nnnn */
func dis_instr__CALL_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* ADC A,nn */
func dis_instr__ADC_A_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 8 */
func dis_instr__RST_8(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "8"
	return result, address + 1, shift
}

/* RET NC */
func dis_instr__RET_NC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "NC"
	return result, address + 1, shift
}

/* POP DE */
func dis_instr__POP_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "POP "
	shift = 0
	result += "DE"
	return result, address + 1, shift
}

/* JP NC,nnnn */
func dis_instr__JP_NC_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "NC" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* OUT (nn),A */
func dis_instr__OUT_iNN_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(nn)" + "," + "A"
	return result, address + 1, shift
}

/* CALL NC,nnnn */
func dis_instr__CALL_NC_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "NC" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* PUSH DE */
func dis_instr__PUSH_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "PUSH "
	shift = 0
	result += "DE"
	return result, address + 1, shift
}

/* SUB nn */
func dis_instr__SUB_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 10 */
func dis_instr__RST_10(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "10"
	return result, address + 1, shift
}

/* RET C */
func dis_instr__RET_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* EXX */
func dis_instr__EXX(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "EXX "
	shift = 0
	return result, address + 1, shift
}

/* JP C,nnnn */
func dis_instr__JP_C_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "C" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* IN A,(nn) */
func dis_instr__IN_A_iNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "A" + "," + "(nn)"
	return result, address + 1, shift
}

/* CALL C,nnnn */
func dis_instr__CALL_C_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "C" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* shift DD */
func dis_instr__SHIFT_DD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "shift "
	shift = SHIFT_0xDD
	result += "DD"
	return result, address + 1, shift
}

/* SBC A,nn */
func dis_instr__SBC_A_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 18 */
func dis_instr__RST_18(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "18"
	return result, address + 1, shift
}

/* RET PO */
func dis_instr__RET_PO(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "PO"
	return result, address + 1, shift
}

/* POP HL */
func dis_instr__POP_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "POP "
	shift = 0
	result += "HL"
	return result, address + 1, shift
}

/* JP PO,nnnn */
func dis_instr__JP_PO_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "PO" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* EX (SP),HL */
func dis_instr__EX_iSP_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "EX "
	shift = 0
	result += "(SP)" + "," + "HL"
	return result, address + 1, shift
}

/* CALL PO,nnnn */
func dis_instr__CALL_PO_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "PO" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* PUSH HL */
func dis_instr__PUSH_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "PUSH "
	shift = 0
	result += "HL"
	return result, address + 1, shift
}

/* AND nn */
func dis_instr__AND_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 20 */
func dis_instr__RST_20(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "20"
	return result, address + 1, shift
}

/* RET PE */
func dis_instr__RET_PE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "PE"
	return result, address + 1, shift
}

/* JP HL */
func dis_instr__JP_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	result += "HL"
	return result, address + 1, shift
}

/* JP PE,nnnn */
func dis_instr__JP_PE_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "PE" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* EX DE,HL */
func dis_instr__EX_DE_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "EX "
	shift = 0
	result += "DE" + "," + "HL"
	return result, address + 1, shift
}

/* CALL PE,nnnn */
func dis_instr__CALL_PE_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "PE" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* shift ED */
func dis_instr__SHIFT_ED(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "shift "
	shift = SHIFT_0xED
	result += "ED"
	return result, address + 1, shift
}

/* XOR A,nn */
func dis_instr__XOR_A_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 28 */
func dis_instr__RST_28(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "28"
	return result, address + 1, shift
}

/* RET P */
func dis_instr__RET_P(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "P"
	return result, address + 1, shift
}

/* POP AF */
func dis_instr__POP_AF(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "POP "
	shift = 0
	result += "AF"
	return result, address + 1, shift
}

/* JP P,nnnn */
func dis_instr__JP_P_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "P" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* DI */
func dis_instr__DI(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DI "
	shift = 0
	return result, address + 1, shift
}

/* CALL P,nnnn */
func dis_instr__CALL_P_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "P" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* PUSH AF */
func dis_instr__PUSH_AF(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "PUSH "
	shift = 0
	result += "AF"
	return result, address + 1, shift
}

/* OR nn */
func dis_instr__OR_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 30 */
func dis_instr__RST_30(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "30"
	return result, address + 1, shift
}

/* RET M */
func dis_instr__RET_M(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RET "
	shift = 0
	result += "M"
	return result, address + 1, shift
}

/* LD SP,HL */
func dis_instr__LD_SP_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "SP" + "," + "HL"
	return result, address + 1, shift
}

/* JP M,nnnn */
func dis_instr__JP_M_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "M" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* EI */
func dis_instr__EI(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "EI "
	shift = 0
	return result, address + 1, shift
}

/* CALL M,nnnn */
func dis_instr__CALL_M_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CALL "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "M" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* shift FD */
func dis_instr__SHIFT_FD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "shift "
	shift = SHIFT_0xFD
	result += "FD"
	return result, address + 1, shift
}

/* CP nn */
func dis_instr__CP_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* RST 38 */
func dis_instr__RST_38(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RST "
	shift = 0
	result += "38"
	return result, address + 1, shift
}

/* RLC B */
func dis_instrCB__RLC_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* RLC C */
func dis_instrCB__RLC_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* RLC D */
func dis_instrCB__RLC_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* RLC E */
func dis_instrCB__RLC_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* RLC H */
func dis_instrCB__RLC_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* RLC L */
func dis_instrCB__RLC_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* RLC (HL) */
func dis_instrCB__RLC_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* RLC A */
func dis_instrCB__RLC_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* RRC B */
func dis_instrCB__RRC_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* RRC C */
func dis_instrCB__RRC_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* RRC D */
func dis_instrCB__RRC_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* RRC E */
func dis_instrCB__RRC_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* RRC H */
func dis_instrCB__RRC_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* RRC L */
func dis_instrCB__RRC_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* RRC (HL) */
func dis_instrCB__RRC_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* RRC A */
func dis_instrCB__RRC_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* RL B */
func dis_instrCB__RL_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* RL C */
func dis_instrCB__RL_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* RL D */
func dis_instrCB__RL_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* RL E */
func dis_instrCB__RL_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* RL H */
func dis_instrCB__RL_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* RL L */
func dis_instrCB__RL_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* RL (HL) */
func dis_instrCB__RL_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* RL A */
func dis_instrCB__RL_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* RR B */
func dis_instrCB__RR_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* RR C */
func dis_instrCB__RR_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* RR D */
func dis_instrCB__RR_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* RR E */
func dis_instrCB__RR_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* RR H */
func dis_instrCB__RR_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* RR L */
func dis_instrCB__RR_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* RR (HL) */
func dis_instrCB__RR_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* RR A */
func dis_instrCB__RR_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* SLA B */
func dis_instrCB__SLA_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* SLA C */
func dis_instrCB__SLA_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* SLA D */
func dis_instrCB__SLA_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* SLA E */
func dis_instrCB__SLA_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* SLA H */
func dis_instrCB__SLA_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* SLA L */
func dis_instrCB__SLA_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* SLA (HL) */
func dis_instrCB__SLA_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* SLA A */
func dis_instrCB__SLA_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* SRA B */
func dis_instrCB__SRA_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* SRA C */
func dis_instrCB__SRA_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* SRA D */
func dis_instrCB__SRA_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* SRA E */
func dis_instrCB__SRA_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* SRA H */
func dis_instrCB__SRA_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* SRA L */
func dis_instrCB__SRA_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* SRA (HL) */
func dis_instrCB__SRA_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* SRA A */
func dis_instrCB__SRA_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* SLL B */
func dis_instrCB__SLL_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* SLL C */
func dis_instrCB__SLL_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* SLL D */
func dis_instrCB__SLL_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* SLL E */
func dis_instrCB__SLL_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* SLL H */
func dis_instrCB__SLL_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* SLL L */
func dis_instrCB__SLL_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* SLL (HL) */
func dis_instrCB__SLL_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* SLL A */
func dis_instrCB__SLL_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* SRL B */
func dis_instrCB__SRL_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "B"
	return result, address + 1, shift
}

/* SRL C */
func dis_instrCB__SRL_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "C"
	return result, address + 1, shift
}

/* SRL D */
func dis_instrCB__SRL_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "D"
	return result, address + 1, shift
}

/* SRL E */
func dis_instrCB__SRL_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "E"
	return result, address + 1, shift
}

/* SRL H */
func dis_instrCB__SRL_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "H"
	return result, address + 1, shift
}

/* SRL L */
func dis_instrCB__SRL_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "L"
	return result, address + 1, shift
}

/* SRL (HL) */
func dis_instrCB__SRL_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "(HL)"
	return result, address + 1, shift
}

/* SRL A */
func dis_instrCB__SRL_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	result += "A"
	return result, address + 1, shift
}

/* BIT 0,B */
func dis_instrCB__BIT_0_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "B"
	return result, address + 1, shift
}

/* BIT 0,C */
func dis_instrCB__BIT_0_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "C"
	return result, address + 1, shift
}

/* BIT 0,D */
func dis_instrCB__BIT_0_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "D"
	return result, address + 1, shift
}

/* BIT 0,E */
func dis_instrCB__BIT_0_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "E"
	return result, address + 1, shift
}

/* BIT 0,H */
func dis_instrCB__BIT_0_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "H"
	return result, address + 1, shift
}

/* BIT 0,L */
func dis_instrCB__BIT_0_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "L"
	return result, address + 1, shift
}

/* BIT 0,(HL) */
func dis_instrCB__BIT_0_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 0,A */
func dis_instrCB__BIT_0_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "0" + "," + "A"
	return result, address + 1, shift
}

/* BIT 1,B */
func dis_instrCB__BIT_1_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "B"
	return result, address + 1, shift
}

/* BIT 1,C */
func dis_instrCB__BIT_1_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "C"
	return result, address + 1, shift
}

/* BIT 1,D */
func dis_instrCB__BIT_1_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "D"
	return result, address + 1, shift
}

/* BIT 1,E */
func dis_instrCB__BIT_1_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "E"
	return result, address + 1, shift
}

/* BIT 1,H */
func dis_instrCB__BIT_1_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "H"
	return result, address + 1, shift
}

/* BIT 1,L */
func dis_instrCB__BIT_1_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "L"
	return result, address + 1, shift
}

/* BIT 1,(HL) */
func dis_instrCB__BIT_1_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 1,A */
func dis_instrCB__BIT_1_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "1" + "," + "A"
	return result, address + 1, shift
}

/* BIT 2,B */
func dis_instrCB__BIT_2_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "B"
	return result, address + 1, shift
}

/* BIT 2,C */
func dis_instrCB__BIT_2_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "C"
	return result, address + 1, shift
}

/* BIT 2,D */
func dis_instrCB__BIT_2_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "D"
	return result, address + 1, shift
}

/* BIT 2,E */
func dis_instrCB__BIT_2_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "E"
	return result, address + 1, shift
}

/* BIT 2,H */
func dis_instrCB__BIT_2_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "H"
	return result, address + 1, shift
}

/* BIT 2,L */
func dis_instrCB__BIT_2_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "L"
	return result, address + 1, shift
}

/* BIT 2,(HL) */
func dis_instrCB__BIT_2_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 2,A */
func dis_instrCB__BIT_2_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "2" + "," + "A"
	return result, address + 1, shift
}

/* BIT 3,B */
func dis_instrCB__BIT_3_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "B"
	return result, address + 1, shift
}

/* BIT 3,C */
func dis_instrCB__BIT_3_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "C"
	return result, address + 1, shift
}

/* BIT 3,D */
func dis_instrCB__BIT_3_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "D"
	return result, address + 1, shift
}

/* BIT 3,E */
func dis_instrCB__BIT_3_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "E"
	return result, address + 1, shift
}

/* BIT 3,H */
func dis_instrCB__BIT_3_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "H"
	return result, address + 1, shift
}

/* BIT 3,L */
func dis_instrCB__BIT_3_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "L"
	return result, address + 1, shift
}

/* BIT 3,(HL) */
func dis_instrCB__BIT_3_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 3,A */
func dis_instrCB__BIT_3_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "3" + "," + "A"
	return result, address + 1, shift
}

/* BIT 4,B */
func dis_instrCB__BIT_4_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "B"
	return result, address + 1, shift
}

/* BIT 4,C */
func dis_instrCB__BIT_4_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "C"
	return result, address + 1, shift
}

/* BIT 4,D */
func dis_instrCB__BIT_4_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "D"
	return result, address + 1, shift
}

/* BIT 4,E */
func dis_instrCB__BIT_4_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "E"
	return result, address + 1, shift
}

/* BIT 4,H */
func dis_instrCB__BIT_4_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "H"
	return result, address + 1, shift
}

/* BIT 4,L */
func dis_instrCB__BIT_4_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "L"
	return result, address + 1, shift
}

/* BIT 4,(HL) */
func dis_instrCB__BIT_4_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 4,A */
func dis_instrCB__BIT_4_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "4" + "," + "A"
	return result, address + 1, shift
}

/* BIT 5,B */
func dis_instrCB__BIT_5_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "B"
	return result, address + 1, shift
}

/* BIT 5,C */
func dis_instrCB__BIT_5_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "C"
	return result, address + 1, shift
}

/* BIT 5,D */
func dis_instrCB__BIT_5_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "D"
	return result, address + 1, shift
}

/* BIT 5,E */
func dis_instrCB__BIT_5_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "E"
	return result, address + 1, shift
}

/* BIT 5,H */
func dis_instrCB__BIT_5_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "H"
	return result, address + 1, shift
}

/* BIT 5,L */
func dis_instrCB__BIT_5_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "L"
	return result, address + 1, shift
}

/* BIT 5,(HL) */
func dis_instrCB__BIT_5_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 5,A */
func dis_instrCB__BIT_5_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "5" + "," + "A"
	return result, address + 1, shift
}

/* BIT 6,B */
func dis_instrCB__BIT_6_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "B"
	return result, address + 1, shift
}

/* BIT 6,C */
func dis_instrCB__BIT_6_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "C"
	return result, address + 1, shift
}

/* BIT 6,D */
func dis_instrCB__BIT_6_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "D"
	return result, address + 1, shift
}

/* BIT 6,E */
func dis_instrCB__BIT_6_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "E"
	return result, address + 1, shift
}

/* BIT 6,H */
func dis_instrCB__BIT_6_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "H"
	return result, address + 1, shift
}

/* BIT 6,L */
func dis_instrCB__BIT_6_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "L"
	return result, address + 1, shift
}

/* BIT 6,(HL) */
func dis_instrCB__BIT_6_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 6,A */
func dis_instrCB__BIT_6_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "6" + "," + "A"
	return result, address + 1, shift
}

/* BIT 7,B */
func dis_instrCB__BIT_7_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "B"
	return result, address + 1, shift
}

/* BIT 7,C */
func dis_instrCB__BIT_7_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "C"
	return result, address + 1, shift
}

/* BIT 7,D */
func dis_instrCB__BIT_7_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "D"
	return result, address + 1, shift
}

/* BIT 7,E */
func dis_instrCB__BIT_7_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "E"
	return result, address + 1, shift
}

/* BIT 7,H */
func dis_instrCB__BIT_7_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "H"
	return result, address + 1, shift
}

/* BIT 7,L */
func dis_instrCB__BIT_7_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "L"
	return result, address + 1, shift
}

/* BIT 7,(HL) */
func dis_instrCB__BIT_7_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "(HL)"
	return result, address + 1, shift
}

/* BIT 7,A */
func dis_instrCB__BIT_7_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	result += "7" + "," + "A"
	return result, address + 1, shift
}

/* RES 0,B */
func dis_instrCB__RES_0_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "B"
	return result, address + 1, shift
}

/* RES 0,C */
func dis_instrCB__RES_0_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "C"
	return result, address + 1, shift
}

/* RES 0,D */
func dis_instrCB__RES_0_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "D"
	return result, address + 1, shift
}

/* RES 0,E */
func dis_instrCB__RES_0_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "E"
	return result, address + 1, shift
}

/* RES 0,H */
func dis_instrCB__RES_0_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "H"
	return result, address + 1, shift
}

/* RES 0,L */
func dis_instrCB__RES_0_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "L"
	return result, address + 1, shift
}

/* RES 0,(HL) */
func dis_instrCB__RES_0_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 0,A */
func dis_instrCB__RES_0_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "0" + "," + "A"
	return result, address + 1, shift
}

/* RES 1,B */
func dis_instrCB__RES_1_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "B"
	return result, address + 1, shift
}

/* RES 1,C */
func dis_instrCB__RES_1_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "C"
	return result, address + 1, shift
}

/* RES 1,D */
func dis_instrCB__RES_1_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "D"
	return result, address + 1, shift
}

/* RES 1,E */
func dis_instrCB__RES_1_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "E"
	return result, address + 1, shift
}

/* RES 1,H */
func dis_instrCB__RES_1_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "H"
	return result, address + 1, shift
}

/* RES 1,L */
func dis_instrCB__RES_1_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "L"
	return result, address + 1, shift
}

/* RES 1,(HL) */
func dis_instrCB__RES_1_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 1,A */
func dis_instrCB__RES_1_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "1" + "," + "A"
	return result, address + 1, shift
}

/* RES 2,B */
func dis_instrCB__RES_2_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "B"
	return result, address + 1, shift
}

/* RES 2,C */
func dis_instrCB__RES_2_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "C"
	return result, address + 1, shift
}

/* RES 2,D */
func dis_instrCB__RES_2_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "D"
	return result, address + 1, shift
}

/* RES 2,E */
func dis_instrCB__RES_2_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "E"
	return result, address + 1, shift
}

/* RES 2,H */
func dis_instrCB__RES_2_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "H"
	return result, address + 1, shift
}

/* RES 2,L */
func dis_instrCB__RES_2_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "L"
	return result, address + 1, shift
}

/* RES 2,(HL) */
func dis_instrCB__RES_2_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 2,A */
func dis_instrCB__RES_2_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "2" + "," + "A"
	return result, address + 1, shift
}

/* RES 3,B */
func dis_instrCB__RES_3_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "B"
	return result, address + 1, shift
}

/* RES 3,C */
func dis_instrCB__RES_3_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "C"
	return result, address + 1, shift
}

/* RES 3,D */
func dis_instrCB__RES_3_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "D"
	return result, address + 1, shift
}

/* RES 3,E */
func dis_instrCB__RES_3_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "E"
	return result, address + 1, shift
}

/* RES 3,H */
func dis_instrCB__RES_3_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "H"
	return result, address + 1, shift
}

/* RES 3,L */
func dis_instrCB__RES_3_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "L"
	return result, address + 1, shift
}

/* RES 3,(HL) */
func dis_instrCB__RES_3_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 3,A */
func dis_instrCB__RES_3_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "3" + "," + "A"
	return result, address + 1, shift
}

/* RES 4,B */
func dis_instrCB__RES_4_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "B"
	return result, address + 1, shift
}

/* RES 4,C */
func dis_instrCB__RES_4_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "C"
	return result, address + 1, shift
}

/* RES 4,D */
func dis_instrCB__RES_4_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "D"
	return result, address + 1, shift
}

/* RES 4,E */
func dis_instrCB__RES_4_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "E"
	return result, address + 1, shift
}

/* RES 4,H */
func dis_instrCB__RES_4_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "H"
	return result, address + 1, shift
}

/* RES 4,L */
func dis_instrCB__RES_4_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "L"
	return result, address + 1, shift
}

/* RES 4,(HL) */
func dis_instrCB__RES_4_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 4,A */
func dis_instrCB__RES_4_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "4" + "," + "A"
	return result, address + 1, shift
}

/* RES 5,B */
func dis_instrCB__RES_5_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "B"
	return result, address + 1, shift
}

/* RES 5,C */
func dis_instrCB__RES_5_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "C"
	return result, address + 1, shift
}

/* RES 5,D */
func dis_instrCB__RES_5_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "D"
	return result, address + 1, shift
}

/* RES 5,E */
func dis_instrCB__RES_5_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "E"
	return result, address + 1, shift
}

/* RES 5,H */
func dis_instrCB__RES_5_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "H"
	return result, address + 1, shift
}

/* RES 5,L */
func dis_instrCB__RES_5_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "L"
	return result, address + 1, shift
}

/* RES 5,(HL) */
func dis_instrCB__RES_5_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 5,A */
func dis_instrCB__RES_5_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "5" + "," + "A"
	return result, address + 1, shift
}

/* RES 6,B */
func dis_instrCB__RES_6_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "B"
	return result, address + 1, shift
}

/* RES 6,C */
func dis_instrCB__RES_6_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "C"
	return result, address + 1, shift
}

/* RES 6,D */
func dis_instrCB__RES_6_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "D"
	return result, address + 1, shift
}

/* RES 6,E */
func dis_instrCB__RES_6_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "E"
	return result, address + 1, shift
}

/* RES 6,H */
func dis_instrCB__RES_6_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "H"
	return result, address + 1, shift
}

/* RES 6,L */
func dis_instrCB__RES_6_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "L"
	return result, address + 1, shift
}

/* RES 6,(HL) */
func dis_instrCB__RES_6_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 6,A */
func dis_instrCB__RES_6_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "6" + "," + "A"
	return result, address + 1, shift
}

/* RES 7,B */
func dis_instrCB__RES_7_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "B"
	return result, address + 1, shift
}

/* RES 7,C */
func dis_instrCB__RES_7_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "C"
	return result, address + 1, shift
}

/* RES 7,D */
func dis_instrCB__RES_7_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "D"
	return result, address + 1, shift
}

/* RES 7,E */
func dis_instrCB__RES_7_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "E"
	return result, address + 1, shift
}

/* RES 7,H */
func dis_instrCB__RES_7_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "H"
	return result, address + 1, shift
}

/* RES 7,L */
func dis_instrCB__RES_7_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "L"
	return result, address + 1, shift
}

/* RES 7,(HL) */
func dis_instrCB__RES_7_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "(HL)"
	return result, address + 1, shift
}

/* RES 7,A */
func dis_instrCB__RES_7_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	result += "7" + "," + "A"
	return result, address + 1, shift
}

/* SET 0,B */
func dis_instrCB__SET_0_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "B"
	return result, address + 1, shift
}

/* SET 0,C */
func dis_instrCB__SET_0_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "C"
	return result, address + 1, shift
}

/* SET 0,D */
func dis_instrCB__SET_0_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "D"
	return result, address + 1, shift
}

/* SET 0,E */
func dis_instrCB__SET_0_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "E"
	return result, address + 1, shift
}

/* SET 0,H */
func dis_instrCB__SET_0_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "H"
	return result, address + 1, shift
}

/* SET 0,L */
func dis_instrCB__SET_0_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "L"
	return result, address + 1, shift
}

/* SET 0,(HL) */
func dis_instrCB__SET_0_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 0,A */
func dis_instrCB__SET_0_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "0" + "," + "A"
	return result, address + 1, shift
}

/* SET 1,B */
func dis_instrCB__SET_1_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "B"
	return result, address + 1, shift
}

/* SET 1,C */
func dis_instrCB__SET_1_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "C"
	return result, address + 1, shift
}

/* SET 1,D */
func dis_instrCB__SET_1_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "D"
	return result, address + 1, shift
}

/* SET 1,E */
func dis_instrCB__SET_1_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "E"
	return result, address + 1, shift
}

/* SET 1,H */
func dis_instrCB__SET_1_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "H"
	return result, address + 1, shift
}

/* SET 1,L */
func dis_instrCB__SET_1_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "L"
	return result, address + 1, shift
}

/* SET 1,(HL) */
func dis_instrCB__SET_1_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 1,A */
func dis_instrCB__SET_1_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "1" + "," + "A"
	return result, address + 1, shift
}

/* SET 2,B */
func dis_instrCB__SET_2_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "B"
	return result, address + 1, shift
}

/* SET 2,C */
func dis_instrCB__SET_2_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "C"
	return result, address + 1, shift
}

/* SET 2,D */
func dis_instrCB__SET_2_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "D"
	return result, address + 1, shift
}

/* SET 2,E */
func dis_instrCB__SET_2_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "E"
	return result, address + 1, shift
}

/* SET 2,H */
func dis_instrCB__SET_2_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "H"
	return result, address + 1, shift
}

/* SET 2,L */
func dis_instrCB__SET_2_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "L"
	return result, address + 1, shift
}

/* SET 2,(HL) */
func dis_instrCB__SET_2_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 2,A */
func dis_instrCB__SET_2_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "2" + "," + "A"
	return result, address + 1, shift
}

/* SET 3,B */
func dis_instrCB__SET_3_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "B"
	return result, address + 1, shift
}

/* SET 3,C */
func dis_instrCB__SET_3_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "C"
	return result, address + 1, shift
}

/* SET 3,D */
func dis_instrCB__SET_3_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "D"
	return result, address + 1, shift
}

/* SET 3,E */
func dis_instrCB__SET_3_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "E"
	return result, address + 1, shift
}

/* SET 3,H */
func dis_instrCB__SET_3_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "H"
	return result, address + 1, shift
}

/* SET 3,L */
func dis_instrCB__SET_3_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "L"
	return result, address + 1, shift
}

/* SET 3,(HL) */
func dis_instrCB__SET_3_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 3,A */
func dis_instrCB__SET_3_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "3" + "," + "A"
	return result, address + 1, shift
}

/* SET 4,B */
func dis_instrCB__SET_4_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "B"
	return result, address + 1, shift
}

/* SET 4,C */
func dis_instrCB__SET_4_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "C"
	return result, address + 1, shift
}

/* SET 4,D */
func dis_instrCB__SET_4_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "D"
	return result, address + 1, shift
}

/* SET 4,E */
func dis_instrCB__SET_4_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "E"
	return result, address + 1, shift
}

/* SET 4,H */
func dis_instrCB__SET_4_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "H"
	return result, address + 1, shift
}

/* SET 4,L */
func dis_instrCB__SET_4_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "L"
	return result, address + 1, shift
}

/* SET 4,(HL) */
func dis_instrCB__SET_4_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 4,A */
func dis_instrCB__SET_4_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "4" + "," + "A"
	return result, address + 1, shift
}

/* SET 5,B */
func dis_instrCB__SET_5_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "B"
	return result, address + 1, shift
}

/* SET 5,C */
func dis_instrCB__SET_5_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "C"
	return result, address + 1, shift
}

/* SET 5,D */
func dis_instrCB__SET_5_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "D"
	return result, address + 1, shift
}

/* SET 5,E */
func dis_instrCB__SET_5_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "E"
	return result, address + 1, shift
}

/* SET 5,H */
func dis_instrCB__SET_5_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "H"
	return result, address + 1, shift
}

/* SET 5,L */
func dis_instrCB__SET_5_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "L"
	return result, address + 1, shift
}

/* SET 5,(HL) */
func dis_instrCB__SET_5_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 5,A */
func dis_instrCB__SET_5_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "5" + "," + "A"
	return result, address + 1, shift
}

/* SET 6,B */
func dis_instrCB__SET_6_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "B"
	return result, address + 1, shift
}

/* SET 6,C */
func dis_instrCB__SET_6_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "C"
	return result, address + 1, shift
}

/* SET 6,D */
func dis_instrCB__SET_6_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "D"
	return result, address + 1, shift
}

/* SET 6,E */
func dis_instrCB__SET_6_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "E"
	return result, address + 1, shift
}

/* SET 6,H */
func dis_instrCB__SET_6_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "H"
	return result, address + 1, shift
}

/* SET 6,L */
func dis_instrCB__SET_6_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "L"
	return result, address + 1, shift
}

/* SET 6,(HL) */
func dis_instrCB__SET_6_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 6,A */
func dis_instrCB__SET_6_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "6" + "," + "A"
	return result, address + 1, shift
}

/* SET 7,B */
func dis_instrCB__SET_7_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "B"
	return result, address + 1, shift
}

/* SET 7,C */
func dis_instrCB__SET_7_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "C"
	return result, address + 1, shift
}

/* SET 7,D */
func dis_instrCB__SET_7_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "D"
	return result, address + 1, shift
}

/* SET 7,E */
func dis_instrCB__SET_7_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "E"
	return result, address + 1, shift
}

/* SET 7,H */
func dis_instrCB__SET_7_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "H"
	return result, address + 1, shift
}

/* SET 7,L */
func dis_instrCB__SET_7_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "L"
	return result, address + 1, shift
}

/* SET 7,(HL) */
func dis_instrCB__SET_7_iHL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "(HL)"
	return result, address + 1, shift
}

/* SET 7,A */
func dis_instrCB__SET_7_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	result += "7" + "," + "A"
	return result, address + 1, shift
}

/* IN B,(C) */
func dis_instrED__IN_B_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "B" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),B */
func dis_instrED__OUT_iC_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "B"
	return result, address + 1, shift
}

/* SBC HL,BC */
func dis_instrED__SBC_HL_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "HL" + "," + "BC"
	return result, address + 1, shift
}

/* LD (nnnn),BC */
func dis_instrED__LD_iNNNN_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "BC"
	return result, address + 1, shift
}

/* NEG */
func dis_instrED__NEG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "NEG "
	shift = 0
	return result, address + 1, shift
}

/* RETN */
func dis_instrED__RETN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RETN "
	shift = 0
	return result, address + 1, shift
}

/* IM 0 */
func dis_instrED__IM_0(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IM "
	shift = 0
	result += "0"
	return result, address + 1, shift
}

/* LD I,A */
func dis_instrED__LD_I_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "I" + "," + "A"
	return result, address + 1, shift
}

/* IN C,(C) */
func dis_instrED__IN_C_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "C" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),C */
func dis_instrED__OUT_iC_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "C"
	return result, address + 1, shift
}

/* ADC HL,BC */
func dis_instrED__ADC_HL_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "HL" + "," + "BC"
	return result, address + 1, shift
}

/* LD BC,(nnnn) */
func dis_instrED__LD_BC_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "BC" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* LD R,A */
func dis_instrED__LD_R_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "R" + "," + "A"
	return result, address + 1, shift
}

/* IN D,(C) */
func dis_instrED__IN_D_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "D" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),D */
func dis_instrED__OUT_iC_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "D"
	return result, address + 1, shift
}

/* SBC HL,DE */
func dis_instrED__SBC_HL_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "HL" + "," + "DE"
	return result, address + 1, shift
}

/* LD (nnnn),DE */
func dis_instrED__LD_iNNNN_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "DE"
	return result, address + 1, shift
}

/* IM 1 */
func dis_instrED__IM_1(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IM "
	shift = 0
	result += "1"
	return result, address + 1, shift
}

/* LD A,I */
func dis_instrED__LD_A_I(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "I"
	return result, address + 1, shift
}

/* IN E,(C) */
func dis_instrED__IN_E_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "E" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),E */
func dis_instrED__OUT_iC_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "E"
	return result, address + 1, shift
}

/* ADC HL,DE */
func dis_instrED__ADC_HL_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "HL" + "," + "DE"
	return result, address + 1, shift
}

/* LD DE,(nnnn) */
func dis_instrED__LD_DE_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "DE" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* IM 2 */
func dis_instrED__IM_2(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IM "
	shift = 0
	result += "2"
	return result, address + 1, shift
}

/* LD A,R */
func dis_instrED__LD_A_R(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "R"
	return result, address + 1, shift
}

/* IN H,(C) */
func dis_instrED__IN_H_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "H" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),H */
func dis_instrED__OUT_iC_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "H"
	return result, address + 1, shift
}

/* SBC HL,HL */
func dis_instrED__SBC_HL_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "HL" + "," + "HL"
	return result, address + 1, shift
}

/* LD (nnnn),HL */
func dis_instrED__LD_iNNNN_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "HL"
	return result, address + 1, shift
}

/* RRD */
func dis_instrED__RRD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRD "
	shift = 0
	return result, address + 1, shift
}

/* IN L,(C) */
func dis_instrED__IN_L_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "L" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),L */
func dis_instrED__OUT_iC_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "L"
	return result, address + 1, shift
}

/* ADC HL,HL */
func dis_instrED__ADC_HL_HL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "HL" + "," + "HL"
	return result, address + 1, shift
}

/* LD HL,(nnnn) */
func dis_instrED__LD_HL_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "HL" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* RLD */
func dis_instrED__RLD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLD "
	shift = 0
	return result, address + 1, shift
}

/* IN F,(C) */
func dis_instrED__IN_F_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "F" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),0 */
func dis_instrED__OUT_iC_0(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "0"
	return result, address + 1, shift
}

/* SBC HL,SP */
func dis_instrED__SBC_HL_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "HL" + "," + "SP"
	return result, address + 1, shift
}

/* LD (nnnn),SP */
func dis_instrED__LD_iNNNN_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "SP"
	return result, address + 1, shift
}

/* IN A,(C) */
func dis_instrED__IN_A_iC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IN "
	shift = 0
	result += "A" + "," + "(C)"
	return result, address + 1, shift
}

/* OUT (C),A */
func dis_instrED__OUT_iC_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUT "
	shift = 0
	result += "(C)" + "," + "A"
	return result, address + 1, shift
}

/* ADC HL,SP */
func dis_instrED__ADC_HL_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "HL" + "," + "SP"
	return result, address + 1, shift
}

/* LD SP,(nnnn) */
func dis_instrED__LD_SP_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "SP" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* LDI */
func dis_instrED__LDI(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LDI "
	shift = 0
	return result, address + 1, shift
}

/* CPI */
func dis_instrED__CPI(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CPI "
	shift = 0
	return result, address + 1, shift
}

/* INI */
func dis_instrED__INI(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INI "
	shift = 0
	return result, address + 1, shift
}

/* OUTI */
func dis_instrED__OUTI(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUTI "
	shift = 0
	return result, address + 1, shift
}

/* LDD */
func dis_instrED__LDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LDD "
	shift = 0
	return result, address + 1, shift
}

/* CPD */
func dis_instrED__CPD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CPD "
	shift = 0
	return result, address + 1, shift
}

/* IND */
func dis_instrED__IND(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "IND "
	shift = 0
	return result, address + 1, shift
}

/* OUTD */
func dis_instrED__OUTD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OUTD "
	shift = 0
	return result, address + 1, shift
}

/* LDIR */
func dis_instrED__LDIR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LDIR "
	shift = 0
	return result, address + 1, shift
}

/* CPIR */
func dis_instrED__CPIR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CPIR "
	shift = 0
	return result, address + 1, shift
}

/* INIR */
func dis_instrED__INIR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INIR "
	shift = 0
	return result, address + 1, shift
}

/* OTIR */
func dis_instrED__OTIR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OTIR "
	shift = 0
	return result, address + 1, shift
}

/* LDDR */
func dis_instrED__LDDR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LDDR "
	shift = 0
	return result, address + 1, shift
}

/* CPDR */
func dis_instrED__CPDR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CPDR "
	shift = 0
	return result, address + 1, shift
}

/* INDR */
func dis_instrED__INDR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INDR "
	shift = 0
	return result, address + 1, shift
}

/* OTDR */
func dis_instrED__OTDR(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OTDR "
	shift = 0
	return result, address + 1, shift
}

/* slttrap */
func dis_instrED__SLTTRAP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "slttrap "
	shift = 0
	return result, address + 1, shift
}

/* ADD ix,BC */
func dis_instrDD__ADD_REG_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "ix" + "," + "BC"
	return result, address + 1, shift
}

/* ADD ix,DE */
func dis_instrDD__ADD_REG_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "ix" + "," + "DE"
	return result, address + 1, shift
}

/* LD ix,nnnn */
func dis_instrDD__LD_REG_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "ix" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* LD (nnnn),ix */
func dis_instrDD__LD_iNNNN_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "ix"
	return result, address + 1, shift
}

/* INC ix */
func dis_instrDD__INC_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "ix"
	return result, address + 1, shift
}

/* INC IXH */
func dis_instrDD__INC_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "IXH"
	return result, address + 1, shift
}

/* DEC IXH */
func dis_instrDD__DEC_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "IXH"
	return result, address + 1, shift
}

/* LD IXH,nn */
func dis_instrDD__LD_REGH_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "IXH" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* ADD ix,ix */
func dis_instrDD__ADD_REG_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "ix" + "," + "ix"
	return result, address + 1, shift
}

/* LD ix,(nnnn) */
func dis_instrDD__LD_REG_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "ix" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* DEC ix */
func dis_instrDD__DEC_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "ix"
	return result, address + 1, shift
}

/* INC IXL */
func dis_instrDD__INC_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "IXL"
	return result, address + 1, shift
}

/* DEC IXL */
func dis_instrDD__DEC_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "IXL"
	return result, address + 1, shift
}

/* LD IXL,nn */
func dis_instrDD__LD_REGL_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "IXL" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* INC (ix+dd) */
func dis_instrDD__INC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* DEC (ix+dd) */
func dis_instrDD__DEC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* LD (ix+dd),nn */
func dis_instrDD__LD_iREGpDD_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	address++
	nn := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* ADD ix,SP */
func dis_instrDD__ADD_REG_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "ix" + "," + "SP"
	return result, address + 1, shift
}

/* LD B,IXH */
func dis_instrDD__LD_B_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "IXH"
	return result, address + 1, shift
}

/* LD B,IXL */
func dis_instrDD__LD_B_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "IXL"
	return result, address + 1, shift
}

/* LD B,(ix+dd) */
func dis_instrDD__LD_B_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "B" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* LD C,IXH */
func dis_instrDD__LD_C_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "IXH"
	return result, address + 1, shift
}

/* LD C,IXL */
func dis_instrDD__LD_C_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "IXL"
	return result, address + 1, shift
}

/* LD C,(ix+dd) */
func dis_instrDD__LD_C_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "C" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* LD D,IXH */
func dis_instrDD__LD_D_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "IXH"
	return result, address + 1, shift
}

/* LD D,IXL */
func dis_instrDD__LD_D_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "IXL"
	return result, address + 1, shift
}

/* LD D,(ix+dd) */
func dis_instrDD__LD_D_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "D" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* LD E,IXH */
func dis_instrDD__LD_E_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "IXH"
	return result, address + 1, shift
}

/* LD E,IXL */
func dis_instrDD__LD_E_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "IXL"
	return result, address + 1, shift
}

/* LD E,(ix+dd) */
func dis_instrDD__LD_E_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "E" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* LD IXH,B */
func dis_instrDD__LD_REGH_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXH" + "," + "B"
	return result, address + 1, shift
}

/* LD IXH,C */
func dis_instrDD__LD_REGH_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXH" + "," + "C"
	return result, address + 1, shift
}

/* LD IXH,D */
func dis_instrDD__LD_REGH_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXH" + "," + "D"
	return result, address + 1, shift
}

/* LD IXH,E */
func dis_instrDD__LD_REGH_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXH" + "," + "E"
	return result, address + 1, shift
}

/* LD IXH,IXH */
func dis_instrDD__LD_REGH_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXH" + "," + "IXH"
	return result, address + 1, shift
}

/* LD IXH,IXL */
func dis_instrDD__LD_REGH_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXH" + "," + "IXL"
	return result, address + 1, shift
}

/* LD H,(ix+dd) */
func dis_instrDD__LD_H_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "H" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* LD IXH,A */
func dis_instrDD__LD_REGH_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXH" + "," + "A"
	return result, address + 1, shift
}

/* LD IXL,B */
func dis_instrDD__LD_REGL_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXL" + "," + "B"
	return result, address + 1, shift
}

/* LD IXL,C */
func dis_instrDD__LD_REGL_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXL" + "," + "C"
	return result, address + 1, shift
}

/* LD IXL,D */
func dis_instrDD__LD_REGL_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXL" + "," + "D"
	return result, address + 1, shift
}

/* LD IXL,E */
func dis_instrDD__LD_REGL_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXL" + "," + "E"
	return result, address + 1, shift
}

/* LD IXL,IXH */
func dis_instrDD__LD_REGL_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXL" + "," + "IXH"
	return result, address + 1, shift
}

/* LD IXL,IXL */
func dis_instrDD__LD_REGL_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXL" + "," + "IXL"
	return result, address + 1, shift
}

/* LD L,(ix+dd) */
func dis_instrDD__LD_L_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "L" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* LD IXL,A */
func dis_instrDD__LD_REGL_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IXL" + "," + "A"
	return result, address + 1, shift
}

/* LD (ix+dd),B */
func dis_instrDD__LD_iREGpDD_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + "B"
	return result, address + 1, shift
}

/* LD (ix+dd),C */
func dis_instrDD__LD_iREGpDD_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + "C"
	return result, address + 1, shift
}

/* LD (ix+dd),D */
func dis_instrDD__LD_iREGpDD_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + "D"
	return result, address + 1, shift
}

/* LD (ix+dd),E */
func dis_instrDD__LD_iREGpDD_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + "E"
	return result, address + 1, shift
}

/* LD (ix+dd),H */
func dis_instrDD__LD_iREGpDD_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + "H"
	return result, address + 1, shift
}

/* LD (ix+dd),L */
func dis_instrDD__LD_iREGpDD_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + "L"
	return result, address + 1, shift
}

/* LD (ix+dd),A */
func dis_instrDD__LD_iREGpDD_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(ix+0x%02x)", ix_dd) + "," + "A"
	return result, address + 1, shift
}

/* LD A,IXH */
func dis_instrDD__LD_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* LD A,IXL */
func dis_instrDD__LD_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* LD A,(ix+dd) */
func dis_instrDD__LD_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* ADD A,IXH */
func dis_instrDD__ADD_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* ADD A,IXL */
func dis_instrDD__ADD_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* ADD A,(ix+dd) */
func dis_instrDD__ADD_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* ADC A,IXH */
func dis_instrDD__ADC_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* ADC A,IXL */
func dis_instrDD__ADC_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* ADC A,(ix+dd) */
func dis_instrDD__ADC_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* SUB A,IXH */
func dis_instrDD__SUB_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* SUB A,IXL */
func dis_instrDD__SUB_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* SUB A,(ix+dd) */
func dis_instrDD__SUB_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* SBC A,IXH */
func dis_instrDD__SBC_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* SBC A,IXL */
func dis_instrDD__SBC_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* SBC A,(ix+dd) */
func dis_instrDD__SBC_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* AND A,IXH */
func dis_instrDD__AND_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* AND A,IXL */
func dis_instrDD__AND_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* AND A,(ix+dd) */
func dis_instrDD__AND_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* XOR A,IXH */
func dis_instrDD__XOR_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* XOR A,IXL */
func dis_instrDD__XOR_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* XOR A,(ix+dd) */
func dis_instrDD__XOR_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* OR A,IXH */
func dis_instrDD__OR_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* OR A,IXL */
func dis_instrDD__OR_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* OR A,(ix+dd) */
func dis_instrDD__OR_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* CP A,IXH */
func dis_instrDD__CP_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "A" + "," + "IXH"
	return result, address + 1, shift
}

/* CP A,IXL */
func dis_instrDD__CP_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "A" + "," + "IXL"
	return result, address + 1, shift
}

/* CP A,(ix+dd) */
func dis_instrDD__CP_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	address++
	ix_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(ix+0x%02x)", ix_dd)
	return result, address + 1, shift
}

/* shift DDFDCB */
func dis_instrDD__SHIFT_DDFDCB(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "shift "
	shift = SHIFT_0xCB
	result += "DDFDCB"
	return result, address + 1, shift
}

/* POP ix */
func dis_instrDD__POP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "POP "
	shift = 0
	result += "ix"
	return result, address + 1, shift
}

/* EX (SP),ix */
func dis_instrDD__EX_iSP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "EX "
	shift = 0
	result += "(SP)" + "," + "ix"
	return result, address + 1, shift
}

/* PUSH ix */
func dis_instrDD__PUSH_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "PUSH "
	shift = 0
	result += "ix"
	return result, address + 1, shift
}

/* JP ix */
func dis_instrDD__JP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	result += "ix"
	return result, address + 1, shift
}

/* LD SP,ix */
func dis_instrDD__LD_SP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "SP" + "," + "ix"
	return result, address + 1, shift
}

/* ADD iy,BC */
func dis_instrFD__ADD_REG_BC(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "iy" + "," + "BC"
	return result, address + 1, shift
}

/* ADD iy,DE */
func dis_instrFD__ADD_REG_DE(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "iy" + "," + "DE"
	return result, address + 1, shift
}

/* LD iy,nnnn */
func dis_instrFD__LD_REG_NNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	nnnn := joinBytes(b2, b1)
	result += "iy" + "," + fmt.Sprintf("0x%04x", nnnn)
	return result, address + 1, shift
}

/* LD (nnnn),iy */
func dis_instrFD__LD_iNNNN_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += fmt.Sprintf("(0x%04x)", _nnnn) + "," + "iy"
	return result, address + 1, shift
}

/* INC iy */
func dis_instrFD__INC_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "iy"
	return result, address + 1, shift
}

/* INC IYH */
func dis_instrFD__INC_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "IYH"
	return result, address + 1, shift
}

/* DEC IYH */
func dis_instrFD__DEC_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "IYH"
	return result, address + 1, shift
}

/* LD IYH,nn */
func dis_instrFD__LD_REGH_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "IYH" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* ADD iy,iy */
func dis_instrFD__ADD_REG_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "iy" + "," + "iy"
	return result, address + 1, shift
}

/* LD iy,(nnnn) */
func dis_instrFD__LD_REG_iNNNN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	b1 := memory.ReadByte(address)
	address++
	b2 := memory.ReadByte(address)
	_nnnn := joinBytes(b2, b1)
	result += "iy" + "," + fmt.Sprintf("(0x%04x)", _nnnn)
	return result, address + 1, shift
}

/* DEC iy */
func dis_instrFD__DEC_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "iy"
	return result, address + 1, shift
}

/* INC IYL */
func dis_instrFD__INC_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	result += "IYL"
	return result, address + 1, shift
}

/* DEC IYL */
func dis_instrFD__DEC_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	result += "IYL"
	return result, address + 1, shift
}

/* LD IYL,nn */
func dis_instrFD__LD_REGL_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	nn := memory.ReadByte(address)
	result += "IYL" + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* INC (iy+dd) */
func dis_instrFD__INC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "INC "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* DEC (iy+dd) */
func dis_instrFD__DEC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "DEC "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* LD (iy+dd),nn */
func dis_instrFD__LD_iREGpDD_NN(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	address++
	nn := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + fmt.Sprintf("(0x%02x)", nn)
	return result, address + 1, shift
}

/* ADD iy,SP */
func dis_instrFD__ADD_REG_SP(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "iy" + "," + "SP"
	return result, address + 1, shift
}

/* LD B,IYH */
func dis_instrFD__LD_B_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "IYH"
	return result, address + 1, shift
}

/* LD B,IYL */
func dis_instrFD__LD_B_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "IYL"
	return result, address + 1, shift
}

/* LD B,(iy+dd) */
func dis_instrFD__LD_B_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "B" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* LD C,IYH */
func dis_instrFD__LD_C_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "IYH"
	return result, address + 1, shift
}

/* LD C,IYL */
func dis_instrFD__LD_C_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "IYL"
	return result, address + 1, shift
}

/* LD C,(iy+dd) */
func dis_instrFD__LD_C_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "C" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* LD D,IYH */
func dis_instrFD__LD_D_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "IYH"
	return result, address + 1, shift
}

/* LD D,IYL */
func dis_instrFD__LD_D_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "IYL"
	return result, address + 1, shift
}

/* LD D,(iy+dd) */
func dis_instrFD__LD_D_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "D" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* LD E,IYH */
func dis_instrFD__LD_E_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "IYH"
	return result, address + 1, shift
}

/* LD E,IYL */
func dis_instrFD__LD_E_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "IYL"
	return result, address + 1, shift
}

/* LD E,(iy+dd) */
func dis_instrFD__LD_E_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "E" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* LD IYH,B */
func dis_instrFD__LD_REGH_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYH" + "," + "B"
	return result, address + 1, shift
}

/* LD IYH,C */
func dis_instrFD__LD_REGH_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYH" + "," + "C"
	return result, address + 1, shift
}

/* LD IYH,D */
func dis_instrFD__LD_REGH_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYH" + "," + "D"
	return result, address + 1, shift
}

/* LD IYH,E */
func dis_instrFD__LD_REGH_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYH" + "," + "E"
	return result, address + 1, shift
}

/* LD IYH,IYH */
func dis_instrFD__LD_REGH_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYH" + "," + "IYH"
	return result, address + 1, shift
}

/* LD IYH,IYL */
func dis_instrFD__LD_REGH_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYH" + "," + "IYL"
	return result, address + 1, shift
}

/* LD H,(iy+dd) */
func dis_instrFD__LD_H_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "H" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* LD IYH,A */
func dis_instrFD__LD_REGH_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYH" + "," + "A"
	return result, address + 1, shift
}

/* LD IYL,B */
func dis_instrFD__LD_REGL_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYL" + "," + "B"
	return result, address + 1, shift
}

/* LD IYL,C */
func dis_instrFD__LD_REGL_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYL" + "," + "C"
	return result, address + 1, shift
}

/* LD IYL,D */
func dis_instrFD__LD_REGL_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYL" + "," + "D"
	return result, address + 1, shift
}

/* LD IYL,E */
func dis_instrFD__LD_REGL_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYL" + "," + "E"
	return result, address + 1, shift
}

/* LD IYL,IYH */
func dis_instrFD__LD_REGL_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYL" + "," + "IYH"
	return result, address + 1, shift
}

/* LD IYL,IYL */
func dis_instrFD__LD_REGL_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYL" + "," + "IYL"
	return result, address + 1, shift
}

/* LD L,(iy+dd) */
func dis_instrFD__LD_L_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "L" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* LD IYL,A */
func dis_instrFD__LD_REGL_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "IYL" + "," + "A"
	return result, address + 1, shift
}

/* LD (iy+dd),B */
func dis_instrFD__LD_iREGpDD_B(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + "B"
	return result, address + 1, shift
}

/* LD (iy+dd),C */
func dis_instrFD__LD_iREGpDD_C(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + "C"
	return result, address + 1, shift
}

/* LD (iy+dd),D */
func dis_instrFD__LD_iREGpDD_D(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + "D"
	return result, address + 1, shift
}

/* LD (iy+dd),E */
func dis_instrFD__LD_iREGpDD_E(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + "E"
	return result, address + 1, shift
}

/* LD (iy+dd),H */
func dis_instrFD__LD_iREGpDD_H(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + "H"
	return result, address + 1, shift
}

/* LD (iy+dd),L */
func dis_instrFD__LD_iREGpDD_L(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + "L"
	return result, address + 1, shift
}

/* LD (iy+dd),A */
func dis_instrFD__LD_iREGpDD_A(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(iy+0x%02x)", iy_dd) + "," + "A"
	return result, address + 1, shift
}

/* LD A,IYH */
func dis_instrFD__LD_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* LD A,IYL */
func dis_instrFD__LD_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* LD A,(iy+dd) */
func dis_instrFD__LD_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* ADD A,IYH */
func dis_instrFD__ADD_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* ADD A,IYL */
func dis_instrFD__ADD_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* ADD A,(iy+dd) */
func dis_instrFD__ADD_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADD "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* ADC A,IYH */
func dis_instrFD__ADC_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* ADC A,IYL */
func dis_instrFD__ADC_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* ADC A,(iy+dd) */
func dis_instrFD__ADC_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "ADC "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* SUB A,IYH */
func dis_instrFD__SUB_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* SUB A,IYL */
func dis_instrFD__SUB_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* SUB A,(iy+dd) */
func dis_instrFD__SUB_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SUB "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* SBC A,IYH */
func dis_instrFD__SBC_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* SBC A,IYL */
func dis_instrFD__SBC_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* SBC A,(iy+dd) */
func dis_instrFD__SBC_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SBC "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* AND A,IYH */
func dis_instrFD__AND_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* AND A,IYL */
func dis_instrFD__AND_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* AND A,(iy+dd) */
func dis_instrFD__AND_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "AND "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* XOR A,IYH */
func dis_instrFD__XOR_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* XOR A,IYL */
func dis_instrFD__XOR_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* XOR A,(iy+dd) */
func dis_instrFD__XOR_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "XOR "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* OR A,IYH */
func dis_instrFD__OR_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* OR A,IYL */
func dis_instrFD__OR_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* OR A,(iy+dd) */
func dis_instrFD__OR_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "OR "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* CP A,IYH */
func dis_instrFD__CP_A_REGH(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "A" + "," + "IYH"
	return result, address + 1, shift
}

/* CP A,IYL */
func dis_instrFD__CP_A_REGL(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	result += "A" + "," + "IYL"
	return result, address + 1, shift
}

/* CP A,(iy+dd) */
func dis_instrFD__CP_A_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "CP "
	shift = 0
	address++
	iy_dd := memory.ReadByte(address)
	result += "A" + "," + fmt.Sprintf("(iy+0x%02x)", iy_dd)
	return result, address + 1, shift
}

/* shift DDFDCB */
func dis_instrFD__SHIFT_DDFDCB(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "shift "
	shift = SHIFT_0xCB
	result += "DDFDCB"
	return result, address + 1, shift
}

/* POP iy */
func dis_instrFD__POP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "POP "
	shift = 0
	result += "iy"
	return result, address + 1, shift
}

/* EX (SP),iy */
func dis_instrFD__EX_iSP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "EX "
	shift = 0
	result += "(SP)" + "," + "iy"
	return result, address + 1, shift
}

/* PUSH iy */
func dis_instrFD__PUSH_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "PUSH "
	shift = 0
	result += "iy"
	return result, address + 1, shift
}

/* JP iy */
func dis_instrFD__JP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "JP "
	shift = 0
	result += "iy"
	return result, address + 1, shift
}

/* LD SP,iy */
func dis_instrFD__LD_SP_REG(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "SP" + "," + "iy"
	return result, address + 1, shift
}

/* LD B,RLC (REGISTER+dd) */
func dis_instrDDCB__LD_B_RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RLC"
	return result, address + 1, shift
}

/* LD C,RLC (REGISTER+dd) */
func dis_instrDDCB__LD_C_RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RLC"
	return result, address + 1, shift
}

/* LD D,RLC (REGISTER+dd) */
func dis_instrDDCB__LD_D_RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RLC"
	return result, address + 1, shift
}

/* LD E,RLC (REGISTER+dd) */
func dis_instrDDCB__LD_E_RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RLC"
	return result, address + 1, shift
}

/* LD H,RLC (REGISTER+dd) */
func dis_instrDDCB__LD_H_RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RLC"
	return result, address + 1, shift
}

/* LD L,RLC (REGISTER+dd) */
func dis_instrDDCB__LD_L_RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RLC"
	return result, address + 1, shift
}

/* RLC (REGISTER+dd) */
func dis_instrDDCB__RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RLC "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RLC (REGISTER+dd) */
func dis_instrDDCB__LD_A_RLC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RLC"
	return result, address + 1, shift
}

/* LD B,RRC (REGISTER+dd) */
func dis_instrDDCB__LD_B_RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RRC"
	return result, address + 1, shift
}

/* LD C,RRC (REGISTER+dd) */
func dis_instrDDCB__LD_C_RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RRC"
	return result, address + 1, shift
}

/* LD D,RRC (REGISTER+dd) */
func dis_instrDDCB__LD_D_RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RRC"
	return result, address + 1, shift
}

/* LD E,RRC (REGISTER+dd) */
func dis_instrDDCB__LD_E_RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RRC"
	return result, address + 1, shift
}

/* LD H,RRC (REGISTER+dd) */
func dis_instrDDCB__LD_H_RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RRC"
	return result, address + 1, shift
}

/* LD L,RRC (REGISTER+dd) */
func dis_instrDDCB__LD_L_RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RRC"
	return result, address + 1, shift
}

/* RRC (REGISTER+dd) */
func dis_instrDDCB__RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RRC "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RRC (REGISTER+dd) */
func dis_instrDDCB__LD_A_RRC_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RRC"
	return result, address + 1, shift
}

/* LD B,RL (REGISTER+dd) */
func dis_instrDDCB__LD_B_RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RL"
	return result, address + 1, shift
}

/* LD C,RL (REGISTER+dd) */
func dis_instrDDCB__LD_C_RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RL"
	return result, address + 1, shift
}

/* LD D,RL (REGISTER+dd) */
func dis_instrDDCB__LD_D_RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RL"
	return result, address + 1, shift
}

/* LD E,RL (REGISTER+dd) */
func dis_instrDDCB__LD_E_RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RL"
	return result, address + 1, shift
}

/* LD H,RL (REGISTER+dd) */
func dis_instrDDCB__LD_H_RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RL"
	return result, address + 1, shift
}

/* LD L,RL (REGISTER+dd) */
func dis_instrDDCB__LD_L_RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RL"
	return result, address + 1, shift
}

/* RL (REGISTER+dd) */
func dis_instrDDCB__RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RL "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RL (REGISTER+dd) */
func dis_instrDDCB__LD_A_RL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RL"
	return result, address + 1, shift
}

/* LD B,RR (REGISTER+dd) */
func dis_instrDDCB__LD_B_RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RR"
	return result, address + 1, shift
}

/* LD C,RR (REGISTER+dd) */
func dis_instrDDCB__LD_C_RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RR"
	return result, address + 1, shift
}

/* LD D,RR (REGISTER+dd) */
func dis_instrDDCB__LD_D_RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RR"
	return result, address + 1, shift
}

/* LD E,RR (REGISTER+dd) */
func dis_instrDDCB__LD_E_RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RR"
	return result, address + 1, shift
}

/* LD H,RR (REGISTER+dd) */
func dis_instrDDCB__LD_H_RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RR"
	return result, address + 1, shift
}

/* LD L,RR (REGISTER+dd) */
func dis_instrDDCB__LD_L_RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RR"
	return result, address + 1, shift
}

/* RR (REGISTER+dd) */
func dis_instrDDCB__RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RR "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RR (REGISTER+dd) */
func dis_instrDDCB__LD_A_RR_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RR"
	return result, address + 1, shift
}

/* LD B,SLA (REGISTER+dd) */
func dis_instrDDCB__LD_B_SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SLA"
	return result, address + 1, shift
}

/* LD C,SLA (REGISTER+dd) */
func dis_instrDDCB__LD_C_SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SLA"
	return result, address + 1, shift
}

/* LD D,SLA (REGISTER+dd) */
func dis_instrDDCB__LD_D_SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SLA"
	return result, address + 1, shift
}

/* LD E,SLA (REGISTER+dd) */
func dis_instrDDCB__LD_E_SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SLA"
	return result, address + 1, shift
}

/* LD H,SLA (REGISTER+dd) */
func dis_instrDDCB__LD_H_SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SLA"
	return result, address + 1, shift
}

/* LD L,SLA (REGISTER+dd) */
func dis_instrDDCB__LD_L_SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SLA"
	return result, address + 1, shift
}

/* SLA (REGISTER+dd) */
func dis_instrDDCB__SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLA "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SLA (REGISTER+dd) */
func dis_instrDDCB__LD_A_SLA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SLA"
	return result, address + 1, shift
}

/* LD B,SRA (REGISTER+dd) */
func dis_instrDDCB__LD_B_SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SRA"
	return result, address + 1, shift
}

/* LD C,SRA (REGISTER+dd) */
func dis_instrDDCB__LD_C_SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SRA"
	return result, address + 1, shift
}

/* LD D,SRA (REGISTER+dd) */
func dis_instrDDCB__LD_D_SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SRA"
	return result, address + 1, shift
}

/* LD E,SRA (REGISTER+dd) */
func dis_instrDDCB__LD_E_SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SRA"
	return result, address + 1, shift
}

/* LD H,SRA (REGISTER+dd) */
func dis_instrDDCB__LD_H_SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SRA"
	return result, address + 1, shift
}

/* LD L,SRA (REGISTER+dd) */
func dis_instrDDCB__LD_L_SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SRA"
	return result, address + 1, shift
}

/* SRA (REGISTER+dd) */
func dis_instrDDCB__SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRA "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SRA (REGISTER+dd) */
func dis_instrDDCB__LD_A_SRA_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SRA"
	return result, address + 1, shift
}

/* LD B,SLL (REGISTER+dd) */
func dis_instrDDCB__LD_B_SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SLL"
	return result, address + 1, shift
}

/* LD C,SLL (REGISTER+dd) */
func dis_instrDDCB__LD_C_SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SLL"
	return result, address + 1, shift
}

/* LD D,SLL (REGISTER+dd) */
func dis_instrDDCB__LD_D_SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SLL"
	return result, address + 1, shift
}

/* LD E,SLL (REGISTER+dd) */
func dis_instrDDCB__LD_E_SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SLL"
	return result, address + 1, shift
}

/* LD H,SLL (REGISTER+dd) */
func dis_instrDDCB__LD_H_SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SLL"
	return result, address + 1, shift
}

/* LD L,SLL (REGISTER+dd) */
func dis_instrDDCB__LD_L_SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SLL"
	return result, address + 1, shift
}

/* SLL (REGISTER+dd) */
func dis_instrDDCB__SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SLL "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SLL (REGISTER+dd) */
func dis_instrDDCB__LD_A_SLL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SLL"
	return result, address + 1, shift
}

/* LD B,SRL (REGISTER+dd) */
func dis_instrDDCB__LD_B_SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SRL"
	return result, address + 1, shift
}

/* LD C,SRL (REGISTER+dd) */
func dis_instrDDCB__LD_C_SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SRL"
	return result, address + 1, shift
}

/* LD D,SRL (REGISTER+dd) */
func dis_instrDDCB__LD_D_SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SRL"
	return result, address + 1, shift
}

/* LD E,SRL (REGISTER+dd) */
func dis_instrDDCB__LD_E_SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SRL"
	return result, address + 1, shift
}

/* LD H,SRL (REGISTER+dd) */
func dis_instrDDCB__LD_H_SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SRL"
	return result, address + 1, shift
}

/* LD L,SRL (REGISTER+dd) */
func dis_instrDDCB__LD_L_SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SRL"
	return result, address + 1, shift
}

/* SRL (REGISTER+dd) */
func dis_instrDDCB__SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SRL "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SRL (REGISTER+dd) */
func dis_instrDDCB__LD_A_SRL_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SRL"
	return result, address + 1, shift
}

/* BIT 0,(REGISTER+dd) */
func dis_instrDDCB__BIT_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "0" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* BIT 1,(REGISTER+dd) */
func dis_instrDDCB__BIT_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "1" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* BIT 2,(REGISTER+dd) */
func dis_instrDDCB__BIT_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "2" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* BIT 3,(REGISTER+dd) */
func dis_instrDDCB__BIT_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "3" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* BIT 4,(REGISTER+dd) */
func dis_instrDDCB__BIT_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "4" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* BIT 5,(REGISTER+dd) */
func dis_instrDDCB__BIT_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "5" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* BIT 6,(REGISTER+dd) */
func dis_instrDDCB__BIT_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "6" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* BIT 7,(REGISTER+dd) */
func dis_instrDDCB__BIT_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "BIT "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "7" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD B,RES 0,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 0,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 0,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 0,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 0,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 0,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 0,(REGISTER+dd) */
func dis_instrDDCB__RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "0" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 0,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,RES 1,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 1,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 1,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 1,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 1,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 1,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 1,(REGISTER+dd) */
func dis_instrDDCB__RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "1" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 1,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,RES 2,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 2,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 2,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 2,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 2,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 2,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 2,(REGISTER+dd) */
func dis_instrDDCB__RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "2" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 2,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,RES 3,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 3,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 3,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 3,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 3,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 3,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 3,(REGISTER+dd) */
func dis_instrDDCB__RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "3" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 3,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,RES 4,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 4,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 4,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 4,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 4,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 4,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 4,(REGISTER+dd) */
func dis_instrDDCB__RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "4" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 4,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,RES 5,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 5,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 5,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 5,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 5,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 5,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 5,(REGISTER+dd) */
func dis_instrDDCB__RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "5" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 5,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,RES 6,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 6,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 6,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 6,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 6,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 6,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 6,(REGISTER+dd) */
func dis_instrDDCB__RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "6" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 6,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,RES 7,(REGISTER+dd) */
func dis_instrDDCB__LD_B_RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "RES"
	return result, address + 1, shift
}

/* LD C,RES 7,(REGISTER+dd) */
func dis_instrDDCB__LD_C_RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "RES"
	return result, address + 1, shift
}

/* LD D,RES 7,(REGISTER+dd) */
func dis_instrDDCB__LD_D_RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "RES"
	return result, address + 1, shift
}

/* LD E,RES 7,(REGISTER+dd) */
func dis_instrDDCB__LD_E_RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "RES"
	return result, address + 1, shift
}

/* LD H,RES 7,(REGISTER+dd) */
func dis_instrDDCB__LD_H_RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "RES"
	return result, address + 1, shift
}

/* LD L,RES 7,(REGISTER+dd) */
func dis_instrDDCB__LD_L_RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "RES"
	return result, address + 1, shift
}

/* RES 7,(REGISTER+dd) */
func dis_instrDDCB__RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "RES "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "7" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,RES 7,(REGISTER+dd) */
func dis_instrDDCB__LD_A_RES_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "RES"
	return result, address + 1, shift
}

/* LD B,SET 0,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 0,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 0,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 0,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 0,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 0,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 0,(REGISTER+dd) */
func dis_instrDDCB__SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "0" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 0,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_0_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}

/* LD B,SET 1,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 1,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 1,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 1,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 1,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 1,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 1,(REGISTER+dd) */
func dis_instrDDCB__SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "1" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 1,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_1_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}

/* LD B,SET 2,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 2,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 2,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 2,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 2,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 2,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 2,(REGISTER+dd) */
func dis_instrDDCB__SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "2" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 2,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_2_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}

/* LD B,SET 3,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 3,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 3,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 3,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 3,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 3,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 3,(REGISTER+dd) */
func dis_instrDDCB__SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "3" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 3,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_3_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}

/* LD B,SET 4,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 4,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 4,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 4,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 4,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 4,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 4,(REGISTER+dd) */
func dis_instrDDCB__SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "4" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 4,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_4_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}

/* LD B,SET 5,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 5,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 5,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 5,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 5,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 5,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 5,(REGISTER+dd) */
func dis_instrDDCB__SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "5" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 5,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_5_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}

/* LD B,SET 6,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 6,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 6,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 6,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 6,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 6,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 6,(REGISTER+dd) */
func dis_instrDDCB__SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "6" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 6,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_6_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}

/* LD B,SET 7,(REGISTER+dd) */
func dis_instrDDCB__LD_B_SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "B" + "," + "SET"
	return result, address + 1, shift
}

/* LD C,SET 7,(REGISTER+dd) */
func dis_instrDDCB__LD_C_SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "C" + "," + "SET"
	return result, address + 1, shift
}

/* LD D,SET 7,(REGISTER+dd) */
func dis_instrDDCB__LD_D_SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "D" + "," + "SET"
	return result, address + 1, shift
}

/* LD E,SET 7,(REGISTER+dd) */
func dis_instrDDCB__LD_E_SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "E" + "," + "SET"
	return result, address + 1, shift
}

/* LD H,SET 7,(REGISTER+dd) */
func dis_instrDDCB__LD_H_SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "H" + "," + "SET"
	return result, address + 1, shift
}

/* LD L,SET 7,(REGISTER+dd) */
func dis_instrDDCB__LD_L_SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "L" + "," + "SET"
	return result, address + 1, shift
}

/* SET 7,(REGISTER+dd) */
func dis_instrDDCB__SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "SET "
	shift = 0
	address++
	register_dd := memory.ReadByte(address)
	result += "7" + "," + fmt.Sprintf("(REGISTER+0x%02x)", register_dd)
	return result, address + 1, shift
}

/* LD A,SET 7,(REGISTER+dd) */
func dis_instrDDCB__LD_A_SET_7_iREGpDD(memory MemoryReader, address uint16, shift int) (string, uint16, int) {
	result := "LD "
	shift = 0
	result += "A" + "," + "SET"
	return result, address + 1, shift
}
