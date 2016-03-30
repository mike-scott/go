// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"log"

	"cmd/compile/internal/gc"
	"cmd/compile/internal/ssa"
	"cmd/internal/obj"
	"cmd/internal/obj/riscv64"
)

// ssaRegToReg maps ssa register numbers to obj register numbers.
var ssaRegToReg = []int16{
	riscv64.REG_X0,
	riscv64.REG_X1,
	riscv64.REG_X2,
	riscv64.REG_X3,
	riscv64.REG_X4,
	riscv64.REG_X5,
	riscv64.REG_X6,
	riscv64.REG_X7,
	riscv64.REG_X8,
	riscv64.REG_X9,
	riscv64.REG_X10,
	riscv64.REG_X11,
	riscv64.REG_X12,
	riscv64.REG_X13,
	riscv64.REG_X14,
	riscv64.REG_X15,
	riscv64.REG_X16,
	riscv64.REG_X17,
	riscv64.REG_X18,
	riscv64.REG_X19,
	riscv64.REG_X20,
	riscv64.REG_X21,
	riscv64.REG_X22,
	riscv64.REG_X23,
	riscv64.REG_X24,
	riscv64.REG_X25,
	riscv64.REG_X26,
	riscv64.REG_X27,
	riscv64.REG_X28,
	riscv64.REG_X29,
	riscv64.REG_X30,
	riscv64.REG_X31,
}

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
	log.Printf("ssaMarkMoves")
}

func ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	log.Printf("ssaGenValue")

	switch v.Op {
	case ssa.OpRISCV64ADD:
		r := v.Reg()
		r1 := v.Args[0].Reg()
		r2 := v.Args[1].Reg()
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: r2})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpRISCV64MOVmem:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpRISCV64LD:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpRISCV64SD:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		gc.AddAux(&p.To, v)
	case ssa.OpRISCV64LoweredNilCheck:
		log.Printf("TODO: LoweredNilCheck")
	default:
		log.Printf("Unhandled op %v", v.Op)
	}
}

func ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	log.Printf("ssaGenBlock")

	switch b.Kind {
	case ssa.BlockExit:
		s.Prog(obj.AUNDEF)
	case ssa.BlockRet:
		s.Prog(obj.ARET)
	default:
		log.Printf("Unhandled kind %v", b.Kind)
	}
}
