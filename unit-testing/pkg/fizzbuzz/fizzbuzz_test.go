package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		total  int64
		fizzAt int64
		buzzAt int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "normaltest1",
			args: args{total: 16, fizzAt: 3, buzzAt: 5},
			want: []string{"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
				"16"}},
		{name: "normaltest2",
			args: args{total: 16, fizzAt: 5, buzzAt: 3},
			want: []string{"1",
				"2",
				"Buzz",
				"4",
				"Fizz",
				"Buzz",
				"7",
				"8",
				"Buzz",
				"Fizz",
				"11",
				"Buzz",
				"13",
				"14",
				"FizzBuzz",
				"16"}},
		{name: "edge case total negative",
			args: args{total: -1, fizzAt: 5, buzzAt: 3},
			want: []string{}},
		{name: "edge case fizzAt negative",
			args: args{total: 16, fizzAt: -3, buzzAt: 5},
			want: []string{"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
				"16"}},
		{name: "edge case buzzAt negative",
			args: args{total: 16, fizzAt: 3, buzzAt: -5},
			want: []string{"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
				"16"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.total, tt.args.fizzAt, tt.args.buzzAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
