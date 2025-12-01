package day01

import "testing"

func TestCalculateExamplePassword(t *testing.T) {
	tests := []struct {
		name    string
		want    uint32
		wantErr bool
	}{
		{
			name: "example",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateExamplePassword()
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateExamplePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateExamplePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}
