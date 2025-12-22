package day07

import "testing"

func TestTeleport(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				str: exampleInput,
			},
			want1: 21,
			want2: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := Teleport(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Teleport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want1 {
				t.Errorf("Teleport() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Teleport() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
