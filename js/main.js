const {GenLocationsRequest} = require('./satsvc_pb.js');
const {PredictionClient} = require('./satsvc_grpc_web_pb.js');

const mapboxgl = require('mapbox-gl');

var map = new mapboxgl.Map({
        container: 'map',
        style: 'https://map.dev.inair.space/osm-liberty-gl.style',
        center: [48.8, 2.2],
        zoom: 4,
        maxZoom: 15,
        minZoom: 2,
        transformRequest: (url, resourceType)=> {
            if(resourceType === 'Tile') {
                return {
                    url: url,
                    headers: { 'X-Key': '999neunb12beafxxxp17'}
                }
            }
        }
    });

var sats = new Map();

var clocation = {
     "latitude": 48.8,
     "longitude": 2.2,
     "altitude":0
};

function updatePositions() {
    var date = new Date();
    const devicebody = document.getElementById('sat_body');
    //devicebody.innerHTML = '';

    for (let [key, p] of sats) {
        var latlng = p.sat.latlng(date);
        // var observe = new Orb.Observation({"observer":clocation, "target":p.sat});
        // var horizontal = observe.azel(date); // horizontal coordinates(azimuth, elevation)
        // if (horizontal.elevation > 0) {
        //     console.log("visible", p.sat, horizontal.elevation);
        // }
        p.marker.setLngLat(new mapboxgl.LngLat(latlng.longitude, latlng.latitude));
    }
}

map.on('load', function () {
    var predictionClient = new PredictionClient('http://pouf.lan.inair.space:9200');
    var request = new GenLocationsRequest();
    request.setCategory(10);

    predictionClient.genLocations(request, {}, function(err, response) {
            console.log("genLocations", response, err);
    });
});

map.on('click', function(e) {
    console.log(e);
    new mapboxgl.Popup()
        .setLngLat(e.lngLat)
        .setHTML(e.title)
        .addTo(map);
});

// Add zoom and rotation controls to the map.
map.addControl(new mapboxgl.NavigationControl());
// Add geolocate control to the map.
map.addControl(new mapboxgl.GeolocateControl({
    positionOptions: {
        enableHighAccuracy: false
    },
    trackUserLocation: false
}));
