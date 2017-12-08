package esolang

const (
	flip  = '*'
	right = '>'
	left  = '<'
	start = '['
	end   = ']'
)

func Interpreter(code, tape string) string {
	var (
		commands, tapeBits = []rune(code), []rune(tape)
		pointer            int
		finish             bool
	)
	for commandPointer := 0; commandPointer < len(code); commandPointer++ {
		command := commands[commandPointer]
		switch command {
		case flip:
			tapeBits[pointer] = flipOnTape[tapeBits[pointer]]
		case right, left:
			if command == right {
				pointer++
			} else {
				pointer--
			}
			// Pointer out of tape?
			finish = pointer < 0 || pointer >= len(tape)
		case start:
			if tapeBits[pointer] == '0' {
				commandPointer = skipLoop(commands, commandPointer)
			}
		case end:
			if tapeBits[pointer] == '1' {
				commandPointer = rewindLoop(commands, commandPointer)
			}
		}

		if finish {
			break
		}
	}
	return string(tapeBits)
}

var flipOnTape = map[rune]rune{
	'0': '1',
	'1': '0',
}

func skipLoop(commands []rune, pointer int) int {
	return findNext(commands, end, pointer, true)
}

func rewindLoop(commands []rune, pointer int) int {
	return findNext(commands, start, pointer, false)
}

func findNext(
	commands []rune,
	lookFor rune,
	pointer int,
	forward bool,
) int {
	direction := -1
	if forward {
		direction = 1
	}
	opposite := end
	if lookFor == end {
		opposite = start
	}
	var nested int
	for pointer += direction; pointer >= 0 && pointer < len(commands); pointer += direction {
		switch commands[pointer] {
		case lookFor:
			if nested == 0 {
				return pointer
			}
			nested--
		case opposite:
			nested++
		}
	}
	return map[int]int{1: len(commands), -1: -1}[direction]
}
