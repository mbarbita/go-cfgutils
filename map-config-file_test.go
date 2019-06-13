package cfgutils

import "testing"

func TestReadCfgFile(t *testing.T) {
	cfgMap := ReadCfgFile("cfg.ini", false)
	v := cfgMap["test1"]
	if v != "test" {
		t.Error("expected test, got:", v)
	}

	v = cfgMap["tEst9"]
	if v != "key with Upper case" {
		t.Error("expected key with Upper case, got:", v)
	}

}
