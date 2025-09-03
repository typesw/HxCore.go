package HxCore

/**
 * Create By J.H.Park (2025-09-03)
 * - string을 클래스처럼 쓰고 싶어서...그런데 실제로 쓸일이 많을까 싶음 ㅜ.ㅜ
 * - string과 HxUtils를 잘 만들어 쓰자~
 * ※ HxUtils를 참조하고 있으므로, 상호 참조가 안되도록 주의
 */

//#region HxString 타입 & 문자처리 함수

type HxString string

// string을 HxString으로 변환하는 함수
func ConvertStringToHxString(s string) HxString {
	// 간단한 타입 캐스팅을 통해 변환
	return HxString(s)
}
func ConvertValueToHxString(value interface{}) HxString {
	strValue, isString := value.(string)

	if isString != true {
		strValue = ""
	}

	return ConvertStringToHxString(strValue)
}

// HxString을 string으로 변환하는 함수 (메서드 방식)
func (s HxString) String() string {
	// 간단한 타입 캐스팅을 통해 변환
	return string(s)
}
func (s HxString) ToString() string {
	return s.String()
}
func (s HxString) Trim() HxString {
	var str string = Trim(s.ToString())
	return ConvertStringToHxString(str)
}
func (s HxString) Lower() HxString {
	var str string = Lower(s.ToString())
	return ConvertStringToHxString(str)
}
func (s HxString) Upper() HxString {
	var str string = Upper(s.ToString())
	return ConvertStringToHxString(str)
}
func (s HxString) SubStr(start int, end int) HxString {
	var str string = SubStr(s.ToString(), start, end)
	return ConvertStringToHxString(str[start:end])
}
func (s HxString) IsRegexpMatch(pattern string) bool {
	return IsRegexpMatch(pattern, s.ToString())
}
func (s HxString) RegexpMatch(pattern string) map[int]string {
	return RegexpMatch(pattern, s.ToString())
}
func (s HxString) ToInt() int {
	return ConvertStringToInt(s.ToString())
}
func (s HxString) ToInt64() int64 {
	return ConvertStringToInt64(s.ToString())
}
func (s HxString) ToFloat() float32 {
	return ConvertStringToFloat32(s.ToString())
}
func (s HxString) ToFloat64() float64 {
	return ConvertStringToFloat64(s.ToString())
}
func (s HxString) ToBool() bool {
	return ConvertStringToBool(s.ToString())
}
func (s HxString) ToNumber() float64 {
	return ConvertStringToNumber(s.ToString())
}
func (s HxString) ConvertInt64ToCommaString() HxString {
	return ConvertStringToHxString(ConvertInt64ToCommaString(s.ToInt64()))
}

//#endregion HxString
