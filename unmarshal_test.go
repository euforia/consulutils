package consulutils

import (
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

	client, err := NewConsulClient(testAddress, testDC)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	pairs, err := GetKVTree(client, "metrilyx/annotations", nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	var ccfg TestConsulConfig
	if err = Unmarshal(pairs, &ccfg); err != nil {
		t.Fatalf("%s\n", err)
	} else {
		t.Logf("%#v\n", ccfg)
	}

	var uCfg TestConsulConfig
	if _, err = GetKVTree(client, "metrilyx/annotations", &uCfg); err != nil {
		t.Fatalf("%s\n", err)
	} else {
		t.Logf("%#v\n", uCfg)
	}
}
