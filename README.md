# 🚀 Benchmarking HTTP Routers in Go

A performance comparison of Go HTTP routers/frameworks (`net/http`, `Gin`, and `Gorilla/mux`) with caching optimizations.

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 📌 Overview

This project benchmarks three popular Go HTTP routing approaches:
1. Standard library `net/http`
2. High-performance `Gin` framework
3. Feature-rich `Gorilla/mux`


## 📊 Key Results

### Baseline Performance (No Caching)
| Framework    | Requests/sec | P99 Latency |
|--------------|-------------|------------|
| `net/http`   | 105,483     | 25ms       |
| `Gin`        | 102,341     | 26ms       |
| `Gorilla`    | 98,765      | 29ms       |

### With Caching (`sync.Map`)
| Framework    | Requests/sec | Improvement |
|--------------|-------------|------------|
| `net/http`   | 850,000     | 8x         |
| `Gin`        | 820,000     | 8x         |

## 🛠️ Setup

### Prerequisites
- Go 1.21+
- `gorilla/mux` (for distributed caching tests)
- `wrk` load testing tool (`go install github.com/tsliwowicz/go-wrk@latest`)

### Installation
```bash
git clone https://github.com/KingBean4903/BenchHTTPRouters
cd BenchHTTPRouters

### Running Tests
`make all`

.
├── nethttp/          # net/http implementation
├── gin/              # Gin framework implementation
├── gorilla/          # Gorilla/mux implementation
├── benchmarks/       # Test scripts and results
├── Makefile          # Build/test automation
