package day03

import "testing"

func TestTotalJoltage(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				str: exampleInput,
			},
			want: 357,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TotalJoltage(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalJoltage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TotalJoltage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
