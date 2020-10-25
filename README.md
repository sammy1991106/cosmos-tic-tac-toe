# cosmos-tic-tac-toe

**cosmos-tic-tac-toe** is a simple blockchain application implementing the game tic-tac-toe built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve -v
```

`serve` command installs dependencies, initializes and runs the application.

## Commands

```
cosmos-tic-tac-toecli tx cosmostictactoe create-game
cosmos-tic-tac-toecli tx cosmostictactoe join-game
cosmos-tic-tac-toecli tx cosmostictactoe create-game-move
cosmos-tic-tac-toecli tx cosmostictactoe challenge-game-timeout
```

## Configure

Initialization parameters of your app are stored in `config.yml`.
