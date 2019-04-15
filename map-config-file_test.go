package cfgutils

import "testing"

func TestReadCfgFile(t *testing.T) {
	cfgMap := ReadCfgFile("cfg.ini", true)
	v := cfgMap["test1"]
	if v != "test" {
		t.Error("expected test, got:", v)
	}
}
