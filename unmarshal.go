package consulutils

import (
	"github.com/hashicorp/consul/api"
	"reflect"
	"strconv"
)

func extractKeyValuePairs(pairs api.KVPairs) (kvpairs map[string]string) {
	kvpairs = map[string]string{}
	for _, pair := range pairs {
		val := string(pair.Value)
		if val != "" {
			kvpairs[pair.Key] = val
		}
	}
	return
}

func Unmarshal(consulKvPairs api.KVPairs, dst interface{}) error {
	// KV pair map
	kvpairs := extractKeyValuePairs(consulKvPairs)

	dstVal := reflect.ValueOf(dst)
	dstType := reflect.TypeOf(dst)
	// tmp obj to get tags.
	nVal := reflect.New(dstType.Elem())
	for i := 0; i < nVal.Elem().NumField(); i++ {

		tagKey := dstType.Elem().Field(i).Tag.Get("consul")
		// Check if key exists
		if v, ok := kvpairs[tagKey]; ok {

			switch dstVal.Elem().Field(i).Kind() {
			case reflect.String:
				dstVal.Elem().Field(i).SetString(v)
				break
			case reflect.Int64:
				i64val, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				dstVal.Elem().Field(i).SetInt(i64val)
				break
			case reflect.Int:
				i32val, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					return err
				}
				dstVal.Elem().Field(i).SetInt(i32val)
				break
			case reflect.Bool:
				boolVal, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				dstVal.Elem().Field(i).SetBool(boolVal)
				break
			}
		}
	}
	return nil
}
