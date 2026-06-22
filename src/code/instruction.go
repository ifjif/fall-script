package code

import (
	"bytes"
	"fmt"
)

type EmittedInstruct struct {
	OpCode OpCode
	Pos    int
}

type Instructions []byte

func (ins Instructions) String() string {
	var buf bytes.Buffer

	i := 0

	for i < len(ins) {
		def, err := LookupDef(ins[i])
		if err != nil {
			fmt.Fprintf(&buf, "Error: %s\n", err)
			continue
		}

		operands, length := def.ReadOperand(ins[i+1:])

		fmt.Fprintf(&buf, "%04d %s\n", i, ins.fmtInstruction(def, operands))

		i += 1 + length

	}

	return buf.String()
}

func (ins Instructions) StringWithIndent(indent string) string {
	var buf bytes.Buffer

	i := 0

	for i < len(ins) {
		def, err := LookupDef(ins[i])
		if err != nil {
			fmt.Fprintf(&buf, "Error: %s\n", err)
			continue
		}

		operands, length := def.ReadOperand(ins[i+1:])

		fmt.Fprintf(&buf, "%s%04d %s\n", indent, i, ins.fmtInstruction(def, operands))

		i += 1 + length

	}

	return buf.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if operandCount != len(operands) {
		return fmt.Sprintf("Error: operand len %d does not match defind %d\n", len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	case 2:
		return fmt.Sprintf("%s %d %d", def.Name, operands[0], operands[1])
	}

	return fmt.Sprintf("unhandled operandCount for %s\n", def.Name)
}
