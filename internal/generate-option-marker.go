package internal

import "google.golang.org/protobuf/compiler/protogen"

func (o *OutputFile) GenerateOptionMarker(msg *protogen.Message, oneof *protogen.Oneof, option *protogen.Field) {
	oneofOptionInterfaceName := OptionsInterfaceName(msg, oneof)

	o.P("// Marks ", option.Message.GoIdent.GoName, " as a valid ", msg.GoIdent.GoName, " event.")
	o.P("func (e *", option.Message.GoIdent.GoName, ") ", oneofOptionInterfaceName, "() {}")
	o.P()
}
