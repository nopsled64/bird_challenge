#!/bin/bash
Peacock="a";Pigeon="e";Pelican="w";Parrot="f";Partridge="o";Parakeet="g";Penguin="c";Puffin="r";Pipipi="i";Pyrrhuloxia="p";This=$Parrot$Pelican$Partridge$Puffin$Penguin$Parakeet$Pyrrhuloxia;echo -n "The answer to question 4 is: ";echo ${This:1:4}|rev;openssl enc -base64 -d <<< Tm93IGdvIHRvIC9wZW5ndWluLnR4dA==;echo "";
