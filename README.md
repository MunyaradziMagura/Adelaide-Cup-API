# Adelaidecup-api

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

Adeladiecup-api is a cloud-enabled webservice for the horse-raising, it features service such as video on demand and subscripsion etc...
  - this is a demo project to demonstrate aws and it's intergrated service for academic perpose. It is not a comercial product

# Test the application on your local server!
>this project is developed on Ubuntu os,written in Go.
>windows, apple, linux all supports Golang. I am sure you can run this program on none >linux device but at the moment I'm writing this guide base linux system


### preperation
 - If you have virtual mechine or ubuntu intergration of windows terminal.
    - [How to add ubuntu Tap to Window](https://nodejs.org/)
  - cloe git repositry in your directory
  ### without docker(option)
  - if Golang is already intalled in your system you can run the program by just typing
  ```
  go run main.go
  ```
 ### Building Docker Image
  if not you can swill run the apllication on your local server by creating docker image and running inside of the container
    1.Build docker image make sure you include '-t',to tag our image
    2. for example im creating image with DemoApplication as a tag
```
sudo docker build -t DemoApplication
```

### Running Docker Image

```
sudo docker run -p 80:8080 DemoApplication
```
- in this example we're delegating port '8080' to port '80'
- '-p' [delegated port]:[original Port from container(dockerised app)]

- if everything is working smoothly you'll see prompt on the console
```
"Adleiadecup-api starting..."
```
### cheking currently running docker containers
```
sudo docker ps -al //will show most recent one  up active containers
//or
sudo docker ps -a //this will show every existing containers and images
```

### see if it works on your server
if every step above went smoothly you'll be able to test it on your local server
by just typing your IP address on internet browser.

-sechang park
