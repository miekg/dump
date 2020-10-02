package dump

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/pkg/replacer"
	"github.com/coredns/coredns/request"

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
	state := request.Request{W: w, Req: r}
	rep := replacer.New()
	trw := dnstest.NewRecorder(w)
	fmt.Fprintln(output, rep.Replace(ctx, state, trw, format))
	return plugin.NextOrFailure(d.Name(), d.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (d Dump) Name() string { return "dump" }
