package dagjose

import (
	"fmt"
	"strconv"

	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/fluent"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/node/mixins"
)

type jwsSignaturesNode struct{ sigs []jwsSignature }

// jwsSignatures Node implementation

func (d *jwsSignaturesNode) Kind() ipld.Kind {
	return ipld.Kind_List
}
func (d *jwsSignaturesNode) LookupByString(key string) (ipld.Node, error) {
	index, err := strconv.Atoi(key)
	if err != nil {
		return nil, err
	}
	return d.LookupByIndex(int64(index))
}
func (d *jwsSignaturesNode) LookupByNode(key ipld.Node) (ipld.Node, error) {
	index, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return d.LookupByIndex(index)
}
func (d *jwsSignaturesNode) LookupByIndex(idx int64) (ipld.Node, error) {
	if int64(len(d.sigs)) > idx {
		return jwsSignatureNode{&d.sigs[idx]}, nil
	}
	return nil, fmt.Errorf("Index %v out of range", idx)
}
func (d *jwsSignaturesNode) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	idx, err := seg.Index()
	if err != nil {
		return nil, err
	}
	return d.LookupByIndex(idx)
}
func (d *jwsSignaturesNode) MapIterator() ipld.MapIterator {
	return nil
}
func (d *jwsSignaturesNode) ListIterator() ipld.ListIterator {
	return &jwsSignaturesIterator{
		sigs:  d.sigs,
		index: 0,
	}
}
func (d *jwsSignaturesNode) Length() int64 {
	return int64(len(d.sigs))
}
func (d *jwsSignaturesNode) IsAbsent() bool {
	return false
}
func (d *jwsSignaturesNode) IsNull() bool {
	return false
}
func (d *jwsSignaturesNode) AsBool() (bool, error) {
	return mixins.List{TypeName: "jose.JWSSignature"}.AsBool()
}
func (d *jwsSignaturesNode) AsInt() (int64, error) {
	return mixins.List{TypeName: "jose.JWSSignature"}.AsInt()
}
func (d *jwsSignaturesNode) AsFloat() (float64, error) {
	return mixins.List{TypeName: "jose.JWSSignature"}.AsFloat()
}
func (d *jwsSignaturesNode) AsString() (string, error) {
	return mixins.List{TypeName: "jose.JWSSignature"}.AsString()
}
func (d *jwsSignaturesNode) AsBytes() ([]byte, error) {
	return mixins.List{TypeName: "jose.JWSSignature"}.AsBytes()
}
func (d *jwsSignaturesNode) AsLink() (ipld.Link, error) {
	return mixins.List{TypeName: "jose.JWSSignature"}.AsLink()
}
func (d *jwsSignaturesNode) Prototype() ipld.NodePrototype {
	return nil
}

// joseSignaturesNode ListIterator implementation

type jwsSignaturesIterator struct {
	sigs  []jwsSignature
	index int
}

func (j *jwsSignaturesIterator) Next() (idx int64, value ipld.Node, err error) {
	if j.Done() {
		return 0, nil, ipld.ErrIteratorOverread{}
	}
	result := &j.sigs[j.index]
	j.index += 1
	return int64(j.index), jwsSignatureNode{result}, nil
}

func (j *jwsSignaturesIterator) Done() bool {
	return j.index >= len(j.sigs)
}

// end ipld.Node implementation

// JOSESignature Node implementation

type jwsSignatureNode struct{ *jwsSignature }

var signatureMixin = mixins.Map{TypeName: "jwsSignature"}

func (d jwsSignatureNode) Kind() ipld.Kind {
	return ipld.Kind_Map
}
func (d jwsSignatureNode) LookupByString(key string) (ipld.Node, error) {
	if key == "signature" {
		return basicnode.NewBytes(d.signature), nil
	}
	if key == "protected" {
		return basicnode.NewBytes(d.protected), nil
	}
	if key == "header" {
		if d.header == nil {
			return nil, nil
		}
		return fluent.MustBuildMap(
			basicnode.Prototype.Map,
			int64(len(d.header)),
			func(ma fluent.MapAssembler) {
				for key, value := range d.header {
					ma.AssembleEntry(key).AssignNode(value)
				}
			},
		), nil
	}
	return nil, nil
}
func (d jwsSignatureNode) LookupByNode(key ipld.Node) (ipld.Node, error) {
	keyString, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return d.LookupByString(keyString)
}
func (d jwsSignatureNode) LookupByIndex(idx int64) (ipld.Node, error) {
	return nil, nil
}

func (d jwsSignatureNode) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return d.LookupByString(seg.String())
}
func (d jwsSignatureNode) MapIterator() ipld.MapIterator {
	return &jwsSignatureMapIterator{sig: d, index: 0}
}
func (d jwsSignatureNode) ListIterator() ipld.ListIterator {
	return nil
}
func (d jwsSignatureNode) Length() int64 {
	return int64(len((&jwsSignatureMapIterator{sig: d, index: 0}).presentKeys()))
}
func (d jwsSignatureNode) IsAbsent() bool {
	return false
}
func (d jwsSignatureNode) IsNull() bool {
	return false
}
func (d jwsSignatureNode) AsBool() (bool, error) {
	return signatureMixin.AsBool()
}
func (d jwsSignatureNode) AsInt() (int64, error) {
	return signatureMixin.AsInt()
}
func (d jwsSignatureNode) AsFloat() (float64, error) {
	return signatureMixin.AsFloat()
}
func (d jwsSignatureNode) AsString() (string, error) {
	return signatureMixin.AsString()
}
func (d jwsSignatureNode) AsBytes() ([]byte, error) {
	return signatureMixin.AsBytes()
}
func (d jwsSignatureNode) AsLink() (ipld.Link, error) {
	return signatureMixin.AsLink()
}
func (d jwsSignatureNode) Prototype() ipld.NodePrototype {
	return nil
}

// end JOSESignature ipld.Node implementation

type jwsSignatureMapIterator struct {
	sig   jwsSignatureNode
	index int
}

func (j *jwsSignatureMapIterator) Next() (key ipld.Node, value ipld.Node, err error) {
	if j.Done() {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	keys := j.presentKeys()
	keyString := keys[j.index]
	value, _ = j.sig.LookupByString(keyString)
	j.index += 1
	return basicnode.NewString(keyString), value, nil
}

func (j *jwsSignatureMapIterator) presentKeys() []string {
	result := []string{"signature"}
	if j.sig.protected != nil {
		result = append(result, "protected")
	}
	if j.sig.header != nil {
		result = append(result, "header")
	}
	return result
}

func (j *jwsSignatureMapIterator) Done() bool {
	return j.index >= len(j.presentKeys())
}
