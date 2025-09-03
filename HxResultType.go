package HxCore

//#region HxResultType

type HxResultType int

const (
	None    HxResultType = iota
	Notice  HxResultType = 1
	Warning HxResultType = 2
	Success HxResultType = 3
	Fail    HxResultType = 4
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

//#endregion HxResultType
