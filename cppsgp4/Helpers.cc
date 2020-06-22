#include "Helpers.h"
#include "Observer.h"
#include "Util.h"
#include "CoordTopocentric.h"
#include "TleException.h"
#include "CoordGeodetic.h"

#include <cmath>
#include <iostream>

std::vector<struct CoordGeodetic> GeneratePosList(
                                               SGP4& sgp4,
                                               const DateTime& start_time,
                                               const DateTime& end_time,
                                               const int time_step)
{
  std::vector<struct CoordGeodetic> pass_list;
  DateTime previous_time(start_time);
  DateTime current_time(start_time);

  while (current_time < end_time)
  {
    /*
     * calculate satellite position
     */
    Eci eci = sgp4.FindPosition(current_time);
    CoordGeodetic geo = eci.ToGeodetic();
    pass_list.push_back(geo);
  }

  return pass_list;
}

std::vector<struct PassDetails> GeneratePassList(
                                               const double lat,
                                               const double lng,
                                               const double alt,
                                               SGP4& sgp4,
                                               const DateTime& start_time,
                                               const DateTime& end_time,
                                               const int time_step)
{
  std::vector<struct PassDetails> pass_list;
  CoordGeodetic user_geo(lat, lng, alt);

  Observer obs(user_geo);

  EventHorizonDetails aos_event;
  EventHorizonDetails los_event;

  bool found_aos = false;

  DateTime previous_time(start_time);
  DateTime current_time(start_time);

  while (current_time < end_time)
  {
    bool end_of_pass = false;

    /*
     * calculate satellite position
     */
    Eci eci = sgp4.FindPosition(current_time);
    CoordTopocentric topo = obs.GetLookAngle(eci);

    if (!found_aos && topo.elevation > 0.0)
    {
      /*
       * aos hasn't occurred yet, but the satellite is now above horizon
       * this must have occurred within the last time_step
       */
      if (start_time == current_time)
      {
        /*
         * satellite was already above the horizon at the start,
         * so use the start time
         */
        aos_event.time = start_time;
        aos_event.azimuth = topo.azimuth;
      }
      else
      {
        /*
         * find the point at which the satellite crossed the horizon
         */
        aos_event = FindCrossingPoint(
                                     user_geo,
                                     sgp4,
                                     previous_time,
                                     current_time,
                                     true);
      }

      found_aos = true;
    }
    else if (found_aos && topo.elevation < 0.0)
    {
      found_aos = false;
      /*
       * end of pass, so move along more than time_step
       */
      end_of_pass = true;
      /*
       * already have the aos, but now the satellite is below the horizon,
       * so find the los
       */
      los_event = FindCrossingPoint(
                                   user_geo,
                                   sgp4,
                                   previous_time,
                                   current_time,
                                   false);

      struct PassDetails pd;
      pd.aos = aos_event.time;
      pd.los = los_event.time;
      pd.aos_azimuth = aos_event.azimuth;
      pd.los_azimuth = los_event.azimuth;
      pd.max_elevation = FindMaxElevation(
                                          user_geo,
                                          sgp4,
                                          aos_event.time,
                                          los_event.time);

      Eci eci = sgp4.FindPosition(aos_event.time);
      CoordTopocentric topo = obs.GetLookAngle(eci);
      pd.aos_range_rate = topo.range_rate;

      eci = sgp4.FindPosition(los_event.time);
      topo = obs.GetLookAngle(eci);
      pd.los_range_rate = topo.range_rate;

      pass_list.push_back(pd);
    }

    /*
     * save current time
     */
    previous_time = current_time;

    if (end_of_pass)
    {
      /*
       * at the end of the pass move the time along by 30mins
       */
      current_time = current_time + TimeSpan(0, 30, 0);
    }
    else
    {
      /*
       * move the time along by the time step value
       */
      current_time = current_time + TimeSpan(0, 0, time_step);
    }

    if (current_time > end_time)
    {
      /*
       * dont go past end time
       */
      current_time = end_time;
    }
  };

  if (found_aos)
  {
    /*
     * satellite still above horizon at end of search period, so use end
     * time as los
     */
    struct PassDetails pd;
    pd.aos = aos_event.time;
    pd.los = los_event.time;
    pd.aos_azimuth = aos_event.azimuth;
    pd.los_azimuth = los_event.azimuth;
    pd.max_elevation = FindMaxElevation(user_geo, sgp4, aos_event.time, los_event.time);

    Eci eci = sgp4.FindPosition(aos_event.time);
    CoordTopocentric topo = obs.GetLookAngle(eci);
    pd.aos_range_rate = topo.range_rate;

    eci = sgp4.FindPosition(los_event.time);
    topo = obs.GetLookAngle(eci);
    pd.los_range_rate = topo.range_rate;
    pass_list.push_back(pd);
  }

  return pass_list;
}


