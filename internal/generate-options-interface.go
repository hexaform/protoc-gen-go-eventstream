package internal

import "google.golang.org/protobuf/compiler/protogen"

func (o *OutputFile) GenerateOptionsInterface(msg *protogen.Message, oneof *protogen.Oneof) {
	oneofOptionInterfaceName := OptionsInterfaceName(msg, oneof)
	oneofWrapperInterfaceName := OptionsWrapperInterfaceName(msg, oneof)
	optionsWrapperMethodName := OptionsWrapperMethodName(msg, oneof)

	o.P("// Implemented by all domain-event messages carried by ", msg.GoIdent.GoName, ".")
	o.P("type ", oneofOptionInterfaceName, " interface {")
	o.P("	", oneofOptionInterfaceName, "()")
	o.P("	", optionsWrapperMethodName, "() ", oneofWrapperInterfaceName)
	o.P("}")
	o.P()
}

func OptionsInterfaceName(msg *protogen.Message, oneof *protogen.Oneof) string {
	return "is" + msg.GoIdent.GoName + "_" + oneof.GoName + "_Option"
}
