package wordAnalysis

import(
  "fmt"
  "io/ioutil"
  "strings"
  "os"
)

//Read text from file and return each a slice of words
func ReadText(filename string) []string{
  data, err := ioutil.ReadFile(filename)
  if err != nil{
    fmt.Println(err)
    return nil
  }
  return seperateIntoWords(toString(data))
}

//convert a byte slice to a string
func toString(value []byte) string{
  return string(value)
}

//create slice of words form string
func seperateIntoWords(value string) []string{
  return strings.Fields(value)
}

//Write text to information in the slice to file
func WriteText(lines []*wordCount){
  file, err := os.Create("results.txt")
  if err != nil{
    fmt.Println(err)
  }

  for _, line := range lines{
    file.WriteString(line.word + " " + fmt.Sprintf("%d\n",line.count))
  }

  if err != nil{
    fmt.Println(err)
  }

  file.Close()
}
