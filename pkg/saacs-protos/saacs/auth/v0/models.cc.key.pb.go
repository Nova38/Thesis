// Code generated by protoc-gen-go-hlf. DO NOT EDIT.
// versions:
// - protoc-gen-cckey v0.0.1
// source: saacs/auth/v0/models.proto

package v0

import (
	v0 "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	strings "strings"
)

// ──────────────────────────────────────────────────
// saacs.auth.v0.KeyAttribute
// Primary Item

// Domain Item
func (m *KeyAttribute) SetKey(key *v0.ItemKey) {
	m.SetKeyAttr(key.ItemKeyParts)
	m.CollectionId = key.GetCollectionId()
	return
}

// SetKeyAttr - Sets the key attributes, returns the number of extra attributes
func (m *KeyAttribute) SetKeyAttr(attr []string) int {
	if len(attr) > 0 {
		m.MspId = attr[0]
	} else {
		return 0
	}
	if len(attr) > 1 {
		m.Oid = attr[1]
	} else {
		return 0
	}
	if len(attr) > 2 {
		m.Value = attr[2]
	} else {
		return 0
	}
	return len(attr) - 3
}

func (m *KeyAttribute) ItemKey() *v0.ItemKey {
	key := &v0.ItemKey{
		CollectionId: m.GetCollectionId(),
		ItemKind:     2,
		ItemType:     "saacs.auth.v0.KeyAttribute",
		ItemKeyParts: m.KeyAttr(),
	}
	return key
}

func (m *KeyAttribute) KeyAttr() []string {
	attr := []string{}
	attr = append(attr, m.GetMspId())
	attr = append(attr, m.GetOid())
	attr = append(attr, m.GetValue())
	return attr
}

func (m *KeyAttribute) ItemKind() v0.ItemKind {
	return v0.ItemKind_ITEM_KIND_PRIMARY_ITEM
}

func (m *KeyAttribute) ItemType() string {
	return "saacs.auth.v0.KeyAttribute"
}

func (m *KeyAttribute) KeySchema() *v0.KeySchema {
	return &v0.KeySchema{
		ItemKind: v0.ItemKind_ITEM_KIND_PRIMARY_ITEM,
		Properties: &fieldmaskpb.FieldMask{Paths: []string{
			"msp_id",
			"oid",
			"value",
		}},
	}
}

// NewFromKey - Creates a new item from a key
func (m *KeyAttribute) NewFromKey(key *v0.ItemKey) *KeyAttribute {
	item := &KeyAttribute{}
	item.SetKey(key)

	return item
}

// StateKey - Returns a composite key for the state
// This follows the same structure as the saacs-cc stub library,
// Main difference is that it doesn't check the key for invalid characters
//
// Example key:= "\u0000auth.Collection\u0000collection0\u0000collection0\u0000"

func (m *KeyAttribute) StateKey() string {

	const sep = string(rune(0))

	attrs := m.ItemKey().GetItemKeyParts()
	if attrs == nil {
		panic("ItemKeyParts is nil")
	}

	collectionId := m.ItemKey().GetCollectionId()
	if collectionId == "" {
		panic("CollectionId is nil")
	}

	if len(attrs) == 0 {
		k := sep + "saacs.auth.v0.KeyAttribute" + sep + collectionId + sep
		return k
	}
	k := sep + "saacs.auth.v0.KeyAttribute" + sep + collectionId + sep + strings.Join(attrs, sep) + sep

	return k
}
