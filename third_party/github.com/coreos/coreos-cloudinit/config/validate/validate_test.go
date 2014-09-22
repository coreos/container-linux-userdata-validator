package validate

/*
import (
	"fmt"
	"strings"
	"testing"
)

func TestCloudConfigUnknownKeys(t *testing.T) {
			contents := `
		coreos:
		  etcd:
		    discovery: "https://discovery.etcd.io/827c73219eeb2fa5530027c37bf18877"
		  coreos_unknown:
		    foo: "bar"
		section_unknown:
		  dunno:
		    something
		bare_unknown:
		  bar
		write_files:
		  - content: fun
		    path: /var/party
		    file_unknown: nofun
		users:
		  - name: fry
		    passwd: somehash
		    user_unknown: philip
		hostname:
		  foo
		`
				cfg, err := NewCloudConfig(contents)
				if err != nil {
					t.Fatalf("error instantiating CloudConfig with unknown keys: %v", err)
				}
				if cfg.Hostname != "foo" {
					t.Fatalf("hostname not correctly set when invalid keys are present")
				}
				if cfg.Coreos.Etcd.Discovery != "https://discovery.etcd.io/827c73219eeb2fa5530027c37bf18877" {
					t.Fatalf("etcd section not correctly set when invalid keys are present")
				}
				if len(cfg.WriteFiles) < 1 || cfg.WriteFiles[0].Content != "fun" || cfg.WriteFiles[0].Path != "/var/party" {
					t.Fatalf("write_files section not correctly set when invalid keys are present")
				}
				if len(cfg.Users) < 1 || cfg.Users[0].Name != "fry" || cfg.Users[0].PasswordHash != "somehash" {
					t.Fatalf("users section not correctly set when invalid keys are present")
				}

				var warnings string
				catchWarn := func(f string, v ...interface{}) {
					warnings += fmt.Sprintf(f, v...)
				}

				warnOnUnrecognizedKeys(contents, catchWarn)

				if !strings.Contains(warnings, "coreos_unknown") {
					t.Errorf("warnings did not catch unrecognized coreos option coreos_unknown")
				}
				if !strings.Contains(warnings, "bare_unknown") {
					t.Errorf("warnings did not catch unrecognized key bare_unknown")
				}
				if !strings.Contains(warnings, "section_unknown") {
					t.Errorf("warnings did not catch unrecognized key section_unknown")
				}
				if !strings.Contains(warnings, "user_unknown") {
					t.Errorf("warnings did not catch unrecognized user key user_unknown")
				}
				if !strings.Contains(warnings, "file_unknown") {
					t.Errorf("warnings did not catch unrecognized file key file_unknown")
				}
}
*/
