package types

import (
	"errors"
	"reflect"
	"strconv"
)

/// Public

type Settable interface {
	SetValueFromString(str_val string) error
}

type Value reflect.Value

/* Convert some stringed value to correct type if possible
 * and set it into the field.
 */
func SetValueFromString(f reflect.Value, str_val string) error {

	var conv_map = map[reflect.Kind]func(reflect.Value, string) error{
		reflect.String:  setValueString,
		reflect.Bool:    setValueBool,
		reflect.Int:     setValueInt,
		reflect.Int8:    setValueInt8,
		reflect.Int16:   setValueInt16,
		reflect.Int32:   setValueInt32,
		reflect.Int64:   setValueInt64,
		reflect.Uint:    setValueUint,
		reflect.Uint8:   setValueUint8,
		reflect.Uint16:  setValueUint16,
		reflect.Uint32:  setValueUint32,
		reflect.Uint64:  setValueUint64,
		reflect.Float32: setValueFloat32,
		reflect.Float64: setValueFloat64,
	}

	var conv_func = conv_map[f.Kind()]

	if conv_func == nil {
		return errors.New(`config: Field of type '` + f.Type().Name() + `' not suported.`)
	}

	return conv_func(f, str_val)
}

/* Convert some stringed value to correct type if possible
 * and set it into the field.
 */
func (f Value) SetValueFromString(str_val string) error {
	return SetValueFromString(reflect.Value(f), str_val)
}

/// Private

func setValueString(f reflect.Value, str_val string) error {
	f.SetString(str_val)
	return error(nil)
}

func setValueBool(f reflect.Value, str_val string) error {
	vl, err := strconv.ParseBool(str_val)
	if err != nil {
		return err
	}

	f.SetBool(vl)
	return nil
}

func setValueUintGeneric(f reflect.Value, str_val string, baseInt int) error {
	vl, err := strconv.ParseUint(str_val, 10, baseInt)
	if err != nil {
		return err
	}

	f.SetUint(vl)
	return nil
}

func setValueUint(f reflect.Value, str_val string) error {
	return setValueUintGeneric(f, str_val, 32)
}

func setValueUint8(f reflect.Value, str_val string) error {
	return setValueUintGeneric(f, str_val, 8)
}

func setValueUint16(f reflect.Value, str_val string) error {
	return setValueUintGeneric(f, str_val, 16)
}

func setValueUint32(f reflect.Value, str_val string) error {
	return setValueUintGeneric(f, str_val, 32)
}

func setValueUint64(f reflect.Value, str_val string) error {
	return setValueUintGeneric(f, str_val, 64)
}

func setValueIntGeneric(f reflect.Value, str_val string, baseInt int) error {
	vl, err := strconv.ParseInt(str_val, 10, baseInt)
	if err != nil {
		return err
	}

	f.SetInt(vl)
	return nil
}

func setValueInt(f reflect.Value, str_val string) error {
	return setValueIntGeneric(f, str_val, 32)
}

func setValueInt8(f reflect.Value, str_val string) error {
	return setValueIntGeneric(f, str_val, 8)
}

func setValueInt16(f reflect.Value, str_val string) error {
	return setValueIntGeneric(f, str_val, 16)
}

func setValueInt32(f reflect.Value, str_val string) error {
	return setValueIntGeneric(f, str_val, 32)
}

func setValueInt64(f reflect.Value, str_val string) error {
	return setValueIntGeneric(f, str_val, 64)
}

func setValueFloatGeneric(f reflect.Value, str_val string, baseInt int) error {
	vl, err := strconv.ParseFloat(str_val, baseInt)
	if err != nil {
		return err
	}

	f.SetFloat(vl)
	return nil
}

func setValueFloat32(f reflect.Value, str_val string) error {
	return setValueFloatGeneric(f, str_val, 32)
}

func setValueFloat64(f reflect.Value, str_val string) error {
	return setValueFloatGeneric(f, str_val, 64)
}
