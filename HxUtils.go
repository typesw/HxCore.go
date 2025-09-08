package HxCore

import (
	"fmt"

	"github.com/godror/godror"
	//"github.com/shopspring/decimal"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//#region 문자처리 함수

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func Lower(str string) string {
	return strings.ToLower(str)
}
func Upper(str string) string {
	return strings.ToUpper(str)
}
func SubStr(str string, start int, end int) string {
	return str[start:end]
}

//#endregion 문자처리 함수

//#region 정규식

func IsRegexpMatch(pattern string, str string) bool {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return r.MatchString(str)
}

func GetRegexpMatch(pattern string, str string) map[int]string {
	regex, _ := regexp.Compile(pattern)
	match := regex.FindStringSubmatch(str)

	Result := make(map[int]string)

	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			Result[i] = name
		}
	}

	return Result
}

//#endregion 정규식

// #region 값 & 타입 체크
func isBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
func IsNullOrEmpty(value string) bool {
	if IsString(value) == true && isBlank(value) == true {
		return true
	}

	switch value {
	case
		strconv.FormatInt(math.MinInt32, 10),                                 // Int32 최솟값
		strconv.FormatInt(math.MinInt64, 10),                                 // Int64 최솟값
		strconv.FormatFloat(math.MaxFloat64*-1, 'f', -1, 64),                 // float64(double) 최솟값
		strconv.FormatFloat(math.MaxFloat32*-1, 'f', -1, 32),                 // float32(float) 최솟값
		time.Time{}.Format("2006-01-02 15:04:05.999999999 -0700 MST"),        // DateTime 최솟값(Go의 제로 값)
		"1900-01-01 00:00:00 +0000 UTC", "1900-01-01 00:00:00", "1900-01-01": // 특정 날짜 (1900년 1월 1일)
		return true
	default:
		//return false
	}

	return false
}

func IsNullOrWhiteSpace(value string) bool {
	return IsNullOrEmpty(value)
}
func IsInt(value string) bool {
	_, err := strconv.ParseInt(value, 10, 64)
	return err == nil
}

func IsFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func IsString(value string) bool {
	_, err := strconv.Unquote(value)
	return err == nil
}

func IsNumber(value string) bool {
	return IsFloat(value)
}

func IsTime(value string) bool {
	_, err := time.Parse(time.RFC3339, value)
	return err == nil
}
func IsDate(value string) bool {
	_, err := time.Parse(time.RFC3339, value)
	return err == nil
}
func IsDatetime(value string) bool {
	_, err := time.Parse(time.RFC3339, value)
	return err == nil
}
func IsDatetimeUnix(value string) bool {
	_, err := time.Parse(time.RFC3339, value)
	return err == nil
}
func IsDatetimeUnixInt(value string) bool {
	return IsInt(value)
}

//#endregion 값 & 타입 체크

// #region 타입 변환 헬퍼
func ConvertValueToInt(v interface{}) int {
	switch n := v.(type) {
	case int:
		return n
	case int8:
		return int(n)
	case int16:
		return int(n)
	case int32:
		return int(n)
	case int64:
		return int(n)
	case uint:
		return int(n)
	case uint8:
		return int(n)
	case uint16:
		return int(n)
	case uint32:
		return int(n)
	case uint64:
		return int(n)
	case string: // 문자열도 숫자로 변환 시도
		i, _ := strconv.ParseInt(n, 10, 64)
		return int(i)
	case godror.Number:
		s := ConvertValueToString(n)
		i, _ := strconv.ParseInt(s, 10, 64)
		return int(i)
	default:
		return 0
	}
}

func ConvertValueToInt64(v interface{}) int64 {

	if v == nil {
		return 0
	}

	switch n := v.(type) {
	case int:
		return int64(n)
	case int8:
		return int64(n)
	case int16:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	case uint:
		return int64(n)
	case uint8:
		return int64(n)
	case uint16:
		return int64(n)
	case uint32:
		return int64(n)
	case uint64:
		return int64(n)
	case string: // 문자열도 숫자로 변환 시도
		i, _ := strconv.ParseInt(n, 10, 64)
		return i
	case godror.Number:
		s := ConvertValueToString(n)
		i, _ := strconv.ParseInt(s, 10, 64)
		return i
	default:
		return 0
	}
}
func ConvertValueToUInt64(v interface{}) uint64 {
	switch n := v.(type) {
	case int:
		return uint64(n)
	case int8:
		return uint64(n)
	case int16:
		return uint64(n)
	case int32:
		return uint64(n)
	case int64:
		return uint64(n)
	case uint:
		return uint64(n)
	case uint8:
		return uint64(n)
	case uint16:
		return uint64(n)
	case uint32:
		return uint64(n)
	case uint64:
		return uint64(n)
	case string: // 문자열도 숫자로 변환 시도
		i, _ := strconv.ParseUint(n, 10, 64)
		return i
	default:
		return 0
	}
}

