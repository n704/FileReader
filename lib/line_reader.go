package lib

import (
	"bufio"
	"os"
)

/*
lineReaderFunc is type output function as part closure response
*/
type lineReaderFunc func() (string, bool)

/*
FileReader Object
*/
type FileReader struct {
	file     *os.File
	ReadLine lineReaderFunc
}

/*
fetchLineFunc is returns a func to read line.
*/
func (fileReader *FileReader) fetchLineFunc() lineReaderFunc {
	scanner := bufio.NewScanner(fileReader.file)
	return func() (string, bool) {
		for scanner.Scan() {
			command := scanner.Text()
			return command, true
		}
		return "", false
	}
}

/*
getReadLineFunc updates func of ReadLine to fetch line.
*/
func (fileReader *FileReader) getReadLineFunc() {
	readerFunc := fileReader.fetchLineFunc()
	fileReader.ReadLine = readerFunc
}

/*
Close function is close underlying file object
*/
func (fileReader *FileReader) Close() {
	fileReader.file.Close()
}

/*
GetLineReader return FileReader objects to read line by line
of the file.

@args file os.File
@returns FileReader object
*/
func GetLineReader(file *os.File) (*FileReader, bool) {
	if file == nil {
		return nil, false
	}
	fileReader := &FileReader{file: file}
	fileReader.getReadLineFunc()
	return fileReader, true
}
