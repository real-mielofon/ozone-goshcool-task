package counter

import "testing"

func TestCounter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test ", args: args{s: "aaabbbccccc"}, want: "a3b3c5"},
		{name: "test ", args: args{s: "aaabbbcccccaaaaa"}, want: "a8b3c5"},
		{name: "test ", args: args{s: "zzzzcccUUUuu"}, want: "U3c3u2z4"},
		{name: "test ", args: args{s: "ЯЯЯБББддд"}, want: "Б3Я3д3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Counter(tt.args.s); got != tt.want {
				t.Errorf("Counter() = (%d) '%v', want (%d) '%v'", len(got), got, len(tt.want), tt.want)
			}
		})
	}
}
