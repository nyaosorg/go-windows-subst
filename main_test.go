package subst

import (
	"os/exec"
	"strings"
	"testing"
)

func TestUtf16toStringArray(t *testing.T) {
	in := []uint16{'A', 0, 'B', 'C', 0, 'D', 'E', 'F'}
	out := utf16toStringArray(in)

	if len(out) != 3 {
		t.Fatal("len(out) != 3")
	}
	if out[0] != "A" {
		t.Fatal("out[0] != \"A\"")
	}
	if out[1] != "BC" {
		t.Fatal("out[1] != \"BC\"")
	}
	if out[2] != "DEF" {
		t.Fatal("out[2] != \"DEF\"")
	}
}

func TestQueryDosDevices(t *testing.T) {
	_, err := queryDosDevices()
	if err != nil {
		t.Fatal(err.Error())
	}
	// for i, s := range list {
	// println(i, s)
	//}
}

const SUBSTDRIVE = `B:`
const SUBSTDIR = `C:\Windows`

func cmdQuery(drive string) (string, error) {
	out, err := exec.Command("subst").Output()
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(out), "\n") {
		pair := strings.SplitN(strings.TrimSpace(line), " => ", 2)
		if len(pair) >= 2 && pair[0] == SUBSTDRIVE+`\:` {
			return pair[1], nil
		}
	}
	return "", nil
}

func cleanup() {
	exec.Command("subst", SUBSTDRIVE, "/D").Run()
}

func TestDefineAndRemove(t *testing.T) {
	cleanup()
	err := Define(SUBSTDRIVE, SUBSTDIR)
	if err != nil {
		t.Fatalf("Define: %s", err.Error())
	}
	result, err := cmdQuery(SUBSTDRIVE)
	if err != nil {
		t.Fatalf("cmdQuery: %s", err.Error())
	}
	if result != SUBSTDIR {
		cleanup()
		t.Fatalf("Define: not match: '%s' != '%s'", result, SUBSTDIR)
	}
	err = Remove(SUBSTDRIVE)
	if err != nil {
		t.Fatalf("Remove: %s", err.Error())
	}
	result, err = cmdQuery(SUBSTDRIVE)
	cleanup()
	if err != nil {
		t.Fatalf("cmdQuery: %s", err.Error())
	}
	if result != "" {
		t.Fatalf("Remove: failed '%s' remains", result)
	}
}

func TestQuery(t *testing.T) {
	cleanup()
	err := exec.Command("subst", SUBSTDRIVE, SUBSTDIR).Run()
	if err != nil {
		t.Fatalf("subst: %s", err.Error())
	}
	result, err := Query(SUBSTDRIVE)
	cleanup()
	if err != nil {
		t.Fatalf("Query: %s", err.Error())
	}
	if result != SUBSTDIR {
		t.Fatalf("Query: '%s' != '%s'", result, SUBSTDIR)
	}
}
