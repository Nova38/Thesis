package common

import (
	"reflect"
	"testing"

	sample "github.com/nova38/thesis/lib/go/gen/sample/v0"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
)

// func TestKeyExists(t *testing.T) {
// 	type args struct {
// 		ctx TxCtxInterface
// 		key string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := KeyExists(tt.args.ctx, tt.args.key); got != tt.want {
// 				t.Errorf("KeyExists() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestMakeItemKeyAttr(t *testing.T) {
	type args struct {
		key *authpb.ItemKey
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 part",
			args: args{
				key: &authpb.ItemKey{
					CollectionId: "collection",
					ItemIdParts:  []string{"id"},
				},
			},
			want: []string{
				"collection",
				"id",
			},
		},
		{
			name: "2 parts",
			args: args{
				key: &authpb.ItemKey{
					CollectionId: "collection",
					ItemIdParts:  []string{"id1", "id2"},
				},
			},
			want: []string{
				"collection",
				"id1",
				"id2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeItemKeyAttr(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeItemKeyAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeItemKeyPrimary(t *testing.T) {
	type args struct {
		key *authpb.ItemKey
	}
	tests := []struct {
		name        string
		args        args
		wantItemKey string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItemKey, err := MakeItemKeyPrimary(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeItemKeyPrimary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotItemKey != tt.wantItemKey {
				t.Errorf("MakeItemKeyPrimary() = %v, want %v", gotItemKey, tt.wantItemKey)
			}
		})
	}
}

func TestMakePrimaryKeyAttr(t *testing.T) {
	type args struct {
		obj *sample.SimpleItem
	}
	tests := []struct {
		name     string
		args     args
		wantAttr []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAttr := MakePrimaryKeyAttr(tt.args.obj); !reflect.DeepEqual(
				gotAttr,
				tt.wantAttr,
			) {
				t.Errorf("MakePrimaryKeyAttr() = %v, want %v", gotAttr, tt.wantAttr)
			}
		})
	}
}

func TestMakePrimaryKey(t *testing.T) {
	type args struct {
		obj *sample.SimpleItem
	}
	tests := []struct {
		name    string
		args    args
		wantKey string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := MakePrimaryKey(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakePrimaryKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("MakePrimaryKey() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func TestMakeSubKeyAtter(t *testing.T) {
	type args struct {
		obj *sample.SimpleItem
	}
	tests := []struct {
		name     string
		args     args
		wantAttr []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAttr := MakeSubKeyAtter(tt.args.obj); !reflect.DeepEqual(gotAttr, tt.wantAttr) {
				t.Errorf("MakeSubKeyAtter() = %v, want %v", gotAttr, tt.wantAttr)
			}
		})
	}
}

func TestMakeSubItemKeyAtter(t *testing.T) {
	type args struct {
		key *authpb.ItemKey
	}
	tests := []struct {
		name     string
		args     args
		wantAttr []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAttr := MakeSubItemKeyAtter(tt.args.key); !reflect.DeepEqual(
				gotAttr,
				tt.wantAttr,
			) {
				t.Errorf("MakeSubItemKeyAtter() = %v, want %v", gotAttr, tt.wantAttr)
			}
		})
	}
}

func TestMakeHiddenKey(t *testing.T) {
	type args struct {
		obj *sample.SimpleItem
	}
	tests := []struct {
		name          string
		args          args
		wantHiddenKey string
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHiddenKey, err := MakeHiddenKey(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeHiddenKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHiddenKey != tt.wantHiddenKey {
				t.Errorf("MakeHiddenKey() = %v, want %v", gotHiddenKey, tt.wantHiddenKey)
			}
		})
	}
}

func TestMakeSuggestionKeyAtter(t *testing.T) {
	type args struct {
		obj          *sample.SimpleItem
		suggestionId string
	}
	tests := []struct {
		name     string
		args     args
		wantAttr []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAttr := MakeSuggestionKeyAtter(tt.args.obj, tt.args.suggestionId); !reflect.DeepEqual(
				gotAttr,
				tt.wantAttr,
			) {
				t.Errorf("MakeSuggestionKeyAtter() = %v, want %v", gotAttr, tt.wantAttr)
			}
		})
	}
}

func TestMakeSuggestionKey(t *testing.T) {
	type args struct {
		obj          *sample.SimpleItem
		suggestionId string
	}
	tests := []struct {
		name              string
		args              args
		wantSuggestionKey string
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuggestionKey, err := MakeSuggestionPrimaryKey(tt.args.obj, tt.args.suggestionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeSuggestionPrimaryKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuggestionKey != tt.wantSuggestionKey {
				t.Errorf(
					"MakeSuggestionPrimaryKey() = %v, want %v",
					gotSuggestionKey,
					tt.wantSuggestionKey,
				)
			}
		})
	}
}

func TestMakeItemKeySuggestion(t *testing.T) {
	type args struct {
		objKey       *authpb.ItemKey
		suggestionId string
	}
	tests := []struct {
		name              string
		args              args
		wantSuggestionKey string
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuggestionKey, err := MakeItemKeySuggestion(tt.args.objKey, tt.args.suggestionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeItemKeySuggestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuggestionKey != tt.wantSuggestionKey {
				t.Errorf(
					"MakeItemKeySuggestion() = %v, want %v",
					gotSuggestionKey,
					tt.wantSuggestionKey,
				)
			}
		})
	}
}

func TestMakeRefKeyAttrs(t *testing.T) {
	type args struct {
		ref *authpb.ReferenceKey
	}
	tests := []struct {
		name        string
		args        args
		wantRefKey1 []string
		wantRefKey2 []string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRefKey1, gotRefKey2, err := MakeRefKeyAttrs(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRefKeyAttrs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRefKey1, tt.wantRefKey1) {
				t.Errorf("MakeRefKeyAttrs() gotRefKey1 = %v, want %v", gotRefKey1, tt.wantRefKey1)
			}
			if !reflect.DeepEqual(gotRefKey2, tt.wantRefKey2) {
				t.Errorf("MakeRefKeyAttrs() gotRefKey2 = %v, want %v", gotRefKey2, tt.wantRefKey2)
			}
		})
	}
}

func TestMakeRefKeys(t *testing.T) {
	type args struct {
		ref *authpb.ReferenceKey
	}
	tests := []struct {
		name        string
		args        args
		wantRefKey1 string
		wantRefKey2 string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRefKey1, gotRefKey2, err := MakeRefKeys(tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRefKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRefKey1 != tt.wantRefKey1 {
				t.Errorf("MakeRefKeys() gotRefKey1 = %v, want %v", gotRefKey1, tt.wantRefKey1)
			}
			if gotRefKey2 != tt.wantRefKey2 {
				t.Errorf("MakeRefKeys() gotRefKey2 = %v, want %v", gotRefKey2, tt.wantRefKey2)
			}
		})
	}
}
