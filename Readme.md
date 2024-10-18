# Creating the same Golang API again and again

I wanted to test out the various different options for creating an API in Go. There were a lot of recommendations and I wanted to try them. This repo holds the source code for all such attempts.

## The project

The API implements a code breaking game called bulls and cows. The classic version of the game uses 4 digit numbers but our version uses random 4 alphabet words.

The api exposes the following endpoints.

### POST /game
/new-game starts a new game creates. It choses a random word and a unique id for the game. The game id and the word are stored in a map. 
- @Todo: All games are deleted after a configurable timeout period or when the limit on the number of games is reached. 

### POST /game/:id/guess
/game/:id/guess submits a new guess for the given game. It stores the guess and returns the response. If the word is an exact match it returns a flag denoting that the game is over.

### GET /game/:id
/game/:id returns the current details of the game. It returns the list of all the guesses that have been made so far.


## A list of tryout: Past and future
- Stdlib