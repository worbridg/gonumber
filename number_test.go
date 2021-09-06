package gonumber

import (
	"testing"
)

func TestInt_Is(t *testing.T) {
	n := New(3)
	if n.Is(2) {
		t.Errorf("n must be 3")
	}
}

func TestInt_IsNot(t *testing.T) {
	n := New(3)
	if n.IsNot(3) {
		t.Errorf("n must be 3")
	}
}

func TestInt_Increment(t *testing.T) {
	n := New(0)
	n.Increment()
	if n.IsNot(1) {
		t.Errorf("n must be 1")
	}
	n.WillBe(2)
	if err := n.Increment(); err != nil {
		t.Error(err)
	}
	n.WillBe(4)
	if err := n.Increment(); err == nil {
		t.Error(err)
	}
	n.WillBe(3)
	if err := n.Increment(); err != nil {
		t.Error(err)
	}
	n.ShouldBe(3)
	if err := n.Increment(); err == nil {
		t.Error(err)
	}
	n.ShouldBe(3, 4)
	if err := n.Increment(); err != nil {
		t.Error(err)
	}
}

func TestInt_Was(t *testing.T) {
	n := New(0)
	n.Increment()
	if n.Was(1) {
		t.Errorf("previous n must be 0")
	}
}

func TestInt_WasNot(t *testing.T) {
	n := New(0)
	n.Increment()
	if n.WasNot(0) {
		t.Errorf("previous n must be 0")
	}
}

func TestInt_Add(t *testing.T) {
	n := New(0)
	n.Add(3)
	if n.IsNot(3) {
		t.Errorf("n must be 3")
	}
	n.ShouldBe(1, 2, 3, 6)
	if err := n.Add(2); err == nil {
		t.Error(err)
	}
	n = New(0)
	n.WillBe(3)
	if err := n.Add(2); err == nil {
		t.Error(err)
	}
}
func TestInt_Sub(t *testing.T) {
	n := New(0)
	n.Sub(3)
	if n.IsNot(-3) {
		t.Errorf("n must be -3")
	}
	n.ShouldBe(-1, -2, -3)
	if err := n.Sub(2); err == nil {
		t.Error(err)
	}
}

func TestNumbers_Are(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name    string
		numbers Numbers
		args    args
		want    bool
	}{
		{
			name:    "Equal",
			numbers: Numbers{New(1), New(2), New(3)},
			args: args{
				n: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name:    "Random",
			numbers: Numbers{New(1), New(2), New(3)},
			args: args{
				n: []int{3, 1, 2},
			},
			want: false,
		},
		{
			name:    "notEqual",
			numbers: Numbers{New(1), New(2), New(3)},
			args: args{
				n: []int{1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.numbers.Are(tt.args.n...); got != tt.want {
				t.Errorf("Numbers.Are() = %v, want %v", got, tt.want)
			}
		})
	}
}
