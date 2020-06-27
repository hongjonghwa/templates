# Docker
docker-machine start
eval $(docker-machine env)

# RUN
docker run --rm --name jekyll-container -v "$PWD/../hongjonghwa.github.io:/srv/jekyll" -p 4000:4000  -it jekyll/jekyll:3.8 bash

# RE-RUN
docker start  `docker ps -f name=jekyll-container -aq` # restart it in the background
docker attach `docker ps -f name=jekyll-container -aq` # reattach the terminal & stdin

#export JEKYLL_VERSION=3.8
#docker run --rm \
#  --volume="$PWD/../hongjonghwa.github.io:/srv/jekyll" \
#  -it jekyll/jekyll:$JEKYLL_VERSION \
#  jekyll build

# JEKYLL
jekyll build
jekyll serve

##
http://$DOCKER_HOST:4000
http://192.168.99.101:4000
