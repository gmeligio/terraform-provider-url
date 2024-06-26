---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netparse_domain Data Source - netparse"
subcategory: ""
description: |-
  Parses Public Suffix List properties from a domain. For more details on the domain parts, see What is a Domain Name? https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_domain_name.
---

# netparse_domain (Data Source)

Parses Public Suffix List properties from a domain. For more details on the domain parts, see [What is a Domain Name?](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_domain_name).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) The host that identifies the domain name

### Read-Only

- `domain` (String) The domain name. It's the tld plus one more label. For example: example.com for host foo.example.com
- `manager` (String) The manager is the entity that manages the domain. It can be one of the following: ICANN, Private, or None.
- `sld` (String) The second-level domain (SLD) is the label to the left of the effective TLD. For example: example for example.com, or foo for foo.co.uk
- `subdomain` (String) The subdomain is the left part of the host that is not the domain. For example: www for www.example.com, mail for mail.foo.org, blog for blog.bar.org
- `tld` (String) The effective top-level domain (eTLD) of the domain. This is the public suffix of the domain. For example: com for example.com, or co.uk for foo.co.uk
