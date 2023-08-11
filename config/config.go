package config

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
