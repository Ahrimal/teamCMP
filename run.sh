#!/bin/bash
go get github.com/Ahrimal/teamCMP/...
cd $GOPATH/src/github.com/Ahrimal/teamCMP/
sudo docker image build -t team_cmp .
read -n 1 -s -r -p "Press any key to continue with load all"
sudo docker run team_cmp -a
read -n 1 -s -r -p "Press any key to continue with load glorf"
sudo docker run team_cmp glorf
read -n 1 -s -r -p "Press any key to continue with load flub"
sudo docker run team_cmp flub
read -n 1 -s -r -p "Press any key to continue with load no-existent"
sudo docker run team_cmp no-existent
