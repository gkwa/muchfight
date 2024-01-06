#!/usr/bin/env bash

/usr/bin/mdfind "kMDItemFSContentChangeDate >= \$time.now(-172800) && kMDItemFSName == '*.go'" -onlyin /Users/mtm/pdev |
    /usr/local/opt/findutils/libexec/gnubin/xargs -d '\n' -a - rg -l Command |
    /usr/local/opt/findutils/libexec/gnubin/xargs -d '\n' -a - rg -l boom
