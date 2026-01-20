package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func FindEnvelopes(file *protogen.File) []*protogen.Message {
	var envelopes []*protogen.Message

	for _, message := range file.Messages {
		if isEnvelope(message) {
			envelopes = append(envelopes, message)
		}
	}

	return envelopes
}

func isEnvelope(message *protogen.Message) bool {
	var hasMetadata bool
	var hasEventOneof bool

	standaloneFields := getStandaloneFields(message)

	if len(standaloneFields) != 1 {
		return false
	}
	field := standaloneFields[0]
	if isMetadataField(field) {
		hasMetadata = true
	}

	if len(message.Oneofs) != 1 {
		return false
	}
	oneof := message.Oneofs[0]
	if oneof.Desc.IsSynthetic() {
		return false
	}
	for _, field := range oneof.Fields {
		if isEventOneof(field) {
			hasEventOneof = true
			continue
		}
	}

	return hasMetadata && hasEventOneof
}

func getStandaloneFields(field *protogen.Message) []*protogen.Field {
	var fields []*protogen.Field

	for _, f := range field.Fields {
		if f.Oneof != nil {
			continue
		}
		fields = append(fields, f)
	}

	return fields
}

func isEventOneof(field *protogen.Field) bool {
	if field.Oneof == nil {
		return false
	}

	for _, subfield := range field.Oneof.Fields {
		if subfield.Message == nil {
			return false
		}
	}
	return true
}
