## SatAAS

Satellites As A Service


## Tech

`cppsgp4` is swig wrapped from the [sgp4 c++ library](https://github.com/dnwrnr/sgp4).  
`sgp4` is a wrapper to handle c++ exceptions into Go errors, and a Goish API.

## Update bindings

```
swig -c++ -intgosize 64 -go SGP4.i
```

## TLE data

https://digitalarsenal.io/data/all.txt  
https://celestrak.com/NORAD/elements/active.txt  
http://www.amsat.org/amsat/ftp/keps/current/nasabare.txt
