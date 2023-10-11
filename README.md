ds2key-srv
==========

A server which listen on a port (9501 by default) for UDP packet sent from
[ds2key](https://code.google.com/p/ds2key/) and convert them to keystroke.

OSX only, you can find server for linux and windows on the [ds2key](https://code.google.com/p/ds2key/)
website.

### Why?
At the time of writing this, I wanted to use my Nintendo DS as a GamedPad on my MacBook.
I found ds2key but there was no version for OSX so I decided to build one in Go to better learn the language.

I have no idea if this still work in 2023 since it relies on Carbon which is pretty old now!

### TODO

- configuration file to personalize key bindings
