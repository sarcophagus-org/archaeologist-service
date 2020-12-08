#!/bin/sh
lsof -ti tcp:$1 | xargs kill