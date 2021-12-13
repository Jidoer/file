package file
import (
	"io/ioutil"
	"os"
)

func Create(fileName string,nr string) {
   content := []byte(nr)
   err := ioutil.WriteFile(fileName, content, 0644)
   if err != nil {
      panic(err)
   }
}
func Reader(fileName string) string{
    
   file, err := os.Open(fileName)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   content, err := ioutil.ReadAll(file)
  //log.Print(string(content))
  result := string(content)
   return result
}
func Exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
