package provider

import (
	"context"

	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// componentsDataSourceModel describes the data source data model.
// References used.
// https://registry.terraform.io/modules/matti/urlparse/external/latest
// https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/url_parse
type componentsDataSourceModel struct {
	Url         types.String `tfsdk:"url"`
	Authority   types.String `tfsdk:"authority"`
	Protocol    types.String `tfsdk:"protocol"`
	Scheme      types.String `tfsdk:"scheme"`
	Credentials types.String `tfsdk:"credentials"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	Host        types.String `tfsdk:"host"`
	Port        types.String `tfsdk:"port"`
	Path        types.String `tfsdk:"path"`
	Search      types.String `tfsdk:"search"`
	Query       types.String `tfsdk:"query"`
	Hash        types.String `tfsdk:"hash"`
	Fragment    types.String `tfsdk:"fragment"`
}

func (d componentsDataSourceModel) validate(_ context.Context) diag.Diagnostics {
	var diags diag.Diagnostics

	// if d.Host.IsUnknown() || d.Host.IsNull() {
	// 	return diags
	// }

	// host := d.Host.ValueString()

	// eTLD, icann := publicsuffix.PublicSuffix(host)

	// manager := findManager(icann, eTLD)

	// if manager == "None" {
	// 	diags.AddAttributeError(
	// 		path.Root("host"),
	// 		"Invalid Attribute Configuration",
	// 		"Expected host to have as a manager either ICANN or Private.",
	// 	)
	// }

	return diags
}

func (d *componentsDataSourceModel) update(_ context.Context) diag.Diagnostics {
	var diags diag.Diagnostics

	// host := d.Host.ValueString()

	// eTLD, icann := publicsuffix.PublicSuffix(host)
	// d.TLD = types.StringValue(eTLD)

	// sld, err := extractSld(host, eTLD)
	// if err != nil {
	// 	diags.AddAttributeError(
	// 		path.Root("sld"),
	// 		"Invalid Attribute Configuration",
	// 		err.Error(),
	// 	)
	// }
	// d.SLD = types.StringValue(sld)

	// domain := sld + "." + eTLD
	// d.Domain = types.StringValue(domain)

	// manager := findManager(icann, eTLD)
	// d.Manager = types.StringValue(manager)

	// subdomain := extractSubdomain(host, domain)
	// d.Subdomain = types.StringValue(subdomain)

	rawURL := d.Url.ValueString()

	parsed, err := url.Parse(rawURL)
	if err != nil {
		diags.AddError(
			"Invalid URL",
			err.Error(),
		)
		return diags
	}

	authority := renderAuthority(parsed)
	d.Authority = types.StringValue(authority)

	scheme := parsed.Scheme
	d.Scheme = types.StringValue(scheme)

	protocol := scheme + ":"
	d.Protocol = types.StringValue(protocol)

	credentials := parsed.User.String()
	d.Credentials = types.StringValue(credentials)

	username := parsed.User.Username()
	d.Username = types.StringValue(username)

	password, _ := parsed.User.Password()
	d.Password = types.StringValue(password)

	host := parsed.Hostname()
	d.Host = types.StringValue(host)

	port := parsed.Port()
	d.Port = types.StringValue(port)

	path := parsed.Path
	d.Path = types.StringValue(path)

	search := renderSearch(parsed)
	d.Search = types.StringValue(search)

	query := parsed.RawQuery
	d.Query = types.StringValue(query)

	fragment := parsed.Fragment
	d.Fragment = types.StringValue(fragment)

	hash := renderHash(parsed)
	d.Hash = types.StringValue(hash)

	return diags
}

func renderHash(parsed *url.URL) string {
	fragment := parsed.Fragment
	if fragment == "" {
		return ""
	}

	return "#" + fragment
}

func renderSearch(parsed *url.URL) string {
	query := parsed.RawQuery
	if query == "" {
		return ""
	}

	return "?" + query
}

func renderAuthority(parsed *url.URL) string {
	credentials := parsed.User.String()
	port := parsed.Port()

	var credentialsComponent string
	if credentials != "" {
		credentialsComponent = credentials + "@"
	}

	var portComponent string
	if port != "" {
		portComponent = ":" + port
	}

	return credentialsComponent + parsed.Hostname() + portComponent
}