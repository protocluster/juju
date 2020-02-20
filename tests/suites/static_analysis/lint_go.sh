run_dep_check() {
  OUT=$(dep check 2>&1 || true)
  if [ -n "${OUT}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "\\n${OUT}" >&2
    exit 1
  fi
}

run_go_vet() {
  OUT=$(go vet -composites=false ./... 2>&1 || true)
  if [ -n "${OUT}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "\\n${OUT}" >&2
    exit 1
  fi
}

run_go_lint() {
  OUT=$(golint -set_exit_status ./ 2>&1 || true)
  if [ -n "${OUT}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "\\n${OUT}" >&2
    exit 1
  fi
}

run_go_imports() {
  FOLDERS="${2}"
  OUT=$(echo "${FOLDERS}" | xargs -I % goimports -l % 2>&1 || true)
  LIST=$(echo "${OUT}" | grep -v "_mock.go" | grep -v "_mock_test.go" | xargs grep -L "MACHINE GENERATED BY")
  GOFILES=$(echo "${LIST}" | xargs head -q -n1 | grep -v "Code generated by" | tr -d '\n')
  if [ -n "${GOFILES}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "\\n${LIST}" >&2
    exit 1
  fi
}

run_deadcode() {
  OUT=$(deadcode ./ 2>&1 || true)
  if [ -n "${OUT}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "\\n${OUT}" >&2
    exit 1
  fi
}

run_misspell() {
  FILES=${2}
  OUT=$(misspell -source=go 2>/dev/null "${FILES}" || true)
  if [ -n "${OUT}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "${OUT}"
    exit 1
  fi
}

run_unconvert() {
  OUT=$(unconvert ./ 2>&1 || true)
  if [ -n "${OUT}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "\\n${OUT}" >&2
    exit 1
  fi
}

run_ineffassign() {
  OUT=$(ineffassign ./.. | grep -v "_test.go" | grep "github.com/juju/juju" | sed -E "s/^(.+src\\/github\\.com\\/juju\\/juju\\/)(.+)/\2/")
  if [ -n "${OUT}" ]; then
    echo ""
    echo "$(red 'Found some issues:')"
    echo "${OUT}"
    exit 1
  fi
}

run_go_fmt() {
  FILES=${2}
  OUT=$(echo "${FILES}" | xargs gofmt -l -s)
  if [ -n "${OUT}" ]; then
    OUT=$(echo "${OUT}" | sed "s/^/  /")
    echo ""
    echo "$(red 'Found some issues:')"
    for ITEM in ${OUT}; do
      echo "gofmt -s -w ${ITEM}"
    done
    exit 1
  fi
}

test_static_analysis_go() {
  if [ "$(skip 'test_static_analysis_go')" ]; then
      echo "==> TEST SKIPPED: static go analysis"
      return
  fi

  (
    set_verbosity

    cd .. || exit

    FILES=$(find ./* -name '*.go' -not -name '.#*' -not -name '*_mock.go' | grep -v vendor/ | grep -v acceptancetests/)
    FOLDERS=$(echo "${FILES}" | sed s/^\.//g | xargs dirname | awk -F "/" '{print $2}' | uniq | sort)

    ## Functions starting by empty line
    # turned off until we get approval of test suite
    # run "func vet"

    ## Check dependency is correct
    if which dep >/dev/null 2>&1; then
      run "run_dep_check"
    else
      echo "dep not found, dep static analysis disabled"
    fi

    ## go vet, if it exists
    if go help vet >/dev/null 2>&1; then
      run "run_go_vet" "${FOLDERS}"
    else
      echo "vet not found, vet static analysis disabled"
    fi

    ## golint
    if which golint >/dev/null 2>&1; then
      run "run_go_lint"
    else
      echo "golint not found, golint static analysis disabled"
    fi

    ## goimports
    if which goimports >/dev/null 2>&1; then
      run "run_go_imports" "${FOLDERS}"
    else
      echo "goimports not found, goimports static analysis disabled"
    fi

    ## deadcode
    if which deadcode >/dev/null 2>&1; then
      run "run_deadcode"
    else
      echo "deadcode not found, deadcode static analysis disabled"
    fi

    ## misspell
    if which misspell >/dev/null 2>&1; then
      run "run_misspell" "${FILES}"
    else
      echo "misspell not found, misspell static analysis disabled"
    fi

    ## unconvert
    if which unconvert >/dev/null 2>&1; then
      run "run_unconvert"
    else
      echo "unconvert not found, unconvert static analysis disabled"
    fi

    ## ineffassign
    if which ineffassign >/dev/null 2>&1; then
      run "run_ineffassign"
    else
      echo "ineffassign not found, ineffassign static analysis disabled"
    fi

    ## go fmt
    run "run_go_fmt" "${FILES}"
  )
}
