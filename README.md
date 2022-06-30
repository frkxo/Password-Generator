# Password-Generator
Password-Generator is a CLI tool that generates password.

## Install

If you have Go installed and configured:

```
go get github.com/frkxo/Password-Generator
```

## Usage

```
▶ Password-Generator --help
Password-Generator is a password generator written in Golang.

Usage:
  Password-Generator [flags]

Flags:
  -c, --character-set string           character set of the password.
  -a, --exclude-ambiguous-characters   exclude characters that hard to remember. (default true)
  -e, --exclude-similar-characters     exclude characters that look the same in the characters. (default true)
  -h, --help                           help for Password-Generator
  -o, --include-lowercase-letters      include lowercase letters in the password. (default true)
  -n, --include-numbers                include numbers in the password. (default true)
  -s, --include-symbols                include symbols in the password. (default true)
  -u, --include-uppercase-letters      include uppercase letters in the password. (default true)
  -l, --length uint                    length of the password. (default 24)
  -t, --times uint                     How many passwords you want to generate. (default 1)
```

## Example Usage

```
▶ Password-Generator
Your Password/s is/are:
XPxW?b!%3Qt2g28P4vYnWusn
```

```
▶ Password-Generator -l=36 
Your Password/s is/are:
XturqAV.Qf_dqkDy5+dRX!*XDgkdgwHHNk48
```

```
▶ Password-Generator -t=3 -c=agbkhjl 
Your Password/s is/are:
aajbgagbajkglllljjllgghb
lhblhlhghlbgajahalhblaaa
ghkhgghlhkghglghlgkgjhgh
```

```
▶ Password-Generator -l=8 -c=aunotypfmg -n=false -s=false
Your Password/s is/are:
tfumpoyu
```
