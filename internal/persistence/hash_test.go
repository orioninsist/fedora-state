package persistence

import "testing"

func TestChanged(t *testing.T) {

	tests := []struct {
		name string
		old  string
		new  string
		want bool
	}{
		{
			name: "first scan",
			old:  "",
			new:  "abc",
			want: true,
		},
		{
			name: "same state",
			old:  "abc",
			new:  "abc",
			want: false,
		},
		{
			name: "changed state",
			old:  "abc",
			new:  "def",
			want: true,
		},
	}

	for _, test := range tests {

		got := Changed(
			test.old,
			test.new,
		)

		if got != test.want {
			t.Fatalf(
				"%s: got %v want %v",
				test.name,
				got,
				test.want,
			)
		}
	}
}
