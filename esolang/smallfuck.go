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
	var nested int
	for pointer++; pointer < len(commands); pointer++ {
		switch commands[pointer] {
		case end:
			if nested == 0 {
				return pointer
			}
			nested--
		case start:
			nested++
		}
	}
	return len(commands)
}

func rewindLoop(commands []rune, pointer int) int {
	var nested int
	for pointer--; pointer >= 0; pointer-- {
		switch commands[pointer] {
		case start:
			if nested == 0 {
				return pointer
			}
			nested--
		case end:
			nested++
		}
	}
	return -1
}
