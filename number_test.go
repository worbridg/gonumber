package gonumber

import (
	"testing"
)

func TestInt_Is(t *testing.T) {
	n := NewInt(3)
	if n.Is(2) {
		t.Errorf("n must be 3")
	}
}

func TestInt_IsNot(t *testing.T) {
	n := NewInt(3)
	if n.IsNot(3) {
		t.Errorf("n must be 3")
	}
}

func TestInt_Increment(t *testing.T) {
	n := NewInt(0)
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
	n := NewInt(0)
	n.Increment()
	if n.Was(1) {
		t.Errorf("previous n must be 0")
	}
}

func TestInt_WasNot(t *testing.T) {
	n := NewInt(0)
	n.Increment()
	if n.WasNot(0) {
		t.Errorf("previous n must be 0")
	}
}

func TestInt_Add(t *testing.T) {
	n := NewInt(0)
	n.Add(3)
	if n.IsNot(3) {
		t.Errorf("n must be 3")
	}
	n.ShouldBe(1, 2, 3, 6)
	if err := n.Add(2); err == nil {
		t.Error(err)
	}
	n = NewInt(0)
	n.WillBe(3)
	if err := n.Add(2); err == nil {
		t.Error(err)
	}
}
func TestInt_Sub(t *testing.T) {
	n := NewInt(0)
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
			numbers: Numbers{NewInt(1), NewInt(2), NewInt(3)},
			args: args{
				n: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name:    "Random",
			numbers: Numbers{NewInt(1), NewInt(2), NewInt(3)},
			args: args{
				n: []int{3, 1, 2},
			},
			want: false,
		},
		{
			name:    "notEqual",
			numbers: Numbers{NewInt(1), NewInt(2), NewInt(3)},
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

func TestNumbers_AreSame(t *testing.T) {
	tests := []struct {
		name    string
		numbers Numbers
		want    bool
	}{
		{
			name:    "Same",
			numbers: Numbers{NewInt(1), NewInt(1), NewInt(1)},
			want:    true,
		},
		{
			name:    "NotSame",
			numbers: Numbers{NewInt(1), NewInt(1), NewInt(2)},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.numbers.AreSame(); got != tt.want {
				t.Errorf("Numbers.AreSame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_UpTo(t *testing.T) {
	number := NewInt(3)
	numbers := number.UpTo(5)
	if numbers.AreNot(3, 4, 5) {
		t.Errorf("3 up to 5")
	}
	numbers = number.UpTo(2)
	if numbers.AreNot(3) {
		t.Errorf("numbers must be 3")
	}
}

func TestNumber_ShouldBePositive(t *testing.T) {
	number, err := NewInt(3).ShouldBePositive()
	if err != nil {
		t.Error(err)
	}
	if err := number.ChangeTo(-1); err == nil {
		t.Errorf("the number should be positive")
	}
	if err := number.ChangeTo(5); err != nil {
		t.Error(err)
	}
}

func TestNumber_ShouldBeBetween(t *testing.T) {
	number, err := NewInt(3).ShouldBeBetween(1, 5)
	if err != nil {
		t.Error(err)
	}
	if err := number.ChangeTo(0); err == nil {
		t.Errorf("the number should be between 1 and 5")
	}
	if err := number.ChangeTo(4); err != nil {
		t.Error(err)
	}
}

func TestNumber_In(t *testing.T) {
	numbers := NewInt(1).UpTo(5)
	number := NewInt(8)
	if number.In(numbers) {
		t.Errorf("8 shouldn't be in numbers")
	}
}

func TestEvery(t *testing.T) {
	numbers := NewInt(0).UpTo(3)
	n := NewInt(1)
	numbers = append(numbers, n)
	if Every(n).In(numbers).AreNot(1, 1) {
		t.Errorf("numbers should be 1 and 1")
	}
}