func ConvertValueToFloat64(v interface{}) float64 {
	switch n := v.(type) {
	case int:
		return float64(n)
	case int8:
		return float64(n)
	case int16:
		return float64(n)
	case int32:
		return float64(n)
	case int64:
		return float64(n)
	case uint:
		return float64(n)
	case uint8:
		return float64(n)
	case uint16:
		return float64(n)
	case uint32:
		return float64(n)
	case uint64:
		return float64(n)
	case float32:
		return float64(n)
	case float64:
		return n
	case string:
		i, _ := strconv.ParseFloat(n, 64)
		return i
	default:
		return 0
	}
}

func ConvertValueToString(v interface{}) string {
	switch n := v.(type) {
	case int:
		return strconv.FormatInt(int64(n), 10)
	case int8:
		return strconv.FormatInt(int64(n), 10)
	case int16:
		return strconv.FormatInt(int64(n), 10)
	case int32:
		return strconv.FormatInt(int64(n), 10)
	case int64:
		return strconv.FormatInt(n, 10)
	case uint:
		return strconv.FormatUint(uint64(n), 10)
	case uint8:
		return strconv.FormatUint(uint64(n), 10)
	case uint16:
		return strconv.FormatUint(uint64(n), 10)
	case uint32:
		return strconv.FormatUint(uint64(n), 10)
	case uint64:
		return strconv.FormatUint(uint64(n), 10)
	case float32:
		return strconv.FormatFloat(float64(n), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(n, 'f', -1, 64)
	case string:
		return n
	case bool:
		return strconv.FormatBool(n)
	case time.Time:
		return n.Format(time.RFC3339)
	case godror.Number:
		return v.(godror.Number).String()
	case godror.NullTime:
		return v.(godror.NullTime).Time.Format(time.RFC3339)
	case godror.Object:
		return v.(string)
	default:
		return ""
	}
}

func ConvertStringToInt(value string) int {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return math.MinInt32
	}
	return int(i)
}
func ConvertStringToInt64(value string) int64 {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return math.MinInt64
	}
	return i
}
func ConvertStringToFloat64(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return math.MaxFloat64 * -1
	}
	return f
}
func ConvertStringToFloat32(value string) float32 {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return math.MaxFloat32 * -1
	}
	return float32(f)
}
func ConvertStringToBool(value string) bool {
	b, err := strconv.ParseBool(Upper(value))
	if err != nil {
		return false
	}
	return b
}
func ConvertStringToTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}
	}
	return t
}
func ConvertStringToTimeUnix(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}
	}
	return t
}
func ConvertStringToTimeUnixInt(value string) int64 {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
func ConvertStringToTimeUnixFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return f
}

func ConvertIntToString(value int) string {
	return strconv.Itoa(value)
}
func ConvertInt64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}
func ConvertIntToBool(value int) bool {
	return value > 0
}
func ConvertIntToFloat(value int) float64 {
	return float64(value)
}
func ConvertIntToTimeUnix(value int) time.Time {
	return time.Unix(int64(value), 0)
}
func ConvertIntToInt64(value int) int64 {
	return int64(value)
}
func ConvertIntToInt32(value int) int32 {
	return int32(value)
}
func ConvertStringToNumber(value string) float64 {
	return ConvertStringToFloat64(value)
}
func ConvertFloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
func ConvertFloatToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return i
}

//#endregion 타입 변환 헬퍼

//#region 숫자를 컴마가 포함된 문자열로 변환

