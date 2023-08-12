package config

import (
	"reflect"
	"regexp"

	"github.com/raulscr/arkhanna-master-server/types"
)

// This package describes how to read some simple config text pattern
// into struct with user defined tags, just as xml or json from encoding package does

// Example:
// some/file.cfg
//|
//|// comments comments
//|; more comments
//|some-config: teste
//|some-config-2: "teste 2" ; comments
//|some-config-3: "test//;e 3" // comments
//|some-config-int: 3423 // comments
//|some-config-bool: false // comments
//|

// some/file.go
// // struct that will be loaded
// type TEST_STR struct {
// 	SomeConfig     string `cfg:"some-config"`      // will have value string("teste")
// 	SomeConfig2    string `cfg:"some-config-2"`    // will have value string("teste 2")
// 	SomeConfig3    string `cfg:"some-config-3"`    // will have value string("test//;e 2")
// 	SomeConfigInt  string `cfg:"some-config-int"`  // will have value int(3423)
// 	SomeConfigBool string `cfg:"some-config-bool"` // will have value boolean(false)
// }

// Note: we should realy invert dependency from regex,
// cause the config pattern may change, and we want it to be a little more generic

const key_tag string = `cfg`

const key_rgx string = `[\w_-]+`

const s_rgx string = `[\t\f ]*`

const config_regex string = `(?m)^(:?(` + key_rgx + `)` + s_rgx + `\:` + s_rgx + `\"(.+?)\"|(` + key_rgx + `)` + s_rgx + `\:` + s_rgx + `(.+?))` + s_rgx + `(:?\/\/|;)?$`

type ConfigMap map[string]string

// Load the configs
func LoadConfigs(cfg_string string) (ConfigMap, error) {
	// TODO maybe it's a good idea to make it more dynamic, for example, to pass the driver
	// then you can define externally if you'll read a json, a .cfg, or wharever regex defined config file (just like this one, i never saw this cfg pattern before)
	var re *regexp.Regexp = regexp.MustCompile(config_regex)

	var conf = make(ConfigMap)

	// TODO: improve loop, maybe improvig regex to avoid empty strings
	for _, match := range re.FindAllStringSubmatch(cfg_string, -1) {
		var key string
		var value string
		for i, _ := range match {
			if i+1 < len(match) && match[i] != "" && match[i+1] != "" {
				key = match[i]
				value = match[i+1]
			}
		}

		if key != "" {
			conf[key] = value
		}
	}

	return conf, nil
}

func (confs ConfigMap) SetValuesFromMap(v any) error {
	// TODO return error if v isn't a reference
	var err error = nil
	var t reflect.Type = reflect.ValueOf(v).Elem().Type()
	for i := 0; i < t.NumField(); i++ {
		tag_name, ok := t.Field(i).Tag.Lookup(key_tag)
		if ok {
			err = types.SetValueFromString(reflect.ValueOf(v).Elem().Field(i), confs[tag_name])
			if err != nil {
				return err
			}
		}
	}

	return err
}

