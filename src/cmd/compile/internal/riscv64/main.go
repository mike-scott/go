// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj/riscv64"
)

func betypeinit() {
	gc.Widthptr = 8
	gc.Widthint = 8
	gc.Widthreg = 8
}

func Main() {
	gc.Thearch.Thechar = 'V'
	gc.Thearch.Thestring = "riscv64"
	gc.Thearch.Thelinkarch = &riscv64.LinkRISCV64
	gc.Thearch.Betypeinit = betypeinit

	gc.Thearch.REGSP = riscv64.REG_SP
	gc.Thearch.REGCTXT = riscv64.REG_CTXT
	gc.Thearch.REGMIN = riscv64.REG_X0
	gc.Thearch.REGMAX = riscv64.REG_X31
	gc.Thearch.FREGMIN = riscv64.REG_F0
	gc.Thearch.FREGMAX = riscv64.REG_F31
	// TODO(prattmic): all the other arches use 50 bits, even though
	// they have 48-bit vaddrs. why?
	gc.Thearch.MAXWIDTH = 1 << 50

	gc.Thearch.Gins = gins

	// TODO(prattmic): other fields?

	gc.Thearch.SSARegToReg = ssaRegToReg
	gc.Thearch.SSAMarkMoves = ssaMarkMoves
	gc.Thearch.SSAGenValue = ssaGenValue
	gc.Thearch.SSAGenBlock = ssaGenBlock

	gc.Main()
	gc.Exit(0)
}
