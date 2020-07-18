package reflecthelper

import (
	"reflect"
	"testing"
)

func TestExtractString(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExtractString(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ExtractString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
