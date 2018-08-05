package parts

// Plugboard represents a plugboard in an Enigma machine. The plugboard receives a signal
// and (maybe) change it to a different letter.
type Plugboard interface {
	Translate(input Signal) Signal
}

// NoPlugboard is an empty plugboard, that does not change the signal
var NoPlugboard Plugboard = &plugboardImpl{}

// CreatePlugboard builds a new plugboard, that swap the inputs according to the pairs of letters specified.
// For example, if plugs  is specified as "ABCD", this means that the plugboard changes 'A' to 'B', 'B' to 'A',
// 'C' to 'D' and 'D' to 'C'
func CreatePlugboard(plugs string) Plugboard {
	m := make(map[int]int)
	runes := []rune(plugs)
	size := len(runes)

	for i := 0; i+1 < size; i += 2 {
		a := charToInt(runes[i])
		b := charToInt(runes[i+1])

		if a != -1 && b != -1 {
			m[a] = b
			m[b] = a
		}
	}

	return &plugboardImpl{m}
}

type plugboardImpl struct {
	plugs map[int]int
}

func (board *plugboardImpl) Translate(input Signal) Signal {
	m := board.plugs
	a := int(input)

	if b, ok := m[a]; ok {
		return Signal(b)
	}

	return input
}
