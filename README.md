# ndsemu

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/rasky/ndsemu?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Status

This emulator is **HIGHLY EXPERIMENTAL**. Almost no games work for now. Please be patient :)

## How to compile

To compile, you must clone into a `ndsemu` subdirectory:

    git clone https://github.com/rasky/ndsemu $GOPATH/src/ndsemu
    cd $GOPATH/src/ndsemu
    go get
    go build

## BIOS

You need access to an official NDS BIOS and firmware. Put them within a "bios" subdirectory, like this:

    ndsemu
      |---bios
           |---- firmware.bin
           |---- biosnds7.rom
           |---- biodnds9.rom

## Run it

At this point, you can just run it with:

    ./ndsemu <path-to-your-rom-file>
    
    
