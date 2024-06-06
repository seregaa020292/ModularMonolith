package errsx

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Error(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		want []string
	}{
		{
			name: "Empty map should return empty string",
			m:    Map{},
			want: []string{""},
		},
		{
			name: "Single error in map should return correct string",
			m: Map{
				"foo": errors.New("bar"),
			},
			want: []string{"foo: bar"},
		},
		{
			name: "Multiple errors in map should return concatenated string",
			m: Map{
				"foo": errors.New("bar"),
				"bar": errors.New("baz"),
			},
			want: []string{"foo: bar", "bar: baz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, strings.Split(tt.m.Error(), "; "))
		})
	}
}

func TestMap_String(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		want []string
	}{
		{
			name: "Empty map should return empty string representation",
			m:    Map{},
			want: []string{""},
		},
		{
			name: "Single error in map should return correct string representation",
			m: Map{
				"foo": errors.New("bar"),
			},
			want: []string{"foo: bar"},
		},
		{
			name: "Multiple errors in map should return correct string representation",
			m: Map{
				"foo": errors.New("bar"),
				"bar": errors.New("baz"),
			},
			want: []string{"foo: bar", "bar: baz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, strings.Split(tt.m.String(), "; "))
		})
	}
}

func TestMap_Get(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		key  string
		want string
	}{
		{
			name: "Getting from empty map should return empty string",
			m:    Map{},
			key:  "foo",
			want: "",
		},
		{
			name: "Getting existing key should return associated error message",
			m: Map{
				"foo": errors.New("bar"),
			},
			key:  "foo",
			want: "bar",
		},
		{
			name: "Getting non-existing key should return empty string",
			m: Map{
				"foo": errors.New("bar"),
				"bar": errors.New("baz"),
			},
			key:  "bar",
			want: "baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.Get(tt.key))
		})
	}
}

func TestMap_Has(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		key  string
		want bool
	}{
		{
			name: "Empty map should not have any keys",
			m:    Map{},
			key:  "foo",
			want: false,
		},
		{
			name: "Map with single key should confirm existence",
			m: Map{
				"foo": errors.New("bar"),
			},
			key:  "foo",
			want: true,
		},
		{
			name: "Map should return false for non-existing key",
			m: Map{
				"foo": errors.New("bar"),
			},
			key:  "bar",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.Has(tt.key))
		})
	}
}

func TestMap_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		want []byte
	}{
		{
			name: "Marshaling empty map should return empty JSON object",
			m:    Map{},
			want: []byte("{}"),
		},
		{
			name: "Marshaling single error should return correct JSON object",
			m: Map{
				"foo": errors.New("bar"),
			},
			want: []byte("{\"foo\":\"bar\"}"),
		},
		{
			name: "Marshaling multiple errors should return correct JSON object",
			m: Map{
				"foo": errors.New("bar"),
				"bar": errors.New("baz"),
			},
			want: []byte("{\"foo\":\"bar\", \"bar\":\"baz\"}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMap_Set(t *testing.T) {
	type args struct {
		key string
		msg any
	}
	tests := []struct {
		name       string
		m          Map
		args       args
		want       string
		wantHas    bool
		wantString []string
	}{
		{
			name:       "Setting key should return empty string",
			m:          Map{},
			args:       args{},
			want:       "",
			wantHas:    false,
			wantString: []string{""},
		},
		{
			name: "Set on empty map should store the error message correctly",
			m:    Map{},
			args: args{
				key: "foo",
				msg: errors.New("bar"),
			},
			want:       "bar",
			wantHas:    true,
			wantString: []string{"foo: bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.msg != nil {
				tt.m.Set(tt.args.key, tt.args.msg)
			}
			assert.Equal(t, tt.wantHas, tt.m.Has(tt.args.key))
			assert.Equal(t, tt.want, tt.m.Get(tt.args.key))
			assert.ElementsMatch(t, tt.wantString, strings.Split(tt.m.String(), "; "))
		})
	}
}
