package tabletest

import (
	"testing"

	"github.com/vyacheslavprod/golibrary/prose"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData {
		{list: []string{}, want: ""},
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}
	for _, test := range tests {
		got := prose.JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}