package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func (o *OutputFile) GenerateMessageWrapper(msg *protogen.Message, oneof *protogen.Oneof) {
	oneofOptionInterfaceName := OptionsInterfaceName(msg, oneof)
	optionsWrapperMethodName := OptionsWrapperMethodName(msg, oneof)

	metadataField := MetadataField(msg)

	o.P("// Construct the envelope message from the concrete ", oneof.GoName, " message.")
	o.P("func Wrap", oneof.GoName, "(e ", oneofOptionInterfaceName, ") *", msg.GoIdent.GoName, " {")
	o.P("  return &", msg.GoIdent.GoName, "{")
	if metadataField != nil {
		metadataIdent := o.QualifiedGoIdent(MetadataImportPath.Ident("NewMessageMetadata"))
		o.P("    ", metadataField.GoName, ": ", metadataIdent, "(),")
	}
	o.P("    ", oneof.GoName, ": e.", optionsWrapperMethodName, "(),")
	o.P("  }")
	o.P("}")
	o.P()
}
