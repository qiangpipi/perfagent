package functions

import (
	//	. "pagent/logs"
	"testing"
)

func Test_procUsage01(t *testing.T) {
	t.Log("Check proc usage of PID 1")
	cu, mu, err := procUsageByPid("1")
	if err != nil {
		t.Log("CPU usage:", cu)
		t.Log("Mem usage:", mu)
		t.Error(err)
	} else {
		if len(cu) < 0 {
			t.Error("Didn't get cpu usage")
		}
		if len(mu) < 0 {
			t.Error("Didn't get mem usage")
		}
	}
}
