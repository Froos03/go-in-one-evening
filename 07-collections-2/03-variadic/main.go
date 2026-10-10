package main

func DebugLog(args ...string) []string {
	prependedMessage := []string{"[DEBUG]"}
	messages := append(prependedMessage, args...)
	return messages
}

func InfoLog(args ...string) []string {
	prependedMessage := []string{"[INFO]"}
	messages := append(prependedMessage, args...)
	return messages
}

func ErrorLog(args ...string) []string {
	prependedMessage := []string{"[ERROR]"}
	messages := append(prependedMessage, args...)
	return messages
}
