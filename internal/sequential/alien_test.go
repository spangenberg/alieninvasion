package sequential

import (
	"testing"

	"github.com/spangenberg/alieninvasion/internal"
)

func TestAlien_move(t *testing.T) {
	type fields struct {
		Name string
		City *internal.City
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Alien should move",
			fields: fields{
				City: &internal.City{
					North: &internal.City{},
				},
			},
			want: true,
		},
		{
			name: "Alien should not move",
			fields: fields{
				City: &internal.City{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &alien{
				city: tt.fields.City,
			}
			if got := a.move(); got != tt.want {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
			if tt.want && a.city == tt.fields.City {
				t.Error("Alien did not change city")
			}
		})
	}
}
