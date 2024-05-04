# go-demo

## Short Go Demo app

In this modified version, the /fibonacci endpoint accepts a query parameter n, which specifies the number of Fibonacci numbers to generate. It then calculates the Fibonacci sequence up to the nth number and returns it as a response.

### Building and running the App

* ```podman build -t my-golang-api .``` **(Or if using docker ```docker build -t my-golang-api .``` )**

* ```podman run -p 8080:8080 my-golang-api``` **(Or if using docker ```docker run -p 8080:8080 my-golang-api``` )**

### Testing

In the test file:

    We import the testing package, which provides support for writing tests.
    We define a test function TestFibonacci that verifies the behavior of the fibonacci function.
    Inside the test function, we define test cases with different input values (n) and their expected results.
    We iterate over the test cases, calling the fibonacci function with each input value and comparing the result with the expected value.
    If the result doesn't match the expected value, we use t.Errorf to report a test failure.

To run the tests, you can use the go test command in the terminal:

bash

```go test```

This command will run all test functions in any files ending with _test.go in the current directory and its subdirectories.

#### Functional test

To use this endpoint, you can make a GET request to <http://localhost:8080/fibonacci?n=(#)> where "n" is the number of Fibonacci numbers you want to generate.

For example, a request to ```http://localhost:8080/fibonacci?n=10``` would return the Fibonacci sequence up to the 10th number.(<http://localhost:8080/fibonacci?n=10>)

This should yeild ```Fibonacci sequence up to 10: [0 1 1 2 3 5 8 13 21 34]``` in the browser.
