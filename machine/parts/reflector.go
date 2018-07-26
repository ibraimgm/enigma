package parts

// Reflector receives a signal, transforms it, so teh signal can be "bounced back" to the rotors,
// who will receive it in reverse order. This is needed because to make the same key cypher and decypher
// the message.
type Reflector interface {
	Reflect(input Signal) Signal
}
