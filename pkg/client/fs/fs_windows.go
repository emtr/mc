// +build windows

/*
 * Minio Client (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this fs except in compliance with the License.
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

package fs

import (
	"path/filepath"
	"syscall"
)

func normalizePath(path string) string {
	if filepath.VolumeName(path) == "" && filepath.HasPrefix(path, "\\") {
		path, err = syscall.FullPath(path)
		if err != nil {
			panic(err)
		}
	}
	return path
}
