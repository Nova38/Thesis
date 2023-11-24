package serializer

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/hyperledger/fabric-contract-api-go/serializer"
	"google.golang.org/protobuf/proto"
)

type TxSerializer struct {
	serializer.JSONSerializer
}

// FromString receives the value in its original string form, the reflected type that the
// new value should be of, the schema defining the rules that the converted value should
// adhere to and components which the schema may point to as a reference. The schema and
// component metadata may be nil. The function should produce a reflect value which matches
// the goal type.
func (s *TxSerializer) FromString(
	param string,
	objType reflect.Type,
	pm *metadata.ParameterMetadata,
	cm *metadata.ComponentMetadata,
) (reflect.Value, error) {
	// Check to see if the type is a protobuf message

	if objType.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		fmt.Println("protobuf message")
		obj := reflect.New(objType)

		err := json.Unmarshal([]byte(param), obj.Interface())
		if err != nil {
			return reflect.Value{}, err
		}

		return obj.Elem(), nil
	}

	return s.JSONSerializer.FromString(param, objType, pm, cm)
}

// ToString receives a reflected value of a value, the reflected type of that that value was
// originally, the schema defining the rules of what that value should meet and components
// which the schema may point to as a reference. The schema and component metadata may be nil
// The function should produce a string which represents the original value
func (s *TxSerializer) ToString(
	v reflect.Value,
	t reflect.Type,
	rm *metadata.ReturnMetadata,
	cm *metadata.ComponentMetadata,
) (string, error) {
	return s.JSONSerializer.ToString(v, t, rm, cm)
}
