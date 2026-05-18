*Baca dalam Bahasa Indonesia* | *[Read in English](#english-version)*

# 04-panic-recover - Panic & Recover

## Overview
Modul ini mendemonstrasikan bagaimana `panic` menghentikan alur eksekusi dan bagaimana `defer` + `recover` dapat digunakan untuk memulihkan sistem secara terkendali.

## File Structure & Responsibilities
- `panic_recover.go`: Entry point berisi `func main()` yang memicu `panic` melalui `accessSecureAsset()` dan menangkapnya menggunakan `recover` dalam fungsi `defer`.

## Key Concepts Learned
- `panic` untuk kondisi fatal
- `defer` untuk menjalankan cleanup/recovery saat fungsi keluar
- `recover` untuk menangkap panic dan menjaga aplikasi tetap stabil

## How to Run
- `go run ./04-panic-recover/panic_recover.go`

---

## English Version

# 04-panic-recover - Panic & Recover

### Overview
This module demonstrates how `panic` halts execution flow and how `defer` + `recover` can be used to recover the system in a controlled manner.

### File Structure & Responsibilities
- `panic_recover.go`: The entry point containing `func main()` that triggers a `panic` via `accessSecureAsset()` and catches it using `recover` in a `defer` function.

### Key Concepts Learned
- `panic` for fatal conditions
- `defer` for running cleanup/recovery when a function exits
- `recover` to catch panics and keep the application stable

### How to Run
- `go run ./04-panic-recover/panic_recover.go`
