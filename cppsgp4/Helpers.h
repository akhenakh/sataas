#include "DateTime.h"
#include "SGP4.h"
#include <vector>


namespace libsgp4
{
  
struct EventHorizonDetails {
  DateTime time;
  double azimuth;
};

struct PassDetails {
  DateTime aos;
  DateTime los;
  double aos_azimuth;
  double los_azimuth;
  double max_elevation;
  double aos_range_rate;
  double los_range_rate;
};

std::vector<struct CoordGeodetic> GeneratePosList(
                                               SGP4& sgp4,
                                               const DateTime& start_time,
                                               const DateTime& end_time,
                                               const int time_step);

std::vector<struct PassDetails> GeneratePassList(
                                               const double lat,
                                               const double lng,
                                               const double alt,
                                               SGP4& sgp4,
                                               const DateTime& start_time,
                                               const DateTime& end_time,
                                               const int time_step);

EventHorizonDetails FindCrossingPoint(
                           const CoordGeodetic& user_geo,
                           SGP4& sgp4,
                           const DateTime& initial_time1,
                           const DateTime& initial_time2,
                           bool finding_aos);

double FindMaxElevation(
                        const CoordGeodetic& user_geo,
                        SGP4& sgp4,
                        const DateTime& aos,
                        const DateTime& los);

};
