#!/bin/sh

CONNLOGPATH="/path/to/src/connlogger"
CONNLOG="$CONNLOGPATH/connlogger"
PORT=33389

$CONNLOG -port 33389 -logfile "$CONNLOGPATH/connlog"
