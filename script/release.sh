#!/bin/sh

check_branch() {
  BRANCH=$(git branch | grep '*' | awk '{print $2}')
  if [ $BRANCH != "master" ]; then
    echo "branch is not master"
    exit 1
  fi
}

check_version_exists() {
  version=$1
  if git tag | grep $version >/dev/null; then
    echo "ERROR: $version already exists in git-tag list"
    exit 1
  fi
}

update_version_file() {
  cat <<__TEXT__ > cmd/version.go
package cmd

var VERSION = "$version"
__TEXT__
}


test_and_build_check() {
  make test
  make && make install
}

push_and_commit() {
  git add cmd/version.go
  git commit -m ":up: bump up $version"
  git tag $version
  git push origin
  git push origin --tags
}

if [ $# -ne 1 ]; then
  echo "usage: <version>"
  exit 1
fi

version=$1
check_branch
check_version_exists $version
update_version_file
test_and_build_check
push_and_commit

