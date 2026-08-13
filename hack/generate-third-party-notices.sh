#!/usr/bin/env bash
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
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

set -euo pipefail

# LC_ALL=C on every sort and grep below: collation and case folding must not vary
# by locale (under tr_TR glibc will not fold I to i, so LICENSE stops matching).

OUTPUT="${OUTPUT:-THIRD_PARTY_NOTICES.md}"
LICENSES_DIR="${LICENSES_DIR:-.licenses-cache}"

# The whole module: go-nvml is consumed as a library, so pkg/, gen/ and examples/
# all reach a consumer's disk as part of the module zip.
PACKAGES=("./...")

# Fixed, so the document does not depend on where it was generated. The only
# build-constraint split in the module is linux vs !linux; verify_platform_coverage
# fails if a file appears for a GOOS this matrix does not include.
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

die() {
    printf 'ERROR: %s\n' "$1" >&2
    shift
    if (( $# > 0 )); then
        printf '%s\n' "$@" >&2
    fi
    exit 1
}

log() {
    printf '%s\n' "$*" >&2
}

# Licenses that are themselves Markdown close a fixed ``` fence early and invert
# every block after it, so open with one backtick more than the file's longest run.
fence_for() {
    local file="$1" longest width
    # -a: a license containing a NUL byte is otherwise treated as binary and
    # grep prints "Binary file ... matches" rather than the matches themselves.
    longest=$(LC_ALL=C grep -oaE '`+' "${file}" 2>/dev/null \
        | awk '{ if (length($0) > m) m = length($0) } END { print m+0 }')
    width=$(( longest + 1 ))
    (( width < 3 )) && width=3
    printf '%*s' "${width}" '' | tr ' ' '`'
}

check_prerequisites() {
    command -v go >/dev/null 2>&1 || die "go is not installed."

    if ./bin/go-licenses --help >/dev/null 2>&1; then
        GO_LICENSES="${PWD}/bin/go-licenses"
    elif command -v go-licenses >/dev/null 2>&1; then
        GO_LICENSES="$(command -v go-licenses)"
    else
        die "go-licenses is not installed, or ./bin/go-licenses cannot run on this platform." \
            "A copy built for another platform cannot be reused: delete ./bin/go-licenses and re-run."
    fi

    local f
    for f in go.mod go.sum; do
        [[ -f "${f}" ]] || die "${f} not found — run 'make third-party-notices' from the repo root."
    done

    LOCAL_MODULE=$(go list -m 2>/dev/null || true)
    [[ -n "${LOCAL_MODULE}" ]] || die "could not determine local module path via 'go list -m'."

    # The repository vendors, so go would default to -mod=vendor and go-licenses
    # would report each license at its vendored path in this repo at HEAD instead
    # of the upstream URL at its version. Reading from the module cache also keeps
    # a stale cache from rewriting go.mod behind the maintainer's back.
    export GOFLAGS="-mod=readonly"

    # CGO stays on because that is how this library is really built: every entry
    # point in pkg/nvml is behind 'import "C"', and 'go build ./...' with
    # CGO_ENABLED=0 fails outright. go-licenses resolves imports rather than
    # type-checking, so it reports the same dependency either way today, but that
    # holds only while no cgo-only file imports something of its own. No C
    # compiler is needed: go-licenses never compiles.
    export CGO_ENABLED=1
}

# Stands in for the released-platform cross-check the vendored repositories do
# against their image matrix; this module has no release matrix to compare with.
# A GOOS named by a file suffix or build constraint but absent from PLATFORMS
# means that file is never looked at and the inventory is quietly short.
verify_platform_coverage() {
    # Space separated, not newline separated: the macOS awk rejects a newline
    # inside a -v assignment.
    local known matrix uncovered
    known=$(go tool dist list | cut -d/ -f1 | LC_ALL=C sort -u | tr '\n' ' ')
    matrix=$(printf '%s\n' "${PLATFORMS[@]}" | cut -d/ -f1 | LC_ALL=C sort -u | tr '\n' ' ')

    # vendor/ is third-party source carrying build tags for platforms this project
    # never builds (google/uuid ships node_js.go); that is not a gap in the matrix.
    # shellcheck disable=SC2016  # $0 belongs to awk, not to the shell.
    uncovered=$(find . -name '*.go' -not -path './.git/*' -not -path './vendor/*' -not -path "./${LICENSES_DIR}/*" -print0 \
        | LC_ALL=C sort -z \
        | xargs -0 awk -v known="${known}" -v matrix="${matrix}" '
            BEGIN {
                split(known, k, " ");  for (i in k) is_goos[k[i]] = 1
                split(matrix, m, " "); for (i in m) covered[m[i]] = 1
            }
            FNR == 1 {
                name = FILENAME
                sub(/.*\//, "", name)
                sub(/\.go$/, "", name)
                n = split(name, part, "_")
                # foo_linux.go, foo_linux_amd64.go: the GOOS never comes first.
                for (i = 2; i <= n; i++) {
                    if (is_goos[part[i]] && !covered[part[i]]) found[part[i]] = 1
                }
            }
            /^\/\/go:build/ || /^\/\/[[:space:]]*\+build / {
                n = split($0, tok, /[^A-Za-z0-9]+/)
                for (i = 1; i <= n; i++) {
                    if (is_goos[tok[i]] && !covered[tok[i]]) found[tok[i]] = 1
                }
            }
            END { for (g in found) print g }
        ' \
        | LC_ALL=C sort -u)

    [[ -z "${uncovered}" ]] || die \
        "sources exist for GOOS values outside the PLATFORMS matrix: $(echo "${uncovered}" | paste -sd ' ' -)" \
        "Add them to the PLATFORMS array in hack/generate-third-party-notices.sh, or those files are never inspected."
}

prepare_workspace() {
    # Guard the override: '', '/', '.' or '..' would make the rm -rf fatal.
    case "${LICENSES_DIR}" in
        ""|"/"|"."|"..")
            die "refusing to 'rm -rf' unsafe LICENSES_DIR='${LICENSES_DIR}'."
            ;;
    esac
    rm -rf "${LICENSES_DIR}"
    mkdir -p "${LICENSES_DIR}"

    # Explicit templates: macOS mktemp ignores TMPDIR without one.
    local t="${TMPDIR:-/tmp}/go-nvml-notices"
    SAVE_ROOT="$(mktemp -d "${t}.XXXXXX")"
    COMBINED_CSV="$(mktemp "${t}-csv.XXXXXX")"
    MODULE_MAP="$(mktemp "${t}-mod.XXXXXX")"
    INDEX_FILE="$(mktemp "${t}-idx.XXXXXX")"

    # Composed next to OUTPUT, not in TMPDIR, so the publish below is a rename.
    local out_dir
    out_dir="$(dirname "${OUTPUT}")"
    mkdir -p "${out_dir}"
    OUT_TMP="$(mktemp "${out_dir}/.$(basename "${OUTPUT}").XXXXXX")"

    trap 'rm -rf "${SAVE_ROOT}"; rm -f "${COMBINED_CSV}" "${MODULE_MAP}" "${INDEX_FILE}" "${OUT_TMP}"' EXIT
}

# module path -> module@version, replacing the vendor/modules.txt lookup the
# vendored repositories use. Built per platform and unioned, because a dependency
# can be reachable on one GOOS only.
build_module_map() {
    log "Resolving module versions..."

    # Intermediates live under SAVE_ROOT so the EXIT trap already covers them.
    local platform goos goarch raw="${SAVE_ROOT}/modules.raw"

    # A replaced module is what actually gets built, so it is what the notices
    # file must name. Field order: original path, effective path, version.
    local template='{{with .Module}}{{if .Replace}}{{.Path}} {{.Replace.Path}} {{.Replace.Version}}{{else}}{{.Path}} {{.Path}} {{.Version}}{{end}}{{end}}'

    : > "${raw}"
    for platform in "${PLATFORMS[@]}"; do
        goos="${platform%/*}"
        goarch="${platform#*/}"
        GOOS="${goos}" GOARCH="${goarch}" go list -deps -f "${template}" "${PACKAGES[@]}" >> "${raw}" \
            || die "'go list -deps' failed for ${goos}/${goarch}." \
                "Module versions come from the module cache, not from the vendor directory." \
                "Re-run with network access or after 'go mod download'."
    done

    # awk, not grep: the module path is compared as a whole field, so its dots
    # cannot act as regex wildcards.
    LC_ALL=C sort -u "${raw}" \
        | awk -v local="${LOCAL_MODULE}" 'NF && $1 != local' \
        > "${MODULE_MAP}"

    [[ -s "${MODULE_MAP}" ]] \
        || die "no module versions resolved from 'go list -deps' — refusing to write an unattributed notices file."

    # An empty version means a filesystem replace, or a module the cache could
    # not describe. Either way the entry cannot be attributed.
    local bad
    bad=$(awk 'NF != 3 || $3 == "" { print $1 }' "${MODULE_MAP}")
    [[ -z "${bad}" ]] || die \
        "no version resolved for: $(echo "${bad}" | paste -sd ' ' -)" \
        "A local 'replace' directive has no version to report; teach hack/generate-third-party-notices.sh how to attribute it."

    awk '{ print $1, $2 }' "${MODULE_MAP}" > "${SAVE_ROOT}/modules.map"
    mv "${SAVE_ROOT}/modules.map" "${MODULE_MAP}"
}

collect_licenses() {
    local platform goos goarch save_dir

    for platform in "${PLATFORMS[@]}"; do
        goos="${platform%/*}"
        goarch="${platform#*/}"
        log "Collecting licenses for ${goos}/${goarch}..."

        save_dir="${SAVE_ROOT}/${goos}_${goarch}"

        # Only the local module: --ignore matches raw string prefixes, not path
        # segments, so a stdlib list adds the token "go" and silently drops
        # golang.org/x/*, google.golang.org/* and gopkg.in/*.
        GOOS="${goos}" GOARCH="${goarch}" "${GO_LICENSES}" save "${PACKAGES[@]}" \
            --save_path="${save_dir}" \
            --force \
            --ignore="${LOCAL_MODULE}"

        GOOS="${goos}" GOARCH="${goarch}" "${GO_LICENSES}" csv "${PACKAGES[@]}" \
            --ignore="${LOCAL_MODULE}" \
            >> "${COMBINED_CSV}"

        # Module cache files are 0444 and cp preserves that, so the next
        # platform's copy fails unless write permission is restored.
        cp -R "${save_dir}/." "${LICENSES_DIR}/"
        chmod -R u+w "${LICENSES_DIR}"
    done
}

# Licenses are joined, not picked: go-licenses emits a row per recognized license,
# so keeping one would hide a second license behind the first.
collapse_index() {
    LC_ALL=C sort -u "$1" | awk -F, '
        {
            pkg = $1
            if (!(pkg in url)) { url[pkg] = $2; order[++n] = pkg }
            if (!((pkg SUBSEP $3) in seen)) {
                seen[pkg SUBSEP $3] = 1
                # Count, do not test "pkg in lic": mawk instantiates the
                # assignment target before evaluating the right-hand side, so
                # that test is true on the first row and BWK awk disagrees.
                lic[pkg] = (cnt[pkg]++ ? lic[pkg] " / " : "") $3
            }
        }
        END { for (i = 1; i <= n; i++) print order[i] "," url[order[i]] "," lic[order[i]] }
    '
}

# go-licenses names a row after the directory that owns the license file, which
# can sit above or below the imported package, so match the longest module path
# that prefixes it.
annotate_modules() {
    awk -v modfile="${MODULE_MAP}" '
        BEGIN {
            FS = OFS = ","
            while ((getline line < modfile) > 0) {
                if (split(line, f, " ") < 2) continue
                mods[++m] = f[1]
                disp[f[1]] = f[2]
            }
            close(modfile)
            # A read error makes getline return -1 and the loop never runs.
            if (m == 0) {
                print "ERROR: no modules read from " modfile > "/dev/stderr"
                exit 1
            }
        }
        {
            best = ""
            for (i = 1; i <= m; i++) {
                mp = mods[i]
                if (($1 == mp || index($1, mp "/") == 1) && length(mp) > length(best)) best = mp
            }
            print $0, (best == "" ? "unknown" : disp[best])
        }
    '
}

build_index() {
    log "Generating dependency index..."
    collapse_index "${COMBINED_CSV}" | annotate_modules > "${INDEX_FILE}"

    [[ -s "${INDEX_FILE}" ]] \
        || die "go-licenses produced no entries for ${PACKAGES[*]} — refusing to write empty notices file."

    if cut -d, -f4 "${INDEX_FILE}" | LC_ALL=C grep -qE '^$|^unknown$'; then
        die "could not resolve module@version for some packages." \
            "Run 'go mod download' and re-run, rather than committing a file with unattributed entries."
    fi

    # An unclassifiable license is reported as "Unknown" with a zero exit. The
    # empty alternative matters because the renderer falls back to "Unknown" too,
    # and the " / " anchors catch a composite like "BSD-3-Clause / Unknown".
    if cut -d, -f3 "${INDEX_FILE}" | LC_ALL=C grep -qE '^$|(^| / )Unknown( / |$)'; then
        die "go-licenses could not classify the license of some packages." \
            "Inspect ${LICENSES_DIR} and attribute them by hand rather than shipping an Unknown."
    fi

    # go-licenses asks non-github hosts over the network (GET <path>?go-get=1) and
    # falls back to "Unknown" with a warning and a zero exit.
    if cut -d, -f2 "${INDEX_FILE}" | LC_ALL=C grep -qE '^$|^Unknown$'; then
        die "go-licenses could not resolve source URLs for some modules." \
            "This usually means the network blocked a '?go-get=1' lookup. Re-run with" \
            "access to the module hosts rather than committing a degraded file."
    fi
}

# Filter by name: for restricted licenses 'go-licenses save' copies the whole
# module source, which does not belong here.
license_files_for() {
    local dir="$1" f
    [[ -d "${dir}" ]] || return 0
    while IFS= read -r -d '' f; do
        if printf '%s' "$(basename "${f}")" \
            | LC_ALL=C grep -qiE '^(licen[cs]e|notice|copying|copyright|authors|patents)([-._].*)?$'; then
            printf '%s\n' "${f}"
        fi
    done < <(find "${dir}" -maxdepth 1 -type f -print0 2>/dev/null | LC_ALL=C sort -z)
}

emit_index_table() {
    local pkg url license module
    printf '| Package | License | Module | Source |\n'
    printf '|---------|---------|--------|--------|\n'

    while IFS=, read -r pkg url license module; do
        [[ -z "${pkg}" ]] && continue
        # go-licenses points at the licence file on the resolved ref. Drop that
        # suffix so the link names the project rather than one version of it.
        url="${url%%/blob/*}"
        # shellcheck disable=SC2016  # backticks are literal markdown here.
        printf '| `%s` | %s | `%s` | %s |\n' \
            "${pkg}" "${license:-Unknown}" "${module:-unknown}" "${url:-n/a}"
    done < "${INDEX_FILE}"
}

emit_sections() {
    local pkg url license module files lf fence

    while IFS=, read -r pkg url license module; do
        [[ -z "${pkg}" ]] && continue
        url="${url%%/blob/*}"

        printf '### %s\n\n' "${pkg}"
        printf '* License: %s\n' "${license:-Unknown}"
        printf '* Module: %s\n' "${module:-unknown}"
        printf '* Source: %s\n\n' "${url:-n/a}"

        files=()
        while IFS= read -r lf; do
            [[ -n "${lf}" ]] && files+=("${lf}")
        done < <(license_files_for "${LICENSES_DIR}/${pkg}")

        if (( ${#files[@]} == 0 )); then
            printf 'License text unavailable. See upstream source for the full license.\n'
        else
            for lf in "${files[@]}"; do
                fence="$(fence_for "${lf}")"
                printf '#### %s\n\n' "$(basename "${lf}")"
                printf '%stext\n' "${fence}"
                cat "${lf}"
                echo
                printf '%s\n' "${fence}"
                echo
            done
        fi
        echo
    done < "${INDEX_FILE}"
}

compose_document() {
    log "Composing ${OUTPUT}..."
    {
        cat <<'EOF'
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

EOF
        emit_index_table

        cat <<'EOF'

## License Texts

EOF
        emit_sections
    } > "${OUT_TMP}"

    # mv, not cp: OUT_TMP is in OUTPUT's directory, so this is a rename(2) and
    # OUTPUT is never a partial write. mktemp creates 0600, hence the chmod.
    chmod 644 "${OUT_TMP}"
    mv "${OUT_TMP}" "${OUTPUT}"
}

main() {
    check_prerequisites
    verify_platform_coverage
    prepare_workspace

    build_module_map
    collect_licenses
    build_index
    compose_document

    # Index rows are per package, not per module: one module can own several.
    local count
    count=$(wc -l < "${INDEX_FILE}" | tr -d ' ')
    log "Wrote ${OUTPUT} (${count} third-party Go packages)"
}

main "$@"
