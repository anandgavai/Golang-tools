# Golang-tools
# Following settings are important to make sure GO is working correctly

## change yoru .bashrc file to

export GOROOT=/usr/local/go  # location where GO is installed

export GOPATH=$GOROOT/bin    # combines GOROOT and points to go execution code

export PATH="/usr/local/bin:$GOROOT:$GOPATH:$PATH"



## Installing packages from github for example golearn (Machine Learning)
go get -t -u -v github.com/sjwhitworth/golearn

## all packages will be installed in

$GOPATH/src


My findings

1. variables and methods are only declared within Structures and Interfaces

2. Remember there will never be a method declaration in a structure

3. Inheritance of a structure for method declaration is done by passing structure argument to a function that is outside the structure

4. variable types are mentioned after variable names
