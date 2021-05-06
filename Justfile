

@_default:
	just _term-wipe
	just --list


# Build the app
build:
	go build -o rotroh cmd/rotroh/*


# Build and install the app
install:
	cd cmd/rotroh; go install


# Run the code
run +args='':
	just _term-wipe
	go run cmd/rotroh/main.go {{args}}


# Run unit tests
@test:
	#!/bin/sh
	just _term-wipe
	echo
	echo '$ go test -coverprofile=c.out'
	go test -coverprofile=c.out
	echo
	echo '$ go tool cover -func=c.out'
	go tool cover -func=c.out
	echo
	rm c.out


_term-wipe:
	#!/bin/sh
	if [[ ${#VISUAL_STUDIO_CODE} -gt 0 ]]; then
		clear
	elif [[ ${KITTY_WINDOW_ID} -gt 0 ]] || [[ ${#TMUX} -gt 0 ]] || [[ "${TERM_PROGRAM}" = 'vscode' ]]; then
		printf '\033c'
	elif [[ "$(uname)" == 'Darwin' ]] || [[ "${TERM_PROGRAM}" = 'Apple_Terminal' ]] || [[ "${TERM_PROGRAM}" = 'iTerm.app' ]]; then
		osascript -e 'tell application "System Events" to keystroke "k" using command down'
	elif [[ -x "$(which tput)" ]]; then
		tput reset
	elif [[ -x "$(which reset)" ]]; then
		reset
	else
		clear
	fi