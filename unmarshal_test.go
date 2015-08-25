package consulutils

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

var (
	testAddress = "app1.inv.toolsash1.cloudsys.tmcs:8080"
	testDC      = "toolsash1"
	testCfgKey  = "metrilyx/annotations"
)

type TestConsulConfig struct {
	Host    string `consul:"metrilyx/annotations/dataprovider/host"`
	Port    int64  `consul:"metrilyx/annotations/dataprovider/port"`
	Index   string `consul:"metrilyx/annotations/dataprovider/index"`
	Enabled bool   `consul:"metrilyx/annotations/enabled"`
}

func Test_Unmarshal(t *testing.T) {

	cfg := api.DefaultConfig()
	cfg.Address = testAddress
	cfg.Datacenter = testDC

	client, err := api.NewClient(cfg)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	kv := client.KV()

	pairs, _, err := kv.List(testCfgKey, nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	var ccfg TestConsulConfig
	if err = Unmarshal(pairs, &ccfg); err != nil {
		t.Fatalf("%s\n", err)
	} else {
		t.Logf("%#v\n", ccfg)
	}
}
