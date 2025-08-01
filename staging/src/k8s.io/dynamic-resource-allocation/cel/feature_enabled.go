/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cel

// Define functions and global variables to set FeatureEnabled field in VersionedOptions
// within this package (compile.go).

import "sync/atomic"

var (
	consumableCapacity atomic.Value
)

func SetDRAConsumableCapacity() {
	consumableCapacity.Store(true)
}

func UnsetDRAConsumableCapacity() {
	consumableCapacity.Store(false)
}

func DRAConsumableCapacity() bool {
	value := consumableCapacity.Load()
	if value == nil {
		return false
	}
	return value.(bool)
}
