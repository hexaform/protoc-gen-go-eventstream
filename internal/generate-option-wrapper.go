package internal

import "google.golang.org/protobuf/compiler/protogen"

func (o *OutputFile) GenerateOptionWrapper(msg *protogen.Message, oneof *protogen.Oneof, option *protogen.Field) {
	oneofWrapperInterfaceName := OptionsWrapperInterfaceName(msg, oneof)
	optionsWrapperMethodName := OptionsWrapperMethodName(msg, oneof)

	o.P("// Wraps ", option.Message.GoIdent.GoName, " in ", msg.GoIdent.GoName+"_"+option.Message.GoIdent.GoName, " oneof wrapper.")
	o.P("func (e *", option.Message.GoIdent.GoName, ") ", optionsWrapperMethodName, "() ", oneofWrapperInterfaceName, " {")
	o.P("    return &", msg.GoIdent.GoName+"_"+option.GoName, "{")
	o.P("        ", option.GoName, ": e,")
	o.P("    }")
	o.P("}")
	o.P()
}

func OptionsWrapperInterfaceName(msg *protogen.Message, oneof *protogen.Oneof) string {
	return "is" + msg.GoIdent.GoName + "_" + oneof.GoName
}

func OptionsWrapperMethodName(msg *protogen.Message, oneof *protogen.Oneof) string {
	return "wrap" + msg.GoIdent.GoName + "_" + oneof.GoName + "_Option"
}
