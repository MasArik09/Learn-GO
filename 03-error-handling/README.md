*Baca dalam Bahasa Indonesia* | *[Read in English](#english-version)*

# 03-error-handling - Penanganan Error

## Overview
Modul ini memperkenalkan arsitektur penanganan error idiomatik di Go melalui pola return `(result, error)` dan guard clause. Modul ini juga membedakan sentinel error dan error dinamis untuk validasi yang lebih kuat.

## File Structure & Responsibilities
- `error_handling.go`: Entry point berisi `func main()` dan fungsi `divideNumbers()` untuk mendemonstrasikan pola `(value, error)` pada skenario sukses dan gagal; mengeksekusi demo custom error melalui `RunCustomErrorsDemo()`.
- `custom_errors.go`: Berisi sentinel error `ErrNegativeInput`, fungsi validasi `validateAge()`, serta `RunCustomErrorsDemo()` untuk demonstrasi `errors.Is` dan `fmt.Errorf`.

## Key Concepts Learned
- Pola `(value, error)` dan pengecekan `if err != nil`
- Guard clause untuk validasi input
- Sentinel error + `errors.Is`
- Error dinamis menggunakan `fmt.Errorf`

## How to Run
- `go run ./03-error-handling/`

---

## English Version

# 03-error-handling - Error Handling

### Overview
This module introduces idiomatic error handling architecture in Go through the `(result, error)` return pattern and guard clauses. It also distinguishes between sentinel errors and dynamic errors for stronger validation.

### File Structure & Responsibilities
- `error_handling.go`: The entry point containing `func main()` and the `divideNumbers()` function to demonstrate the `(value, error)` pattern in both success and failure scenarios; executes the custom error demo via `RunCustomErrorsDemo()`.
- `custom_errors.go`: Contains the sentinel error `ErrNegativeInput`, the `validateAge()` validation function, and `RunCustomErrorsDemo()` to demonstrate `errors.Is` and `fmt.Errorf`.

### Key Concepts Learned
- `(value, error)` pattern and `if err != nil` checking
- Guard clause for input validation
- Sentinel error + `errors.Is`
- Dynamic errors using `fmt.Errorf`

### How to Run
- `go run ./03-error-handling/`
