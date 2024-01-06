#!/usr/bin/env bash

mdfind "kMDItemFSContentChangeDate >= \$time.now(-172800) && kMDItemFSName == '*.go'" -onlyin /Users/mtm/pdev |
    xargs -d '\n' -a - rg -l Command |
    xargs -d '\n' -a - rg -l bloom
