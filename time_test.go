package reflecthelper

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	type args struct {
		timeTxt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				timeTxt: "",
			},
			wantErr: true,
		},
		{
			name: "unknown date format",
			args: args{
				timeTxt: "19/09/2012 07:56",
			},
			wantErr: true,
		},
		{
			name: "unknown date format",
			args: args{
				timeTxt: "19-09-2012 07:56",
			},
			wantErr: true,
		},
		{
			name: "valid stamp format",
			args: args{
				timeTxt: "Jan 3 15:04:05",
			},
			wantErr: false,
		},
		{
			name: "valid RFC3339 format",
			args: args{
				timeTxt: "2019-10-20T00:00:00.00Z",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseTime(tt.args.timeTxt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
