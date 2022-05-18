# Running Atlantis
Atlantis is a tools to run terraform automation.

To run atlantis locally
```shell
docker run -p 4141:4141 \ 
    atlantis server \
    --config /usr/local/bin/config.yml
```