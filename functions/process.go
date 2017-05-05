package functions

import (
	"errors"
	"net/http"
	"os/exec"
	. "perfagent/logs"
	"strings"
)

func Ps(w http.ResponseWriter, r *http.Request) {
	cu := ""
	mu := ""
	if err := r.ParseForm(); err != nil {
		Error(err)
	} else {
		for k, v := range r.Form {
			if k == "pid" {
				cu, mu, err = procUsageByPid(strings.Join(v, ""))
				if err != nil {
					w.Write([]byte(err.Error()))
				} else {
					tmp := "{\"CPU%\":" + cu + "\"MEM%\":" + mu + "}"
					w.Write([]byte(tmp))
				}
			} else if k == "pname" {
				cu, mu, err = procUsageByName(strings.Join(v, ""))
				if err != nil {
					w.Write([]byte(err.Error()))
				} else {
					tmp := "{\"CPU%\":" + cu + "\"MEM%\":" + mu + "}"
					w.Write([]byte(tmp))
				}
			} else {
				w.Write([]byte("Please provide params!"))
			}
		}
	}
}

func Plist(w http.ResponseWriter, r *http.Request) {
}

func Monitor(w http.ResponseWriter, r *http.Request) {
}

func procUsageByName(pName string) (cpuUsage string, memUsage string, err error) {
	if pid, err := getPidByName(pName); err != nil {
		return "0", "0", err
	} else {
		return procUsageByPid(pid)
	}
}

func getPidByName(pName string) (pid string, err error) {
	cmd := exec.Command("pgrep", pName)
	out, err := cmd.Output()
	if err != nil {
		Error("Cmd exec failed")
		return "0", err
	} else {
		tmp := strings.Fields(string(out))
		if len(tmp) > 1 {
			return "0", errors.New("Multiple process running")
		} else {
			return string(out), nil
		}
	}
}

func procUsageByPid(pid string) (cpuUsage string, memUsage string, err error) {
	cmd := exec.Command("ps", "h", "o", "pcpu,pmem", "-q", pid)
	out, err := cmd.Output()
	if err != nil {
		Error("Cmd exec failed")
		return "0", "0", err
	} else {
		tmp := strings.Fields(string(out))
		cpuUsage = tmp[0]
		memUsage = tmp[1]
		Debug(cpuUsage)
		Debug(memUsage)
		return cpuUsage, memUsage, nil
	}
}
