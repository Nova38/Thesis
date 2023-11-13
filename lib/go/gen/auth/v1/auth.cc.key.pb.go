// Code generated by protoc-gen-go-hlf. DO NOT EDIT.
// versions:
// - protoc-gen-cckey v0.0.1
// source: auth/v1/auth.proto

package v1

import (
	errors "errors"
	lo "github.com/samber/lo"
	strings "strings"
)

func (m *Collection) Namespace() string {
	return "auth.Collection"
}
func (m *Collection) Key() ([]string, error) {
	attr := []string{m.GetCollectionId()}
	ok := lo.Try(func() error {
		attr = append(attr, m.GetCollectionId())
		return nil
	})
	if !ok {
		return nil, errors.New("Key is nil")
	}
	return attr, nil
}
func (m *Collection) FlatKey() string {
	attr, err := m.Key()
	if err != nil {
		return ""
	}
	attr = attr[1:]
	return strings.Join(attr, "-")
}
func (m *User) Namespace() string {
	return "auth.User"
}
func (m *User) Key() ([]string, error) {
	attr := []string{m.GetCollectionId()}
	ok := lo.Try(func() error {
		attr = append(attr, m.GetMspId())
		attr = append(attr, m.GetUserId())
		return nil
	})
	if !ok {
		return nil, errors.New("Key is nil")
	}
	return attr, nil
}
func (m *User) FlatKey() string {
	attr, err := m.Key()
	if err != nil {
		return ""
	}
	attr = attr[1:]
	return strings.Join(attr, "-")
}
func (m *Role) Namespace() string {
	return "auth.Role"
}
func (m *Role) Key() ([]string, error) {
	attr := []string{m.GetCollectionId()}
	ok := lo.Try(func() error {
		attr = append(attr, m.GetRoleId())
		return nil
	})
	if !ok {
		return nil, errors.New("Key is nil")
	}
	return attr, nil
}
func (m *Role) FlatKey() string {
	attr, err := m.Key()
	if err != nil {
		return ""
	}
	attr = attr[1:]
	return strings.Join(attr, "-")
}
func (m *Membership) Namespace() string {
	return "auth.Membership"
}
func (m *Membership) Key() ([]string, error) {
	attr := []string{m.GetCollectionId()}
	ok := lo.Try(func() error {
		attr = append(attr, m.GetRoleId())
		return nil
	})
	if !ok {
		return nil, errors.New("Key is nil")
	}
	return attr, nil
}
func (m *Membership) FlatKey() string {
	attr, err := m.Key()
	if err != nil {
		return ""
	}
	attr = attr[1:]
	return strings.Join(attr, "-")
}