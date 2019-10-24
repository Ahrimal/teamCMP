#Installation steps
go get github.com/Ahrimal/teamCMP/...

#How to run your code / tests
Build:
1. cd $GOPATH/src/github.com/Ahrimal/teamCMP/
2. go build .
3. teamCMP [-a | <source_name>]

-a means import all videos from all sources

Test: go run ./test

For docker: 
sudo docker image build -t team_cmp .
sudo docker run team_cmp [-a | <source_name>]
docker run -it testTeamCMP 

#Where to find your code
You can find the code at: 

#Was it your first time writing a unit test, using a particular framework, etc?

#What would you have done differently if you had had more time
1. Added a help command message
2. Better testing (the tests are quite simple)
3. Implement the DB connections
4. Better error handlers
5. Making a backup and deletion of read files
6. Maybe a rethink of the actual design... (it has an initial thinking about FTP, but has not been done at all)
7. Remove println from unit testing in go calls
8. Defer close opened file should be better handled if needed