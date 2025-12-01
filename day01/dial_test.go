package day01

import "testing"

func TestCalculatePasswords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want1   uint32
		want2   uint32
		wantErr bool
	}{
		{
			name:  "example",
			args:  args{str: exampleInput},
			want1: 3,
			want2: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := CalculatePasswords(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatePasswords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want1 {
				t.Errorf("CalculatePasswords() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CalculatePasswords() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
