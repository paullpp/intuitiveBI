#!/bin/bash

DIR="$(pwd)"
FRONTEND_DIR="$DIR/src/frontend"
SERVER_DIR="$DIR/src/backend"

cd $SERVER_DIR
air &

cd $FRONTEND_DIR
bun run dev &
