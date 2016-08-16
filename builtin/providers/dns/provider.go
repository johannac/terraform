package dns

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for DNS dynamic updates.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DNS_SERVER", nil),
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  53,
			},
			"key_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DNS_KEY_NAME", nil),
			},
			"key_algorithm": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DNS_KEY_ALGORITHM", nil),
			},
			"key_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DNS_KEY_SECRET", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"dns_a_record_set": resourceDnsARecordSet(),
			"dns_ptr_record":   resourceDnsPtrRecord(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		server:    d.Get("server").(string),
		port:      d.Get("port").(int),
		keyname:   d.Get("key_name").(string),
		keyalgo:   d.Get("key_algorithm").(string),
		keysecret: d.Get("key_secret").(string),
	}

	return config.Client()
}
