#!/bin/bash

function preprelease() {
  version=${VERSION:?"VERSION must be set"}

  sed -i.bak "s/SDKVersion = .*/SDKVersion = \"$version\"/" harness/version.go && rm harness/version.go.bak
}

$1
