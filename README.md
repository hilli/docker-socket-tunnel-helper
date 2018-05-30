# docker-socket-tunnel-helper
Need access to remote docker? Don't want to configure/expose docker on TCP? Then this is the tool.

Coding commencing on the 4th of June 2018...


# Having the docker daemon running on a remote host is great.
Getting access to it from you local machine is a little less fun.

 

## Current options:
- Creating your docker host with docker-machine - Fine, but difficult to move credentials around.
- SSHing to the host, running your code editing and docker commands remote - Losing access to local editors and files
- Exposing the docker daemon on TCP (after configuring it listen on TCP), potentially giving root access to the world

## The solution to the above problems:
- Tunnel the docker socket file over SSH
- Having a helper tool, that can automate the job in an easy configurable way.
- Not making the docker daemon listen publicly on TCP.
 

*The intensions are to create said tool, easing the setup process, and continuing maintenance of the tunnel process over SSH.*
