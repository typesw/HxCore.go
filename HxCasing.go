package HxCore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// Casing은 JSON 키 표기법을 나타내는 타입입니다.
type HxCasing int

const (
	NormalCase HxCasing = iota
	PascalCase          // 파스칼 케이스 (e.g., UserName)
	CamelCase           // 카멜 케이스 (e.g., userName)
	SnakeCase           // 스네이크 케이스 (e.g., user_name)
	JsonCase            // json 태그가 있으면 그 값을, 없으면 Normal처럼 처리
)

var stringToHxCasingMap = map[string]HxCasing{
	"default": JsonCase,

	"normalcase": NormalCase,
	"pascalcase": PascalCase,
	"camelcase":  CamelCase,
	"snakecase":  SnakeCase,
	"jsoncase":   JsonCase,

	"normal": NormalCase,
	"pascal": PascalCase,
	"camel":  CamelCase,
	"snake":  SnakeCase,
	"json":   JsonCase,
}

// ParseHxCasing은 주어진 문자열을 HxCasing 타입으로 변환합니다.
// 일치하는 타입이 없으면 에러를 반환합니다.
func GetHxCasingFromString(s string) (HxCasing, error) {
	// 1. 입력받은 문자열의 공백을 제거하고 소문자로 변환하여 비교 준비를 합니다.
	lowerStr := strings.ToLower(strings.TrimSpace(s))

	// 2. 맵에서 해당하는 HxCasing 타입을 찾습니다.
	casing, ok := stringToHxCasingMap[lowerStr]

	// 3. 맵에 값이 있으면(ok == true) 찾은 타입을 반환합니다.
	if ok {
		return casing, nil
	}

	// 4. 맵에 값이 없으면, 기본값(NormalCase)과 함께 에러를 반환합니다.
	return NormalCase, fmt.Errorf("알 수 없는 HxCasing 타입입니다: '%s'", s)
}

func GetNameingCase(s string, casing HxCasing) string {
	Result := s

	switch casing {
	case PascalCase:
		Result = GetPascalCase(s)
	case CamelCase:
		Result = GetCamelCase(s)
	case SnakeCase:
		Result = GetSnakeCase(s)
	case NormalCase, JsonCase:
		fallthrough
	default:
		Result = s // Normal이 기본값
	}
	return Result
}

// GetCamelCase는 문자열을 카멜 케이스로 변환합니다.
func GetCamelCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// GetPascalCase는 문자열을 파스칼 케이스로 변환합니다.
func GetPascalCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// GetSnakeCase는 문자열을 스네이크 케이스로 변환합니다. (e.g., UserID -> user_id)
func GetSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// MarshalWithCasing은 주어진 구조체를 선택한 표기법의 JSON으로 변환합니다.
func GetMarshalWithCasing(data any, casing HxCasing) ([]byte, error) {
	newMap, err := GetDataByKeyNameWithCasing(data, casing)
	if err != nil {
		return nil, err
	}
	return json.Marshal(newMap)
}
func GetMarshalIndentWithCasing(data any, casing HxCasing) ([]byte, error) {
	newMap, err := GetDataByKeyNameWithCasing(data, casing)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(newMap, "", " ")
}

func GetDataByKeyNameWithCasing(data any, casing HxCasing) (map[string]any, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("구조체 타입만 지원합니다")
	}

	newMap := make(map[string]any)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if !field.IsExported() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		keyName := field.Name
		if jsonTag != "" && jsonTag != "-" {
			keyName = strings.Split(jsonTag, ",")[0]
		}

		var finalKeyName string
		switch casing {
		case JsonCase:
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" && jsonTag != "-" {
				finalKeyName = strings.Split(jsonTag, ",")[0]
			} else {
				finalKeyName = keyName // 태그 없으면 Normal처럼
			}
		case PascalCase:
			finalKeyName = GetPascalCase(keyName)
		case CamelCase:
			finalKeyName = GetCamelCase(keyName)
		case SnakeCase:
			finalKeyName = GetSnakeCase(keyName)
		case NormalCase:
			fallthrough
		default:
			finalKeyName = keyName // Normal이 기본값
		}

		newMap[finalKeyName] = value.Interface()
	}

	return newMap, nil
}

// 주어진 구조체를 선택한 표기법과 '원본 순서'의 JSON으로 변환합니다.
func GetJsonWithCasing(data any, casing HxCasing) ([]byte, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("구조체 타입만 지원합니다")
	}

	var builder strings.Builder
	builder.WriteString("{") // JSON 객체 시작

	typ := val.Type()
	firstField := true

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if !field.IsExported() {
			continue
		}

		// 쉼표(,) 추가 로직
		if !firstField {
			builder.WriteString(",")
		}
		firstField = false

		var finalKeyName string
		fieldName := field.Name

		switch casing {
		case JsonCase:
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" && jsonTag != "-" {
				finalKeyName = strings.Split(jsonTag, ",")[0]
			} else {
				finalKeyName = fieldName
			}
		default:
			finalKeyName = GetNameingCase(fieldName, casing)
		}
		// 키(key) 추가 (큰따옴표 포함)
		builder.WriteString(fmt.Sprintf(`"%s":`, finalKeyName))

		// 값(value)을 표준 json 마샬러를 이용해 변환하여 추가
		jsonValue, err := json.Marshal(value.Interface())
		if err != nil {
			return nil, fmt.Errorf("필드 '%s' 마샬링 실패: %w", fieldName, err)
		}
		builder.Write(jsonValue)
	}

	builder.WriteString("}") // JSON 객체 종료

	// 보기 좋게 Indent 처리 (선택 사항)
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, []byte(builder.String()), "", "  ")
	if err != nil {
		return nil, err
	}

	return prettyJson.Bytes(), nil
}
