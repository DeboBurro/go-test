# go-test
## Summary
A project to understand some parts of golang's philosophy
---
---
## Prereqisite
1. Install Golang [link](https://go.dev/doc/install)

---
## Run
1. `go run .` under the repo root folder

---
## Review

### Pros
1. Static datatype that we can get information when compiling.
2. Nice interface to spin up Goroutine.
3. Feel easy to understand how to write and run go programs.
4. Pointer and reference.
5. In some parts it smells cleaner
### Cons
1. Feel something weird about `fmt` library, like `Scanf` will pipe string back to stdin which might cause issues.
2. Controversial error handling. I feel good to be able to explicitly know the error message but I need to bypass the errors by several if or switch statements to handle each case which feels a little bit too much.
---
## TODO
1. Kinda wondering the performance on CPU, RAM, compile time, process time between c++ and go.
