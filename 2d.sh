#!/usr/bin/env bash

mdfind -onlyin /Users/mtm/pdev 'kMDItemFSContentChangeDate >= $time.now(-172800) && kMDItemFSName == "*.go"' |
    xargs -d '\n' -a - rg -l Command |
    xargs -d '\n' -a - rg -l bloom
