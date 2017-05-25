# tlc

Count lines per unit time.

## Installation

You need to set up Go and $GOPATH.

```sh
$ export PATH=$PATH:$GOPATH/bin
$ go get github.com/shrhdk/tlc
```

## Example

Prepare sample log file.

```sh
$ wget ftp://ita.ee.lbl.gov/traces/NASA_access_log_Aug95.gz
$ gzip -d NASA_access_log_Aug95.gz
$ head NASA_access_log_Aug95
in24.inetnebr.com - - [01/Aug/1995:00:00:01 -0400] "GET /shuttle/missions/sts-68/news/sts-68-mcc-05.txt HTTP/1.0" 200 1839
uplherc.upl.com - - [01/Aug/1995:00:00:07 -0400] "GET / HTTP/1.0" 304 0
uplherc.upl.com - - [01/Aug/1995:00:00:08 -0400] "GET /images/ksclogo-medium.gif HTTP/1.0" 304 0
uplherc.upl.com - - [01/Aug/1995:00:00:08 -0400] "GET /images/MOSAIC-logosmall.gif HTTP/1.0" 304 0
uplherc.upl.com - - [01/Aug/1995:00:00:08 -0400] "GET /images/USA-logosmall.gif HTTP/1.0" 304 0
ix-esc-ca2-07.ix.netcom.com - - [01/Aug/1995:00:00:09 -0400] "GET /images/launch-logo.gif HTTP/1.0" 200 1713
uplherc.upl.com - - [01/Aug/1995:00:00:10 -0400] "GET /images/WORLD-logosmall.gif HTTP/1.0" 304 0
slppp6.intermind.net - - [01/Aug/1995:00:00:10 -0400] "GET /history/skylab/skylab.html HTTP/1.0" 200 1687
piweba4y.prodigy.com - - [01/Aug/1995:00:00:10 -0400] "GET /images/launchmedium.gif HTTP/1.0" 200 11853
slppp6.intermind.net - - [01/Aug/1995:00:00:11 -0400] "GET /history/skylab/skylab-small.gif HTTP/1.0" 200 9202
```

Count lines per 2 days.

```sh
$ cat NASA_access_log_Aug95 | tlc -f 'dd/mmm/yyyy:HH:MM:SS' -p '48h'
1995-07-31 00:00:00 +0000 UTC   33997
1995-08-02 00:00:00 +0000 UTC   41388
1995-08-04 00:00:00 +0000 UTC   91450
1995-08-06 00:00:00 +0000 UTC   89782
1995-08-08 00:00:00 +0000 UTC   120615
1995-08-10 00:00:00 +0000 UTC   122494
1995-08-12 00:00:00 +0000 UTC   74551
1995-08-14 00:00:00 +0000 UTC   118725
1995-08-16 00:00:00 +0000 UTC   115641
1995-08-18 00:00:00 +0000 UTC   88340
1995-08-20 00:00:00 +0000 UTC   88503
1995-08-22 00:00:00 +0000 UTC   115859
1995-08-24 00:00:00 +0000 UTC   109873
1995-08-26 00:00:00 +0000 UTC   64431
1995-08-28 00:00:00 +0000 UTC   123484
1995-08-30 00:00:00 +0000 UTC   170765
```

Count lines per 10 days. (with grep)

```sh
$ cat NASA_access_log_Aug95 | grep '^piweba1y\.prodigy\.com' | tlc -f 'dd/mmm/yyyy:HH:MM:SS' -p '120h'
1995-08-01 00:00:00 +0000 UTC   190
1995-08-06 00:00:00 +0000 UTC   391
1995-08-11 00:00:00 +0000 UTC   348
1995-08-16 00:00:00 +0000 UTC   523
1995-08-21 00:00:00 +0000 UTC   555
1995-08-26 00:00:00 +0000 UTC   877
1995-08-31 00:00:00 +0000 UTC   73
```
