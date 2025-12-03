package day02

import "testing"

func TestSumInvalidProductIIDs(t *testing.T) {
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
			want: 1227775554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInvalidProductIIDs(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInvalidProductIIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumInvalidProductIIDs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
