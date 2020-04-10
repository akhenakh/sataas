%module cppsgp4
%{
#include "SGP4.h"
#include "Observer.h"
#include "Eci.h"
#include "CoordGeodetic.h"
#include "TleException.h"
%}


%include <typemaps.i>
%include "std_string.i"
%include "exception.i"
%exception {
    try {
        $action;
    } catch (std::runtime_error &e) {
        _swig_gopanic(e.what());
    }
}

class SGP4 {
public:
    SGP4(const Tle& tle);
    Eci FindPosition(const DateTime& date) const;
};
class Tle {
public:
    Tle(const std::string& line_one, const std::string& line_two);
};

class Observer {
public:
    Observer(const double latitude, const double longitude, const double altitude);
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
};

struct CoordGeodetic {
    /** latitude in radians (-PI >= latitude < PI) */
    double latitude;
    /** latitude in radians (-PI/2 >= latitude <= PI/2) */
    double longitude;
    /** altitude in kilometers */
    double altitude;
};


