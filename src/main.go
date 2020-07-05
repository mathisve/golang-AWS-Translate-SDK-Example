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

func init() {
  translateSession = translate.New(session.Must(session.NewSession(&aws.Config{
    Region: aws.String("eu-central-1"), // Frankfurt
    })))
}

func translateAndWrite (text []byte, TARGET_LANGUAGE string) {
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

func main() {
  text, err := ioutil.ReadFile("lyrics.txt")
  if err != nil {
    panic(err)
  }

  for _, TARGET_LANGUAGE := range []string{"en", "nl", "de", "es"} {
    translateAndWrite(text, TARGET_LANGUAGE)
  }

}
