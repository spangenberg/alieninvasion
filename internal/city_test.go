package internal

import (
	"reflect"
	"testing"
)

type alien struct{}

func (a *alien) String() string {
	return "alien"
}

func TestCity_AddAlien(t *testing.T) {
	c := &City{}
	c.AddAlien(&alien{})
	if len(c.Aliens) != 1 {
		t.Error("Alien was not added")
	}
}

func TestCity_MapFormat(t *testing.T) {
	type fields struct {
		Name  string
		North *City
		East  *City
		South *City
		West  *City
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "All neighbors",
			fields: fields{
				Name:  "City",
				North: &City{name: "North"},
				East:  &City{name: "East"},
				South: &City{name: "South"},
				West:  &City{name: "West"},
			},
			want: "City north=North east=East south=South west=West",
		},
		{
			name: "Few neighbors",
			fields: fields{
				Name:  "City",
				North: &City{name: "North"},
				West:  &City{name: "West"},
			},
			want: "City north=North west=West",
		},
		{
			name: "No neighbors",
			fields: fields{
				Name: "City",
			},
			want: "City",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				name:  tt.fields.Name,
				North: tt.fields.North,
				East:  tt.fields.East,
				South: tt.fields.South,
				West:  tt.fields.West,
			}
			if got := c.MapFormat(); got != tt.want {
				t.Errorf("MapFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCity_RandomNeighbor(t *testing.T) {
	type fields struct {
		North *City
		East  *City
		South *City
		West  *City
	}
	tests := []struct {
		name   string
		fields fields
		want   *City
	}{
		{
			name: "All neighbors",
			fields: fields{
				North: &City{name: "City"},
				East:  &City{name: "City"},
				South: &City{name: "City"},
				West:  &City{name: "City"},
			},
			want: &City{name: "City"},
		},
		{
			name: "Few neighbors",
			fields: fields{
				North: &City{name: "City"},
				West:  &City{name: "City"},
			},
			want: &City{name: "City"},
		},
		{
			name: "No neighbors",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				North: tt.fields.North,
				East:  tt.fields.East,
				South: tt.fields.South,
				West:  tt.fields.West,
			}
			if got := c.RandomNeighbor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomNeighbor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCity_RemoveAlien(t *testing.T) {
	c := &City{
		Aliens: []Alien{
			&alien{},
		},
	}
	c.RemoveAlien()
	if len(c.Aliens) != 0 {
		t.Error("Alien was not removed")
	}
}
