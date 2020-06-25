const {GenLocationsRequest} = require('./satsvc_pb.js');
const {PredictionClient} = require('./satsvc_grpc_web_pb.js');

var predictionClient = new PredictionClient('http://localhost:8081');
var request = new GenLocationsRequest();
request.setCategory(10);

predictionClient.genLocations(request, {}, function(err, response) {
    console.log(response, err);
});
