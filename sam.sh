#!/bin/sh
./samdict $* | xclip -sel clipboard
xclip -o -sel clipboard
