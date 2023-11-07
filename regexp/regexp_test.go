package regexp

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/wujiu2020/lux/request"
)

func TestRegexp(t *testing.T) {
	html, err := request.Get("https://v.ifeng.com/c/8O1MxOEuQ4k", "", nil)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(html)
	re := regexp.MustCompile(`"docData":(.*?)\,"parentColumn"`)
	match := re.FindStringSubmatch(html)
	t.Log(len(match))
	t.Log(match[1])
}
