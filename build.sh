#!/bin/bash -e

COMMAND="${1}"

function fail() {
	exit 1
}

function build() {
	go build
}

function test() {
	go fmt ./...
	go test ./...
}

function unknown() {
	echo "Unknown command '${COMMAND}'." >&2
}

function update() {
	GOPROXY=direct go get -u ./...
	go mod tidy
}

function usage() {
	echo "$0: [command]"
	echo
	echo "Commands:"
	echo -e "\tbuild"
	echo -e "\ttest"
	echo -e "\tupdate"
	echo
}

case "${COMMAND}" in
	"")
		test
		;;
	build)
		build
		;;
	test)
		test
		;;
	update)
		update
		;;
	*)
		unknown
		echo
		usage
		fail
		;;
esac
