.PHONY: all build run benchmark clean

# Variables
BIN_DIR := ./bin
BENCH_DIR := ./benchmarks
SERVERS := nethttp gin gorilla
PORTS  := 8700 8800 8900

# build all implementations
build:
	@mkdir -p $(BIN_DIR)
	@for server in $(SERVERS); do \
				echo "Building $$server... "; \
				go build -o $(BIN_DIR)/$$server ./$$server; \
	done

#RUN  all servers in background
run:
	@for i in 1 2 3; do \
					server=$$(echo $(SERVERS) | cut -d' ' -f$$i); \
					port=$$(echo $(PORTS) | cut -d' ' -f$$i); \
					echo "Starting server on :$$port"; \
					$(BIN_DIR)/$$server -port $$port & \
	done

# Benchmark with wrk
benchmark:
		@mkdir -p $(BENCH_DIR)/results
		@for i in 1 2 3; do \
						server=$$(echo $(SERVERS) | cut -d' ' -f$$i); \
						port=$$(echo $(PORTS) | cut -d' ' -f$$i); \
						echo "Benchmarking $$server (http://localhost:$$port/stocks/AAPL)..."; \
						wrk -t12 -c1000 -d30s http://localhost:$$port/stocks/AAPL > $(BENCH_DIR)/results/$$server-wrk.txt; \
		done
		@echo "Results saved to $(BENCH_DIR)/results/"

# Clean up
clean:
		@killall -9 nethttp gin gorilla 2>/dev/null || true
		@rm -rf $(BIN_DIR) $(BENCH_DIR)/results
		@echo "Cleaned up."

# Shortcut: build + run + benchmark
all: build run benchmark

