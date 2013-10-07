You can download Go from [http://golang.org/](http://golang.org/)

# Database
You will need to install MongoDB. You will find it at [https://www.mongodb.com/](https://www.mongodb.com/). You will also need to create the path C:\data\db (Windows) or /data/db (other). 

# Dependencies
First of all, you will need to install Bazaar, which you can get at [http://wiki.bazaar.canonical.com/Download](http://wiki.bazaar.canonical.com/Download). Next you will have to add the path to bzr.exe (or bzr) to your search path. The last step is to issue the command

go get labix.org/v2/mgo

# Running

Add your Git's path to the GOPATH enviroment variable. Then, to build the web server, enter:

go build gomgoweb.go

You might need more parameters if you're using Linux. To start the server, simply execute the "gomgoweb" command.  