package HxCore

import (
	"net/http"
	"strings"
)

// GetFormValueIgnoreCase는 http.Request에서 대소문자를 구분하지 않고 파라미터 값을 찾습니다.
// 일치하는 키가 없으면 빈 문자열을 반환합니다.
func GetRequestFormValueIgnoreCase(r *http.Request, key string) string {
	// 1. r.Form을 채우기 위해 ParseForm을 호출합니다. (GET, POST 모두 처리)
	//    여러 번 호출해도 안전하게 한 번만 파싱됩니다.
	r.ParseForm()

	// 2. r.Form 맵의 모든 키를 순회합니다.
	for k, v := range r.Form {
		// 3. strings.EqualFold 함수로 대소문자 무시하고 비교합니다.
		if strings.EqualFold(k, key) {
			// 4. 일치하는 키를 찾으면, 그 키의 첫 번째 값을 반환합니다.
			if len(v) > 0 {
				return v[0]
			}
		}
	}

	// 5. 일치하는 키가 없으면 빈 문자열을 반환합니다.
	return ""
}

func GetRequestFormValue(r *http.Request, key string) string {
	return GetRequestFormValueIgnoreCase(r, key)
}
