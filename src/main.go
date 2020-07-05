package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/translate"
)

const (
  SOURCE_LANGUAGE = "fr"
)

var translateSession *translate.Translate
var TARGET_LANGUAGES = []string{"en", "nl", "de", "es"}


func init() {
  translateSession = translate.New(session.Must(session.NewSession(&aws.Config{
    Region: aws.String("eu-central-1"), // Frankfurt
    })))
}

func main() {
  text, err := ioutil.ReadFile("lyrics.txt")
  if err != nil {
    panic(err)
  }

  fmt.Println(string(text))

  for _, TARGET_LANGUAGE := range TARGET_LANGUAGES {
    response, err := translateSession.Text(&translate.TextInput{
      SourceLanguageCode: aws.String(SOURCE_LANGUAGE),
      TargetLanguageCode: aws.String(TARGET_LANGUAGE),
      Text: aws.String(string(text)),
    })
    if err != nil {
      panic(err)
    }
    // fmt.Println(*response.TranslatedText)

    f, err := os.Create(fmt.Sprintf("lyrics_%s.txt", TARGET_LANGUAGE))
    if err != nil {
      panic(err)
      f.Close()
    }

    _, err = f.WriteString(*response.TranslatedText)
    if err != nil {
      panic(err)
    }

    f.Close()
  }

}
