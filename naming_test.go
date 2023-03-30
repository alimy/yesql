package yesql

import "testing"

func TestFieldUp(t *testing.T) {
	ns := NewFieldUpNamingStrategy()
	for name, expect := range map[string]string{
		"abc_de__fg":         "AbcDeFg",
		"_abc__de_fg_":       "AbcDeFg",
		"____ab__c_de_fg___": "AbCDeFg",
		"Abc_de_fg_":         "AbcDeFg",
		"AbC__De_fG_":        "AbCDeFG",
		"_Abc_De_fG_":        "AbcDeFG",
	} {
		if value := ns.FiledNaming(name); value != expect {
			t.Errorf("want %s bug got %s when namming %s", expect, value, name)
		}
	}
}
