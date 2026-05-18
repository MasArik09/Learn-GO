*Baca dalam Bahasa Indonesia* | *[Read in English](#english-version)*

# 05-arrays-slices - Array & Slice

## Overview
Modul ini membahas perbedaan arsitektur array (ukuran tetap) dan slice (dinamis), termasuk strategi alokasi memori, pertumbuhan kapasitas, dan risiko kebocoran perubahan karena slice berbagi underlying array.

## File Structure & Responsibilities
- `arrays_slices.go`: Entry point berisi `func main()` yang mendemonstrasikan array fixed-size, slice dinamis dengan `make`/`append`, observasi `len`/`cap`, validasi boundary, lalu memanggil `RunAdvancedSlicesDemo()`.
- `advanced_slices.go`: Berisi `RunAdvancedSlicesDemo()` untuk menunjukkan shared underlying array pada subslice, efek modification leakage, serta deep copy menggunakan `make` + `copy` untuk isolasi memori.

## Key Concepts Learned
- Perbedaan array vs slice (fixed-size vs dynamic)
- `len` dan `cap`, serta perilaku re-alokasi saat `append`
- Subslice berbagi underlying array (reference semantics)
- Deep copy dengan `copy()` untuk isolasi data

## How to Run
- `go run ./05-arrays-slices/`

---

## English Version

# 05-arrays-slices - Array & Slice

### Overview
This module discusses the architectural differences between arrays (fixed size) and slices (dynamic), including memory allocation strategies, capacity growth, and the risk of modification leakage due to slices sharing the underlying array.

### File Structure & Responsibilities
- `arrays_slices.go`: The entry point containing `func main()` that demonstrates fixed-size arrays, dynamic slices with `make`/`append`, observing `len`/`cap`, boundary validation, and calls `RunAdvancedSlicesDemo()`.
- `advanced_slices.go`: Contains `RunAdvancedSlicesDemo()` to show shared underlying arrays in subslices, modification leakage effects, and deep copy using `make` + `copy` for memory isolation.

### Key Concepts Learned
- Differences between arrays and slices (fixed-size vs dynamic)
- `len` and `cap`, and re-allocation behavior when using `append`
- Subslice sharing underlying array (reference semantics)
- Deep copy with `copy()` for data isolation

### How to Run
- `go run ./05-arrays-slices/`
