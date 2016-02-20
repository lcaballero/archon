#!/bin/bash

function tests() {
	go test --race $(go list ./... | grep -v vendor)
}

function benches() {
	go test -bench=. $(go list ./... | grep -v vendor)
}

$1
