package HxCore

import (
	"fmt"
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

func RegexpMatch(pattern string, str string) map[int]string {
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

//#region 타입 변환 헬퍼

func ConvertValueToInt64(v interface{}) int64 {
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
