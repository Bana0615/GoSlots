# GoSlots

GoSlots is a simple slot machine game written in Go. It's a command-line based game where you can spin the slots and try your luck!

## Features

*   **Simple Gameplay:**  Easy-to-understand slot machine mechanics.
*   **Randomized Results:** Each spin generates a random outcome.
*   **Command-Line Interface:** Play directly in your terminal.

## How to Run

1.  **Prerequisites:**
    *   Make sure you have Go installed on your system. You can check by running `go version` in your terminal. If it's not installed, download and install it from [https://go.dev/dl/](https://go.dev/dl/).

2.  **Clone the Repository:**
```
bash
    git clone <repository_url>
    cd GoSlots

```
3. **Run**
```
bash
    go run .

```
4. **Gameplay Instructions**
* Follow the on-screen prompts to play the game.
* Starting Balance: You begin the game with a virtual balance of $200.
* Placing Bets: Each spin requires a bet. The standard cost is $1 per spin (ensure you have enough balance).
* Spinning the Reels: Use the spin mechanism to start the round. The three slot machine wheels will spin and stop randomly.
* Winning Condition: You win the round if all three symbols match along the center payline.
* Payouts: When you win, your $200 balance will be updated based on the specific winning combination (payout details might be listed elsewhere).
* Continuing Play: Keep spinning as long as you have sufficient balance.
* Quitting the Game: To exit the slot machine, simply set your bet amount to $0 when prompted.