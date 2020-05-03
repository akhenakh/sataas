## SatAAS

Satellites As A Service

## API WIP

API is exported with gRPC

```proto
service Prediction {
    rpc SatInfos (SatRequest) returns (SatInfosResponse) {}
    rpc SatLocation (SatLocationRequest) returns (Location) {}
    rpc GenPasses(GenPassesRequest) returns (Passes) {}
}
```

## Example

Start the server
```sh
make sataas && ./cmd/sataas/sataas   
```

Getting International space station location & compute passes for the next 24h
```sh
make satcli && ./cmd/satcli/satcli -lat=46.83 -lng=-71.25 -noradNumber=25544 -duration=24h
2020/05/03 11:50:14 Sat norad_number:25544 name:"ISS (ZARYA)" tle1:"1 25544U 98067A   20124.53871913  .00001280  00000-0  31012-4 0  9998" tle2:"2 25544  51.6445 209.4500 0001204 226.6962 260.4054 15.49345471225099" update_time:<seconds:1588520809 nanos:14221000 > 
2020/05/03 11:50:14 Location latitude:51.31580710155846 longitude:-171.35883090195048 altitude:425.2472690866516 
2020/05/03 11:50:14 Passes to 2020-05-04 11:50:14.583622 -0400 EDT m=+86400.006103692:
"passes:<aos:<seconds:1588542559 > los:<seconds:1588542559 > aos_azimuth:193.6166342690838 los_azimuth:75.12442655371434 max_elevation:13.973425085394087 aos_range_rate:-5.9086092980506475 los_range_rate:5.934551649343886 > passes:<aos:<seconds:1588548290 > los:<seconds:1588548290 > aos_azimuth:238.63668816979606 los_azimuth:64.79444652669332 max_elevation:77.72179738878219 aos_range_rate:-6.898678801702919 los_range_rate:6.894572757238934 > passes:<aos:<seconds:1588554115 > los:<seconds:1588554115 > aos_azimuth:271.98516512139787 los_azimuth:70.15630435247343 max_elevation:37.74944455778578 aos_range_rate:-6.699542208763361 los_range_rate:6.696085832024824 > passes:<aos:<seconds:1588559948 > los:<seconds:1588559948 > aos_azimuth:291.2628290549357 los_azimuth:91.29589726355236 max_elevation:40.17197266804935 aos_range_rate:-6.723881101236595 los_range_rate:6.72763328794447 > passes:<aos:<seconds:1588565761 > los:<seconds:1588565761 > aos_azimuth:294.84959038621196 los_azimuth:126.07096318781281 max_elevation:64.11644195745436 aos_range_rate:-6.881086501296848 los_range_rate:6.88455836307821 > passes:<aos:<seconds:1588571598 > los:<seconds:1588571598 > aos_azimuth:282.11272970999846 los_azimuth:172.93577008949484 max_elevation:11.120639347177876 aos_range_rate:-5.614553944558927 los_range_rate:5.586085658736835 > "
```

## Tech

`cppsgp4` is swig wrapped from the [sgp4 c++ library](https://github.com/dnwrnr/sgp4).  
`sgp4` is a wrapper to handle c++ exceptions into Go errors, and a Goish API.

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