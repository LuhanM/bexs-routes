package scale

import (
	"reflect"
	"testing"
)

func TestGetScale(t *testing.T) {
	type args struct {
		origin      string
		destination string
	}
	tests := []struct {
		name    string
		args    args
		want    TypeScale
		wantErr bool
	}{
		// TODO: Add test cases.
		{"TestGetScale_sucesso", args{"GRU", "CDG"}, TypeScale{Origin: "GRU", Destination: "CDG", Cost: 100}, false},
		{"TestGetScale_sucesso", args{"FLN", "ORL"}, TypeScale{Origin: "FLN", Destination: "ORL", Cost: 2500}, false},
		{"TestGetScale_sucesso", args{"ITA", "FRC"}, TypeScale{}, true},
	}
	Scales = append(Scales, TypeScale{Origin: "GRU", Destination: "CDG", Cost: 100})
	Scales = append(Scales, TypeScale{Origin: "BRA", Destination: "ARG", Cost: 500})
	Scales = append(Scales, TypeScale{Origin: "FLN", Destination: "ORL", Cost: 2500})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetScale(tt.args.origin, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetScale() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
