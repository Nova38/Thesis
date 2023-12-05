package generators

import (
	auth_pb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/proto"

	_ "strings"

	_ "google.golang.org/protobuf/types/dynamicpb"
)

const (
	fmtPkg     = protogen.GoImportPath("fmt")
	gatewayPkg = protogen.GoImportPath(
		"github.com/hyperledger/fabric-gateway/pkg/client",
	)
	generatedGatewayPackageSuffix = "gateway"
)

type ServiceGenerator struct{}

func (sg *ServiceGenerator) GenerateFile(
	gen *protogen.Plugin,
	file *protogen.File,
) (*protogen.GeneratedFile, error) {
	filename := file.GeneratedFilenamePrefix + ".cc.service.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	// Generate the header
	g.P("// Code generated by proto-gen-go-auth_pb. DO NOT EDIT.")
	g.P("// versions:")

	g.P("// - protoc-gen-cckey v0.0.1")
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	if file.Services == nil {
		g.Skip()
	}

	for _, sv := range file.Services {
		// g.P("type AuthContractImpl struct{}")
		sg.GenerateService(gen, g, sv)

		// sg.GenerateGoGatewayHandler(gen, g, sv)
	}

	// sg.GenerateGatewayFile(gen, file)

	return g, nil
}

func (sv *ServiceGenerator) GenerateGatewayFile(
	gen *protogen.Plugin,
	file *protogen.File,
) {
	filename := file.GeneratedFilenamePrefix + ".cc.gateway.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	// Generate the header
	g.P("// Code generated by proto-gen-go-auth_pb. DO NOT EDIT.")
	g.P("// versions:")

	g.P("// - protoc-gen-cckey v0.0.1")
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, s := range file.Services {
		sv.GenerateGoGatewayHandler(gen, g, s)
	}
}

func (sv *ServiceGenerator) GenerateService(
	gen *protogen.Plugin,
	g *protogen.GeneratedFile,
	v *protogen.Service,
) {
	g.P("// Service ", v.GoName)

	sv.GenerateInterface(gen, g, v)
	g.P()
	g.P()
	sv.GenerateStruct(gen, g, v)
	g.P()
	g.P()
}

func (sv *ServiceGenerator) GenerateInterface(
	gen *protogen.Plugin,
	g *protogen.GeneratedFile,
	v *protogen.Service,
) {
	ctx := g.QualifiedGoIdent(
		protogen.GoIdent{
			GoName: "GenericTxCtxInterface",

			GoImportPath: "github.com/nova38/thesis/packages/saacs/common",
		},
	)
	// shortName, _ := strings.CutSuffix(v.GoName, "Service")
	// ctx := shortName + "TxCtx"

	g.P("type ", v.GoName, "Interface[T ", ctx, "] interface{")
	defer g.P("}")

	for _, m := range v.Methods {

		mComments := m.Comments.Leading.String()

		if len(mComments) == 0 {
			// add the fn name to the comments
			mComments += "// " + m.GoName + "\n // \n"
		}

		op, ok := proto.GetExtension(m.Desc.Options(), auth_pb.E_Operation).(*auth_pb.Operation)

		if !ok {
			mComments += "// No operation defined for " + m.GoName + "\n"
		} else if op != nil {
			mComments += "// # Operation: \n"
			mComments += "//   - Domain: " + op.GetAction().String() + "\n"
			// mComments += "//   - Action: " + op.ItemTypeName + "\n"
		}

		if m.Input.Desc.Name() == "Empty" {
			mComments += "// " + "\n"
			mComments += "// req is empty\n"
			g.P(
				mComments, m.GoName,
				"(ctx T) ",
				"(res *", m.Output.GoIdent, ",err error)",
			)
		} else {
			g.P(
				mComments, m.GoName,
				"(ctx T, req *", m.Input.GoIdent, ") ",
				"(res *", m.Output.GoIdent, ",err error)",
			)
		}
		g.P()
	}
}

func (sv *ServiceGenerator) GenerateStruct(
	gen *protogen.Plugin,
	g *protogen.GeneratedFile,
	v *protogen.Service,
) {
	g.P("type ", v.GoName, "Base struct{")
	g.P("}")

	sv.GenerateStructEvaluateTransactions(gen, g, v)
	g.P()
	GenerateOperationLookup(gen, g, v)
}

func (sv *ServiceGenerator) GenerateStructEvaluateTransactions(
	gen *protogen.Plugin,
	g *protogen.GeneratedFile,
	v *protogen.Service,
) {
	g.P("func (s *", v.GoName, "Base) GetEvaluateTransactions() []string {")
	// g.P("return []string{")

	var fns []string

	for _, m := range v.Methods {
		tt, ok := proto.GetExtension(m.Desc.Options(), auth_pb.E_TransactionType).(auth_pb.TransactionType)

		if !ok {
			continue
		}

		if tt == auth_pb.TransactionType_TRANSACTION_TYPE_QUERY {
			fns = append(fns, m.GoName)
		}
	}
	if len(fns) == 0 {
		g.P("return []string{}")
	} else {
		g.P("return []string{")
		for _, fn := range fns {
			g.P("\"", fn, "\",")
		}
		g.P("}")
	}
	g.P("}")
}

