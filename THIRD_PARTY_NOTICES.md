# Third-Party Notices

NVIDIA go-nvml

This file covers the **Go dependencies** of go-nvml, with the verbatim text of
each one's license.

Scope is the whole module, `./...`. go-nvml is consumed as a library rather than
as a binary: what is distributed is the module itself, so `pkg/`, `gen/` and
`examples/` are all inventoried here alongside the library proper. The inventory
is the union across `linux/amd64`, `linux/arm64`, `darwin/amd64` and
`darwin/arm64`, because the module selects different files per GOOS.

Only non-test imports are followed, so test-only dependencies —
`github.com/stretchr/testify` and the modules it pulls in — are not listed.
They are vendored under `vendor/` and so travel in the module zip, but a
consumer never links them.

Go standard library packages are excluded; they are covered by the license of
the Go distribution itself.

This document inventories Go dependencies only. The C headers in the tree are
NVIDIA's own, not third-party content: `pkg/nvml/cgo_helpers.h` is under this
repository's `LICENSE`, and `gen/nvml/nvml.h` and `pkg/nvml/nvml.h` are under the
notice each carries. The NVML library these bindings load at runtime,
`libnvidia-ml.so.1`, is not distributed here — it comes from the NVIDIA driver
installation on the host and carries its own license.

## Dependency Index

| Package | License | Module | Source |
|---------|---------|--------|--------|
| `github.com/google/uuid` | BSD-3-Clause | `github.com/google/uuid` | https://github.com/google/uuid |

## License Texts

### github.com/google/uuid

* License: BSD-3-Clause
* Module: github.com/google/uuid
* Source: https://github.com/google/uuid

#### LICENSE

```text
Copyright (c) 2009,2014 Google Inc. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

```


