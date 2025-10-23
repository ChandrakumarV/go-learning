                       --------------- GO Toolchain ----------------

> go get -u github.com/gorilla/mux              // install or update the library
> go env                                        // view environment variables (GOPATH, GOROOT, etc.)
> go build .                                    // compile the Go code in the current directory

                        ----------------- GO Modules ------------------

> go mod init github.com/<username>/<module_name>  
            // initialize a new module

> go mod verify                                 // verify dependencies against their checksums
> go mod tidy                                   // add missing and remove unused modules (use carefully)
> go mod why github.com/gorilla/mux             // explain why a module is needed
> go mod graph                                  // show the module dependency graph

> go mod edit -go=1.25.1                        // update the Go version in go.mod (note the '=')
> go mod vendor                                 // copy dependencies to the vendor/ folder
> go run -mod=vendor main.go                    // use vendor directory for dependencies

                        ----------------- GO List ------------------

> go list                                       // list packages in the current module
> go list all                                   // list all packages (including dependencies)
> go list -m all                                // list all modules (including dependencies)
> go list -m -versions github.com/gorilla/mux   // show all available versions of this module