func ConvertInt64ToCommaString(value int64) string {
	// 음수 처리
	isNegative := false
	if value < 0 {
		isNegative = true
		value = -value // 양수로 변환하여 처리
	}
	s := ConvertInt64ToString(value)
	length := len(s)
	// 세 자리 이하는 쉼표가 필요 없음
	if length <= 3 {
		if isNegative {
			return "-" + s
		}
		return s
	}
	// strings.Builder를 사용해 효율적으로 문자열을 조립
	var sbr strings.Builder

	// 첫 부분 (1~3자리)을 먼저 처리
	firstPartLen := length % 3
	if firstPartLen == 0 {
		firstPartLen = 3
	}
	sbr.WriteString(s[:firstPartLen])

	// 나머지 부분을 3자리씩 잘라 쉼표와 함께 추가
	for i := firstPartLen; i < length; i += 3 {
		sbr.WriteByte(',')
		sbr.WriteString(s[i : i+3])
	}

	// 음수였다면 앞에 '-' 부호를 붙여줌
	if isNegative {
		return "-" + sbr.String()
	}

	return sbr.String()
}
func ConvertIntToCommaString(value int) string {
	i := ConvertIntToInt64(value)
	return ConvertInt64ToCommaString(i)
}
func ConvertFloatToCommaString(value float64) string {
	// 1. 실수를 문자열로 변환 ('f'는 일반적인 실수 표현, -1은 필요한 만큼의 소수점 자리)
	s := strconv.FormatFloat(value, 'f', -1, 64)

	// 2. 소수점을 기준으로 정수부와 소수부로 분리
	parts := strings.SplitN(s, ".", 2)
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = parts[1]
	}

	i := ConvertStringToInt64(integerPart)
	// 3.정수 쉼표
	Result := ConvertInt64ToCommaString(i)

	//4. 실수일 경우 소수점 아래자리 추가
	if decimalPart != "" {
		Result += "." + decimalPart
	}

	return Result
}

// AddCommas 함수는 어떤 숫자 타입이든 받아 쉼표를 추가한 문자열로 반환합니다.
func ConvertToCommaString(value interface{}) string {
	switch v := value.(type) {
	// 정수형 타입들 처리
	case int, int8, int16, int32, int64:
		// 모든 정수형을 int64로 변환하여 처리
		n := ConvertValueToInt64(v)
		return ConvertInt64ToCommaString(n) // 아래에 정의된 정수 처리 헬퍼 함수 호출

	// 실수형 타입들 처리
	case float32, float64:
		// 모든 실수형을 float64로 변환하여 처리
		n := ConvertValueToFloat64(v)
		return ConvertFloatToCommaString(n) // 아래에 정의된 실수 처리 헬퍼 함수 호출

	// 지원하지 않는 타입은 기본 문자열로 변환
	default:
		return fmt.Sprintf("%v", v)
	}
}

//#endregion 숫자를 컴마가 포함된 문자열로 변환

func GetNowDateOnlyString() string {
	return time.DateOnly
}
func GetNowDateString() string {
	return time.Now().Format("2006-01-02")
}
func GetNowTimeOnlyString() string {
	return time.Now().Format("15:04:05")
}
func GetNowDateTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// add 함수는 맵, 키, 추가할 값을 받아 알아서 처리해줍니다.
func SetValueAppendFromMap(m map[string][]any, key string, value interface{}) {
	// 맵에서 키에 해당하는 슬라이스를 가져옵니다.
	// 키가 존재하지 않으면, Go는 슬라이스의 제로 값(zero value)인 nil을 반환합니다.
	// append 함수는 nil 슬라이스에도 안전하게 동작합니다.
	m[key] = append(m[key], value)
}

// 슬라이스에서 특정 조건에 맞는 맵을 찾아 값을 설정합니다.
// findKey, findValue: 찾으려는 맵을 식별하기 위한 조건 (예: "id", 10)
// setKey, setValue: 추가하거나 수정하려는 키와 값
// 반환값: 해당 맵을 찾아 수정했는지 여부 (bool)
func SetValueFindKeyFromMapSlice(slice []map[string]any, findKey string, findValue any, setKey string, setValue any) bool {
	for _, item := range slice {
		// 맵 안에 찾는 키가 있고, 그 값이 일치하는지 확인
		if val, ok := item[findKey]; ok && val == findValue {
			// 조건을 만족하는 맵을 찾았으면, 새로운 값을 설정
			item[setKey] = setValue
			return true // 성공
		}
	}
	return false // 해당 맵을 찾지 못함
}

