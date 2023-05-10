# mtemp
Create temp email and read received mails

# Installation

```bash
$ go install github.com/nsecho/mtemp@latest
```

# Usage

```bash
$ mtemp help
Create and monitor temp emails

Usage:
  mtemp [command]

Available Commands:
  help        Help about any command
  new         Create new email

Flags:
  -h, --help   help for mtemp

Use "mtemp [command] --help" for more information about a command.
```

By default, `new` will only create new temp email on the `1secmail.com` website. If you pass `-m` flag it will monitor mailbox for new mails every 5 seconds and print the content to the screen.

```bash
$ mtemp new -m
Created new "ozsgslct@wuuvo.com" mail
New mail Subject from xdaemonx@proton.me
Content:
<div style="font-family: Arial, sans-serif; font-size: 14px;">Test email</div><div style="font-family: Arial, sans-serif; font-size: 14px;"><br></div>
<div class="protonmail_signature_block" style="font-family: Arial, sans-serif; font-size: 14px;">
    <div class="protonmail_signature_block-user protonmail_signature_block-empty">

            </div>

            <div class="protonmail_signature_block-proton">
        Sent with <a target="_blank" href="https://proton.me/" rel="noopener noreferrer">Proton Mail</a> secure email.
    </div>
</div>
```
