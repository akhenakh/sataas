## SatAAS

Satellites As A Service

## API WIP

API exported via gRPC.

```proto
service Prediction {
    // Categories returns all known categories with norad numbers
    rpc Categories (google.protobuf.Empty) returns (CategoriesResponse) {}

    // SatsInfos returns norad number and tles for active sats
    rpc SatsInfos (SatsRequest) returns (SatsInfosResponse) {}

    // SatsLocations returns a stream of locations for a list of sats or a category
    rpc SatsLocations (SatsRequest) returns (stream SatsLocationsResponse) {}

    // SatLocationFromObs returns a stream of observations from a location for a list of sats or a category
    rpc SatLocationFromObs (SatLocationFromObsRequest) returns (stream Observation) {}

    // GenLocations returns the predicted locations for a list of sats or a category
    rpc GenLocations(GenLocationsRequest) returns (GenLocationsResponse) {}

    // GenPasses returns the predicted passes for a sat
    rpc GenPasses(GenPassesRequest) returns (Passes) {}
}
```

## Example

Start the server
```sh
make sataas && ./cmd/sataas/sataas   
```

## Example tools provided

- Getting International space station location & compute passes for the next 2h & get a live stream of the position
```sh
 ./cmd/satcli/satcli -lat=46.83 -lng=-71.25 -noradNumber=25544 -duration=2h
2020/05/03 22:18:35 Sat norad_number:25544 name:"ISS (ZARYA)" tle1:"1 25544U 98067A   20124.73419492  .00000752  00000-0  21557-4 0  9996" tle2:"2 25544  51.6446 208.4853 0001149 231.1370 266.9719 15.49344194225126" update_time:<seconds:1588558517 nanos:377783000 > 
2020/05/03 22:18:35 Location latitude:-1.023585595854026 longitude:-51.195516207358374 altitude:419.98553195727436 
2020/05/03 22:18:35 Passes to 2020-05-04 00:18:35.698505 -0400 EDT m=+7200.005970072:
"passes:<aos:<seconds:1588559950 > los:<seconds:1588559950 > aos_azimuth:291.29683639401554 los_azimuth:91.28806140765137 max_elevation:40.168053684422446 aos_range_rate:-6.72311220577045 los_range_rate:6.727437896026004 > passes:<aos:<seconds:1588565761 > los:<seconds:1588565761 > aos_azimuth:294.8517672977992 los_azimuth:91.28806140765137 max_elevation:-5.729577951308175e+15 aos_range_rate:-6.8810715601595875 los_range_rate:6.727437896026004 > "
2020/05/03 22:18:36 Latitude : -0.94 Longitude : -51.13 Altitude: 420.0km
Azimuth : 154 Elevation -21.6 Range: 5654.2km RangeRage: -3.859895
2020/05/03 22:18:37 Latitude : -0.89 Longitude : -51.10 Altitude: 420.0km
Azimuth : 154 Elevation -21.6 Range: 5650.4km RangeRage: -3.856864
^C

```

- Get all Starlink satellites positions
```sh
 ./cmd/starlink/starlink 
2020/06/15 22:19:18 ID: 44724 Latitude : -23.21 Longitude : 119.74 Altitude: 556.5km
2020/06/15 22:19:18 ID: 44723 Latitude : 9.01 Longitude : -72.19 Altitude: 550.0km
2020/06/15 22:19:18 ID: 44722 Latitude : 52.66 Longitude : 0.14 Altitude: 555.4km
...
^C
```

## Tech

`cppsgp4` is swig wrapped from the [sgp4 c++ library](https://github.com/dnwrnr/sgp4).  
`sgp4` is a wrapper to handle c++ exceptions into Go errors, and a Goish API.

The gRPC is exposed as pure gRPC and as grpc-web on the same port.

## Update bindings

```
swig -c++ -intgosize 64 -go SGP4.i
```

## Self Promo

Check my iOS app [SatSat](https://satsat.inair.space/), it's free and offers Satellites passes predictions.

## TLE data

https://digitalarsenal.io/data/all.txt  
https://celestrak.com/NORAD/elements/active.txt  
http://www.amsat.org/amsat/ftp/keps/current/nasabare.txt

Beacons: http://www.ne.jp/asahi/hamradio/je9pel/satslist.csv

https://www.ucsusa.org/resources/satellite-database

https://www.wmo-sat.info/oscar/satellitefrequencies

https://db.satnogs.org/api/

## On going
https://gis.stackexchange.com/questions/77651/how-to-find-ring-of-coverage-of-gps-satellite-on-wgs-84-ellipsoid

https://rogchap.com/2019/07/26/in-process-grpc-web-proxy/