package subst

import (
	"os/exec"
	"strings"
	"testing"
)

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
