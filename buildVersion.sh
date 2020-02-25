#!/bin/bash

cat <<EOF > cmd/main/buildVersion.go
package main

const buildVersion = "$(git describe)"
EOF