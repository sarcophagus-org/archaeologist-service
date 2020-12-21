#!/bin/sh
pgrep -f $1 | xargs kill
#pid=$(ps | grep $1 | grep node | cut -d' ' -f1); kill -TERM $pid || kill -KILL $pid