install:
	GO114MODULE=on go mod tidy
build: install
	GO114MODULE=on go build -o simple-apm.bin .

run_server: build
	./simple-apm.bin \
	-run_mode=server

run_worker: build
	./simple-apm.bin \
	-run_mode=worker \
	-target_queue=${target_queue} \
	-batch_size=${batch_size} \
	-job_type=${job_type}

# Example command
# make run run_mode=server
