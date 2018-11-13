package helpers

import (
	"testing"

	"github.com/bclicn/color"
)

func TestColorizedDenied(t *testing.T) {
	var testCases = []struct {
		input string
		want  string
	}{
		{"string", "string"},
		{"any_string", "any_string"},
		{"начало отклонен конец", color.Red("начало отклонен конец")},
		{"начало Запрет конец", color.Red("начало Запрет конец")},
	}
	for _, test := range testCases {
		if got := ColorizedDenied(test.input); got != test.want {
			t.Errorf("Event(%q) is %q. Need %q", test.input, got, test.want)
		}
	}
}

func TestColorizedWorker(t *testing.T) {
	var testCases = []struct {
		fullName string
		substr string
		want  string
	}{
		{"Иванов", "Иванов", color.Yellow("Иванов")},
		{"Иванов", "иванов", color.Yellow("Иванов")},
		{"Петров", "Иванов", "Петров"},
		{"Петров", "", "Петров"},
	}
	for _, test := range testCases {
		if got := ColorizedWorker(test.fullName, test.substr); got != test.want {
			t.Errorf("Worker(%q) is %q. Need %q", test.fullName, got, test.want)
		}
	}

}