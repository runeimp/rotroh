

@_default:
	cls
	just --list


# Build the app
build:
	go build -o rotroh cmd/rotroh/*


# Build and install the app
install:
	cd cmd/rotroh; go install


# Run the code
run +args='':
	cls
	go run cmd/rotroh/main.go {{args}}

