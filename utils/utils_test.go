package utils

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestGetPidFile(t *testing.T) {

	_, e := GetPidFile()
	if e != ErrPidDirConfig {
		t.Errorf("Expexted %v, but got %v", ErrPidDirConfig, e)
	}

	viper.SetDefault(":pid_dir", "/opt/codedeploy-agent/state/.pid/")
	_, e = GetPidFile()
	if e != ErrProgramName {
		t.Errorf("Expexted %v, but got %v", ErrProgramName, e)
	}

	viper.SetDefault(":program_name", "codedeploy-agent")

	r, e := GetPidFile()

	if e != nil {
		t.Errorf("Expexted no error, but got %v", e)
	}

	expectedPid := "/opt/codedeploy-agent/state/.pid/codedeploy-agent.pid"
	if strings.Compare(r, expectedPid) != 0 {
		t.Errorf("expected '%s', but got '%s'", expectedPid, r)
	}
	viper.Reset()

}

// TODO improve this test ... Add tables
func TestGetPidID(t *testing.T) {

	getPidIDTests := []struct {
		IntValue int
		StrValue string
		Err      bool
	}{
		{
			123,
			"123",
			false,
		},
		{
			124,
			"124\n",
			false,
		},
		{
			-1,
			"someting",
			true,
		},
	}

	for _, v := range getPidIDTests {
		r, e := GetPidID(bytes.NewBufferString(v.StrValue))
		if v.Err && e == nil {
			t.Error("Expexted error, but got no error")
		}
		if !v.Err && e != nil {
			t.Errorf("Expexted no error, but got %v", e)
		}
		if !v.Err && r != v.IntValue {
			t.Errorf("Expexted %d, but got %d", v.IntValue, r)
		}
	}

}
