package main
 
import (
    "net/mail"
    "bytes"
    "io"
    "io/ioutil"
    "os"
    "github.com/jhillyerd/go.enmime"
    "fmt"
    "encoding/json"
)

type Attachment struct {
    ContentType string
    Filename string
    Disposition string
}

type Email struct {
    Text string
    Html string
    Headers map[string]string
    From []*mail.Address
    To []*mail.Address
    Cc []*mail.Address
    Priority string
    Attachments []Attachment
}

func StreamToString(stream io.Reader) string {
    buf := new(bytes.Buffer)
    buf.ReadFrom(stream)

    return buf.String()
}
 
func ParseEmail(content string) Email {
    email := Email{}

    msg, err := mail.ReadMessage(bytes.NewBuffer([]byte(content)))
    if err != nil {
        return email
    }
    mime, _ := enmime.ParseMIMEBody(msg)

    headers := map[string]string{}
    for key,values := range msg.Header {
        headers[key] = values[0]
    }

    fromAddrs,_ := msg.Header.AddressList("From")
    toAddrs,_ := msg.Header.AddressList("To")
    ccAddrs,_ := msg.Header.AddressList("Cc")

    var attachments []Attachment
    for _,attach := range mime.Attachments {
        a := Attachment{}
        a.ContentType = attach.ContentType()
        a.Filename = attach.FileName()
        a.Disposition = attach.Disposition()
        attachments = append(attachments, a)
    }

    return Email{
        Html: mime.Html,
        Text: mime.Text,
        Headers: headers,
        Priority: "normal",
        From: fromAddrs,
        To: toAddrs,
        Cc: ccAddrs,
        Attachments: attachments}
}

func main() {
    bytes,_ := ioutil.ReadAll(os.Stdin)

    email := ParseEmail(string(bytes))

    b, err := json.MarshalIndent(email, "", "  ")
    if err != nil {
        fmt.Println("error:", err)
    }
    os.Stdout.Write(b)
}
