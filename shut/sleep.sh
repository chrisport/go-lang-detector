#!/usr/bin/env bash
delayInMin=$1

if [[ -z "$delayInMin" ]]; then
    delayInMin=6
fi

delayInSeconds=$(($delayInMin*60))
reminder=$(($delayInSeconds-60))

echo Sleeping in $delayInMin minutes...
(sleep $reminder; echo Sleeping in 1 minute...) &

sleep $delayInSeconds;
pmset sleepnow
