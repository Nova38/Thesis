package serializer

import (
	"encoding/json"
	"log/slog"

	"github.com/nova38/saacs/pkg/saacs-cc/config"
	"google.golang.org/protobuf/proto"
)

func Marshal[M proto.Message](in M) ([]byte, error) {
	switch config.GetSerialization() {
	case "proto":
		// 'proto' is the serialization format for protobuf messages
		return proto.MarshalOptions{Deterministic: true}.Marshal(in)
	case "json":
		return json.Marshal(in)
	default:
		// 'json' is the default serialization format
		slog.
			With("serialization", config.GetSerialization()).
			Warn("unknown serialization format, using 'json'")

		return json.Marshal(in)
	}
}

func Unmarshal[M proto.Message](in []byte, out M) error {

	switch config.GetSerialization() {
	case "proto":
		// 'proto' is the serialization format for protobuf messages
		return proto.Unmarshal(in, out)
	case "json":
		return json.Unmarshal(in, out)
	default:
		// 'json' is the default serialization format
		slog.
			With("serialization", config.GetSerialization()).
			Warn("unknown serialization format, using 'json'")

		return json.Unmarshal(in, out)
	}
}
