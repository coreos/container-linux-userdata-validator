package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/coreos/coreos-cloudinit-validate/third_party/github.com/coreos/coreos-cloudinit/third_party/gopkg.in/yaml.v1"
)

// CloudConfig encapsulates the entire cloud-config configuration file and maps directly to YAML
type CloudConfig struct {
	SSHAuthorizedKeys []string `yaml:"ssh_authorized_keys"`
	Coreos            struct {
		Etcd   Etcd   `yaml:"etcd"`
		Fleet  Fleet  `yaml:"fleet"`
		OEM    OEM    `yaml:"oem"`
		Update Update `yaml:"update"`
		Units  []Unit `yaml:"units"`
	} `yaml:"coreos"`
	WriteFiles        []File   `yaml:"write_files"`
	Hostname          string   `yaml:"hostname"`
	Users             []User   `yaml:"users"`
	ManageEtcHosts    EtcHosts `yaml:"manage_etc_hosts"`
	NetworkConfigPath string   `yaml:"-"`
	NetworkConfig     string   `yaml:"-"`
}

// NewCloudConfig instantiates a new CloudConfig from the given contents (a
// string of YAML), returning any error encountered. It will ignore unknown
// fields but log encountering them.
func NewCloudConfig(contents string) (*CloudConfig, error) {
	var cfg CloudConfig
	err := yaml.Unmarshal([]byte(contents), &cfg)
	if err != nil {
		return &cfg, err
	}
	return &cfg, nil
}

func (cc CloudConfig) String() string {
	bytes, err := yaml.Marshal(cc)
	if err != nil {
		return ""
	}

	stringified := string(bytes)
	stringified = fmt.Sprintf("#cloud-config\n%s", stringified)

	return stringified
}

func IsEmpty(c interface{}) bool {
	cv := reflect.ValueOf(c)

	for i := 0; i < cv.NumField(); i++ {
		if cv.Field(i).String() != "" {
			return false
		}
	}
	return true
}

// AssertValid checks the fields in the structure and makes sure that they
// contain valid values as specified by the 'valid' flag. Empty fields are
// implicitly valid.
func AssertValid(c interface{}) error {
	ct := reflect.TypeOf(c)
	cv := reflect.ValueOf(c)
	for i := 0; i < ct.NumField(); i++ {
		ft := ct.Field(i)
		valid := ft.Tag.Get("valid")
		name := ft.Name
		val := cv.Field(i).String()

		if !isFieldValid(valid, val) {
			return fmt.Errorf("invalid value %q for option %q (valid options: %q)", val, name, valid)
		}
	}
	return nil
}

func isFieldValid(valid, val string) bool {
	if valid == "" || val == "" {
		return true
	}
	for _, v := range strings.Split(valid, ",") {
		if val == v {
			return true
		}
	}
	return false
}
