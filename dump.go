package dump

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/pkg/replacer"

	"github.com/caddyserver/caddy"
	"github.com/miekg/dns"
)

// Dump implement the plugin interface.
type Dump struct {
	Next plugin.Handler
}

func init() { plugin.Register("dump", setup) }

func setup(c *caddy.Controller) error {
	for c.Next() {
		if c.NextArg() {
			return plugin.Error("dump", c.ArgErr())
		}
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Dump{Next: next}
	})

	return nil
}

const format = `{remote} ` + replacer.EmptyValue + ` {>id} {type} {class} {name} {proto} {port}`

var output io.Writer = os.Stdout

// ServeDNS implements the plugin.Handler interface.
func (d Dump) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	rrw := dnstest.NewRecorder(w)
	rep := replacer.New(r, rrw, replacer.EmptyValue)
	fmt.Fprintln(output, rep.Replace(format))

	return plugin.NextOrFailure(d.Name(), d.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (d Dump) Name() string { return "dump" }
