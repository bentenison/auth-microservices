# Testing in golang
- Unit Test
- BenchMarks tets
- profiling
- escape analysis
# Some usefull commands
- go test -benchmem -run=none -bench ^BenchmarkTokenGeneration$ grpc-auth-app/auth-server/pkg/auth -memprofile p.out -gcflags="-m -m" -coverprofile c.out
- go test -run ^regex
- go run -run=none -bench ^regex <package_name> // run benchmark from a package
- go run -run=none -bench ^regex <package_name> -memprofile p.out // run benchmark with memory profile
- go run -run=none -bench ^regex <package_name> -memprofile p.out -gcflags="-m -m" // run test with escape analysis
- go run -run=none -bench ^regex <package_name> -coverprofile c.out // run test with coverage profile
- go tool pprof p.out // start memory profile analysis with pprof
- top  // list top allocating functions
- list <function_name> // list apecific fn with allocations
- go run -run=none -bench ^regex <package_name> -cpu 1 -benchtime 3s // run cpu profile with 1 cpu