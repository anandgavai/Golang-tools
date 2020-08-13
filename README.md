# Golang-tools
# Following settings are important to make sure GO is working correctly

## change yoru .bashrc file to

export GOROOT=/usr/local/go  # This is the place where go gets insalled

export GOPATH=$GOROOT/bin    # This is the place where you can create different projects and each project has 3 sub folder by default src, pkg and bin

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

## Dealing with Golang Modules follow this link here for a clear instruction

https://golangbot.com/go-packages/
