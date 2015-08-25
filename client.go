package consulutils

import (
	"github.com/hashicorp/consul/api"
)

/*
 *  Helper consul client functions
 */

func NewConsulClient(address, datacenter string) (*api.Client, error) {

	cfg := api.DefaultConfig()
	cfg.Address = testAddress
	cfg.Datacenter = testDC

	return api.NewClient(cfg)
}

/*
   Get full tree under a key and optionally unmarshal.

    Args:
        client : Consul client
        key    : Key to query for.
        output : Unmarshal data to this interface{} if non-nil
*/
func GetKVTree(client *api.Client, key string, output interface{}) (pairs api.KVPairs, err error) {
	kv := client.KV()
	pairs, _, err = kv.List(key, nil)
	if output != nil {
		err = Unmarshal(pairs, output)
	}
	return
}