EventHorizonDetails FindCrossingPoint(
                           const CoordGeodetic& user_geo,
                           SGP4& sgp4,
                           const DateTime& initial_time1,
                           const DateTime& initial_time2,
                           bool finding_aos)
{
  Observer obs(user_geo);
  EventHorizonDetails result;
  bool running;
  int cnt;

  DateTime time1(initial_time1);
  DateTime time2(initial_time2);
  DateTime middle_time;

  running = true;
  cnt = 0;
  while (running && cnt++ < 16)
  {
    middle_time = time1.AddSeconds((time2 - time1).TotalSeconds() / 2.0);
    /*
     * calculate satellite position
     */
    Eci eci = sgp4.FindPosition(middle_time);
    CoordTopocentric topo = obs.GetLookAngle(eci);
    result.azimuth = topo.azimuth;

    if (topo.elevation > 0.0)
    {
      /*
       * satellite above horizon
       */
      if (finding_aos)
      {
        time2 = middle_time;
      }
      else
      {
        time1 = middle_time;
      }
    }
    else
    {
      if (finding_aos)
      {
        time1 = middle_time;
      }
      else
      {
        time2 = middle_time;
      }
    }

    if ((time2 - time1).TotalSeconds() < 1.0)
    {
      /*
       * two times are within a second, stop
       */
      running = false;
      /*
       * remove microseconds
       */
      int us = middle_time.Microsecond();
      middle_time = middle_time.AddMicroseconds(-us);
      /*
       * step back into the pass by 1 second
       */
      middle_time = middle_time.AddSeconds(finding_aos ? 1 : -1);
    }
  }

  /*
   * go back/forward 1second until below the horizon
   */
  running = true;
  cnt = 0;
  while (running && cnt++ < 6)
  {
    Eci eci = sgp4.FindPosition(middle_time);
    CoordTopocentric topo = obs.GetLookAngle(eci);
    result.azimuth = topo.azimuth;
    if (topo.elevation > 0)
    {
      middle_time = middle_time.AddSeconds(finding_aos ? -1 : 1);
    }
    else
    {
      running = false;
    }
  }
  result.time = middle_time;
  return result;
}


double FindMaxElevation(
                        const CoordGeodetic& user_geo,
                        SGP4& sgp4,
                        const DateTime& aos,
                        const DateTime& los)
{
  Observer obs(user_geo);

  bool running;

  double time_step = (los - aos).TotalSeconds() / 9.0;
  DateTime current_time(aos); //! current time
  DateTime time1(aos); //! start time of search period
  DateTime time2(los); //! end time of search period
  double max_elevation; //! max elevation

  running = true;

  do
  {
    running = true;
    max_elevation = -99999999999999.0;
    while (running && current_time < time2)
    {
      /*
       * find position
       */
      Eci eci = sgp4.FindPosition(current_time);
      CoordTopocentric topo = obs.GetLookAngle(eci);

      if (topo.elevation > max_elevation)
      {
        /*
         * still going up
         */
        max_elevation = topo.elevation;
        /*
         * move time along
         */
        current_time = current_time.AddSeconds(time_step);
        if (current_time > time2)
        {
          /*
           * dont go past end time
           */
          current_time = time2;
        }
      }
      else
      {
        /*
         * stop
         */
        running = false;
      }
    }

    /*
     * make start time to 2 time steps back
     */
    time1 = current_time.AddSeconds(-2.0 * time_step);
    /*
     * make end time to current time
     */
    time2 = current_time;
    /*
     * current time to start time
     */
    current_time = time1;
    /*
     * recalculate time step
     */
    time_step = (time2 - time1).TotalSeconds() / 9.0;
  }
  while (time_step > 1.0);

  return max_elevation;
}
