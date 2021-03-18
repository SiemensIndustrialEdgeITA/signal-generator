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

## Info

This project is still under construction.
As a first step a simple connection and streaming of linear data to the broker will be possible as follows:

- 1 datapoint/sec
- simpleJson schema:
``json
{
  key: "",
  value: 0 
}
``
- broker: **ie-databus:1883**
- topic:  **/signal-generator/simplejson**  
- user/password: **simatic/simatic**  

PS: **user, password and topic** must exist on the **databus** in order to properly connect and publish the data 

## Direct usage

Can be launched as follows:
```bash
./signal-generator
```

## Build the app

Run:
```bash
./build.sh
```
Upload docker-compose.yml through the industrial-edge-publisher


## Precompiled

Find precompiled version of the app under 
```bash
./app
```
This can be directly imported inside your Industrial Edge Management
