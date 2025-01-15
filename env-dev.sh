#!/bin/bash

# sessions names

session="api"

# start new sessions with names

tmux new-session -d -s $session

tmux rename-window -t 1 'nvim'
tmux send-keys -t 'nvim' 'nvim .' C-m

tmux new-window -t $session:2 -n 'server'
tmux send-keys -t 'server' 'go run main.go' C-m

tmux attach-session -t $session:1
