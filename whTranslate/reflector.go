package whtranslate

import (
    "errors"
    "reflect"
    "strings"
)

// structToString uses the magic of reflection to turn structs containing webhook-provided event data into strings!
//
// Struct fields with the `readable` tag set will use the value in that tag as the readable field name. Else, this is
// derived from the JSON tag.
// Struct fields with the `prefix` tag will have the value within it prepended to any value within the field
// post-conversion.
func structToString(s interface{}) (string, error) {

    ct := reflect.TypeOf(s)
    if ct == nil {
        return "", errors.New("structToString: no type data available")
    }

    if ct.Kind() != reflect.Struct {
        return "", errors.New("structToString: invalid top-level type, must be struct")
    }

    var (
        sb strings.Builder
        ctv = reflect.ValueOf(s)
    )

    for i := 0; i < ct.NumField(); i += 1 {

        // deal with nested privacy structs
        if field := ctv.Field(i); field.Type().Kind() == reflect.Struct {
            x, err := structToString(field.Interface())
            if err != nil {
                return "", err
            }
            sb.WriteString(x)
            continue
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
            if x != "" {
                formatted = formatString(x)
            }
        case *bool:
            if x != nil {
                formatted = formatBool(*x)
            }
        }

        if formatted != "" {
            sb.WriteString(
                formatUpdateMessage(readableName, prefix + formatted),
            )
        }

    }

    return sb.String(), nil
}