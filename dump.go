package dump

import (
	"context"
	"fmt"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	corelog "github.com/coredns/coredns/plugin/log"
	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/pkg/replacer"

	"github.com/mholt/caddy"
	"github.com/miekg/dns"
)

// Dump implement the plugin interface.
type Dump struct {
	Next plugin.Handler
}

func init() {
	caddy.RegisterPlugin("dump", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

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

const format = `{remote} ` + corelog.CommonLogEmptyValue + ` [{when}] {>id} {type} {class} {name} {proto} {port}`

// ServeDNS implements the plugin.Handler interface.
func (d Dump) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	rrw := dnstest.NewRecorder(w)
	rep := replacer.New(r, rrw, corelog.CommonLogEmptyValue)
	fmt.Println(rep.Replace(format))

	return plugin.NextOrFailure(d.Name(), d.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (d Dump) Name() string { return "dump" }
