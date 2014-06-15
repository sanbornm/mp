package main

import(
    "testing"
    "io/ioutil"
)

func TestParseEmailContainsHtml(t *testing.T) {
    content, err := ioutil.ReadFile("fixtures/test.eml")
    if err != nil {
        //Do something
    }
    data := ParseEmail(string(content))
    expected := "<b>Hello world!</b>"
    if data.Html !=  expected {
        t.Errorf("Expected html %s got: '%s'", expected, data.Html)
    }
}

func TestParseEmailContainsText(t *testing.T) {
    content, err := ioutil.ReadFile("fixtures/test.eml")
    if err != nil {
        //Do something
    }
    data := ParseEmail(string(content))
    expected := "Hello world!"
    if data.Text !=  expected {
        t.Errorf("Expected text: '%s' got: '%s'", expected, data.Text)
    }
}

func TestHeaderIsPresent(t *testing.T) {
    content, err := ioutil.ReadFile("fixtures/test.eml")
    if err != nil {
        //Do something
    }
    data := ParseEmail(string(content))
    expected := "Nodemailer 1.0"
    if data.Headers["X-Mailer"] != expected {
        t.Errorf("Expected: '%s' got: '%s'", expected, data.Headers["x-mailer"])
    }
}

func TestFromName(t *testing.T) {
    content, err := ioutil.ReadFile("fixtures/test.eml")
    if err != nil {
        //Do something
    }
    data := ParseEmail(string(content))
    expected := "Me"
    if data.From[0].Name != expected {
        t.Errorf("Expected: '%s' got: '%s'", expected, data.From[0].Name)
    }

    expected = "me@domain.com"
    if data.From[0].Address != expected {
        t.Errorf("Expected: '%s' got: '%s'", expected, data.From[0].Address)
    }
}
