# questionable-cli
![image](https://user-images.githubusercontent.com/24460340/182623974-eb0641b1-c686-4261-9cca-1a57c4077f1d.png)
## A very questionable cli indeed. Written in Go.

This is a cli written in Go with a lot of help from Cobra. I wrote it to have some practice writing in Go, and writing CLI's.
It's public but I don't really see what use anyone else could get out of it lol. Enjoy.

Example: `questionable-cli <command>`

## Requirements
- golang

## How to use

1.  Clone this repo `git clone https://github.com/mrintern/questionable-cli.git`
2.  Build the executable `go build`
3.  Add to your $PATH `sudo mv questionable-cli /usr/bin`

## (Optional) create qcli alias
5.  Add the following line to .bashrc, .zshrc, or whatever file governs your shells profile `alias qcli='questionable-cli'` it should look something like this:

![image](https://user-images.githubusercontent.com/24460340/183295674-881d8113-d79f-4dd7-af9e-b046ef4796b0.png)

6.  restart your shell and run `qcli`

Now, you can just type `qcli` to run the tool, which is more convenient than typing `questionable-cli`
