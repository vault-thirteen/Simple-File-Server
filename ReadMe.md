# Simple File Server

A Go language library for a simple file server.

## Why ?
A file server built into the standard library of the Go language is so poor 
that it does not even provide any way of configuring it. 

## Functionality

* This library provides a simple file server which uses caching, i.e. it does not 
kill your HDD (Hard Disk Drive) if you ask for the same file several thousand 
times in a second. If for some reason you need to disable cache, the library 
provides such a functionality.


* This file server does not look for an index file when you have not asked for it 
intentionally. So, you do not have to put thousands of `index.html` files all 
over your storage in order to disable file listing. 
