%module cppsgp4
%{
#include "SGP4.h"
#include "Observer.h"
#include "Eci.h"
#include "CoordGeodetic.h"
#include "TleException.h"
#include "CoordTopocentric.h"
#include "Helpers.h"
%}

%{
using namespace libsgp4;
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"
%include "exception.i"

%exception {
    try {
        $action;
    } catch (std::runtime_error &e) {
        _swig_gopanic(e.what());
    }
}

namespace std {
   %template(PassDetailsVector) vector<PassDetails>;
   %template(GeosVector) vector<CoordGeodetic>;
}


class SGP4 {
public:
    SGP4(const Tle& tle);
    Eci FindPosition(const DateTime& date) const;
};
class Tle {
public:
    Tle(const std::string& name, const std::string& line_one, const std::string& line_two);
    std::string Line1() const;
    std::string Line2() const;
    std::string Name() const;
    unsigned int NoradNumber() const;
};

class Observer {
public:
    Observer(const double latitude, const double longitude, const double altitude);
    CoordTopocentric GetLookAngle(const Eci &eci);
};


class Eci {
public:
    Eci(const DateTime& dt, const double latitude, const double longitude, const double altitude);
    CoordGeodetic ToGeodetic() const;
};

class DateTime {
public:
    DateTime(int year, int month, int day, int hour, int minute, int second);
    static DateTime Now(bool useMicroseconds = false);
    double ToJulian() const;
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

struct CoordGeodetic {
    /** latitude in radians (-PI >= latitude < PI) */
    double latitude;
    /** latitude in radians (-PI/2 >= latitude <= PI/2) */
    double longitude;
    /** altitude in kilometers */
    double altitude;
};

struct CoordTopocentric {
    double azimuth;
    double elevation;
    double range;
    double range_rate;
};

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
