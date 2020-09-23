package date_agent

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	hostname, err := GetHostName()
	if err != nil {
		t.Errorf(err.Error())
	}
	type args struct {
		commands []string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name:    "case_1",
			args:    args{commands: []string{"hostname"}},
			wantOut: fmt.Sprintf("%s\n", hostname),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := Exec(tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("Exec() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
