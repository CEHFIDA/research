package core

import "testing"
import "reflect"
import "time"
import "runtime"

func TestGeneralOptions(t *testing.T) {
	opts := &GeneralOptions{}
	var units = []struct {
		got interface{}
		exp interface{}
	}{
		{opts.Verbose, false},
		{opts.ColorSupport, false},
		{opts.AvailableCores, 0},
		{opts.Recursive, false},
		{opts.PassiveOnly, false},
		{opts.IgnoreErrors, false},
		{opts.OutputType, ""},
		{opts.OutputDir, ""},
		{len(opts.TargetDomains), 0},
		{len(opts.Sources), 0},
		{len(opts.Resolvers), 0},
	}
	for _, u := range units {
		if !reflect.DeepEqual(u.exp, u.got) {
			t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
		}
	}
}

func TestDefaultDNSResolvers(t *testing.T) {
	var units = []struct {
		got interface{}
		exp interface{}
	}{
		{len(defaultDNSResolvers), 8},
	}
	for _, u := range units {
		if !reflect.DeepEqual(u.exp, u.got) {
			t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
		}
	}
}

func TestNewGeneralOptions(t *testing.T) {
	opts := NewDefaultGeneralOptions()
	var units = []struct {
		got interface{}
		exp interface{}
	}{
		{opts.Verbose, false},
		{opts.ColorSupport, true},
		{opts.AvailableCores, runtime.NumCPU()},
		{opts.DefaultTimeout, time.Duration(5 * time.Second)},
		{opts.Recursive, false},
		{opts.PassiveOnly, false},
		{opts.IgnoreErrors, false},
		{opts.OutputType, "plaintext"},
		{opts.OutputDir, ""},
		{len(opts.TargetDomains), 0},
		{len(opts.Sources), 0},
		{len(opts.Resolvers), 8},
	}
	for _, u := range units {
		if !reflect.DeepEqual(u.exp, u.got) {
			t.Fatalf("expected '%v', got '%v'", u.exp, u.got)
		}
	}
}

func ExampleGeneralOptions() {
	opts := GeneralOptions{}
	opts.Print()
	// Output:
	// Verbose:	 'false'
	// ColorSupport:	 'false'
	// AvailableCores:	 '0'
	// DefaultTimeout:	 '0s'
	// TargetDomains:	 '[]'
	// Recursive:	 'false'
	// PassiveOnly:	 'false'
	// IgnoreErrors:	 'false'
	// OutputType:	 ''
	// Sources:	 '[]'
	// OutputDir:	 ''
	// Resolvers:	 '[]'
}

func ExampleNewDefaultGeneralOptions() {
	opts := NewDefaultGeneralOptions()
	opts.Print()
	// Output:
	// Verbose:	 'false'
	// ColorSupport:	 'true'
	// AvailableCores:	 '8'
	// DefaultTimeout:	 '5s'
	// TargetDomains:	 '[]'
	// Recursive:	 'false'
	// PassiveOnly:	 'false'
	// IgnoreErrors:	 'false'
	// OutputType:	 'plaintext'
	// Sources:	 '[]'
	// OutputDir:	 ''
	// Resolvers:	 '[1.1.1.1 1.0.0.1 8.8.8.8 8.8.4.4 9.9.9.9 149.112.112.112 208.67.222.222 208.67.220.220]'
}
