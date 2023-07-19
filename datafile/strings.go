package datafile

import (
	"bufio"
	"os"
)

func GetStrings (fileName string) ([] string, error) {
var lines []string
file, err := os.Open(fileName)
if err != nil {
	return nil, err
}
scanner := bufio.NewScanner(file)
if scanner.Err() != nil {
 return nil, scanner.Err()
}
for scanner.Scan() {
	line := scanner.Text()
	lines = append(lines, line )

}
err = file.Close()
if err != nil {
return nil, err
}
return lines, nil
}