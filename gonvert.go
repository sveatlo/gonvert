package gonvert

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ToString(value interface{}) (string, error) {
	switch v := value.(type) {
	case nil:
		return "", nil
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case float32, float64:
		intVal, err := ToInt(value)
		if err == nil {
			return fmt.Sprintf("%d", intVal), nil
		}

		return fmt.Sprintf("%g", value), nil
	default:
		return fmt.Sprintf("%v", value), nil
	}
}

func ToInt(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil

	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case uint:
		return int64(v), nil

	case float32, float64:
		var float64Value float64
		switch value.(type) {
		case float32:
			float64Value = float64(value.(float32))
		case float64:
			float64Value = value.(float64)
		}

		var intValue = int64(float64Value)
		if float64Value == float64(intValue) {
			return intValue, nil
		}

		return 0, fmt.Errorf("Cannot truncate float to int. Data loss")

	case string:
		n, err := strconv.ParseInt(v, 10, 64)
		return n, err
	}

	return 0, fmt.Errorf("unsupported value type in ToInt conversion. type = `%T`", value)
}

func ToFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int:
		return float64(v), nil

	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uint:
		return float64(v), nil

	case float32:
		return float64(v), nil
	case float64:
		return v, nil

	case string:
		return strconv.ParseFloat(value.(string), 64)
	}

	return 0, fmt.Errorf("unsupported value type in ToFloat conversion. type = `%T`", value)
}

func ToBool(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case int, int8, int16, int32, int64, float32, float64:
		return v != 0, nil
	case string:
		v = strings.ToLower(v)

		if v == "true" || v == "yes" {
			return true, nil
		} else if v == "false" || v == "no" {
			return false, nil
		}
	}

	return false, fmt.Errorf("unsupported value type in ToBool conversion. type = `%T`", value)
}

func ToSlice(value interface{}) ([]interface{}, error) {
	switch value.(type) {
	case nil:
		return []interface{}{}, nil
	case bool, float32, float64, int, int8, int16, int32, int64, string, uint, uint8, uint16, uint32, uint64:
		return []interface{}{value}, nil
	default:
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.Slice:
			res := make([]interface{}, v.Len())
			for i := 0; i < v.Len(); i++ {
				res[i] = v.Index(i).Interface()
			}

			return res, nil
		}
	}

	return nil, fmt.Errorf("unsupported value type in ToSlice conversion. type = `%T`", value)
}

func ToMapString(value interface{}) (map[string]interface{}, error) {
	switch val := value.(type) {
	case map[string]interface{}:
		return val, nil
	case map[interface{}]interface{}:
		var (
			res = map[string]interface{}{}
			err error
		)

		for k, v := range val {
			switch v.(type) {
			case map[interface{}]interface{}:
				v, err = ToMapString(v)
				if err != nil {
					return nil, err
				}
			case []interface{}:
				var arr []interface{}
				for _, w := range v.([]interface{}) {
					switch w.(type) {
					case map[interface{}]interface{}:
						i, err := ToMapString(w)
						if err != nil {
							return nil, err
						}
						arr = append(arr, i)
					default:
						arr = append(arr, w)
					}
				}

				v = arr
			}

			if ks, ok := k.(string); ok {
				res[ks] = v
			} else {
				res[fmt.Sprintf("%v", k)] = v
			}
		}

		return res, nil
		// TODO: handle other map types via reflect
	}

	return nil, fmt.Errorf("unsupported value type in ToMapString conversion. type = `%T`", value)
}
