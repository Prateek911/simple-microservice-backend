package mapper

import (
	"fmt"
	"reflect"
)

// doesn't work!!
func Map[T any, R any](source T, target *R) error {
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
			}
		}
	}

	return nil
}
