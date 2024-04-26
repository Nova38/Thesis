package serializer

import (
	"fmt"
	"log/slog"
	"reflect"

	"github.com/bufbuild/protovalidate-go"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/hyperledger/fabric-contract-api-go/serializer"
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

type PbTxSerializer struct {
	serializer.JSONSerializer
}

var validator *protovalidate.Validator

func init() {

	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	validator = v

}

// FromString receives the value in its original string form, the reflected type that the
// new value should be of, the schema defining the rules that the converted value should
// adhere to and components which the schema may point to as a reference. The schema and
// component metadata may be nil. The function should produce a reflect value which matches
// the goal type.
func (s *PbTxSerializer) FromString(
	param string,
	objType reflect.Type,
	pm *metadata.ParameterMetadata,
	cm *metadata.ComponentMetadata,
) (reflect.Value, error) {
	// Check to see if the type is a protobuf message

	if objType.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		fmt.Println("protobuf message")
		obj := reflect.New(objType.Elem()).Interface().(proto.Message)

		if err := Unmarshal([]byte(param), obj); err != nil {
			return reflect.Value{}, err
		}

		if err := validator.Validate(obj); err != nil {
			e := common.WrapError(oops.
				In("serializer").
				With("param", param).
				With("Should be of type:", objType.String()).
				With("Error", err.Error()).
				Wrap(common.RequestInvalid),
			)

			slog.Default().Error(
				e.Error(),
				slog.Any("error", e),
				slog.Any("error", err),
				slog.Any("param", param),
			)
			return reflect.Value{}, e
		}

		return reflect.ValueOf(obj), nil
	}

	return s.JSONSerializer.FromString(param, objType, pm, cm)
}

// ToString receives a reflected value of a value, the reflected type of that that value was
// originally, the schema defining the rules of what that value should meet and components
// which the schema may point to as a reference. The schema and component metadata may be nil
// The function should produce a string which represents the original value
func (s *PbTxSerializer) ToString(
	v reflect.Value,
	t reflect.Type,
	rm *metadata.ReturnMetadata,
	cm *metadata.ComponentMetadata,
) (string, error) {
	var str string

	if t.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		// fmt.Println("protobuf message")
		fmt.Println(v.Interface())
		pb, ok := v.Interface().(proto.Message)
		if !ok {
			return "", fmt.Errorf("not a proto message")
		}

		bytes, err := Marshal(pb)
		str = string(bytes)
		if err != nil {
			return "", err
		}

		return str, nil
	}

	return s.JSONSerializer.ToString(v, t, rm, cm)
}
