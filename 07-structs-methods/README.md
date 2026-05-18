*Baca dalam Bahasa Indonesia* | *[Read in English](#english-version)*

# 07-structs-methods - Struct, Method, & Interface

## Overview
Modul ini memperkenalkan custom data type dengan `struct` dan prinsip enkapsulasi melalui field exported vs unexported. Selain itu, modul ini membangun arsitektur method receiver dan interface untuk polymorphism, dependency injection, serta inspeksi tipe yang aman.

## File Structure & Responsibilities
- `structs_basic.go`: Entry point berisi `func main()`, definisi `Student` (field publik dan `gpa` unexported), contoh instansiasi struct, serta orkestrasi pemanggilan seluruh demo.
- `structs_methods.go`: Berisi method `Introduce()` (value receiver) dan `UpdateGPAPointer()` (pointer receiver) serta `RunMethodsDemo()` untuk membandingkan copy-by-value vs mutation.
- `structs_constructor.go`: Berisi factory constructor `NewStudent()` dengan validasi (guard clause) serta `RunConstructorDemo()` untuk alur input invalid vs valid.
- `interfaces_basic.go`: Berisi interface `NotificationProvider`, implementasi `EmailProvider`/`SMSProvider` beserta `Send()`, dan `RunInterfacesBasicDemo()` untuk demonstrasi dynamic dispatch melalui slice interface.
- `interfaces_engine.go`: Berisi `NotificationEngine` (orchestrator) dengan `RegisterProvider()` dan `Broadcast()` serta `RunInterfacesEngineDemo()` untuk dependency injection dan broadcast terpusat.
- `interfaces_assertion.go`: Berisi `InspectAndReport()` untuk type assertion aman (comma-ok) dan type switch, serta `RunInterfacesAssertionDemo()` untuk verifikasi runtime.

## Key Concepts Learned
- Struct literal (named vs positional) dan enkapsulasi field (exported vs unexported)
- Value receiver vs pointer receiver
- Factory constructor + validasi guard clause (`(*Student, error)`)
- Implicit interface implementation dan polymorphism
- Dependency Injection melalui slice interface dan orchestrator
- Type assertion (comma-ok) dan type switch untuk inspeksi tipe aman

## How to Run
- `go run ./07-structs-methods/`

---

## English Version

# 07-structs-methods - Struct, Method, & Interface

### Overview
This module introduces custom data types with `struct` and the principle of encapsulation through exported vs unexported fields. It also builds method receiver and interface architecture for polymorphism, dependency injection, and safe type inspection.

### File Structure & Responsibilities
- `structs_basic.go`: The entry point containing `func main()`, the `Student` definition (public fields and unexported `gpa`), struct instantiation examples, and orchestration of all demo calls.
- `structs_methods.go`: Contains the `Introduce()` (value receiver) and `UpdateGPAPointer()` (pointer receiver) methods, as well as `RunMethodsDemo()` to compare copy-by-value vs mutation.
- `structs_constructor.go`: Contains the factory constructor `NewStudent()` with validation (guard clause) and `RunConstructorDemo()` for invalid vs valid input flows.
- `interfaces_basic.go`: Contains the `NotificationProvider` interface, `EmailProvider`/`SMSProvider` implementations with `Send()`, and `RunInterfacesBasicDemo()` to demonstrate dynamic dispatch via interface slices.
- `interfaces_engine.go`: Contains the `NotificationEngine` (orchestrator) with `RegisterProvider()` and `Broadcast()`, and `RunInterfacesEngineDemo()` for dependency injection and centralized broadcast.
- `interfaces_assertion.go`: Contains `InspectAndReport()` for safe type assertion (comma-ok) and type switch, and `RunInterfacesAssertionDemo()` for runtime verification.

### Key Concepts Learned
- Struct literal (named vs positional) and field encapsulation (exported vs unexported)
- Value receiver vs pointer receiver
- Factory constructor + guard clause validation (`(*Student, error)`)
- Implicit interface implementation and polymorphism
- Dependency Injection via interface slice and orchestrator
- Type assertion (comma-ok) and type switch for safe type inspection

### How to Run
- `go run ./07-structs-methods/`
