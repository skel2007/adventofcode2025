package day07

import "testing"

func TestCountSplits(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				str: exampleInput,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountSplits(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountSplits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountSplits() got = %v, want %v", got, tt.want)
			}
		})
	}
}
