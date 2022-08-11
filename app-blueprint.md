# Blueprint

To keep binaries as small as possible seperate out functionality
- generation of cmd string to evade IDS
- it is a hack tool so automate generation of binaries for easy use with a selector
	RATHER than bundle everything always because

## Ratherless of subcommand:
1. main 
-> reformcmd
-> argparsing
-> "handle command implications.." - naming..
1. Then share directory contains any functionailty referenced in model:
	Protocol?
	Encrypted?
	...
	extendable evasion mechanicisms go here

## Either subcommand

For selectivity in compilation
Directory structure should be /model/protocol + protocol test

- This will allow for long term supprt and modernisation of protocol

-> Listener
	
-> Client
	


-> File Transfer inclusive but potential in future a Client - still questioning myself on this one, but
Rational:
	1. Seperate command line interaction to make steathly process exploitation an extend feature
	1. Requirements would bloat the binary so almost certianly in the future this will have seperate binary that was built along side Client so that Client can like evil-winrm only upload and download this binary to perform more compute and time intensive and stealthier exfiltration, add a flag that builds he file transfer functionality if required into main client binary.
	1. More Control over how exfiltration and evasion occurs


