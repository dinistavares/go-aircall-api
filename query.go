package aircall

import (
	"fmt"
	"reflect"
)

type QueryValues map[string]string

func (v *QueryValues) getQueryValues() QueryValues {
	return *v
}

func (v QueryValues) encode() string {
	var query string

	count := 0

	for key, value := range v {
		if count > 0 {
			query += "&"
		} else {
			query = "?"
		}

		query += key + "=" + value

		count++
	}

	return query
}

func (v QueryValues) from(value string) {
	v["from"] = value
}

func (v QueryValues) to(value string) {
	v["to"] = value
}

func (v QueryValues) order(value string) {
	v["order"] = value
}

func (v QueryValues) orderBy(value string) {
}

func (v QueryValues) page(page int, perPage int) {
	v["page"] = fmt.Sprintf("%d", page)
	v["per_page"] = fmt.Sprintf("%d", perPage)
}

func (v QueryValues) set(key string, value string) {
	v[key] = value
}

func isPointerWithQueryValues(i interface{}) (interface{}, bool) {
	if i == nil {
		return nil, false
	}

	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		field := val.FieldByName("QueryValues")

		if field.IsValid() {
			return field.Interface(), true
		}
	}

	return nil, false
}
