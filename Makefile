MAIN_FILE := ./cli/cmd.go
BUILD_DIR := ./build/

run: $(MAIN_FILE)
	@ go run $(MAIN_FILE)

