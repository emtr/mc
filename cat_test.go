/*
 * Minio Client (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/minio/mc/pkg/console"
	. "gopkg.in/check.v1"
)

func (s *CmdTestSuite) TestCatCmd(c *C) {
	/// filesystem
	root, err := ioutil.TempDir(os.TempDir(), "cmd-")
	c.Assert(err, IsNil)
	defer os.RemoveAll(root)

	objectPath := filepath.Join(root, "object1")
	objectPathServer := server.URL + "/bucket/object1"
	data := "hello"
	dataLen := len(data)
	err = putTarget(objectPath, int64(dataLen), bytes.NewReader([]byte(data)))
	c.Assert(err, IsNil)
	err = putTarget(objectPathServer, int64(dataLen), bytes.NewReader([]byte(data)))
	c.Assert(err, IsNil)

	var sourceURLs []string
	sourceURLs = append(sourceURLs, objectPath)
	sourceURLs = append(sourceURLs, objectPathServer)
	for _, sourceURL := range sourceURLs {
		c.Assert(doCatCmd(sourceURL), IsNil)
	}

	objectPath = filepath.Join(root, "object2")
	c.Assert(doCatCmd(objectPath), Not(IsNil))
}

func (s *CmdTestSuite) TestCatContext(c *C) {
	err := app.Run([]string{os.Args[0], "cat", server.URL + "/bucket/object1"})
	c.Assert(err, IsNil)
	c.Assert(console.IsExited, Equals, false)

	err = app.Run([]string{os.Args[0], "cat", server.URL + "/invalid"})
	c.Assert(err, IsNil)
	c.Assert(console.IsExited, Equals, true)
	// reset back
	console.IsExited = false
}
