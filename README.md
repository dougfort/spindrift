# spindrift

## a distributed standards system

From Wikipedia:

Spindrift usually refers to spray, particularly to the spray blown from cresting waves during a gale. This spray, which "drifts" in the direction of the gale, is one of the characteristics of a wind speed of 8 Beaufort and higher at sea.

In Great Britain "spindrift" is the telltale sign used by mariners to define a Force 8 (Beaufort Scale) wind (not higher) when observed at sea.

The above definition seems to vaguely relate to the design of this system. But the name actually comes from my favorite drink: http://spindriftfresh.com/

### purpose
This is my approach to a coding challenge proposed by the folks at [Koding, Inc.](http://www.koding.com/)

```
Let’s assume you have hundreds of thousands of servers under your command and
all servers must be within the following standards:

* Given some files, must be in defined paths
* Given some strings, must be in given path's content (eg: in a log file)
* Given some process names, must be running

Write an example of how you would meet these standards (see operation example
below). How to distribute these operations will be your choice of distribution
strategy. Your solution should come with scaling/distribution in-place.  
Eg: your solution should be able to send M requests from N machines to Y
destinations. M, N and Y respectively would be millions.

You should know which clients are online/sending results.
You should persist the responses coming from clients.

Operation example:
{
  "check_etc_hosts_has_4488": {
    "path": "/etc/hosts",
    "type": "file_contains",
    "check": "4.4.8.8"
  },
  "check_virus_file_exists": {
    "path": "/var/log/virus.txt",
    "type": "file_exists"
  }
}

This example is here only for demo purposes, and does not mean that the server
will only get these checks, we could send many different  checks in any unsorted
order.  
```
