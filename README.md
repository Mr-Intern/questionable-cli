# questionable-cli
![image](https://user-images.githubusercontent.com/24460340/182623974-eb0641b1-c686-4261-9cca-1a57c4077f1d.png)
## A very questionable cli indeed. Written in Go.

This is a cli written in Go with a lot of help from Cobra. I wrote it to have some practice writing in Go, and writing CLI's.
It's public but I don't really see what use anyone else could get out of it lol. Enjoy.

Example: `qcli <command>`

## Requirements
- golang

## How to use

1.  Clone this repo `git clone https://github.com/mrintern/questionable-cli.git`
2.  Build the executable `go build`
3.  Use the following command to find your `/go/bin` direcotry
 `go list -f '{{.Target}}'`
4.  Add your `/go/bin` directory from the above command to your `$PATH`

for example, if the output is `export PATH=$PATH:/home/intern/go/bin/questionable-cli`
your command would be `export PATH=$PATH:/home/intern/go/bin`

5.  Run `alias qcli='questionable-cli'`
6.  Run `go install`

- test your installation by running `qcli` ,  you should see the following output:
![image](https://user-images.githubusercontent.com/24460340/183295404-2a162fd0-b7c1-4399-ada5-84a2b65911a7.png)

7. To persist these changes, add the commands from step 4 and 5 to the end of your .bashrc or .zshrc (depending on what terminal you use). It should look something like this:

![image](https://user-images.githubusercontent.com/24460340/183295674-881d8113-d79f-4dd7-af9e-b046ef4796b0.png)
