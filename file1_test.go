package ebitin_init_failure

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reflect"
	"testing"
)

func TestIdentity(t *testing.T) {
	type args struct {
		i *ebiten.Image
	}
	tests := []struct {
		name string
		args args
		want *ebiten.Image
	}{
		{
			name: "nil",
			args: args{
				i: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Identity(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Identity() = %v, want %v", got, tt.want)
			}
		})
	}
}
