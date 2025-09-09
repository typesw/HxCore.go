package HxCore

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

//import "reflect"

//#region HxResultValue

type HxResultValue struct {
	Result     string       `json:"result"`
	Values     any          `json:"values"`
	Message    string       `json:"message"`
	Remark     string       `json:"remark"`
	ResultType HxResultType `json:"ResultType"`
	ValueType  string       `json:"ValueType"`
	ValueCount int          `json:"ValueCount"`
}

func CreateHxResultValue(resultType HxResultType, value any, message string, remark string) HxResultValue {
	result := HxResultValue{
		Result:     resultType.String(),
		ResultType: resultType,
		Values:     value,
		Message:    message,
		Remark:     remark,
	}
	//result.Result = resultType.String()
	result.ValueCount = result.CountEx()
	result.ValueType = result.TypeEx()
	return result
}

func CreateHxResponseResult(resultType HxResultType, value any, message string, remark string) HxResultValue {
	return CreateHxResultValue(resultType, value, message, remark)
}

// 옵션 함수의 타입을 정의
type ResultOption func(*HxResultValue)

// 각 옵션을 설정하는 ResultOption 타입을 반환
func WithValue(value any) ResultOption {
	return func(r *HxResultValue) {
		r.Values = value
	}
}

func WithMessage(message string) ResultOption {
	return func(r *HxResultValue) {
		r.Message = message
	}
}

func WithOptionString(remark string) ResultOption {
	return func(r *HxResultValue) {
		r.Remark = remark
	}
}

func (r *HxResultValue) TypeEx() string {
	// Value가 nil이면 0을 반환
	if r.Values == nil {
		return ""
	}

	// reflect 패키지를 사용해 Value의 실제 타입을 알아냄
	val := reflect.ValueOf(r.Values)
	Result := val.Kind().String()
	switch Result {
	case "slice":
		Result = "array"
	}
	return Result

}

func (r *HxResultValue) CountEx() int {
	// Value가 nil이면 0을 반환
	if r.Values == nil {
		return 0
	}

	// reflect 패키지를 사용해 Value의 실제 타입을 알아냄
	val := reflect.ValueOf(r.Values)

	// 타입에 따라 길이를 반환
	switch val.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
		return val.Len() // 슬라이스, 배열, 맵, 문자열 등은 Len()으로 길이를 알 수 있음
	default:
		// 단일 값(int, struct 등)은 길이가 없으므로,
		// 값이 존재하면 1, nil이면 0을 반환하도록 처리할 수 있음.
		if val.IsValid() && !val.IsNil() {
			return 1
		}
		return 0
	}
}

/*
// ResponseResult를 생성하는 함수
func CreateResponseResult(resultType HxResultType, resultValue any, resultMessage string, resultRemark string) HxResultValue {
	var count int = -1
	var valType reflect.TypeEx
	if resultValue != nil {
		val := reflect.ValueOf(resultValue)
		valType := val.Kind()
		switch valType {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
			count = val.Len() // 슬라이스, 배열, 맵, 문자열 등은 Len()으로 길이를 알 수 있음
		default:
			// 단일 값(int, struct 등)은 길이가 없으므로,
			// 값이 존재하면 1, nil이면 0을 반환하도록 처리할 수 있음.
			if val.IsValid() && !val.IsNil() {
				count = 1
			}
			count = 0
		}
	}

	Result := HxResultValue{
		ResultType: resultType,
		Values:      resultValue,
		ValueCount: -1,
		Message:    resultMessage,
		Remark:     resultRemark,
	}
	Result.ValueType = valType.String()
	Result.ValueCount = count
	return Result
}*/

func (res *HxResultValue) ToJsonResponseWriter(w http.ResponseWriter) bool {

	if res.ValueCount > -1 {
		res.ValueCount = res.CountEx()
	}
	if res.ValueType == "" {
		res.ValueType = res.TypeEx()
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("JSON 인코딩 에러: %v", err)
		//fmt.Println(err.Error())
		//fmt.Println(res)
		res.ResultType = Fail
		res.Message += " / ERROR : " + err.Error()
		return false
	} else {
		//w.Header().Set("Content-TypeEx", "application/json; charset=utf-8")
		//w.WriteHeader(http.StatusOK)
		//w.WriteHeader(http.StatusOK)
		return true
	}
}
func (res *HxResultValue) ToJsonString() (string, error) {
	jsonDataBytes, err := res.ToJsonBytes()

	if err != nil {
		log.Fatalf("JSON 마샬링 에러: %v", err)
		return "", err
	}
	jsonString := string(jsonDataBytes)
	return jsonString, err
}

func (res *HxResultValue) ToJsonBytes() ([]byte, error) {
	//jsonDataBytes, err := json.Marshal(res)
	jsonDataBytes, err := json.MarshalIndent(res, "", " ")
	//jsonDataBytes, err := res.ToJsonCaseBytes()
	if err != nil {
		log.Fatalf("JSON 마샬링 에러: %v", err)
		return nil, err
	}
	return jsonDataBytes, err
}

func (res *HxResultValue) ToJsonBytesFromCaseingString(s string) ([]byte, error) {
	casing, _ := GetHxCasingFromString(s)
	return res.ToJsonBytesFromCaseingType(casing)
}

func (res *HxResultValue) ToJsonBytesFromCaseingType(casing HxCasing) ([]byte, error) {
	jsonDataBytes, err := GetJsonWithCasing(res, casing)
	//jsonDataBytes, err := GetDataByKeyNameWithCasing(res, casing)
	if err != nil {
		log.Fatalf("JSON 마샬링 에러: %v", err)
		return nil, err
	}
	return jsonDataBytes, err
}

func (res *HxResultValue) ToJsonPascalCasingBytes() ([]byte, error) {
	return GetMarshalIndentWithCasing(res, PascalCase)
}
func (res *HxResultValue) ToJsonCamelCasingBytes() ([]byte, error) {
	return GetMarshalIndentWithCasing(res, CamelCase)
}
func (res *HxResultValue) ToJsonSnakeCasingBytes() ([]byte, error) {
	return GetMarshalIndentWithCasing(res, SnakeCase)
}
func (res *HxResultValue) ToJsonCaseBytes() ([]byte, error) {
	return GetMarshalIndentWithCasing(res, JsonCase)
}
func (res *HxResultValue) ToJsonNormalCaseBytes() ([]byte, error) {
	return GetMarshalIndentWithCasing(res, NormalCase)
}

/*
func CreateResponseResult2(resultType HxResultType, options ...ResultOption) HxResultValue {
	// 기본값을 가진 ResponseResult를 먼저 생성합니다.
	var res HxResultValue = CreateResponseResult(resultType, nil, "", "")

	// 받아온 옵션들을 순서대로 적용합니다.
	for _, opt := range options {
		opt(&res)
	}
	return res
}
*/
//#endregion HxResultValue
