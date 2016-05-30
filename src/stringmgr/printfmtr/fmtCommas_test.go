package printfmtr

import (
	"strconv"
	"testing"
)

func TestCommasInt(t *testing.T) {
	var a = 6100200300
	expected := "6,100,200,300"
	aa := CommasInt(a)
	errMsg := `Expected value "%s" IS NOT EQUAL to CommasInt() Result: %s`
	successMsg := `Expected value "%s" is equal to CommasInt() Result: %s`
	t.Log("Given the need to test formatting and comma insertion for int:")
	{
		t.Logf(`When checking formatted comma insertion for int, %s`, strconv.Itoa(a))
		{
			if expected != aa {
				t.Errorf(errMsg, expected, aa)
			} else {
				t.Logf(successMsg, expected, aa)
			}
		}
	}
}
