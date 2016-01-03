# goMongodbUI

goMongodbUI is a web based user interface for [MongoDB] database. The main goal for this project is to create free UI that is very easy and OS independent. This project was inspired by [MailHog].

It uses beego web framework for API calls and AngularJS for frontend webapp. It could be done all with beego using templates but this way I feel it is more modular and flexible.

### Version
0.0.1

### Tech

* [Go] - an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Beego] - An open source framework to build and develop your applications in the Go way.
* [Yeoman] - The web's scaffolding tool for modern webapps.
* [AngularJS] - frontend web framework.

### Installation

You need Golang installed:

```sh
$ sudo apt-get install golang
```

```sh
$ git clone https://github.com/karantan/goMongodbUI goMongodbUI 
$ cd goMongodbUI
$ export GOPATH=<your_path>/goMongodbUI
$ go get github.com/astaxie/beego
$ go get gopkg.in/mgo.v2
$ go build
```

### Todos
 - drop collection
 - remove document
 - update document
 - AngularJS frontend
 - Write Tests


   [MongoDB]: <https://www.mongodb.org/>
   [MailHog]: <https://github.com/mailhog/MailHog>
   [Go]: <https://golang.org/>
   [Beego]: <http://beego.me/>
   [Yeoman]: <http://yeoman.io/>
   [AngularJS]: <https://angularjs.org/>
