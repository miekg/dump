package dump

import (
	"context"
	"log"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"

	"github.com/mholt/caddy"
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

// ServeDNS implements the plugin.Handler interface.
func (d Dump) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Printf("[DEBUG] %d %s %d\n", r.Id, r.Question[0].Name, r.Question[0].Qtype)
	return plugin.NextOrFailure(d.Name(), d.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (d Dump) Name() string { return "dump" }
