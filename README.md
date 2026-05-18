# learn-golang

> **A practical, module-based introduction to Go.**
> 
> **Pengenalan Go yang praktis dan berbasis modul.**

---

## Syllabus / Progress Tracker

| #   | Module (EN)                                             | Module (ID)                                   | Folders                                      | Status        |
|-----|---------------------------------------------------------|-----------------------------------------------|-----------------------------------------------|---------------|
| 01  | Basic Syntax & Variables                                | Sintaks Dasar & Variabel                      | `01-basic-syntax`, `01-variables-operators`   | ✅ Completed   |
| 02  | Control Flow (Conditionals)                             | Alur Kontrol (Pengkondisian)                  | `02-conditionals-if-else`                     | ✅ Completed   |
| 03  | Error Handling Foundations                              | Dasar Penanganan Error                        | `03-error-handling`                           | ✅ Completed   |
| 04  | Panic & Recover Mechanics                               | Mekanisme Panic & Recover                     | `04-panic-recover`                            | ✅ Completed   |
| 05  | Collective Data: Arrays & Advanced Slices               | Data Kolektif: Array & Slice Tingkat Lanjut   | `05-arrays-slices`                            | ✅ Completed   |
| 06  | Key-Value Data: Maps & Deletion                        | Data Key-Value: Map & Penghapusan             | `06-maps`                                     | ✅ Completed   |
| 07  | Advanced Blueprint: Structs, Pointer Methods & Interfaces| Cetak Biru Lanjutan: Struct, Method Pointer & Interface | `07-structs-methods`                          | ✅ Completed   |

---

## How to Run / Cara Menjalankan

**EN:**
Due to package-level refactoring (to resolve duplicate `main` functions), you must now run each module as a full package directory, not as a single file. Use the following command format:

```bash
go run ./07-structs-methods/
```

Replace `07-structs-methods` with the desired module folder.

**ID:**
Karena refaktorisasi tingkat paket (untuk mengatasi duplikasi fungsi `main`), Anda sekarang harus menjalankan setiap modul sebagai direktori paket lengkap, bukan sebagai file tunggal. Gunakan format perintah berikut:

```bash
go run ./07-structs-methods/
```

Ganti `07-structs-methods` dengan folder modul yang diinginkan.

