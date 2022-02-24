package models

import (
	"reflect"
	"testing"
)

func TestFindPeopleByPhoneNumber(t *testing.T) {
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name string
		args args
		want []*Person
	}{
		{name: "Find John Joe", args: args{phoneNumber: "+1 (800) 555-1212"}, want: people[0:1]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPeopleByPhoneNumber(tt.args.phoneNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPeopleByPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
