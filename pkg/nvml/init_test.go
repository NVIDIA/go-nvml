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

package nvml

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// Shutdown on a library that was never loaded must return rather than reach
// nvmlShutdown, whose symbol is unresolved until Init has opened the library.
func TestShutdownWithoutLoadIsUninitialized(t *testing.T) {
	l := newTestLibrary(&dynamicLibraryMock{})

	require.Equal(t, refcount(0), l.refcount)
	require.Equal(t, ERROR_UNINITIALIZED, l.Shutdown())
	require.Equal(t, refcount(0), l.refcount)
}

// The same holds after a failed Init, which is the Shutdown a caller pairs
// with ERROR_LIBRARY_NOT_FOUND.
func TestShutdownAfterFailedInitIsUninitialized(t *testing.T) {
	l := newTestLibrary(&dynamicLibraryMock{
		OpenFunc: func() error {
			return errors.New("no such library")
		},
	})

	require.Equal(t, ERROR_LIBRARY_NOT_FOUND, l.Init())
	require.Equal(t, refcount(0), l.refcount)
	require.Equal(t, ERROR_UNINITIALIZED, l.Shutdown())
}
