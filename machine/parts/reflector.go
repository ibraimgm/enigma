package parts

// Reflector receives a signal, transforms it, so the signal can be "bounced back" to the rotors,
// who will receive it in reverse order. This is needed to make the same key cypher and decypher
// the message.
type Reflector interface {
	ID() string
	Reflect(input Signal) Signal
}

func (board *plugboardImpl) ID() string {
	return board.id
}

func (board *plugboardImpl) Reflect(input Signal) Signal {
	return board.Translate(input)
}

// Reflectors is a map with default implementations of the historical reflectos used by the Enigma machine.
// The valid keys are "B", "C", "B Dünn" and "C Dünn".
var Reflectors = map[string]Reflector{
	"B":      Reflector(createPlugboardImpl("B", "AYBRCUDHEQFSGLIPJXKNMOTZVW")),
	"C":      Reflector(createPlugboardImpl("C", "AFBVCPDJEIGOHYKRLZMXNWTQSU")),
	"B Dünn": Reflector(createPlugboardImpl("B Dünn", "AEBNCKDQFUGYHWIJLOMPRXSZTV")),
	"C Dünn": Reflector(createPlugboardImpl("C Dünn", "ARBDCOEJFNGTHKIVLMPWQZSXUY")),
}
