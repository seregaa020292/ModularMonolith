package tmplreplacer

import (
	"testing"
)

func TestTmplReplacer_Replace(t *testing.T) {
	type args struct {
		src      string
		varTable map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Replace text",
			args: args{
				src: "Hello ${name}, you balance money equal ${price}, in you wallet ${wallet}",
				varTable: map[string]string{
					"name":   "Denis",
					"price":  "500",
					"wallet": "x000000000000000000",
				},
			},
			want: "Hello Denis, you balance money equal 500, in you wallet x000000000000000000",
		},
		{
			name: "Empty varName",
			args: args{
				src: "Hello ${name}, you balance money equal ${price}",
				varTable: map[string]string{
					"price": "500",
				},
			},
			want: "Hello, you balance money equal 500",
		},
		{
			name: "Normalise text",
			args: args{
				src: " Hello Denis .  You balance money equal 500  , in you wallet x000000000000000000 ",
			},
			want: "Hello Denis. You balance money equal 500, in you wallet x000000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := New(tt.args.src)
			if got := tr.Replace(tt.args.varTable); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}
