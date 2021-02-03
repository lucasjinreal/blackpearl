#!/usr/bin/env bash

sudo chmod -R 777 *
if [ "$(uname)" == "Darwin" ]; then
    # Do something under Mac OS X platform
    sudo cp bin/blackpearl /usr/local/bin
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    # Do something under GNU/Linux platform
    echo "this is linux"
    sudo cp bin/blackpearl /usr/local/bin
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
    # Do something under 32 bits Windows NT platform
    echo "this is windows 32bit, we are provide binary file, build it."
    
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
    # Do something under 64 bits Windows NT platform
    echo "this is windows 64bit, we are provide binary file, build it."
fi

echo "You are just installed blackpearl! call `blackpearl` at anytime you need!"
