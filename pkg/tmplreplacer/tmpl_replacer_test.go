package tmplreplacer

import (
	"testing"
)

func TestTmplReplacer_Replace(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		varTable map[string]string
		want     string
	}{
		{
			name: "Replace text case",
			src:  "Hello ${name}, you balance money equal ${price}, in you wallet ${wallet}",
			varTable: map[string]string{
				"name":   "Denis",
				"price":  "500",
				"wallet": "x000000000000000000",
			},
			want: "Hello Denis, you balance money equal 500, in you wallet x000000000000000000",
		},
		{
			name: "Empty varName case",
			src:  "Hello ${name}, you balance money equal ${price}",
			varTable: map[string]string{
				"price": "500",
			},
			want: "Hello , you balance money equal 500",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := New(tt.src)
			if got := tr.Replace(tt.varTable); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}
