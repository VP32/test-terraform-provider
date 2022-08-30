package hashicups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VPFOO_HOST", "http://localhost:8080"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"vpfoo_fooserver": resourceServer(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"vpfoo_fooserver": dataSourceOrder(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := make(map[string]string)
	c["host"] = d.Get("host").(string)

	// todo Добавить логику подключения к API

	return c, diags
}
