package basics

import "testing"

// TestDivide 使用表驱动测试覆盖正常与异常分支。
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{name: "正常除法", a: 10, b: 2, want: 5},
		{name: "整除", a: 9, b: 3, want: 3},
		{name: "除零", a: 1, b: 0, want: 0, wantErr: true},
		{name: "负数", a: -8, b: 2, want: -4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Divide(%d, %d) error = %v, wantErr = %v", tt.a, tt.b, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
