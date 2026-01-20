package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
)

const MetadataImportPath = protogen.GoImportPath("pkg.hexaform.dev/protogen/envelope")

func MetadataField(msg *protogen.Message) *protogen.Field {
	for _, field := range msg.Fields {
		if field.Message == nil {
			return nil
		}

		if field.Message.GoIdent.GoName != msg.GoIdent.GoName {
			return nil
		}

		if field.Message.GoIdent.GoImportPath != MetadataImportPath {
			return nil
		}

		return field
	}

	return nil
}
