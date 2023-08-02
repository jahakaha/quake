# Quake Parser


### Table of contents
<details open="open">
  <summary>Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About project</a>
      <ol>
        <li><a href="#stack">Stack</a></li>
        <li><a href="#stack">Installation</a></li>
      </ol>
    </li>
    <li><a href="#requirements">Requirements</a></li>
    <li><a href="#running">Running</a></li>
    <li><a href="#reports">Reports</a></li>
    <ol>
        <li><a href="#notes">Notes</a></li>
    </ol>
    <li><a href="#maintainers">Maintainer</a></li>
  </ol>
</details>

## About The Project
 The Quake Log Parser is a Go project that parses Quake 3 Arena server log files, extracting information about each match and the kills that occurred during those matches.
### Stack
The project is made using Golang 
- Golang 1.19
### Installation
git clone https://github.com/jahakaha/quake.git 

## Running
1. Run
```
go run cmd/main.go
```

## Reports
* The Quake Log Parser generates the following reports:

*  Match Report: For each match, it provides the total number of kills, a list of players who participated in the match, and the number of kills each player made.

*  Player Ranking: It shows a ranking of players based on the total number of kills they made across all matches.

*  Deaths by Cause Report: For each match, it groups deaths by their cause and provides the count of deaths for each cause.

## Notes
*  When <'world'> kills a player, that player loses -1 kill score. <'world'> is not considered a player and will not appear in the player list or kill count.
*  The total_kills counter includes both player and <'world'> deaths.

## Maintainers
- Developed by @jahakaha
- Maintained by @jahakaha