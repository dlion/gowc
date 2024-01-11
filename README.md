## Gowc
Unix command line tool WC written in Go

## Author
Domenico Luciani   
https://domenicoluciani.com

## How to build and run this tool

`go build`   
`./gowc <flag> <file>`

For example `./gowc -l resources/test.txt`

## Default behavior - No flags

By default the tool will count `bytes`, `lines` and `words` of the specified file

```bash
❯ ./gowc resources/test.txt
   342190 58164 7145 resources/test.txt
```

## Optional flags available
* `-c`: count bytes
* `-l` count lines
* `-w` count words
* `-m` count characters -multibyte characters supported-

## Idea
This is the implementation of the [WC Coding Challenge](https://codingchallenges.fyi/challenges/challenge-wc) with Go, following those requirements step-by-step   
Any idea or feedback to make it better are always welcomed!