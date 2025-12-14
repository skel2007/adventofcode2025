package day06

import "testing"

func TestGrandTotal(t *testing.T) {
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
			want: 4277556,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GrandTotal(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrandTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GrandTotal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
