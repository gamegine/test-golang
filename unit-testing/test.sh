# https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package

echo there are no test files in the root package, only the main.go file. You will get "no test files." 
go test -v
echo The solution is to test all packages within the current working directory, recursively
go test -v ./...
echo "Or you can specify which package (directory) to test"
go test -v ./math

echo add coverage test
go test -v ./math -coverprofile=coverage.out
echo convert to html coverage test
go tool cover -html=coverage.out -o=coverage.html.out

echo benchmark
go test ./math -bench=Min