package state

import (
    authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
    "testing"
)

func TestMakeCompositeKey(t *testing.T) {
    type args[T Object] struct {
        obj T
    }
    type testCase[T Object] struct {
        name    string
        args    args[T]
        wantKey string
        wantErr bool
    }
    tests := []testCase[*authpb.User, *authpb.Collection]{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotKey, err := MakeCompositeKey(tt.args.obj)
            if (err != nil) != tt.wantErr {
                t.Errorf("MakeCompositeKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if gotKey != tt.wantKey {
                t.Errorf("MakeCompositeKey() gotKey = %v, want %v", gotKey, tt.wantKey)
            }
        })
    }
}

func TestMakeHiddenKey(t *testing.T) {
    type args[T Object] struct {
        obj T
    }
    type testCase[T Object] struct {
        name          string
        args          args[T]
        wantHiddenKey string
        wantErr       bool
    }
    tests := []testCase[ /* TODO: Insert concrete types here */ ]{
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
                t.Errorf("MakeHiddenKey() gotHiddenKey = %v, want %v", gotHiddenKey, tt.wantHiddenKey)
            }
        })
    }
}

func TestMakeSuggestionKey(t *testing.T) {
    type args[T Object] struct {
        obj          T
        suggestionId string
    }
    type testCase[T Object] struct {
        name              string
        args              args[T]
        wantSuggestionKey string
        wantErr           bool
    }
    tests := []testCase[ /* TODO: Insert concrete types here */ ]{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotSuggestionKey, err := MakeSuggestionKey(tt.args.obj, tt.args.suggestionId)
            if (err != nil) != tt.wantErr {
                t.Errorf("MakeSuggestionKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if gotSuggestionKey != tt.wantSuggestionKey {
                t.Errorf("MakeSuggestionKey() gotSuggestionKey = %v, want %v", gotSuggestionKey, tt.wantSuggestionKey)
            }
        })
    }
}
