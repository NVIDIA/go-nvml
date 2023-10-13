/**
# Copyright 2023 NVIDIA CORPORATION
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package nvml

import "fmt"

// LookupSymbol provides top-level access to symbol lookup in the configured library.
// Note that this requires that the library be loaded as is done in a call to nvml.Init().
//
// Deprecated: LookupSymbol represents an unstable API and should be considered experimental.
func LookupSymbol(name string) error {
	if nvml == nil {
		return fmt.Errorf("error looking up %s: library %s not initialized", name, nvmlLibraryName)
	}
	return nvml.Lookup(name)
}
