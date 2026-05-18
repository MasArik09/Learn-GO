*Baca dalam Bahasa Indonesia* | *[Read in English](#english-version)*

# 06-maps - Map & Manajemen Key

## Overview
Modul ini mempraktikkan struktur data map di Go, mulai dari lookup aman dengan idiom `value, ok` hingga iterasi `range` dan penghapusan key secara dinamis menggunakan `delete()`.

## File Structure & Responsibilities
- `maps_basic.go`: Entry point berisi `func main()` yang membuat map dengan `make`, menambahkan pasangan key-value, melakukan lookup aman (comma-ok), lalu memanggil `RunMapsAdvancedDemo()`.
- `maps_advanced.go`: Berisi `RunMapsAdvancedDemo()` dan helper `printSessions()` untuk demonstrasi map literal, iterasi terformat, penghapusan key yang ada, serta penghapusan key yang tidak ada (idempotent).

## Key Concepts Learned
- Inisialisasi map: `make` vs map literal
- Lookup aman dengan idiom `value, ok`
- Iterasi map menggunakan `for key, value := range ...`
- `delete()` bersifat idempotent pada key yang tidak ada

## How to Run
- `go run ./06-maps/`

---

## English Version

# 06-maps - Map & Key Management

### Overview
This module practices the map data structure in Go, from safe lookups using the `value, ok` idiom to `range` iteration and dynamic key deletion using `delete()`.

### File Structure & Responsibilities
- `maps_basic.go`: The entry point containing `func main()` that creates a map with `make`, adds key-value pairs, performs safe lookups (comma-ok), and calls `RunMapsAdvancedDemo()`.
- `maps_advanced.go`: Contains `RunMapsAdvancedDemo()` and the `printSessions()` helper to demonstrate map literals, formatted iteration, deletion of existing keys, and deletion of non-existent keys (idempotent).

### Key Concepts Learned
- Map initialization: `make` vs map literal
- Safe lookup with the `value, ok` idiom
- Map iteration using `for key, value := range ...`
- `delete()` is idempotent for non-existent keys

### How to Run
- `go run ./06-maps/`
