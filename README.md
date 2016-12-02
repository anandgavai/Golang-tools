# Golang-tools
# Following settings are important to make sure GO is working working

## change yoru .bashrc file to 

export GOROOT=/usr/local/go  # location where GO is installed
export GOPATH=$GOROOT/bin    # combines GOROOT and points to go execution code
export PATH="/usr/local/bin:$GOROOT:$GOPATH:$PATH"
