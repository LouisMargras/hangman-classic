package main

import (
	"Hangman_Project/CreateList"
	"Hangman_Project/Game"
	"Hangman_Project/RequestUsr"
)

func main() { //fonction principale
	var ListLetterUsed []string
	CharList := []string{" ", "!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";", "<", "=", ">", "?", "@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "[", "\\", "]", "^", "_", "'", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "{", "|", "}", "~"} //liste de tous les charactères existant dans pour le ASCII art
	ASCII := RequestUsr.AskASCII()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    // Demande quel ASCII art  afficher
	for ASCII == "" {
		ASCII = RequestUsr.AskASCII()
	}
	ListASCII := CreateList.ReadFile(ASCII)
	IndexOfDeath, LineHangman := 0, 0
	LevelName := RequestUsr.Level(ListASCII, CharList) // demande le niveau de difficulté
	for LevelName == "" {
		LevelName = RequestUsr.Level(ListASCII, CharList)
	}
	ListWord := CreateList.ReadFile(LevelName)
	ListWordCap := CreateList.CreateListWordCap(ListWord[RequestUsr.Random(ListWord)])
	DashList := CreateList.CreateDashList(ListWordCap)
	Game.Game(LineHangman, ListLetterUsed, DashList, IndexOfDeath, ListWordCap, ListASCII, CharList) //lance le jeu
}
