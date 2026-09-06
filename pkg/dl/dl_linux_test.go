/**
# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
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

package dl

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPathNotOpened(t *testing.T) {
	t.Parallel()
	dl := New("libdl.so.2", RTLD_LAZY|RTLD_GLOBAL)

	_, err := dl.Path()
	if err == nil {
		t.Fatal("Path() on an unopened library should have errored but did not")
	}
	if !strings.Contains(err.Error(), "not opened") {
		t.Errorf("unexpected error from Path() on an unopened library: %v", err)
	}
}

// TestPathFromSoname covers the dlinfo(RTLD_DI_ORIGIN) branch a name without a
// slash takes.
func TestPathFromSoname(t *testing.T) {
	skipOnMacOS(t)

	t.Parallel()
	dl := New("libdl.so.2", RTLD_LAZY|RTLD_GLOBAL)

	if err := dl.Open(); err != nil {
		t.Fatalf("error opening shared lib: %v", err)
	}

	path, err := dl.Path()
	if err != nil {
		t.Fatalf("error getting library path: %v", err)
	}
	if filepath.Base(path) != "libdl.so.2" {
		t.Errorf("Path() = %q, want a path ending in libdl.so.2", path)
	}
	if !filepath.IsAbs(path) {
		t.Errorf("Path() = %q, want an absolute path", path)
	}

	if err := dl.Close(); err != nil {
		t.Errorf("error closing shared lib: %v", err)
	}
}

// TestPathFromExplicitPath covers the short-circuit for a name carrying a slash.
func TestPathFromExplicitPath(t *testing.T) {
	skipOnMacOS(t)

	t.Parallel()
	resolver := New("libdl.so.2", RTLD_LAZY|RTLD_GLOBAL)
	if err := resolver.Open(); err != nil {
		t.Fatalf("error opening shared lib: %v", err)
	}
	explicit, err := resolver.Path()
	if err != nil {
		t.Fatalf("error getting library path: %v", err)
	}
	if err := resolver.Close(); err != nil {
		t.Fatalf("error closing shared lib: %v", err)
	}

	dl := New(explicit, RTLD_LAZY|RTLD_GLOBAL)
	if err := dl.Open(); err != nil {
		t.Fatalf("error opening shared lib by path: %v", err)
	}
	defer dl.Close()

	path, err := dl.Path()
	if err != nil {
		t.Fatalf("error getting library path: %v", err)
	}
	if path != explicit {
		t.Errorf("Path() = %q, want %q", path, explicit)
	}
}
