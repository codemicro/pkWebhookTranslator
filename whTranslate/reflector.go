/*
 *  pkWebhookTranslate, https://github.com/codemicro/pkWebhookTranslate
 *  Copyright (c) 2021 codemicro and contributors
 *
 *  SPDX-License-Identifier: MIT
 */

package whtranslate

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

// structToString uses the magic of reflection to turn structs containing webhook-provided event data into strings!
//
// If useStatementFormat is true, formatStatementMessage is used to format output messages in place of
// formatUpdateMessage.
//
// Struct fields with the `readable` tag set will use the value in that tag as the readable field name. Else, this is
// derived from the JSON tag.
// Struct fields with the `prefix` tag will have the value within it prepended to any value within the field
// post-conversion.
func structToString(s interface{}, useStatementFormat bool) (string, error) {

	ct := reflect.TypeOf(s)
	if ct == nil {
		return "", errors.New("structToString: no type data available")
	}

	if ct.Kind() != reflect.Struct {
		return "", errors.New("structToString: invalid top-level type, must be struct")
	}

	var (
		sb             strings.Builder
		formatFunction = formatUpdateMessage
		ctv            = reflect.ValueOf(s)
	)

	if useStatementFormat {
		formatFunction = formatStatementMessage
	}

	for i := 0; i < ct.NumField(); i += 1 {

		// deal with nested privacy structs
		{
			field := ctv.Field(i)
			if p, ok := field.Interface().(privacy); ok {
				x, err := structToString(p, useStatementFormat)
				if err != nil {
					return "", err
				}
				sb.WriteString(x)
				continue
			}
		}

		field := ct.Field(i)
		readableName := field.Tag.Get("readable")

		if readableName == "" {
			// convert the JSON tag to a readable string
			readableName = snakeToReadable(strings.Split(field.Tag.Get("json"), ",")[0])
		}

		prefix := field.Tag.Get("prefix")
		value := ctv.Field(i).Interface()

		var formatted string
		switch x := value.(type) {
		case string:
			panic(errors.New("string types should use `nullableString` in place of `string`"))
		case *string:
			panic(errors.New("string types should use `nullableString` in place of `*string`"))
		case nullableString:
			if x.HasValue {
				formatted = formatString(x.Value)
			}
		case *bool:
			if x != nil {
				formatted = formatBool(*x)
			}
		case *time.Time:
			if x != nil {
				formatted = formatTime(*x)
			}
		case []string:
			if len(x) != 0 {
				formatted = formatStringSlice(x)
			}
		}

		if formatted != "" {
			sb.WriteString(
				formatFunction(readableName, prefix+formatted),
			)
		}

	}

	return sb.String(), nil
}
