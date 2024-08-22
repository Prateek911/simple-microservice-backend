package mapper

import (
	"fmt"
	"reflect"
)

// Map copies fields from source to target where field names and types match.
// It handles nested structs and custom types if they are valid.
func Map(source interface{}, target interface{}) error {
	sourceVal := reflect.ValueOf(source)
	targetVal := reflect.ValueOf(target).Elem()

	if sourceVal.Kind() != reflect.Struct || targetVal.Kind() != reflect.Struct {
		return fmt.Errorf("source and target must be structs")
	}

	for i := 0; i < sourceVal.NumField(); i++ {
		sourceField := sourceVal.Type().Field(i)
		sourceValue := sourceVal.Field(i)
		targetField := targetVal.FieldByName(sourceField.Name)

		if targetField.IsValid() && targetField.CanSet() {
			if targetField.Type() == sourceField.Type {
				targetField.Set(sourceValue)
			} else if targetField.Type().Kind() == reflect.Struct && sourceField.Type.Kind() == reflect.Struct {
				// Recursive mapping for nested structs
				err := Map(sourceValue.Interface(), targetField.Addr().Interface())
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
