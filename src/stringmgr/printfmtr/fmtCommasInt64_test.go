package printfmtr

import (
	"strconv"
	"testing"
)

func TestCommasInt64(t *testing.T) {
	var a int64 = 8100200300
	expected := "8,100,200,300"
	aa := CommasInt64(a)
	errMsg := `Expected value "%s" IS NOT EQUAL to CommasInt64() Result: %s`
	successMsg := `Expected value "%s" is equal to CommasInt64() Result: %s`
	t.Log("Given the need to test formatting and comma insertion for int64:")
	{
		t.Logf(`When checking formatted comma insertion for int64, %s`, strconv.FormatInt(a, 10))
		{
			if expected != aa {
				t.Errorf(errMsg, expected, aa)
			} else {
				t.Logf(successMsg, expected, aa)
			}
		}
	}
}
