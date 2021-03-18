# Signal generator (WIP)

This is a simple lighweight Simatic Edge App written in go which generates customizable streams of data suitable for local testing.

It will allow to chose some generator function for the desired signal behavior, i.e.:
- linearSignal
- constSignal (wip)
- sineSignal (wip)
..

Additional components can be added to the signal, i.e.:
- noiseComponent
- periodicFaultComponent (wip)
..

Output sinks can be selected for the different apps among the following:
- simpleJson
- dataService (wip)
- traceConnector (wip)
- freqAnalyzer (wip)
..


## Usage

This project is still under construction.
As a first step a simple connection and streaming of constant data to the broker will be possible as follows:

- 1 datapoint/sec
- simpleJson schema:
``json
{
  key: "",
  value: 0 
}
``
- topic:  **/signal-generator/simplejson**  
- user/password: **simatic/simatic**  


Can be launched as follows:
```
./signal-generator
```
PS: **user, password and topic** must exist in order to properly connect and publish the data 
default databus hostname will be ```ie-databus:1883




