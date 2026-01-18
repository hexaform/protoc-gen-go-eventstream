package internal

import "google.golang.org/protobuf/compiler/protogen"

func (o *OutputFile) GenerateMessageUnwrapper(msg *protogen.Message, oneof *protogen.Oneof, fields []*protogen.Field) {
	oneofOptionInterfaceName := OptionsInterfaceName(msg, oneof)

	o.P("// Returns the concrete event message stored in ", oneof.GoName, " oneof.")
	o.P("func (m *", msg.GoIdent.GoName, ") Unwrap", oneof.GoName, "() ", oneofOptionInterfaceName, " {")
	o.P("  if m == nil {")
	o.P("    return nil")
	o.P("  }")
	o.P("  switch e := m.Get", oneof.GoName, "().(type) {")
	for _, field := range fields {
		oneofWrapperType := msg.GoIdent.GoName + "_" + field.GoName
		o.P("  case *", oneofWrapperType, ":")
		o.P("    return e.", field.GoName)
	}
	o.P("  default:")
	o.P("    return nil")
	o.P("  }")
	o.P("}")
	o.P()
}
