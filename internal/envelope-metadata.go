package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
)

const MetadataImportPath = protogen.GoImportPath("pkg.hexaform.dev/protogen/envelope")
const MetadataGoName = "Metadata"

func MetadataField(msg *protogen.Message) *protogen.Field {
	for _, field := range msg.Fields {
		if isMetadataField(field) {
			return field
		}
	}

	return nil
}

func isMetadataField(field *protogen.Field) bool {
	if field.Message == nil {
		return false
	}

	if field.Message.GoIdent.GoName != MetadataGoName {
		return false
	}

	if field.Message.GoIdent.GoImportPath != MetadataImportPath {
		return false
	}

	return true
}
