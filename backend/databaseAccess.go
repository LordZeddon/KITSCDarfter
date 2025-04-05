package backend

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func SetUpDB() {
	db, err := sql.Open("sqlite3", "file:./KITSCDrafterDB.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully!")

	// Set up the database schema
	createTables()
}

func createTables() {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Players (
    bTag VARCHAR(50) ,
    nickname VARCHAR(50),
    discordtag VARCHAR(50),
    PRIMARY KEY (bTag)
	);`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating Players table:", err)
	}

	_, err = db.Exec(`CREATE TABLE if NOT EXISTS Teams (
    teamID INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50)
	);`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating Teams table:", err)
	}

	_, err = db.Exec(`CREATE TABLE if NOT EXISTS Tournaments (
    tournamentID INT,
    name VARCHAR(50),
    PRIMARY KEY (tournamentID)
	);`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating Tournaments table:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS playerRanks (
    bTag VARCHAR(50),
    tournamentID INT,
    role INT,
    rankCurrent INT,
    rankPeak INT,
    PRIMARY KEY (bTag, tournamentID, role),
    FOREIGN KEY (bTag) REFERENCES Players(bTag)
    FOREIGN KEY (tournamentID) REFERENCES Tournaments(tournamentID)
    );`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating playerRanks table:", err)
	}

	_, err = db.Exec(`CREATE TABLE if NOT EXISTS Captains (
    teamID INTEGER,
    bTag VARCHAR(50),
    tournamentID INT,
	PRIMARY KEY (teamID),
    FOREIGN KEY (teamID) REFERENCES Teams(teamID),
    FOREIGN KEY (bTag) REFERENCES Players(bTag),
    FOREIGN KEY (tournamentID) REFERENCES Tournaments(tournamentID)
	);`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating Captains table:", err)
	}

	_, err = db.Exec(`CREATE TABLE if NOT EXISTS PlayerAssignments (
    teamID INTEGER,
    bTag VARCHAR(50),
    PRIMARY KEY (teamID, bTag),
    FOREIGN KEY (teamID) REFERENCES Teams(teamID),
    FOREIGN KEY (bTag) REFERENCES Players(bTag)
	);`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating PlayerAssignments table:", err)
	}

	//Unhappy with this, boosterID is functionally dependant on bTag

	_, err = db.Exec(`CREATE TABLE if NOT EXISTS Boosters (
    boosterID INT,
    bTag VARCHAR(50),
    PRIMARY KEY (boosterID, bTag),
    FOREIGN KEY (bTag) REFERENCES Players(bTag)
	);`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating Boosters table:", err)
	}

	//Unhappy with this: TournamentID has to be queried with high effort

	_, err = db.Exec(`CREATE TABLE if NOT EXISTS Matches (
    matchID INT, 
    teamIDHome INTEGER, 
    teamIDAway INTEGER, 
    pointsHome INT,
    pointsAway INT, 
    promotesTo INT,
    demotesTo INT,
    PRIMARY KEY (matchID),
    FOREIGN KEY (teamIDHome, teamIDAway) REFERENCES Teams(teamID, teamID),
    FOREIGN KEY (promotesTo, demotesTo) REFERENCES Matches(matchID, matchID)
	);`)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating Matches table:", err)
	}
}

/*
ranks[][] int: dim 3x2 dim of size 3 represents the roles, rank[i][0] is the current rank, rank[i][1] is the peak rank
ranks[0][j] is the Tank rank
ranks[1][j] is the DPS rank
ranks[2][j] is the Support rank
*/
func insertPlayer(bTag string, nickname string, discordtag string, ranks [][]int, tournamentID int) {
	if db == nil {
		log.Fatal("Database connection is not initialized.")
		return
	}

	_, err := db.Exec("INSERT INTO Players (bTag, nickname, discordtag) VALUES (?, ?, ?)", bTag, nickname, discordtag)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		if ranks[i][0] != -1 {
			_, err = db.Exec("INSERT INTO PlayerRanks (bTag, torunamentID, role, rankCurrent, rankPeak) VALUES (?,?,?,?,?)", bTag, tournamentID, i, ranks[i][0], ranks[i][1])
		}

		if err != nil {
			log.Fatal(err)
		}

	}
}

func insertCaptain(bTag string, tournamentID int) {
	if db == nil {
		log.Fatal("Database connection is not initialized.")
		return
	}

	_, err := db.Exec("INSERT INTO Captains (bTag, tournamentID) VALUES (?, ?)", bTag, tournamentID)
	if err != nil {
		log.Fatal(err)
	}
}

func pickPlayer(bTag string, teamID int) {
	if db == nil {
		log.Fatal("Database connection is not initialized.")
		return
	}

	_, err := db.Exec("INSERT INTO PlayerAssignments (bTag, teamID) VALUES (?, ?)", bTag, teamID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM Boosters WHERE bTag == ?", bTag)

	if err != nil {
		log.Fatal(err)
	}
}
