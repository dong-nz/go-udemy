package toolkit

import "testing"

func TestTools_RandomString(t1 *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"generate string has 10 length", args{10}, 10},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tools{}
			if got := t.RandomString(tt.args.n); len(got) != tt.want {
				t1.Errorf("RandomString() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
