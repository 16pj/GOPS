# Game of Pure Strategy Simulation (GOPS) in Go and CLJ

## *Golang*


Program that runs the GOPS game based on provided count of cards and runs. It compares two strategies, random selection and picking the sames as the deck card.

Run it as:

    go run ./main.go

Additional Options to add to simulation

    -games <num of game runs> 
    -cards <num of cards in deck> 
    -verbose <add this flag to print game logs>


## *Clojure*

The CLJ code can be run with lein run

    lein run

Additionally if true is added to the args, then verbosity is increased

    lein run true