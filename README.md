# Tournament-Scoring

## Project Overview
This is the backend API designed to facilitate the organization and tracking of Elevate 4 Epilepsy basketball tournaments. It provides both GET routes and POST routes to view and manage each tournament division. 


## Usage
### GET Requests
#### GET `/api/home` returns a list of all the events as well as information about the tournament.
Example Request: `GET http://localhost:8000/api/home`

Example Response:
```
{
  "year": 2024,
  "events": [
    {
      "_id": "66988d7a98e653bd9fc78d84",
      "name": "3rd/4th Boys",
      "slug": "3rd-4th-boys",
      "time": "0001-01-01T00:00:00Z",
    },
    {
      "_id": "66988d7a98e653bd9fc78d85",
      "name": "5th/6th Boys",
      "slug": "5th-6th-boys",
      "time": "0001-01-01T00:00:00Z",
    },
    ...
  ]
}
```

#### GET `/api/{event_slug}/teams` returns a list of teams for the event.
Example Request: `GET http://localhost:8000/api/3rd-4th-boys/teams`

Example Response: 
```
{
  [
  {
    "_id": "66988da349c67a7b33acfef2",
    "name": "Team A",
    "event": "3rd-4th-boys"
  },
  {
    "_id": "66988da349c67a7b33acfef3",
    "name": "Team B",
    "event": "3rd-4th-boys"
  },
  ...
  ]
}
```

#### GET `/api/{event_slug}/pools` returns a map of pool rounds and games for the event. If pools haven't started, it will return an error.
Example Request: `GET http://localhost:8000/api/3rd-4th-boys/pools`

Example Response: 
```
{
  "1": [
    {
      "_id": "66988df949c67a7b33acfefa",
      "event": "3rd-4th-boys",
      "round": 1,
      "court": "A",
      "team1Id": "66988da349c67a7b33acfef2",
      "team2Id": "66988da349c67a7b33acfef4",
      "team1Name": "Team A",
      "team2Name": "Team C"
    },
    ...
  ],
  "2": [
    {
      "_id": "66988df949c67a7b33acfefe",
      "event": "3rd-4th-boys",
      "round": 2,
      "court": "A",
      "team1Id": "66988da349c67a7b33acfef4",
      "team2Id": "66988da349c67a7b33acfef9",
      "team1Name": "Team C",
      "team2Name": "Team H"
    },
    ...
  ]
}
```

#### GET `/api/{event_slug}/seeding` returns an ordered list of all the seeded teams for an event. If pools haven't finished, it will return an error.
Example Request: `GET http://localhost:8000/api/3rd-4th-boys/seeding`

Example Response: 
```
{
  "1": [
    {
      "_id": "66988df949c67a7b33acfefa",
      "event": "3rd-4th-boys",
      "round": 1,
      "court": "A",
      "team1Id": "66988da349c67a7b33acfef2",
      "team2Id": "66988da349c67a7b33acfef4",
      "team1Name": "Team A",
      "team2Name": "Team C"
    },
    ...
  ],
  "2": [
    {
      "_id": "66988df949c67a7b33acfefe",
      "event": "3rd-4th-boys",
      "round": 2,
      "court": "A",
      "team1Id": "66988da349c67a7b33acfef4",
      "team2Id": "66988da349c67a7b33acfef9",
      "team1Name": "Team C",
      "team2Name": "Team H"
    },
    ...
  ]
}
```

#### GET `/api/{event_slug}/bracket` returns the elimination bracket as a binary tree. If elimination hasn't started yet, it will return an error.
Example Request: `GET http://localhost:8000/api/3rd-4th-boys/bracket`

Example Response: 
```
{
  "event": "3rd-4th-boys",
  "courts": ["A", "B"],
  "rounds": 4,
  "root": {
    "teamId": "000000000000000000000000",
    "left": {
      "team": "Team A",
      "teamId": "669893052342d778f9b4358a",
      "seeding": 1,
    },
    "right": {
      "team": "Team B",
      "teamId": "669893052342d778f9b4358b",
      "seeding": 2,
    },
    "court": "A" 
  }
}
```

## Contributing
Brian Zhou
