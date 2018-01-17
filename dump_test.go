package dump

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"
	"github.com/mholt/caddy"
	"github.com/miekg/dns"
)

func TestSetup(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	tests := []struct {
		input     string
		shouldErr bool
	}{
		// positive
		{
			`dump`, false,
		},
		// negative
		{
			`dump blah`, true,
		},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.input)
		err := setup(c)

		if test.shouldErr && err == nil {
			t.Fatalf("Test %d: Expected error but found %s for input %s", i, err, test.input)
		}

		if err != nil {
			if !test.shouldErr {
				t.Fatalf("Test %d: Expected no error but found one for input %s. Error was: %v", i, test.input, err)
			}
		}
	}
}

func TestDump(t *testing.T) {
	dump := Dump{}

	var f bytes.Buffer
	output = &f

	ctx := context.TODO()
	r := new(dns.Msg)
	r.SetQuestion("example.org.", dns.TypeA)
	r.Id = 42053

	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	dump.ServeDNS(ctx, rec, r)

	dumped := f.String()
	if !strings.Contains(dumped, "42053 A IN example.org. udp") {
		t.Errorf("Expected it to be dumped. Dumped string: %s", dumped)
	}
}