/*
// AddValueToNestedSlice 함수는 중첩된 구조에 안전하게 값을 추가합니다.
// slice: 전체 데이터 ([])
// findKey, findValue: 수정할 맵을 찾는 조건 (map)
// sliceKey: 값을 추가할 대상 슬라이스의 키 (string)
// valueToAdd: 최종적으로 추가할 값 ([]any)
func AddValueToNestedSlice(slice []map[string][]any, findKey string, findValue any, sliceKey string, valueToAdd any) bool {
	if slice == nil || findValue == nil || findKey == "" || sliceKey == "" || valueToAdd == nil || len(slice) == 0 {
		return false
	}
	// 1. 전체 슬라이스를 순회하며 올바른 맵을 찾습니다.
	for _, itemMap := range slice {
		if sliceVal, ok := itemMap[findKey]; ok && len(sliceVal) > 0 && sliceVal[0] == findValue {

			// 2. 맵을 찾았으면, append로 하위 슬라이스에 값을 추가합니다.
			// itemMap[sliceKey]가 nil(없는 경우)이더라도 append는 안전하게 동작하여
			// 새로운 슬라이스를 만들어줍니다.
			itemMap[sliceKey] = append(itemMap[sliceKey], valueToAdd)
			return true // 성공
		}
	}
	return false // 해당 맵을 찾지 못함
}
*/

func GetStringJoinFindKeyFromMapSlice(slice []map[string]any, findKey string, sep string) string {
	Result := ""
	var names []string
	for _, item := range slice {
		if val, ok := item[findKey]; ok {
			v := ConvertValueToString(val)
			names = append(names, v)
			//names = append(names, val.(string))
		}
	}
	Result = strings.Join(names, sep)
	return Result
}
func GetStringJoinFindKey2FromMapSlice(slice []map[string]any, findKey string, sep string) string {
	Result := ""
	var names []string
	for _, item := range slice {
		if val, ok := item[findKey]; ok {
			//str := val.(string)
			//v := val.(decimal.Decimal)
			//v := val.(godror.Number)
			v := ConvertValueToString(val)
			names = append(names, v)
		}
	}
	Result = strings.Join(names, sep)
	return Result
}
func GetStringArrayFindKeyFromMapSlice(slice []map[string]any, findKey string) []string {
	var names []string
	for _, item := range slice {
		if val, ok := item[findKey].(string); ok {
			names = append(names, val)
		}
	}
	return names
}

// 슬라이스에서 특정 키와 값이 일치하는 맵의 개수를 셉니다.
func GetCountWithFilter(slice []map[string]any, key string, value any) int {
	count := 0
	// 1. 슬라이스의 모든 맵을 순회합니다.
	for _, item := range slice {
		// 2. 맵에 해당 키가 있고, 그 값이 원하는 값과 일치하는지 확인합니다.
		if val, ok := item[key]; ok && val == value {
			// 3. 조건이 맞으면 카운트를 1 증가시킵니다.
			count++
		}
	}
	return count
}

// 슬라이스에서 특정 키를 기준으로 그룹화하여 각 그룹의 개수를 셉니다.
// data: 원본 데이터 슬라이스 ([]map[string]any)
// groupByKey: 그룹화의 기준이 될 맵의 키 (string)
// 반환값: 그룹별 개수를 담은 맵 (map[any]int)
func GetSelectGroupByCount(slice []map[string]any, groupByKey string) map[any]int {
	// 1. 결과를 담을 맵을 생성합니다. (키: 그룹 값, 값: 개수)
	counts := make(map[any]int)

	// 2. 전체 데이터를 순회합니다.
	for _, item := range slice {
		// 3. 그룹화할 키가 현재 맵에 존재하는지 확인합니다.
		if groupValue, ok := item[groupByKey]; ok {
			// 4. 존재한다면, 해당 값을 키로 사용하여 counts 맵의 값을 1 증가시킵니다.
			counts[groupValue]++
		}
	}
	return counts
}

// 상품 슬라이스와 '필터 함수'를 인자로 받습니다.
// 이 '필터 함수'는 상품 하나를 받아 조건을 만족하면 true를 반환합니다.
func GetSelectFilter(data []map[string]any, filter func(map[string]any) bool) []map[string]any {
	/** ex) 외부 호출시 사용 법

	//var val map[string]any
	tmp := HxCore.GetSelectFilter(data, func(r map[string]any) bool {
		return r["no"] == val["no"]
	})

	*/
	var result []map[string]any
	for _, p := range data {
		if filter(p) { // 전달받은 필터 함수로 조건을 검사
			result = append(result, p)
		}
	}
	return result
}

func GetNumberStringFromDateStr(value string) string {
	s := value
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, ".", "")
	return s
}
