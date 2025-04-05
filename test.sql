CREATE TABLE IF NOT EXISTS Players (
    bTag VARCHAR(50) ,
    nickname VARCHAR(50),
    discordtag VARCHAR(50),
    PRIMARY KEY (bTag)
);

CREATE TABLE IF NOT EXISTS playerRanks (
    bTag VARCHAR(50),
    tournamentID INT,
    role INT,
    rankCurrent INT,
    rankPeak INT,
    PRIMARY KEY(bTag, tournamentID, role),
    FOREIGN KEY (bTag) REFERENCES Players(bTag)
    FOREIGN KEY (tournamentID) REFERENCES Tournaments(tournamentID)
    );

CREATE TABLE if NOT EXISTS Teams (
    teamID INT,
    name VARCHAR(50)
    PRIMARY KEY (teamID)
);

CREATE TABLE if NOT EXISTS Captains (
    teamID INT,
    bTag VARCHAR(50),
    tournamentID INT,
    PRIMARY KEY (teamID),
    FOREIGN KEY (teamID) REFERENCES Teams(teamID),
    FOREIGN KEY (bTag) REFERENCES Players(bTag),
    FOREIGN KEY (tournamentID) REFERENCES Tournaments(tournamentID)
);

CREATE TABLE if NOT EXISTS PlayerAssignments (
    teamID INT,
    bTag VARCHAR(50),
    PRIMARY KEY (teamID, bTag),
    FOREIGN KEY teamID REFERENCES Teams(teamID),
    FOREIGN KEY bTag REFERENCES Players(bTag)
);

CREATE TABLE if NOT EXISTS Boosters (
    boosterID INT,
    bTag VARCHAR(50),
    PRIMARY KEY (boosterID, bTag),
    FOREIGN KEY bTag REFERENCES Players(bTag)
);

CREATE TABLE if NOT EXISTS Matches (
    matchID INT, 
    teamIDHome INT, 
    teamIDAway INT, 
    pointsHome INT,
    pointsAway INT, 
    promotesTo INT,
    demotesTo INT,
    PRIMARY KEY (matchID)
    FOREIGN KEY teamIDHome, teamIDAway REFERENCES Teams(teamID),
    FOREIGN KEY promotesTo, demotesTo REFERNCES Matches(matchID)
);

CREATE TABLE if NOT EXISTS Tournaments (
    tournamentID INT,
    name VARCHAR(50),
    PRIMARY KEY (tournamentID)
);

