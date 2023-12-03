package common

import (
	"reflect"
	"testing"

	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
)

func TestItemKeyToItemType(t *testing.T) {
	type args struct {
		key *authpb.ItemKey
	}
	tests := []struct {
		name     string
		args     args
		wantItem ItemInterface
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := ItemKeyToItemType(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemKeyToItemType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("ItemKeyToItemType() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func TestItemKeyToItem(t *testing.T) {
	type args struct {
		key *authpb.ItemKey
	}
	tests := []struct {
		name     string
		args     args
		wantItem ItemInterface
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := ItemKeyToItem(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemKeyToItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("ItemKeyToItem() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func TestUnPackItem(t *testing.T) {
	type args struct {
		obj *authpb.Item
	}
	tests := []struct {
		name     string
		args     args
		wantItem ItemInterface
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := UnPackItem(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnPackItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("UnPackItem() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func TestPackItem(t *testing.T) {
	type args struct {
		item ItemInterface
	}
	tests := []struct {
		name    string
		args    args
		wantObj *authpb.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObj, err := PackItem(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("PackItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObj, tt.wantObj) {
				t.Errorf("PackItem() = %v, want %v", gotObj, tt.wantObj)
			}
		})
	}
}

func TestListItemToProtos(t *testing.T) {
	type args struct {
		list []ItemInterface
	}
	tests := []struct {
		name     string
		args     args
		wantObjs []*authpb.Item
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObjs, err := ListItemToProtos(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListItemToProtos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObjs, tt.wantObjs) {
				t.Errorf("ListItemToProtos() = %v, want %v", gotObjs, tt.wantObjs)
			}
		})
	}
}

func TestItemToSuggestion(t *testing.T) {
	type args struct {
		obj ItemInterface
	}
	tests := []struct {
		name           string
		args           args
		wantSuggestion *authpb.Suggestion
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuggestion, err := ItemToSuggestion(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemToSuggestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSuggestion, tt.wantSuggestion) {
				t.Errorf("ItemToSuggestion() = %v, want %v", gotSuggestion, tt.wantSuggestion)
			}
		})
	}
}

func TestSuggestionToItem(t *testing.T) {
	type args struct {
		s *authpb.Suggestion
	}
	tests := []struct {
		name    string
		args    args
		wantObj ItemInterface
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObj, err := SuggestionToItem(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("SuggestionToItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObj, tt.wantObj) {
				t.Errorf("SuggestionToItem() = %v, want %v", gotObj, tt.wantObj)
			}
		})
	}
}

func TestReferenceKeyToItems(t *testing.T) {
	type args struct {
		ref *authpb.ReferenceKey
	}
	tests := []struct {
		name      string
		args      args
		wantItem1 ItemInterface
		wantItem2 ItemInterface
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem1, gotItem2, err := ReferenceKeyToItems(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReferenceKeyToItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem1, tt.wantItem1) {
				t.Errorf("ReferenceKeyToItems() gotItem1 = %v, want %v", gotItem1, tt.wantItem1)
			}
			if !reflect.DeepEqual(gotItem2, tt.wantItem2) {
				t.Errorf("ReferenceKeyToItems() gotItem2 = %v, want %v", gotItem2, tt.wantItem2)
			}
		})
	}
}

func TestPackReference(t *testing.T) {
	type args struct {
		ref   *authpb.ReferenceKey
		item1 ItemInterface
		item2 ItemInterface
	}
	tests := []struct {
		name          string
		args          args
		wantReference *authpb.Reference
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReference, err := PackReference(tt.args.ref, tt.args.item1, tt.args.item2)
			if (err != nil) != tt.wantErr {
				t.Errorf("PackReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReference, tt.wantReference) {
				t.Errorf("PackReference() = %v, want %v", gotReference, tt.wantReference)
			}
		})
	}
}
