package article

import (
	"clx/constants/messages"
	"strconv"
	"strings"

	text "github.com/MichaelMure/go-term-text"

	. "github.com/logrusorgru/aurora/v3"
)

const (
	newLine     = "\n"
	indentBlock = "   "
)

func Parse(title, domain, article, references string) string {
	wrappedTitle, _ := text.Wrap(title, 73)
	truncatedDomain := text.TruncateMax(domain, 73)

	wrappedTitle += newLine
	wrappedTitle += Faint(truncatedDomain).String() + newLine
	wrappedTitle += Faint(messages.LessScreenInfo).String() + newLine
	separator := messages.GetSeparator(73)
	wrappedTitle += separator + newLine + newLine

	lines := strings.Split(article, newLine)
	formattedArticle := ""

	for i, line := range lines {
		isOnFirstOrLastLine := i == 0 || i == len(lines)-1

		if isOnFirstOrLastLine {
			formattedArticle += line + newLine

			continue
		}

		if line == "References" {
			break
		}

		lineIsHeader := isHeader(lines, i, line)

		if lineIsHeader {
			formattedArticle += Bold(line).String() + newLine

			continue
		}

		formattedArticle += line + newLine
	}

	formattedArticle = highlightReferences(formattedArticle)
	formattedReferences := newLine + formatReferences(references)

	return wrappedTitle + formattedArticle + formattedReferences
}

func isHeader(lines []string, i int, line string) bool {
	previousLine := lines[i-1]
	previousLineIsEmpty := len(previousLine) == 0

	nextLine := lines[i+1]
	nextLineLineIsEmpty := len(nextLine) == 0

	currentLineIsNotIndented := !strings.HasPrefix(line, " ")

	lineIsHeader := currentLineIsNotIndented && previousLineIsEmpty && nextLineLineIsEmpty

	return lineIsHeader
}

func highlightReferences(input string) string {
	input = strings.ReplaceAll(input, "[1]", "["+Red("1").String()+"]")
	input = strings.ReplaceAll(input, "[2]", "["+Yellow("2").String()+"]")
	input = strings.ReplaceAll(input, "[3]", "["+Green("3").String()+"]")
	input = strings.ReplaceAll(input, "[4]", "["+Blue("4").String()+"]")
	input = strings.ReplaceAll(input, "[5]", "["+Cyan("5").String()+"]")
	input = strings.ReplaceAll(input, "[6]", "["+Magenta("6").String()+"]")
	input = strings.ReplaceAll(input, "[7]", "["+White("7").String()+"]")
	input = strings.ReplaceAll(input, "[8]", "["+BrightRed("8").String()+"]")
	input = strings.ReplaceAll(input, "[9]", "["+BrightYellow("9").String()+"]")
	input = strings.ReplaceAll(input, "[10]", "["+BrightGreen("10").String()+"]")
	input = strings.ReplaceAll(input, "[11]", "["+BrightBlue("11").String()+"]")
	input = strings.ReplaceAll(input, "[12]", "["+BrightCyan("12").String()+"]")
	input = strings.ReplaceAll(input, "[13]", "["+White("13").String()+"]")
	input = strings.ReplaceAll(input, "[14]", "["+Red("14").String()+"]")
	input = strings.ReplaceAll(input, "[15]", "["+Yellow("15").String()+"]")
	input = strings.ReplaceAll(input, "[16]", "["+Green("16").String()+"]")

	return input
}

func formatReferences(references string) string {
	lines := strings.Split(references, newLine)
	formattedReferences := Faint("References").String() + newLine + newLine

	if len(lines) == 1 {
		return ""
	}

	for i, line := range lines {
		isOnLastLine := i == len(lines)-1

		if isOnLastLine {
			break
		}

		if i == 16 {
			break
		}

		number := strconv.Itoa(i + 1)

		formattedReferences += indentBlock + "[" + number + "] " + line + newLine
	}

	formattedReferences = highlightReferences(formattedReferences)

	return formattedReferences
}
