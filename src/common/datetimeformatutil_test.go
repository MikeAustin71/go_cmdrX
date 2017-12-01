package common

import (
	"fmt"
	"testing"
)

func TestDateTimeFormatUtility_ParseDateTimeStrings(t *testing.T) {
	dtf := DateTimeFormatUtility{}

	dtf.CreateAllFormatsInMemory()

	dtSamples := GetXDateTimeSamples()

	ls := len(dtSamples)
	fmtDateTimeEverything := "Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	for i := 0; i < ls; i++ {

		fmt.Printf("Item: %v  Date Time String : %v \n", i, dtSamples[i][0][0])
		ti, err := dtf.ParseDateTimeString(dtSamples[i][0][0], "")

		if err != nil {
			t.Errorf("Error on dtf.ParseDateTimeString() - Sample Format: %v  Error: %v", dtSamples[i][0][0], err.Error())
		}

		tiStr := ti.Format(fmtDateTimeEverything)
		expected := dtSamples[i][0][1]
		if tiStr != expected {
			t.Errorf("dtSample== %v. Expected: %v - Received %v", dtSamples[i][0][0], expected, tiStr)
		}

	}

	fmt.Printf("Processed %v sample dates successfully!\n", ls)

}

func GetXDateTimeSamples() [][][]string {
	d := make([][][]string, 0)
	// FmtDateTimeEverything := "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	d = append(d, [][]string{{"Saturday 11/12/2016 4:26 PM", "Saturday November 12, 2016 16:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"7-6-16 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"7-6-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"7-06-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"07-6-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"07-06-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"12 Nov 2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"12 November 2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 11:26pm -0600 CST 2016", "Saturday November 12, 2016 23:26:00.000000000 -0600 CST"}})
	d = append(d, [][]string{{"12 November 2016 23:26:00 -0600 CST", "Saturday November 12, 2016 23:26:00.000000000 -0600 CST"}})
	d = append(d, [][]string{{"November 12, 2016 11:6pm +0000 UTC", "Saturday November 12, 2016 23:06:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 2016 11:6 p m +0000 UTC", "Saturday November 12, 2016 23:06:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 2016 1:6pm +0000 UTC", "Saturday November 12, 2016 13:06:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 2016 1:06pm -0500 EST", "Saturday November 12, 2016 13:06:00.000000000 -0500 EST"}})
	d = append(d, [][]string{{"2016-11-12 13:6 -0500 EST", "Saturday November 12, 2016 13:06:00.000000000 -0500 EST"}})
	d = append(d, [][]string{{"5/31/2017 23:2:17 -0700 PDT", "Wednesday May 31, 2017 23:02:17.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"2016-11-12 23:26:00 +0000 UTC", "Saturday November 12, 2016 23:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2016-11-12 23:26:00Z", "Saturday November 12, 2016 23:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2017-6-12 11:26 p.m. Z", "Monday June 12, 2017 23:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2017-11-26 16:26 -0600 CST", "Sunday November 26, 2017 16:26:00.000000000 -0600 CST"}})
	d = append(d, [][]string{{"2017-6-5 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST"}})
	d = append(d, [][]string{{"2017-6-05 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST"}})
	d = append(d, [][]string{{"2017-06-5 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST"}})
	d = append(d, [][]string{{"2017-06-05 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST"}})
	d = append(d, [][]string{{"11/12/16 4:26 PM", "Saturday November 12, 2016 16:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 4:4 P.M.", "Saturday November 12, 2016 16:04:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 4:4:4.012 AM", "Saturday November 12, 2016 04:04:04.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/1/16 4:4:04.012 A.M.", "Tuesday November 1, 2016 04:04:04.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/1/2016 4:4:04.012 A.M.", "Tuesday November 1, 2016 04:04:04.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/1/2016 4:4:04.012 A.M. ", "Tuesday November 1, 2016 04:04:04.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"Monday June 5, 2017 17:24:46.064223400 -0500 CDT", "Monday June 5, 2017 17:24:46.064223400 -0500 CDT"}})
	d = append(d, [][]string{{"6-5-2017 17:30:17 -0700 PDT", "Monday June 5, 2017 17:30:17.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"06-05-2017 17:30:17 -0700 PDT", "Monday June 5, 2017 17:30:17.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"06-5-2017 17:30:17 -0700 PDT", "Monday June 5, 2017 17:30:17.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"11/12/16 4:04:0.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/2/16 04:04:0.012 PM", "Wednesday November 2, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 4:04:00.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:4:0.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:04:00.012 AM", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:04:00.012 am", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:04:00.012 A.M.", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:4:0.012 P.M.", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/2016 04:4:0.012 P.M.", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/2016 04:4:0.012 PM.", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"5/27/2017 11:42PM CDT", "Saturday May 27, 2017 23:42:00.000000000 -0500 CDT"}})
	d = append(d, [][]string{{"06/1/2017 11:42 -0700 PDT", "Thursday June 1, 2017 11:42:00.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"2016-11-26 16:26 CDT -0600", "Saturday November 26, 2016 16:26:00.000000000 -0600 CDT"}})
	d = append(d, [][]string{{"2016/11/26 16:2:3 PDT -0700", "Saturday November 26, 2016 16:02:03.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"June 12th, 2016 4:26 PM", "Sunday June 12, 2016 16:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"05.03.2017", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"5.03.2017", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"5.3.2017", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"5.3.'17", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2017.3.5", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"6/27/2017 23:26:01 -0500 CDT", "Tuesday June 27, 2017 23:26:01.000000000 -0500 CDT"}})
	d = append(d, [][]string{{"23:26:01 -0500 CDT", "Saturday January 1, 0000 23:26:01.000000000 -0500 CDT"}})
	d = append(d, [][]string{{"11-26-2016 16:26 -0600 CST", "Saturday November 26, 2016 16:26:00.000000000 -0600 CST"}})
	d = append(d, [][]string{{"11-26-2016 16:26:0 -0600 CST", "Saturday November 26, 2016 16:26:00.000000000 -0600 CST"}})
	d = append(d, [][]string{{"Monday June 5th2017 17:24:46.064223400 -0500 CDT", "Monday June 5, 2017 17:24:46.064223400 -0500 CDT"}})
	d = append(d, [][]string{{"5/27/2017 11:42PMCDT", "Saturday May 27, 2017 23:42:00.000000000 -0500 CDT"}})
	d = append(d, [][]string{{"06/1/2017 11:42 PM-0700 PDT", "Thursday June 1, 2017 23:42:00.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"06/1/2017 11:42:00   PM  -0700 PDT", "Thursday June 1, 2017 23:42:00.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"June 1st, 2017 11:42:00PM -0700 PDT", "Thursday June 1, 2017 23:42:00.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"June 2nd 2017 11:42:00PM -0700 PDT", "Friday June 2, 2017 23:42:00.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"June 3rd, 2017 11:42:00PM -0700 PDT", "Saturday June 3, 2017 23:42:00.000000000 -0700 PDT"}})

	return d
}
