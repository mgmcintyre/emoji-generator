# emoji-generator
Generate random emojis, because 🐼

## Usage
```
// Installation
% go get -u github.com/mgmcintyre/emoji-generator

// Write a random emoji to ~/.emoji every 1 second
% emoji-generator run -s ./emoji.html -d ~/.emoji -i 1s
```

## How it works
+ Reads the source file
+ Parses out unicode character references & stores them in memory
+ Runs a ticker based on the provided duration
+ On each tick, selects a random emoji and writes it to the destination file

## Flags
+ `--source` the source (html) file to parse for emojis
+ `--destination` the destination file to write to
+ `--interval` the interval to tick on

## Dependencies
Uses the fantastic [github.com/urfave/cli]() library.

There's a copy of this HTML page in the repo: http://unicode.org/emoji/charts/emoji-list.html

## Want emojis in your prompt?
Use the provided .plist to launch emoji-generator at startup
```
% cp uk.co.mgmcintyre.emoji-generator ~/Library/LaunchAgents/
```

Read the destination file into your prompt (e.g. from my ~/.bash_profile)
**bash**
```
% PS1='[\t]\[\e[0m\]$(cat ~/.emoji) \[\e[36m\][\w]\[\e[0m\]\$ '
                    ^ where emojis come from
```
**zsh**
```
% PROMPT='%K{red}%F{white} %* %f%k%K{white}%S  $(cat ~/.emoji)  %s%k%K{blue} %~ %k %# '
                                               ^ where emojis come from
```

## Contributing
Yes, please do.
