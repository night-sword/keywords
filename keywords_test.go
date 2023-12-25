package keywords

import (
	"reflect"
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	K := NewKeywords([]string{"black", "pretty"})

	type args struct {
		str string
	}
	tests := []struct {
		name            string
		args            args
		wantAfterFilter string
		wantIsChange    bool
	}{
		{"case.1", args{"test black test"}, "test ***** test", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAfterFilter, gotIsChange := K.Filter(tt.args.str)
			if gotAfterFilter != tt.wantAfterFilter {
				t.Errorf("Filter() gotAfterFilter = %v, want %v", gotAfterFilter, tt.wantAfterFilter)
			}
			if gotIsChange != tt.wantIsChange {
				t.Errorf("Filter() gotIsChange = %v, want %v", gotIsChange, tt.wantIsChange)
			}
		})
	}
}

func TestFind(t *testing.T) {
	K := NewKeywords([]string{"hello", "fun"})
	K.AutoRefreshKeywords(func() (keywords []string, err error) {
		return []string{"hello", "fun", "good"}, nil
	}, time.Second)

	time.Sleep(time.Second)

	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		found []string
	}{
		{"case.1", args{"test test hello test fun test good"}, []string{"hello", "fun", "good"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := K.Find(tt.args.str)
			if !reflect.DeepEqual(found, tt.found) {
				t.Errorf("Filter() found = %v, want %v", found, tt.found)
			}
		})
	}
}
