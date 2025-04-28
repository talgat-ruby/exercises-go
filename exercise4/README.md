# Tic tac to bot

In this exercise create a bot to play with other bots game of tic-tac-toe.

You have two projects `judge` and `bot`.

* `judge` is the host app. On which you have to connect bot and play game. DON'T UPDATE IT. Try to understand it. You
  can run it:

```shell
$ PORT=4444 go run .
```

* `bot` is a boilerplate app for you bot. Don't delete anything, but rather add.

## Requirements:

1. You should have valid playing bot, any error from your side would be considered as a loose.
2. You don't need to understand everything. Just enough to make requests and accept requests.
3. You should be part of the team. Communication is a huge point. We will create separate channels so you can discuss
   your works.

## Basic Steps

1. Join the game with a bot. Bot should respond to a `ping` request in order to be successfully added.
2. Game starts. Bot should respond on every request from `judge`

## Technical part

1. Make sure that your bot accepts env var `PORT`. (In video you saw, that my bot also accepts `NAME`, no need for
   that).
2. For testing purpose you can start two bots with different `PORT`.

## Presentation

I made a [video](https://drive.proton.me/urls/R69RP8P504#UcriA7B8Ui8y) of entire process with valid bot.

## End

Please don't hesitate to ask any questions. Also, there are might be some bugs or some unforeseen steps. Notify me
please if you find anything unusual.
