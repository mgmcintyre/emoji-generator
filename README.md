# emoji-generator
Generate random emojis, because 🐼

## Usage
```
// Installation
$ go get -u github.com/mgmcintyre/emoji-generator

// Write a random emoji to ~/.emoji every 1 second
$ emoji run -s ./emoji.html -d ~/.emoji -i 1s
```

## How it works
+ Reads the source file
+ Parses out unicode character references & stores them in memory
+ Runs a ticker based on the provided duration
+ On each tick, selects a random emoji and writes it to the destination file

## Flags
+ `--source|-s` the source (html) file to parse for emojis
+ `--destination|-d` the destination file to write to
+ `--interval|-i` the interval to tick on

## Dependencies
Uses the fantastic github.com/urfave/cli library.
There's a copy of this HTML page in the repo: http://unicode.org/emoji/charts/emoji-list.html

## Contributing
Yes, please do.
