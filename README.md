# Discord Message Counter

A simple Go program to count how many messages you've sent in Discord, using the data package you can request from Discord.

## What is this?

This tool scans your exported Discord data and counts the number of messages you have sent across all channels. It works by reading the `messages.json` files inside each channel folder in your Discord data package.

## What does it do?

- Counts every message you've sent in every channel.
- Returns the total number of messages and the number of channels you've sent messages in.

## How does it work?

1. You request your data from Discord (see below).
2. Place the extracted `package` folder in this project's directory.
3. Run the Go program.
4. The program scans each channel folder (folders starting with `c`) and counts the messages in each `messages.json` file.

## How to use

### 1. Request your Discord data

- Go to **User Settings** > **Privacy & Safety** > **Request all of my data**.
- Wait for Discord to email you a download link (this may take a while).
- Download and extract the ZIP file from Discord.
- Inside, you'll find a `package/Messages` folder with many folders named like `cXXXXXXXXXXXXXXX`.

### 2. Set up and run the program

1. **Install Go** (if you don't have it):  
   ```sh
   brew install go
   ```
   Or download from [https://golang.org/dl/](https://golang.org/dl/).

2. **Clone or download this repository**.

3. **Place your `package` folder** (from Discord) in the project directory, so the structure looks like:
   ```
   Message Counter/
   ├── src/
   │   └── main.go
   ├── package/
   │   └── Messages/
   │       └── cXXXXXXXXXXXXXXX/
   │           └── messages.json
   └── go.mod
   ```

4. **Run the program**:
   ```sh
   go run src/main.go
   ```

5. **See your results!**
   ```
   You have sent 1234 messages across 56 channels
   ```

## Notes

- Only your own messages are counted.
- The program ignores files that are not `messages.json`.
- No data is uploaded or shared; everything runs locally.
# message-counter
