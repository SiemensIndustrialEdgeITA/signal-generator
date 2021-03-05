# Signal generator (WIP)

This is a simple lighweight Simatic Edge App written in go which generates customizable streams of data suitable for local testing.

It will allow to chose some generator function for the desired signal behavior, i.e.:
- constSignal
- linearSignal
- sineAignal
..

Additional components can be added to the signal, i.e.:
- noiseComponent
- periodicFaultComponent
..

Output sinks can be selected for the different apps among the following:
- simpleJson
- dataService 
- traceConnector 
- freqAnalyzer 
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
- topic:    /generator/simplejson
- user/password: from cli options


Can be launched as follows:
```
./signal-generator --user simatic --password simatic
```
PS: **user, password and topic** must exist in order to properly connect and publish the data 
default databus hostname will be ```ie-databus:1883




