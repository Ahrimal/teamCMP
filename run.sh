go get github.com/Ahrimal/testCMP
cd $GOPATH/src/github.com/Ahrimal/teamCMP
sudo docker image build -t team_cmp .
sudo docker run team_cmp -a
sudo docker run team_cmp glorf
sudo docker run team_cmp flub
sudo docker run team_cmp adsf
