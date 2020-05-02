## SatAAS

Satellites As A Service

## API WIP

API is exported with gRPC

```proto
service Prediction {
    rpc SatInfos (SatRequest) returns (SatInfosResponse) {}
    rpc SatLocation (SatLocationRequest) returns (Location) {}
}
```

## Example

Getting International space station location
```sh
 ./cmd/satcli/satcli -noradNumber=25544 
2020/05/01 21:29:57 Sat norad_number:25544 name:"ISS (ZARYA)" tle1:"1 25544U 98067A   20122.80026326  .00001880  00000-0  41787-4 0  9992" tle2:"2 25544  51.6449 218.0536 0001301 215.9785 288.2587 15.49341642224825" 
2020/05/01 21:29:57 Location latitude:9.6322477287762 longitude:146.1607791576601 altitude:418.0752680819114 
```

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
