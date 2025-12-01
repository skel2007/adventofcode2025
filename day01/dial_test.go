package day01

import "testing"

func TestCalculateExamplePasswords(t *testing.T) {
	tests := []struct {
		name    string
		want1   uint32
		want2   uint32
		wantErr bool
	}{
		{
			name:  "example",
			want1: 3,
			want2: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := CalculateExamplePasswords()
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateExamplePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want1 {
				t.Errorf("CalculateExamplePassword() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CalculateExamplePassword() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
