package internal

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_processLine(t *testing.T) {
	type args struct {
		line  string
		apply func(name string, directions *Directions) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "All neighbors",
			args: args{
				line: "Foo north=Bar east=Baz south=Qux west=Quux",
				apply: func(name string, directions *Directions) error {
					assert.Equal(t, name, "Foo")
					assert.Equal(t, directions.North, "Bar")
					assert.Equal(t, directions.East, "Baz")
					assert.Equal(t, directions.South, "Qux")
					assert.Equal(t, directions.West, "Quux")
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "Some neighbors",
			args: args{
				line: "Foo north=Bar east=Baz",
				apply: func(name string, directions *Directions) error {
					assert.Equal(t, name, "Foo")
					assert.Equal(t, directions.North, "Bar")
					assert.Equal(t, directions.East, "Baz")
					assert.Equal(t, directions.South, "")
					assert.Equal(t, directions.West, "")
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "No neighbors",
			args: args{
				line: "Foo",
				apply: func(name string, directions *Directions) error {
					assert.Equal(t, name, "Foo")
					assert.Equal(t, directions.North, "")
					assert.Equal(t, directions.East, "")
					assert.Equal(t, directions.South, "")
					assert.Equal(t, directions.West, "")
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "Wrong format",
			args: args{
				line: "Foo foo=Bar east=Baz south=Qux west=Quux",
				apply: func(name string, directions *Directions) error {
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "Empty line",
			args: args{
				line: "",
				apply: func(name string, directions *Directions) error {
					return nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := processLine(tt.args.line, tt.args.apply); (err != nil) != tt.wantErr {
				t.Errorf("processLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
