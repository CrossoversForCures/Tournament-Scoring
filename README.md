# Tournament-Scoring

## Project Overview
This project is a Tournament Management System designed to facilitate the organization and tracking of Elevate 4 Epilepsy basketball tournaments. It caters to spectators and administrators, providing features to track game schedules and locations, score management, and tournament bracket visualization.

## Features
- For each division in the tournament (i.e. 5th-6th Boys, 7th-8th Girls, etc.), there will be a public-facing interface with:
  - game schedules
  - game scores
  - game locations
  - team standings
  - tournament brackets
- Administrator interface for score input and bracket management

## User Stories

### Site Visitor - Player
As a player, I want to view where my games are located, so I know where my next court is and when I should go there.

### Site Visitor - Spectator
As a spectator, I want to view the final scores of the pool play games, the standings after each pool play game, and the single elimination bracket, so I know how my team is doing.

### Administrator
1. As an administrator, I want to input the final scores of the pool play games, so I can ensure the proper seeding of teams in the single elimination bracket.
2. As an administrator, I want to input the final scores of the single elimination bracket games, so I can commemorate the winning team.

## API Requirements
GET `/api/events` returns a list of all the events as well as informatino about the tournament  
POST `/api/start` sorts all the teams in each event into pools  
GET `/api/events/:id/teams` returns list of teams for the event as JSON  
GET `/api/events/:id/pools` returns all pool information for the event as JSON  
POST `/api/events/:id/
POST `//teams` accepts a new team to be added  
GET `/api/game/:id` returns information about a game as JSON  
PUT `/api/game/:id` updates the results of a game  

The backend will also create a new tournament in the database after the current one finishes and generate the pool and elimination bracket based on standings.


## Usage
[Instructions on how to use the application, including any command-line interfaces, APIs, etc.]

## Technologies Used
[List of technologies, frameworks, and libraries used in the project]

## Contributing
Brian Zhou, Dustin He
