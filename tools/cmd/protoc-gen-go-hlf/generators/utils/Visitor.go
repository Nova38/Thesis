package utils

import "google.golang.org/protobuf/compiler/protogen"

type (
	Visitor struct {
		file *protogen.File
		gen  *protogen.Plugin
		g    *protogen.GeneratedFile
	}
)

// VisitMessages visits all messages in the file in depth-first order.
// It calls f for each message. If f returns false, VisitMessages skips
// the message's children.
func (v *Visitor) VisitMessages(fn func(m *protogen.Message)) {

	for _, msg := range v.file.Messages {
		v.visitMessages(fn, msg)
	}

}

// visitMessages is the recursive implementation of VisitMessages.
func (v *Visitor) visitMessages(fn func(m *protogen.Message), m *protogen.Message) {
	fn(m)
	for _, message := range m.Messages {
		v.visitMessages(fn, message)
	}

}
