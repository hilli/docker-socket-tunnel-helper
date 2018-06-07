

## SSH Options
-f # Fork
-N # No command in the remote end
-oExitOnForwardFailure=yes # Wait to bg until forwards have been setup


CTRLSOCK=/tmp/my-ctrl-socket
HOST=ping.vin.dk
FORWARD="/Users/hilli/git-projects/docker-ping/docker.sock:/var/run/docker.sock"

ssh -M -S ${CTRLSOCK} -oExitOnForwardFailure=yes -fnNT -L ${FORWARD} ${HOST}
ssh -S ${CTRLSOCK} -O check ${HOST}
ssh -S ${CTRLSOCK} -O exit ${HOST}
