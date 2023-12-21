# Simple File Server

A _Go_ language library for a simple file server.

## Why ?

A file server built into the standard library of the _Go_ language is so poor 
that it does not even provide any way of configuring itself. 

## Functionality

* This library provides a simple file server which uses caching, i.e. it does 
not kill your HDD (Hard Disk Drive) if you ask for the same file several 
thousand times in a second. If for some reason you need to disable cache, the 
library provides such a functionality.


* This file server does not look for an index file when you have not asked for 
it intentionally. So, you do not have to put thousands of `index.html` files 
all over your storage in order to disable file listing. 


* It is possible to configure a list of default files for a directory. When a 
directory is requested, the server searches the list for a default file and 
returns the first existing file found. To disable this feature, the list of 
default files should be empty. To request a folder, the file path should end 
with a slash symbol.
