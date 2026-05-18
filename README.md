# learn-golang

> **A practical, module-based introduction to Go.**
> 
> **一个模块化、实用导向的 Go 语言入门教程。**

---

## Syllabus / Progress Tracker

| #   | Module (EN)                                             | Module (ZH)                                   | Folders                                      | Status        |
|-----|---------------------------------------------------------|-----------------------------------------------|-----------------------------------------------|---------------|
| 01  | Basic Syntax & Variables                                | 基础语法与变量                                | `01-basic-syntax`, `01-variables-operators`   | ✅ Completed   |
| 02  | Control Flow (Conditionals)                             | 流程控制（条件语句）                          | `02-conditionals-if-else`                     | ✅ Completed   |
| 03  | Error Handling Foundations                              | 错误处理基础                                  | `03-error-handling`                           | ✅ Completed   |
| 04  | Panic & Recover Mechanics                               | Panic 与 Recover 机制                        | `04-panic-recover`                            | ✅ Completed   |
| 05  | Collective Data: Arrays & Advanced Slices               | 集合数据：数组与高级切片                      | `05-arrays-slices`                            | ✅ Completed   |
| 06  | Key-Value Data: Maps & Deletion                        | 键值数据：Map 与删除操作                      | `06-maps`                                     | ✅ Completed   |
| 07  | Advanced Blueprint: Structs, Pointer Methods & Interfaces| 高级蓝图：结构体、指针方法与接口              | `07-structs-methods`                          | ✅ Completed   |

---

## How to Run / 运行方式

**EN:**
Due to package-level refactoring (to resolve duplicate `main` functions), you must now run each module as a full package directory, not as a single file. Use the following command format:

```bash
go run ./07-structs-methods/
```

Replace `07-structs-methods` with the desired module folder.

**ZH:**
由于已重构为包级目录（为解决重复的 `main` 函数），请以包目录方式运行模块，而非单文件。例如：

```bash
go run ./07-structs-methods/
```

将 `07-structs-methods` 替换为你要运行的模块目录。
