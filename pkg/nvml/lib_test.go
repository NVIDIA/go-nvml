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

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLookupFromDefault(t *testing.T) {
	errClose := errors.New("close error")
	errOpen := errors.New("open error")
	errLookup := errors.New("lookup error")

	testCases := []struct {
		description          string
		library              dynamicLibrary
		skipLoadLibrary      bool
		expectedLoadError    error
		expectedLookupErrror error
		expectedCloseError   error
	}{
		{
			description:          "library not loaded yields error",
			library:              &dynamicLibraryMock{},
			skipLoadLibrary:      true,
			expectedLookupErrror: errLibraryNotLoaded,
		},
		{
			description: "open error is returned",
			library: &dynamicLibraryMock{
				OpenFunc: func() error {
					return errOpen
				},
			},

			expectedLoadError:    errOpen,
			expectedLookupErrror: errLibraryNotLoaded,
		},
		{
			description: "lookup error is returned",
			library: &dynamicLibraryMock{
				OpenFunc: func() error {
					return nil
				},
				LookupFunc: func(s string) error {
					return fmt.Errorf("%w: %s", errLookup, s)
				},
				CloseFunc: func() error {
					return nil
				},
			},

			expectedLookupErrror: errLookup,
		},
		{
			description: "lookup succeeds",
			library: &dynamicLibraryMock{
				OpenFunc: func() error {
					return nil
				},
				LookupFunc: func(s string) error {
					return nil
				},
				CloseFunc: func() error {
					return nil
				},
			},
		},
		{
			description: "lookup succeeds",
			library: &dynamicLibraryMock{
				OpenFunc: func() error {
					return nil
				},
				LookupFunc: func(s string) error {
					return nil
				},
				CloseFunc: func() error {
					return nil
				},
			},
		},
		{
			description: "close error is returned",
			library: &dynamicLibraryMock{
				OpenFunc: func() error {
					return nil
				},
				LookupFunc: func(s string) error {
					return nil
				},
				CloseFunc: func() error {
					return errClose
				},
			},
			expectedCloseError: errClose,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			defer setNewDynamicLibraryDuringTest(tc.library)()
			defer resetLibrary()
			l := GetLibrary()
			if !tc.skipLoadLibrary {
				require.ErrorIs(t, libnvml.load(), tc.expectedLoadError)
			}
			require.ErrorIs(t, l.Lookup("symbol"), tc.expectedLookupErrror)
			require.ErrorIs(t, libnvml.close(), tc.expectedCloseError)
			if tc.expectedCloseError == nil {
				require.Nil(t, libnvml.dl)
			} else {
				require.Equal(t, tc.library, libnvml.dl)
			}
		})
	}
}

func setNewDynamicLibraryDuringTest(dl dynamicLibrary) func() {
	original := newDynamicLibrary
	newDynamicLibrary = func(string, int) dynamicLibrary {
		return dl
	}

	return func() {
		newDynamicLibrary = original
	}
}

func resetLibrary() {
	libnvml = library{
		path:  defaultNvmlLibraryName,
		flags: defaultNvmlLibraryLoadFlags,
	}
}
