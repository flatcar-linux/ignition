// Copyright 2017 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package files

import (
	"github.com/coreos/ignition/v2/tests/register"
	"github.com/coreos/ignition/v2/tests/types"
)

func init() {
	register.Register(register.NegativeTest, InvalidHash())
	register.Register(register.NegativeTest, InvalidHashFromHTTPURL())
}

func InvalidHash() types.Test {
	name := "files.verification.badhash.dataurl"
	in := types.GetBaseDisk()
	out := in
	mntDevices := []types.MntDevice{
		{
			Label:        "EFI-SYSTEM",
			Substitution: "$DEVICE",
		},
	}

	config := `{
		"ignition": {"version": "$version" },
		"storage": {
			"filesystems": [{
				"path": "/tmp0",
				"device": "$DEVICE",
				"format": "ext4",
				"wipeFilesystem": true
			}],
			"files": [{
				"path": "/tmp0/test",
				"contents": {
					"source": "data:,asdf", "verification": {"hash": "sha512-1a04c76c17079cd99e688ba4f1ba095b927d3fecf2b1e027af361dfeafb548f7f5f6fdd675aaa2563950db441d893ca77b0c3e965cdcb891784af96e330267d7"}}
			}]}
	}`
	configMinVersion := "3.0.0"

	return types.Test{
		Name:             name,
		In:               in,
		Out:              out,
		MntDevices:       mntDevices,
		Config:           config,
		ConfigMinVersion: configMinVersion,
	}
}

func InvalidHashFromHTTPURL() types.Test {
	name := "files.verification.badhash.http"
	in := types.GetBaseDisk()
	out := in
	config := `{
	  "ignition": { "version": "$version" },
	  "storage": {
	    "files": [{
	      "path": "/foo/bar",
	      "contents": {
	        "source": "http://127.0.0.1:8080/contents",
			"verification": {"hash": "sha512-807e8ff949e61d23f5ee42a629ec96e9fc526b62f030cd70ba2cd5b9d97935461eacc29bf58bcd0426e9e1fdb0eda939603ed52c9c06d0712208a15cd582c60e"}
	      }
	    }]
	  }
	}`
	configMinVersion := "3.0.0"
	out[0].Partitions.AddFiles("ROOT", []types.File{
		{
			Node: types.Node{
				Name:      "bar",
				Directory: "foo",
			},
			Contents: "asdf\nfdsa",
		},
	})

	return types.Test{
		Name:             name,
		In:               in,
		Out:              out,
		Config:           config,
		ConfigMinVersion: configMinVersion,
	}
}
