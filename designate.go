package openstack

import (
	designate "github.com/PurplePlane897/libdns-openstack-designate"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *designate.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.openstack-designate",
		New: func() caddy.Module {
			return &Provider{new(designate.Provider)}
		},
	}
}

// Provision implements the Provisioner interface to initialize the OpenStack Designate client
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.AuthOpenStack.RegionName = repl.ReplaceAll(p.Provider.AuthOpenStack.RegionName, "")
	p.Provider.AuthOpenStack.AuthURL = repl.ReplaceAll(p.Provider.AuthOpenStack.AuthURL, "")
	p.Provider.AuthOpenStack.AuthType = repl.ReplaceAll(p.Provider.AuthOpenStack.AuthType, "")
	p.Provider.AuthOpenStack.ApplicationCredentialId = repl.ReplaceAll(p.Provider.AuthOpenStack.ApplicationCredentialId, "")
	p.Provider.AuthOpenStack.ApplicationCredentialSecret = repl.ReplaceAll(p.Provider.AuthOpenStack.ApplicationCredentialSecret, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// openstack-designate {
//     region_name <string>
//     tenant_id <string>
//     identity_api_version <string>
//     password <string>
//     username <string>
//     tenant_name <string>
//     endpoint_type <string>
//     auth_url <string>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "region_name":
				if d.NextArg() {
					p.Provider.AuthOpenStack.RegionName = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "auth_url":
				if d.NextArg() {
					p.Provider.AuthOpenStack.AuthURL = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "auth_type":
				if d.NextArg() {
					p.Provider.AuthOpenStack.AuthType = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "app_credential_id":
				if d.NextArg() {
					p.Provider.AuthOpenStack.ApplicationCredentialId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "app_credential_secret":
				if d.NextArg() {
					p.Provider.AuthOpenStack.ApplicationCredentialSecret = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