func GenerateOperationLookup(
	gen *protogen.Plugin,
	g *protogen.GeneratedFile,
	v *protogen.Service,
) {
	opImport := g.QualifiedGoIdent(
		protogen.GoIdent{
			GoName:       "Operation",
			GoImportPath: "github.com/nova38/thesis/packages/saacs/gen/auth/v1",
		},
	)
	// fmtImport := (protogen.GoIdent{GoImportPath: "fmt"})
	g.P("//")

	g.P("func ", v.GoName, "GetTxOperation(txName string)", "( op *", opImport, ", err error) {")

	g.P("switch txName {")
	for _, m := range v.Methods {
		g.P("case \"", m.GoName, "\":")
		op, ok := proto.GetExtension(m.Desc.Options(), auth_pb.E_Operation).(*auth_pb.Operation)
		g.P("//", op)
		if !ok {
			g.P("// No operation defined for ", m.GoName)
			g.P(
				"return nil,",
				g.QualifiedGoIdent(fmtPkg.Ident(".Errorf")),
				"(\"No operation defined for ",
				m.GoName,
				"\")",
			)
		} else if op != nil {
			g.P("return &", opImport, "{")
			g.P("Action: ", op.GetAction().Number(), ",")
			// if op.ItemTypeName != "" {
			// 	g.P("ItemTypeName: ", op.ItemTypeName, ",")
			// }
			g.P("}, nil")

		}
	}
	g.P("default:")
	g.P(
		"return nil,",
		g.QualifiedGoIdent(fmtPkg.Ident("Errorf")),
		"(\"No operation defined for \"+txName)",
	)
	g.P("}")
	g.P("return nil, nil")
	g.P("}")

	// Gen GetIgnoredFunctions
	g.P()
	g.P("func (s *", v.GoName, "Base) GetIgnoredFunctions() []string {")
	g.P("return []string{\"GetTxOperation\"}")
	g.P("}")
}

func (sv *ServiceGenerator) GenerateGoGatewayHandler(
	gen *protogen.Plugin,
	g *protogen.GeneratedFile,
	v *protogen.Service,
) {
	contract := g.QualifiedGoIdent(gatewayPkg.Ident("Contract"))
	json := g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "encoding/json"})

	g.P("// ", v.GoName, "Handler is used to interact with the service with the gateway")

	g.P("type ", v.GoName, "Handler struct {")
	g.P("ChaincodeName string")
	g.P("contract *",
		contract,
	)
	g.P("}")

	for _, m := range v.Methods {

		in := m.Input.GoIdent.GoName
		out := m.Output.GoIdent.GoName
		tt, ok := proto.GetExtension(m.Desc.Options(), auth_pb.E_TransactionType).(auth_pb.TransactionType)

		if !ok {
			continue
		}
		if in == "Empty" {
			g.P("func (s *", v.GoName, "Handler)", m.GoName, "( )(out *", out, "){")
			if tt == auth_pb.TransactionType_TRANSACTION_TYPE_QUERY {
				g.P(
					"evaluateResult, err:= s.contract.EvaluateTransaction(s.ChaincodeName+\":",
					m.GoName,
					"\")",
				)
			} else {
				g.P("evaluateResult, err:= s.contract.SubmitTransaction(s.ChaincodeName+\":", m.GoName, "\")")
			}
		} else {
			g.P("func (s *", v.GoName, "Handler)", m.GoName, "(in *", in, ")(out *", out, "){")

			g.P("inBytes, err := ", json, "Marshal(in)")
			g.P("if err != nil {")
			g.P("return nil")
			g.P("}")

			if tt == auth_pb.TransactionType_TRANSACTION_TYPE_QUERY {
				g.P(
					"evaluateResult, err:= s.contract.EvaluateTransaction(s.ChaincodeName\":",
					m.GoName, "\", string(inBytes) )")
			} else {
				g.P(
					"evaluateResult, err:= s.contract.SubmitTransaction(s.ChaincodeName\":",
					m.GoName, "\", string(inBytes) )")
			}
		}

		g.P("if err != nil {")
		g.P("return nil")
		g.P("}")

		g.P("err =", json, "Unmarshal(evaluateResult, &out)")
		g.P("if err != nil {")
		g.P("return nil")
		g.P("}")

		g.P("return out")
		g.P("}")
		// g.P("return []string{")

		// tt, ok := proto.GetExtension(m.Desc.Options(), auth_pb.E_TransactionType).(auth_pb.TransactionType)

		// if !ok {
		// 	continue
		// }

		// if tt == auth_pb.TransactionType_TRANSACTION_TYPE_QUERY {
		// 	g.P("\"", m.GoName, "\",")
		// }
	}
	// g.P("}")
	// g.P("}")
}
