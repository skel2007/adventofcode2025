package day03

import "testing"

func TestTotalJoltage(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want1   uint64
		want2   uint64
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				str: exampleInput,
			},
			want1: 357,
			want2: 3121910778619,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := TotalJoltage(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalJoltage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want1 {
				t.Errorf("TotalJoltage() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("TotalJoltage() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
