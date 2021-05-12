#!/bin/bash

T_API='YOUR_TELEGRAM_API'
CHAT_ID='YOUR_CHAT_ID'
BINARYFILE="server-watcher"

function printASCIIArt(){
    echo '''
    █▀ █▀▀ █▀█ █░█ █▀▀ █▀█   █░█░█ ▄▀█ ▀█▀ █▀▀ █░█ █▀▀ █▀█
    ▄█ ██▄ █▀▄ ▀▄▀ ██▄ █▀▄   ▀▄▀▄▀ █▀█ ░█░ █▄▄ █▀█ ██▄ █▀▄
    '''
}

function installBinary(){
    echo "Downloading Binary!!"
    wget -q 'https://github.com/stefins/server-watcher/releases/download/v1.0.0/server-watcher_1.0.0_Linux_x86_64.tar.gz'
    tar -xvf server-watcher_1.0.0_Linux_x86_64.tar.gz
    rm -rf server-watcher_1.0.0_Linux_x86_64.tar.gz
}

function fileCheck(){
    if test -f "$BINARYFILE"; then
        echo "$BINARYFILE Exists!!"
    else
        installBinary
    fi
}

function runBinary(){
    export T_API
    export CHAT_ID
    nohup "./$BINARYFILE" &
}

function killAllServerWatcher(){
    killall "$BINARYFILE"
}

printASCIIArt
killAllServerWatcher
fileCheck
runBinary