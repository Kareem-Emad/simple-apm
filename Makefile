install:
	GO114MODULE=on go mod tidy
build: install
	GO114MODULE=on go build -o new-new-relic.bin .

run: build
	./new-new-relic.bin \
	-run_mode=$(run_mode)


# Example command
# make run run_mode=server
