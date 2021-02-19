package reflecthelper

import (
	"reflect"
	"testing"
)

func TestCastStruct(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes ReflectStruct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := CastStruct(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("CastStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CastStruct() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
