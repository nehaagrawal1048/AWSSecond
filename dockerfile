#################
FROM golang:1.24-alpine


#### Create directory ##########
WORKDIR /awssecond

######### copy all your code in the image
COPY . .

###### create build ####
RUN go build -o application .

##### RUN THE BUILD ####
CMD ["./application"]


# #### command to craete image
 #docker build -t testimage .

# ####### check the docker image
# docker image ls


##### container out of image
# docker run testimage

# docker run -p 8050:8050 testimage

# docker run -p 9000:8050 testimage

# 9000(outside world ) -> (inside docker image)
# when outside app has to access the inside app then expose it with the port
# outside -> inside
# -t is for naming the docker image to ur custom name


#docker run --name 8julycontainer -d 8july // change the name of the container

# -d detach flag to make the terminal free and not stuck

# docker stop 8julycontainer
# //docker remove 8julycontainer

#docker run --name containername -d -p 8050:8080 imagename //expose the port to access the app in docker(outside to inside port)

#docker compose up -d       //to run the yml file

#docker compose down


# //create the repo
# // create the image with the same name
# //push the image
