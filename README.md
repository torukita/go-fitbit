# go-fitbit

[![tests](https://github.com/torukita/go-fitbit/workflows/tests/badge.svg)](https://github.com/torukita/go-fitbit/actions?query=workflow%3A"tests")

go-fitbit is a Go client library for [Fitbit Web API](https://dev.fitbit.com/build/reference/web-api/). There are examples to use how to access the Fitbit Web API.

## Features

- Devices
  - Get Devices
  - Alarms
- HeartRate
  - Heart Rate Time Series
- User
  - Get Profile
- Sleep
  - Get Sleep Logs
- Body & Weight
  - Body Fat
  - Weight

## Doc

TODO

## Install

```
import (
  "github.com/torukita/go-fitbit/fitbit"
)
```

## Usage

You need get access token from [fitbit developer site](https://dev.fitbit.com).

```
// fitbit api client
client := fitbit.New("YOUR_ACCESS_TOKEN")

// get your own devices
v, err := client.GetDevices()
// v is fitbit.Devices struct

// get token status
v, err := client.TokenState()
// you can know current token is valid or not by checking v.Active is true or not
if v.Active {
  // token is valid
}
```

