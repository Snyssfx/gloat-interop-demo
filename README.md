# Gloat Interop Demo
Demo of how we want to work with Gloat and Glojure for Go interop

## Rationale
There is a go module named `demo` with useful functions,
and the goal is to be able to call them from glojure's REPL,
so that glojure could be a good scripting shell for the task.

- Right now how Go community does that is they create test files or scripts and use `go test` or `go run` to see what the function returns.

- The problem is that tests (and Go semantics and static types in general) requires more boilerplate and state management, and glojure could help to experiment with data flow more rapidly here.

## Scenario

1. The user creates a directory near `demo` go module, in our example it's `gloat`.

2. Then in the `gloat` directory she launches `gloat --repl` and connects to it from the editor.

3. Then she could call the functions from `demo` module, for example `demo.Sum` or `demo.GetDummyJSON`.
   - In Go for working with local modules there are 2 tools: go workspaces and `replace` directive in `go.mod`. How could we work with local Go packages in Glojure?
     - If it's possible to edit `go.mod` in `gloat/gljrepl` it would be cool.

   - There could be Go version mismatch between gloat and `demo` project, for example now `demo` project requires Go 1.26.1 and Gloat runs Glojure with 1.24.0. There might be a problem with feature disparity.

