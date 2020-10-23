package wordAnalysis

import(
  "fmt"
)

type wordCount struct{
  word string
  count int
}

//Print out information
func PrintPairs(values []*wordCount){
  for _, pair := range values{
    fmt.Println(pair.word + " " + string(pair.count))
  }
}

//create a slice of string with their correspoding number of occurances
func CreatePairs( values []string ) []*wordCount{
  var pairs []*wordCount

  for _, word := range values{
    if !doesPairsContain(pairs, word){

      pairs = append(pairs, &wordCount{word, 1})
    }else{
      pairs = incPairValue(pairs, word)
    }
  }
  return pairs
}

//Check if a pair exists in a slice
func doesPairsContain(pairs []*wordCount, val string) bool {
  for _, pair := range pairs{
    if val == pair.word{
      return true;
    }
  }
  return false;
}

//increment the number of occurances of word 
func incPairValue(pairs []*wordCount, key string) []*wordCount{
  for _, pair := range pairs{
    if pair.word == key{
      pair.count += 1;
      return pairs
    }
  }
  return pairs
}
