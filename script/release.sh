#!/bin/sh

if [ $# -ne 1 ]; then
  echo "usage: <version>"
  exit 1
fi

BRANCH=$(git branch | grep '*' | awk '{print $2}')
if [ $BRANCH != "master" ]; then
  echo "branch is not master"
  exit 1
fi

version=$1

if git tag | grep $version >/dev/null; then
  echo "ERROR: $version already exists in git-tag list"
  exit 1
fi

cat <<__TEXT__ > cmd/version.go
package cmd

var VERSION = "$version"
__TEXT__


git add cmd/version.go
git commit -m ":up: bump up $version"
git tag $version
git push origin
git push origin --tags
