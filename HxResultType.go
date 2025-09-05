package HxCore

import "encoding/json"

//#region HxResultType

type HxResultType int

const (
	None    HxResultType = iota
	Notice               = 1
	Warning              = 2
	Success              = 3
	Fail                 = 4
)

func (r HxResultType) String() string {
	/*
		switch r {
		case None:
			return "None"
		case Notice:
			return "Notice"
		case Warning:
			return "Warning"
		case Success:
			return "Success"
		case Fail:
			return "Fail"
		default:
			return "Unknown : " + string(r) //return "Unknown"
		}
	*/
	var names = [...]string{
		"None",
		"Notice",
		"Warning",
		"Success",
		"Fail",
	}

	// 범위 체크
	if r < None || r > Fail {
		return "Unknown"
	}
	return names[r]
}

func (r HxResultType) MarshalJSON() ([]byte, error) {
	// 1. String() 메서드를 호출하여 "Success" 같은 문자열을 얻습니다.
	stringValue := r.String()

	// 2. 이 문자열을 유효한 JSON 문자열(즉, 큰따옴표로 감싸진)로 변환하여 반환합니다.
	//    json.Marshal 함수에 문자열을 넘기면 알아서 따옴표를 붙여줍니다.
	return json.Marshal(stringValue)
}

//#endregion HxResultType
