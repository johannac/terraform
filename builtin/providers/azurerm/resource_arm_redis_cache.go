package azurerm

import (
	"log"

	"fmt"

	"github.com/Azure/azure-sdk-for-go/arm/redis"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jen20/riviera/azure"
)

func resourceArmRedisCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmRedisCacheCreate,
		Read:   resourceArmRedisCacheRead,
		Delete: resourceArmRedisCacheDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"location": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				StateFunc: azureRMNormalizeLocation,
			},

			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"size": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "C1",
			},

			"sku": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Standard",
			},

			"enable_non_ssl_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"shard_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"redis_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"tags": tagsSchema(),
		},
	}
}

func resourceArmRedisCacheCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient)
	redisClient := client.redisClient

	log.Printf("[INFO] preparing arguments for Azure ARM Redis Cache creation.")
	name := d.Get("name").(string)

	location := d.Get("location").(string)
	resGroup := d.Get("resource_group_name").(string)
	tags := d.Get("tags").(map[string]interface{})
	size := d.Get("size").(string)
	enableNonSslPort := d.Get("enable_non_ssl_port").(bool)
	configOptions := d.Get("redis_configuration").(map[string]interface{})

	properties := redis.Properties{
		Sku: &redis.Sku{
			Name:     redis.SkuName(d.Get("sku").(string)),
			Family:   redis.SkuFamily(size[0]),
			Capacity: azure.Int32(int32(size[1])),
		},
		EnableNonSslPort: &enableNonSslPort,
	}

	if v, ok := d.GetOk("shard_count"); ok {
		shardCount := v.(int32)
		properties.ShardCount = &shardCount
	}

	if v, ok := d.GetOk("virtual_network_id"); ok {
		vNetId := v.(string)
		properties.VirtualNetwork = &vNetId
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		subnetId := v.(string)
		properties.Subnet = &subnetId
	}

	if v, ok := d.GetOk("static_ip"); ok {
		staticIp := v.(string)
		properties.StaticIP = &staticIp
	}

	if v, ok := d.GetOk("redis_configuration"); ok {
		config := expandRedisConfiguraton(v.(map[string]interface{}))
		properties.RedisConfiguration = config
	}

	params := redis.CreateOrUpdateParameters{
		Name:       &name,
		Location:   &location,
		Properties: &properties,
		Tags:       expandTags(tags),
	}

	_, err := redisClient.CreateOrUpdate(resGroup, name, params)
	if err != nil {
		return err
	}

	read, err := redisClient.Get(resGroup, name)
	if err != nil {
		return err
	}
	if read.ID == nil {
		return fmt.Errorf("Cannot read Redis Cache %s (resource group %s) ID", name, resGroup)
	}

	d.SetId(*read.ID)

	return resourceArmRedisCacheRead(d, meta)
}

func resourceArmRedisCacheRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceArmRedisCacheDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func expandRedisConfiguraton(configMap map[string]interface{}) *map[string]*string {
	output := make(map[string]*string, len(configMap))

	for i, v := range configMap {
		value, _ := v.(string)
		output[i] = &value
	}

	return &output
}
