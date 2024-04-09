# dependency-injection-test

This project is a test for a dependency injection in go.

No much effort was put on it to make the code the cleanest nor the pretties. Moreover, it does not follow the best
practicies in the word regardig Golang.

## How it works:
The so called "injector", located at the injector.go file, registers all sort of interfaces in it.

Then, it is able to
receive "build functions" that are functions which use the implementation of those interfaces to build something and
identify which parameters each build function needs and pass them properly.

TODO:
- [ ] Enable the Injector to identify the name of the interfaces it is receiving.
- [ ] Allow the Injector to receive build functions that could return all kind of values in all sort of orders.
- [ ] Implement tests.
