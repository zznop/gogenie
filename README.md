# gogenie ![build status](https://github.com/zznop/gogenie/actions/workflows/build.yml/badge.svg)

## Description

gogenie is a utility for making persistent patches to SEGA Genesis ROMs using Game Genie codes

## What is a Game Genie?

A Game Genie is a cheat system and catridge connector for SEGA Genesis and other consoles. It acts as a pass-thru and
sits between the console and the game cartridge. On boot, the Game Genie first presents its own ROM and prompts the
player to enter alphanumeric cheat codes. Each cheat code is nothing more than an encoded 32-bit address and 16-bit
value to apply at the address. After inputing the cheat codes and booting the game, the Game Genie intercepts CPU reads
for patched regions and returns the patched data as a response instead of the legitimate data from the game cart.

gogenie works by decoding Game Genie codes to recover the encoded address and value. It takes a clean ROM as input and
outputs a new ROM with the applied patch. It also re-computes and fixes up the patched ROM's checksum so it can boot
on real hardware.

## Usage

```
Usage: gogenie <code> <in> <outt>

Required:
  code    SEGA Genesis Game Genie alphanumeric code
  in      File path to input ROM
  out     File path to write out patched ROM

example:
  $ gogenie RFAA-A6VR StreetsOfRage.bin StreetsOfRagePatched.bin
```

## License

(C) 2021 Brandon Miller (zznop)

This code is licensed under MIT license (see LICENSE for details)
