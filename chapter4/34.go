package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

func main(){
	filename := "neko.txt.mecab"
	file, err := os.Open(filename)
	if err != nil{
		fmt.Printf("os.Open: %#v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sentence := make([]map[string]string, 0)
	sentences := make([][]map[string]string, 0)
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		if string(b) != "EOS" {
			txt := string(b)
			txt = strings.Replace(txt, "\t", ",", -1)
			elements := strings.Split(txt, ",")
			morpheme := make(map[string]string)
			morpheme["surface"] = elements[0]
			morpheme["base"] = elements[7]
			morpheme["pos"] = elements[1]
			morpheme["pos1"] = elements[2]
			sentence = append(sentence, morpheme)
		}else{
			if len(sentence) > 0{
				sentences = append(sentences, sentence)
				sentence = make([]map[string]string, 0)
			}
		}
	}
	
	no := []string{}
	for _, sentence := range sentences{
		for i, morpheme := range sentence{
			if morpheme["surface"] == "の" && i > 0 && i < len(sentence)-1{
				if sentence[i-1]["pos"] == "名詞" && sentence[i+1]["pos"] == "名詞"{
					no = append(no, sentence[i-1]["surface"] + morpheme["surface"] + sentence[i+1]["surface"])
				}
			}
		}
	}
	fmt.Println(no)
}