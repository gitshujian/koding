package metadata

import (
	"bytes"
	"errors"

	yaml "gopkg.in/yaml.v2"
)

var header = []byte("#cloud-config")

var ErrNotCloudInit = errors.New("not a cloud-init content")

type CloudInit map[string]interface{}

func (ci CloudInit) MergeIn(mixin CloudInit) {

}

func (ci CloudInit) String() string {
	// header + marshal
	return ""
}

type CloudConfig struct {
}

func NewCloudInit(cfg *CloudConfig) CloudInit {
	return metadata.CloudInit{
		"users": []interface{}{
			"default",
			map[interface{}]interface{}{
				"name":        "{{.Username}}",
				"lock_passwd": bool(true),
				"gecos":       "Koding",
				"groups": []interface{}{
					"sudo",
				},
				"sudo": []interface{}{
					"ALL=(ALL) NOPASSWD:ALL",
				},
				"shell": "/bin/bash",
			},
		},
		"write_files": []interface{}{
			map[interface{}]interface{}{
				"path":    "/etc/kite/kite.key",
				"content": "{{.KiteKey}}\n",
			},
			map[interface{}]interface{}{
				"content":     "TODO",
				"path":        "/var/lib/koding/metadata.json",
				"permissions": "0644",
			},
			map[interface{}]interface{}{
				"path":        "/var/lib/koding/user-data.sh",
				"permissions": "0755",
				"encoding":    "b64",
				"content":     "{{.UserData}}",
			},
		},
	}
}

func ParseCloudInit(p []byte) (CloudInit, error) {
	if !bytes.HasPrefix(p, header) {
		return nil, ErrNotCloudInit
	}

	var ci CloudInit

	if err := yaml.Unmarshal(p, &ci); err != nil {
		return nil, err
	}

	return ci, nil
}
