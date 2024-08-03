Hurricane Electric module for Caddy
===================================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy).
It can be used to manage DNS records with [Hurricane Electric](https://dns.he.net/).

## Caddy module name

```
dns.providers.he
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "he",
				"api_key": "API_KEY"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns he API_KEY
}
```

```
# one site
tls {
	dns he {
            api_key API_KEY
        }
}
```
