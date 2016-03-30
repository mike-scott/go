// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import "cmd/internal/obj/riscv64"

func init() {
	var regNamesRISCV64 []string
	var gpMask regMask
	var fpMask regMask

	// Build the list of register names, creating an appropriately indexed
	// regMask for the gp and fp registers as we go.
	for r := riscv64.REG_X0; r <= riscv64.REG_X31; r++ {
		gpMask |= regMask(1) << uint(len(regNamesRISCV64))
		regNamesRISCV64 = append(regNamesRISCV64, "."+riscv64.RegNames[int16(r)])
	}
	for r := riscv64.REG_F0; r <= riscv64.REG_F31; r++ {
		fpMask |= regMask(1) << uint(len(regNamesRISCV64))
		regNamesRISCV64 = append(regNamesRISCV64, "."+riscv64.RegNames[int16(r)])
	}

	if len(regNamesRISCV64) > 64 {
		// regMask is only 64 bits.
		panic("Too many RISCV64 registers")
	}


	gp := regInfo{inputs: []regMask{gpMask}, outputs: []regMask{gpMask}}

	RISCV64ops := []opData{
		{name: "ADD", argLength: 2, reg: gp, asm: "ADD", commutative: true, resultInArg0: true}, // arg0 + arg1
	}

	RISCV64blocks := []blockData{
		{name: "EQ"},
		{name: "NE"},
		{name: "LT"},
		{name: "GE"},
		{name: "LTU"},
		{name: "GEU"},
	}

	archs = append(archs, arch{
		name:     "RISCV64",
		pkg:      "cmd/internal/obj/riscv64",
		genfile:  "../../riscv64/ssa.go",
		ops:      RISCV64ops,
		blocks:   RISCV64blocks,
		regnames: regNamesRISCV64,
	})
}
