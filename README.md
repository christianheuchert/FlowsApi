# FlowsApi
Backend for AiristaFlows

Inspired by Flogo

### FLow Logic
flows.json runs the trigger and takes in input and gives out in the format dicated by the settings. Settings are set on AiristaFlows angular front end. 
The output from the trigger is then set to each function. 
Funtions are run in order. The input comes the previous piece and the output is passed to the next function. 