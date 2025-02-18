# routeros_ip_dns_record (Resource)
Creates a DNS record on the MikroTik device.

## Example Usage
```terraform
resource "routeros_ip_dns_record" "name_record" {
  name    = "router.lan"
  address = "192.168.88.1"
}

resource "routeros_ip_dns_record" "regexp_record" {
  regexp  = ".*pool.ntp.org"
  address = "192.168.88.1"
}

resource "routeros_dns_record" "aaaa_record" {
  name            = "ipv6.lan"
  address         = "ff00::1"
  type            = "AAAA"
}
  
resource "routeros_dns_record" "cname_record" {
  name            = "cname.lan"
  cname           = "ipv4.lan"
  type            = "CNAME"
}
  
resource "routeros_dns_record" "fwd_record" {
  name            = "fwd.lan"
  forward_to      = "127.0.0.1"
  type            = "FWD"
}
  
resource "routeros_dns_record" "mx_record" {
  name            = "mx.lan"
  mx_exchange     = "127.0.0.1"
  mx_preference   = 10
  type            = "MX"
}
  
resource "routeros_dns_record" "ns_record" {
  name            = "ns.lan"
  ns              = "127.0.0.1"
  type            = "NS"
}
  
resource "routeros_dns_record" "nxdomain_record" {
  name            = "nxdomain.lan"
  type            = "NXDOMAIN"
}
  
resource "routeros_dns_record" "srv_record" {
  name            = "srv.lan"
  srv_port        = 8080
  srv_priority    = 10
  srv_target      = "127.0.0.1"
  srv_weight      = 100
  type            = "SRV"
}

resource "routeros_dns_record" "txt_record" {
  name            = "_acme-challenge.yourwebsite.com"
  text            = "dW6MrI3nBy3eJgYWH3QAg1Cwk_TvjFESOuKo+mp6nm1"
  type            = "TXT"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address` (String) The A record to be returend from the DNS hostname.
- `address_list` (String) Name of the Firewall address list to which address must be dynamically added when some request matches the entry.
- `cname` (String) Alias name for a domain name.
- `comment` (String)
- `disabled` (Boolean)
- `forward_to` (String) The IP address of a domain name server to which a particular DNS request must be forwarded.
- `match_subdomain` (Boolean) Whether the record will match requests for subdomains.
- `mx_exchange` (String) The domain name of the MX server.
- `mx_preference` (Number) Preference of the particular MX record.
- `name` (String) The name of the DNS hostname to be created.
- `ns` (String) Name of the authoritative domain name server for the particular record.
- `regexp` (String) DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, RouterOS converts DNS names to lowercase, you should write regex only with lowercase letters.
- `srv_port` (Number) The TCP or UDP port on which the service is to be found.
- `srv_priority` (Number) Priority of the particular SRV record.
- `srv_target` (String) The canonical hostname of the machine providing the service ends in a dot.
- `srv_weight` (String) Weight of the particular SRC record.
- `text` (String) Textual information about the domain name.
- `ttl` (String) The ttl of the DNS record.
- `type` (String) Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT

### Read-Only

- `dynamic` (Boolean) Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dns/static get [print show-ids]]
terraform import routeros_ip_dns_record.name_record "*0"
```
