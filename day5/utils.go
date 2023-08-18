package main

func constructStacks(lines []string) []Stack {
	// top to bottom of the stacks
	// going by row
	var rows [][]string
	for _, line := range lines {
		if line == "" || line[1] == '1' {
			break
		}
		var lineValues []string
		for i := 1; i < len(line); i += 4 {
			// last line, pad with spaces
			if i+4 > len(line) {
				lineValues = append(lineValues, string(line[i]))
				remainingColumns := 9 - len(lineValues)
				for j := 0; j < remainingColumns; j++ {
					lineValues = append(lineValues, " ")
				}
				break
			}
			lineValues = append(lineValues, string(line[i]))
		}
		rows = append(rows, lineValues)
	}

	// Create []Stack of columns
	columns := make([]Stack, len(rows[0]))

	for _, row := range rows {
		// the first row is the top of the stack
		for i, v := range row {
			if v != " " {
				columns[i].PushBottom(v)
			}
		}
	}

	return columns
}

func parseInstructions(lines []string) []Instruction {
	var instructionLines []string
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			instructionLines = lines[i+1:]
		}
	}
	var instructions []Instruction
	for _, line := range instructionLines {
		instruction, err := NewInstruction(line)
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, *instruction)
	}
	return instructions
}
