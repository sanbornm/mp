mp - a simple mail parser [![Build Status](https://travis-ci.org/sanbornm/mp.svg?branch=master)](https://travis-ci.org/sanbornm/mp.svg?branch=master)
==

mp is a simple cli email parser.

It currently takes stdin and outputs JSON.

Example:

```
cat fixtures/test.eml | mp
```

```
{
  "Text": "Hello world!",
  "Html": "\u003cb\u003eHello world!\u003c/b\u003e",
  "Headers": {
    "Content-Type": "multipart/mixed; boundary=\"----mailcomposer-?=_1-1402581589619\"",
    "From": "\"Me\" \u003cme@domain.com\u003e",
    "Mime-Version": "1.0",
    "To": "\"First Receiver\" \u003cfirst@domain.com\u003e, second@domain.com",
    "X-Mailer": "Nodemailer 1.0"
  },
  "From": [
    {
      "Name": "Me",
      "Address": "me@domain.com"
    }
  ],
  "To": [
    {
      "Name": "First Receiver",
      "Address": "first@domain.com"
    },
    {
      "Name": "",
      "Address": "second@domain.com"
    }
  ],
  "Cc": null,
  "Priority": "normal",
  "Attachments": [
    {
      "ContentType": "text/plain",
      "Filename": "dummyFile.txt",
      "Disposition": "attachment"
    }
  ]
}
```
