package dagjose

import (
	"io"

	"github.com/ipld/go-ipld-prime/codec/cbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/multicodec"
)

func init() {
	multicodec.RegisterDecoder(0x85, Decode)
	multicodec.RegisterEncoder(0x85, Encode)
}

// Decode deserializes data from the given io.Reader and feeds it into the
// given datamodel.NodeAssembler. Decode fits the codec.Decoder function
// interface.
func Decode(na datamodel.NodeAssembler, r io.Reader) error {
	return cbor.Decode(na, r)
}

// Encode walks the given datamodel.Node and serializes it to the given
// io.Writer. Encode fits the codec.Encoder function interface.
func Encode(n datamodel.Node, w io.Writer) error {
	// Use `datamodel.Copy` to convert the passed `Node` into a `dagJOSENode`
	// via `dagJOSENodeBuilder`, which applies all the necessary validations
	// required to construct a proper DAG-JOSE IPLD Node.
	dagJoseBuilder := NewBuilder()
	err := datamodel.Copy(n, dagJoseBuilder)
	if err != nil {
		return err
	}
	// CBOR is a superset of DAG-JOSE and, as such, can be used to encode valid
	// DAG-JOSE objects without encoding the CID (as expected by the DAG-JOSE
	// spec: https://specs.ipld.io/block-layer/codecs/dag-jose.html).
	return cbor.Encode(dagJoseBuilder.Build(), w)
}
