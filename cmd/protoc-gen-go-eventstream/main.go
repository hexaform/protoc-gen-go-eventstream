package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"pkg.hexaform.dev/protogen/internal"
)

func main() {
	var flags flag.FlagSet
	opts := protogen.Options{ParamFunc: flags.Set}

	opts.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		
		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}

			envelopeMessages := internal.FindEnvelopes(file)
			if len(envelopeMessages) == 0 {
				continue
			}

			filename := file.GeneratedFilenamePrefix + "-eventstream.pb.go"
			output := &internal.OutputFile{
				GeneratedFile: gen.NewGeneratedFile(filename, file.GoImportPath),
			}

			output.GenerateHeader(file)

			for _, msg := range envelopeMessages {
				if len(msg.Oneofs) == 0 {
					continue
				}

				for _, oneof := range msg.Oneofs {
					if oneof.Desc.IsSynthetic() {
						continue
					}

					var eventFields []*protogen.Field
					for _, field := range oneof.Fields {
						if field.Desc.Kind() == protoreflect.MessageKind {
							eventFields = append(eventFields, field)
						}
					}
					if len(eventFields) == 0 {
						continue
					}

					output.GenerateOptionsInterface(msg, oneof)

					// --- Marker methods on event messages ---
					for _, field := range eventFields {
						output.GenerateOptionMarker(msg, oneof, field)
						output.GenerateOptionWrapper(msg, oneof, field)
					}

					output.GenerateMessageWrapper(msg, oneof)
					output.GenerateMessageUnwrapper(msg, oneof, eventFields)
					output.P()
				}
			}
		}
		return nil
	})
}
