package games

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var player1 = ""
var player2 = ""
var activePlayer = ""
var playersFull bool = false
var lastPieceX int
var lastPieceY int
var gameWin bool = false

const ROWS = 6
const COLS = 6

var xChoice []string = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣"}
var formatBoard [ROWS][COLS]int
var emptyPiece = "⚪"
var p1Piece = "🔴"
var p2Piece = "🔵"
var boardMessage = ""

//ConnectFour driver
func ConnectFour(s *discordgo.Session, m *discordgo.MessageCreate, connectFourRunning bool, playerStart string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	activePlayer = playerStart

	if !connectFourRunning {
		player1 = playerStart

		s.ChannelMessageSend(testing, "Lets Play ConnectFour!")
		//	playerJoin(s, m) //loops until player 2 joins
		boardToString() //string representation of board
		s.ChannelMessageSend(testing, boardMessage)

	} else {
		if !gameWin {
			if !playersFull {
				playerJoin(s, m)
			} else {
				dropPiece(s, m, player1, player2)
				boardToString()
				s.ChannelMessageSend(testing, boardMessage)
				s.ChannelMessageSend(testing, "Ending turn, Switching to Player: "+activePlayer)
			}
		}
	}
	return
}

func checkWin(x int, y int, lastValue int) bool {
	// horizontalCheck
	for j := 0; j < COLS-4; j++ {
		for i := 0; i < ROWS; i++ {
			if formatBoard[i][j] == lastValue && formatBoard[i][j+1] == lastValue && formatBoard[i][j+2] == lastValue && formatBoard[i][j+3] == lastValue {
				return true
			}
		}
	}
	// verticalCheck
	for i := 0; i < COLS-4; i++ {
		for j := 0; j < ROWS; j++ {
			if formatBoard[i][j] == lastValue && formatBoard[i+1][j] == lastValue && formatBoard[i+2][j] == lastValue && formatBoard[i+3][j] == lastValue {
				return true
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 3; i < COLS; i++ {
		for j := 0; j < ROWS-4; j++ {
			if formatBoard[i][j] == lastValue && formatBoard[i-1][j+1] == lastValue && formatBoard[i-2][j+2] == lastValue && formatBoard[i-3][j+3] == lastValue {
				return true
			}
		}
	}
	// descendingDiagonalCheck
	for i := 3; i < COLS; i++ {
		for j := 3; j < ROWS; j++ {
			if formatBoard[i][j] == lastValue && formatBoard[i-1][j-1] == lastValue && formatBoard[i-2][j-2] == lastValue && formatBoard[i-3][j-3] == lastValue {
				return true
			}
		}
	}
	return false
}

//func boardFull() {
///
//}

func setActive(player1 string, player2 string, activePlayer string) {
	if activePlayer == player1 {
		activePlayer = player2
	} else if activePlayer == player2 {
		activePlayer = player1
	}
}

func checkSpace(input int, pieceVal int) bool {
	i := ROWS - 1
	var emptySpace bool = false
	for i > 0 {
		if formatBoard[i][input] != 0 { //checks the input column, row by row
			i--
			if i == 0 {
				return emptySpace //false if no empty pieces in column
			}
		} else {
			formatBoard[i][input] = pieceVal //sets empty piece to activeplayer piece
			checkWin(i, input, pieceVal)
			emptySpace = true
			break
		}
	}
	return emptySpace
}

func dropPiece(s *discordgo.Session, m *discordgo.MessageCreate, player1 string, player2 string) {
	s.ChannelMessageSend(testing, "Player: "+activePlayer+" turn")
	if m.Author.Username != activePlayer {
		s.ChannelMessageSend(testing, "Error: You are not the active Player!")
	} else {
		input, err := strconv.Atoi(m.Content)
		input--
		if err != nil {
			s.ChannelMessageSend(testing, "Error: input not a number")
			log.Fatal(err)
		}
		if input < 0 || input > COLS {
			s.ChannelMessageSend(testing, "Error: input must be in range 0 to "+strconv.Itoa(COLS))
			log.Fatal(err)
		}
		var pieceVal int = 0
		if activePlayer == player1 {
			pieceVal = 1
		} else {
			pieceVal = 2
		}

		//input to change piece on board to activePlayer color
		check := checkSpace(input, pieceVal)

		fmt.Println(player1)
		// fmt.Println(activePlayer + " ")
		// fmt.Println(check)
		if !check {
			s.ChannelMessageSend(testing, "Error: Column Full input another column")
		} else {
			setActive(player1, player2, activePlayer)

		}
	}

}

func playerJoin(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(testing, "Player2 opt in with g!gameJoin")
	if m.Content == "g!gameJoin" {
		player2 = m.Author.Username
		playersFull = true
		s.ChannelMessageSend(testing, "Added player 2: "+player2)
		return
	}
	if m.Content != "g!gameJoin" {
		s.ChannelMessageSend(testing, "Error No Player 2")
		s.ChannelMessageSend(testing, "Exit with g!stop")
	}

}

func boardToString() {
	boardMessage = ""
	var piece int

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			piece = formatBoard[i][j]
			switch piece {
			case 0:
				boardMessage += emptyPiece
			case 1:
				boardMessage += p1Piece
			case 2:
				boardMessage += p2Piece
			}
		}
		boardMessage += "\n"
	}
	boardMessage += "1️⃣" + "2️⃣" + "3️⃣" + "4️⃣" + "5️⃣" + "6️⃣"
	boardMessage += "\n"
}

func player() {

}

type gameBoard struct {
	LastPiece uint8
	Turn      uint8
	Board     string
}

// TODO:
// not changing players after each turn when it says "ending turn switching to player"
// currently does not check for win
