package metadata_test

import (
	"flag"
	"io/ioutil"
	"koding/kites/kloud/metadata"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

var update = flag.Bool("update-golden", false, "Update golden files.")

var testdata = metadata.NewCloudInit(&metadata.CloudConfig{}) // TODO

func parseCloudInit(file string) (metadata.CloudInit, error) {
	p, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return metadata.ParseCloudInit(p)
}

func TestCloudInit(t *testing.T) {
	cases := map[string]string{
		"testdata/cloud-init.yml":          "",
		"testdata/ssh_authorized_keys.yml": "testdata/ssh_authorized_keys.yml.golden",
		"testdata/users.yml":               "testdata/users.yml.golden",
		"testdata/write_files.yml":         "testdata/write_files.yml.golden",
	}

	for yml, golden := range cases {
		t.Run(filepath.Base(yml), func(t *testing.T) {
			if *update {
				// TODO
				return
			}

			got, err := parseCloudInit(yml)
			if err != nil {
				t.Fatalf("ParseCloudInit()=%s", err)
			}

			pretty.Println(got)

			want, err := parseCloudInit(golden)
			if err != nil {
				t.Fatalf("ParseCloudInit()=%s", err)
			}

			got.MergeIn(testdata)

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %+v, want %+v", got, want)
			}
		})
	}
}
