[![Build Status](https://travis-ci.org/johnnye/replicator.svg?branch=master)](https://travis-ci.org/johnnye/replicator)

Replicator
=========

Replicator is a WIP peice of [negroni](https://github.com/codegangsta/negroni) middleware allowing you to replicate requests onto a different URL. 

Why?
-------
You might wat to send real traffic to a staging environment or a beta version allowing you to moitor development with real traffic. Replicator allows you to control the amount of traffic you send. 

## Configuration 
Replicator expects 3 arguments: 

* `newURL` this is the new base URL to send traffic to 
* `meh` meh mode is set to `true` by default, it doesn't make sure that these requests are sent. `false` will guarantee that requests are replicated.
* `percentage` integer from 0 to 100 the percentage of requests that are sent

## Running
Include it in your list of middleware, preferably near the top.