(ns cljgops.core
  (:gen-class))

(defn pickRandomCard [cards]
  (if (> (count cards) 0) (rand-nth cards)))

(defn selectRandomCard [cards]
  (let [randomCard (pickRandomCard cards)] 
    [randomCard (remove #(= % randomCard) cards)]))

(defn removeCard [cards card-to-remove] 
  (remove #(= % card-to-remove) cards))

(defn decision [deck-selection player1-sel player2-sel]
  (cond 
    (or (nil? player1-sel) (nil? player1-sel)) [0 0]
    (> player1-sel player2-sel) [deck-selection 0]
    (> player2-sel player1-sel) [0 deck-selection]
    :else [0 0]))

(defn scream-whats-happening [round deckSelection deckCards player1Selection player1Cards player2Selection player2Cards player1-score player2-score]
          (println "Round:" round)
          (println "Deck selection is:" deckSelection)
          (println "Deck cards left are:" deckCards)
          (println "P1 selection is:" player1Selection)
          (println "P1 cards left are:" player1Cards)
          (println "P2 selection is:" player2Selection)
          (println "P2 cards left are:" player2Cards)
          (println "Player1 score:" player1-score "and Player2 score:" player2-score))


(defn -main
  "Run the game with lein run"
  [& args]
(let [verbosity (= (first args) "true")]
  (loop  [round 0
        deckSelection 0
        player1Selection 0
        player2Selection 0
        player1-score 0
        player2-score 0
        deckCards [1 2 3 4 5 6 7 8]
        player1Cards [1 2 3 4 5 6 7 8]
        player2Cards [1 2 3 4 5 6 7 8]]
  (if (> round 8)
    (do (println "Player1 final score:" player1-score "and Player2 final score:" player2-score))
      (do
        (when (not= round 0)
          (when verbosity (scream-whats-happening round deckSelection deckCards player1Selection player1Cards player2Selection player2Cards player1-score player2-score)))
          (let [selectedCard (pickRandomCard deckCards)
            selectedCardByPlayer1 (pickRandomCard player1Cards)
            selectedCardByPlayer2 selectedCard
            [player1Score, player2Score] (decision selectedCard selectedCardByPlayer1 selectedCardByPlayer2)]
            (recur (+ round 1) 
              selectedCard
              selectedCardByPlayer1
              selectedCardByPlayer2
              (+ player1Score player1-score)
              (+ player2Score player2-score)
              (removeCard deckCards selectedCard)
              (removeCard player1Cards selectedCardByPlayer1)
              (removeCard player2Cards selectedCardByPlayer2))))))))