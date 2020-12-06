(ns cljgops.core-test
  (:require [clojure.test :refer :all]
            [cljgops.core :refer :all]))

(deftest remove-card-test
  "Removal of card from list"
  (let [card 19
        cards [1 2 3 4 5 19 11]]
    (is (removeCard cards card) [1 2 3 4 5 11]))
  (let [card 22
        cards [1 2 3 4 5 19 11]]
    (is (removeCard cards card) cards)))

(deftest pick-random-card-test
  "Select a random card from list"
  (let [cards [1 2 3 4 5 19 11]]
    (is (contains? cards (pickRandomCard cards))))
  (let [cards []]
    (is (nil? (pickRandomCard cards)))))

(deftest decision-test
  "Test to cover the decision making"
  (let [deck-card 7
        p1-card 2
        p2-card 3]
    (is (decision deck-card p1-card p2-card) [0 7]))
  (let [deck-card 2
        p1-card 2
        p2-card 1]
    (is (decision deck-card p1-card p2-card) [2 0])))