# dependency-injection-test

This project is a test of the [dependecy injection](https://en.wikipedia.org/wiki/Dependency_injection) in golang.

This is more a POC, the name with the suffix -test was a mistake, please disconsider.

***Disclaimer***: No much effort was put on it to make the code the cleanest nor the pretties. Moreover, it does not
follow the best practicies in the word regardig Golang.

## How it works:
The so called "injector", located at the injector.go file, registers all sort of interfaces in it.

Then, it is able to receive "build functions" that are functions which use the implementation of those interfaces to
build something and identify which parameters each build function needs and pass them properly.

## Tests
### Boot time difference
It is not only important to look at how ease the code becomes with an improvement in it, but also the side effects of
using any approach to solve an issue. The dependecy injection tries to simplify the code and avoid too much boiler
plate, but in the other hand, it may affect the code performance at least in a slightly way.

### Tests performed so far
This projects has a small example inside the /example folder to help to understand how it works. This example also
measures the difference in time to initialize the workers it is initializig between the regular way (initializing and
passingn the dependences by hand) and the dependecy injection way.

It goes without saying that by using dependency injection which relies on reflection, naturaly the boot time is longer.
The problem come even more perceptible as the project size grows.

Usually the results for the exampes inside the /example folder are in between:
- The values being almost the same (with a small edge to the regular way).
- The depedenncy injection taking the double of time than its counterpart.

There must be some room to improve the code but there is no miracle, it all comes down to how necessery a dependency
injection is and how much it worth to have some extra seconds of minutes of initializing time.

## TODO:
- [ ] Enable the Injector to identify the name of the interfaces it is receiving (is is possible?).
- [ ] Allow the Injector to receive build functions that could return all kind of values in all sort of orders.
- [ ] Implement tests (why not?).
