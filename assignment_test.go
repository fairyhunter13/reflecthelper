package reflecthelper

import (
	"reflect"
	"testing"
)

func TestAssignReflect(t *testing.T) {
	type args struct {
		assigner reflect.Value
		val      reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AssignReflect(tt.args.assigner, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("AssignReflect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
