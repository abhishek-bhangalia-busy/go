package main

import (
	"fmt"
	"reflect"
)

func merger(arr1 interface{}, arr2 interface{}) (interface{}, error) {

	if arr1 == nil && arr2 == nil { //if both array are received as empty, there is nothing to merge, so error is sent
		return nil, fmt.Errorf("no value to merge")
	}

	if arr1 == nil || arr2 == nil { //if any one array is received as empty, so it is merged with nothing and sent back as one
		if arr1 == nil {
			return arr2, nil
		} else {
			return arr1, nil
		}
	} else { //both arrays have received something, so merge should happen accordingly

		var result interface{}
		var final []interface{}

		first := reflect.ValueOf(arr1)

		switch first.Kind() {
		case reflect.Slice:
			switch elemKind := first.Type().Elem().Kind(); elemKind {
			case reflect.Interface:
				// Handle slice of interface{}
				a := first.Interface().([]interface{})
				for _, val := range a {
					final = append(final, val)
				}
			default:
				// Handle other types
				for i := 0; i < first.Len(); i++ {
					final = append(final, first.Index(i).Interface())
				}
			}
			break

		case reflect.Int:
			a := first.Interface().(int) //if first array is just 1 int value
			final = append(final, a)
			break
		case reflect.String:
			a := first.Interface().(string) // if first array is just 1 string value
			final = append(final, a)
			break

		case reflect.Float64:
			a := first.Interface().(float64) // if first array is just 1 string value
			final = append(final, a)
			break

		case reflect.Bool:
			a := first.Interface().(bool) // if first array is just 1 string value
			final = append(final, a)
			break
		}

		second := reflect.ValueOf(arr2)

		switch second.Kind() {
		case reflect.Slice:
			switch elemKind := second.Type().Elem().Kind(); elemKind {
			case reflect.Interface:
				// Handle slice of interface{}
				a := second.Interface().([]interface{})
				for _, val := range a {
					final = append(final, val)
				}
			default:
				// Handle other types
				for i := 0; i < second.Len(); i++ {
					final = append(final, second.Index(i).Interface())
				}
			}
			break

		case reflect.Int:
			a := second.Interface().(int)
			final = append(final, a)
			break
		case reflect.String:
			a := second.Interface().(string)
			final = append(final, a)
			break
		case reflect.Float64:
			a := second.Interface().(float64) // if first array is just 1 string value
			final = append(final, a)
			break
		case reflect.Bool:
			a := second.Interface().(bool) // if first array is just 1 string value
			final = append(final, a)
			break
		}

		result = final
		return result, nil
	}
}

func main() {
	var arr1 interface{}
	// arr1 = []interface{}{1, "bcd", 3, 4.23, 5}
	// arr1 = true
	arr1=325.234
	var arr2 interface{} = []interface{}{"st", 6,true, 7, 3.14}

	mergedarr, err := merger(arr1, arr2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mergedarr)
	}

	fmt.Println(reflect.TypeOf(mergedarr))
}
