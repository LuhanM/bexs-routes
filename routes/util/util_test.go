package util

import (
	"reflect"
	"testing"
)

func TestSplitRoute(t *testing.T) {
	type args struct {
		route string
	}
	tests := []struct {
		name            string
		args            args
		wantOrigin      string
		wantDestination string
		wantErr         bool
	}{
		// TODO: Add test cases.
		{"testeSucessoComMinusculas", args{"gru-cdg"}, "GRU", "CDG", false},
		{"testeSucessoComMaiusculas", args{"GRU-CDG"}, "GRU", "CDG", false},
		{"testeFalha", args{"grucdg"}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrigin, gotDestination, err := SplitRoute(tt.args.route)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOrigin != tt.wantOrigin {
				t.Errorf("SplitRoute() gotOrigin = %v, want %v", gotOrigin, tt.wantOrigin)
			}
			if gotDestination != tt.wantDestination {
				t.Errorf("SplitRoute() gotDestination = %v, want %v", gotDestination, tt.wantDestination)
			}
		})
	}
}

func TestRoute_addStep(t *testing.T) {
	type args struct {
		steps []string
		step  string
	}
	tests := []struct {
		name  string
		route *Route
		args  args
		want  *Route
	}{
		{
			"teste1",
			&Route{
				Steps:       []string{},
				CompleteWay: "",
			},
			args{[]string{}, "ABC"},
			&Route{
				Steps:       []string{"ABC"},
				CompleteWay: "ABC",
			},
		},
		{
			"teste2",
			&Route{
				Steps:       []string{"ABC"},
				CompleteWay: "ABC",
			},
			args{[]string{"ABC"}, "DEF"},
			&Route{
				Steps:       []string{"ABC", "DEF"},
				CompleteWay: "ABC - DEF",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.route.addStep(tt.args.steps, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Route.addStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
