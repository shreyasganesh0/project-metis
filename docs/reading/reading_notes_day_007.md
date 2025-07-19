# Web Servers and Docker

## Net/http Package Go

### Quick Clients
- http.Get(url)
    - get the response from a request to url
- http.Post(url, "response/type", &buf)
- http.PostForm(url , "url.Values{"key": {"Values"}})

- the called must close the response body
    - resp.Body.Close()

### Clients and Transports
```
client := &http.Client {
    CheckRedirect: redirectPolicyFunc,
}
client.Get(url)

req = http.NewRequest("GET", url, nil);
req.Header.Add("Header_key", "header-value")
client.Do(req)
```
- this is a little more fine grained control over a client
    - lets you set header and redirect policy etc

```
tr := http.Transport {
    MaxIdelConns: 10,
    IdleConnTimeout: 30 * time.Second,
    DisableCompression: true,
}
```
- to configure tls keep alive and compresson


### Servers
```
http.Handle("/path/to/resource", handler)

handler(w http.ResponseWriter, r *http.Request) {}

log.Fatal(http.ListenAndServe(":8080", nil)
```
- the Handle funciton reads the request line parses it via a mux
    - the mux is principle of exact match where / will match every single
      path
- listen and serve starts a port listener on the hostname:port


## Dockerfile best practices
1. Multi stage builds reduce size of final image
    - dockerfile instructions into stages
        - each stage must only create files that are necessary
        - stages can also be run in parallel
    - create reusable stages
        - multiple images can share the outputs of common stages

2. Choose the right base image
    - choose base image from a secure source
        - Docker official images
        - Verified publisher
        - Docker sponsored open source
    - choose a minimal base image that matched requirements when building your own
    - seperate image for buliding and e unit testing (heavier) and prod (ligther)
        - lighter images wont contain the build tools and debugging tools

3. Rebuild images often
    - building an image is a snapshot of the image
    - to update dependencies we need to rebuild the image
    - --no-cache option with docker build to avoid caching old dependency versions
        - FROM ubuntu:24.04
            - this uses the 24.04 version of the ubuntu image
            - the underlying dependencies may change with time
            - this is where --no-cache helps
4. Exclude with .dockerignore
    - exclude files not relevant for building

5. Create ephemeral containers
    - the image should support ephemeral containers
    - should be able to be stopped and rebuild with minimum configuration

6. Dont install unnecessary packages
    - dont include packages that are just for nice to have sake
    - reduces file size and build times

7. Decouple Applications
    - put applications with different concerns into different containers
        - webapp, database and cache each running in their own containers
    - one processes per container is ideal but not hard and fast
        - celery and apache can create processes per request or worker processes
8. Sort multi-line arguments
    - alphanumerically sort multi-line arguments
    - helps avoid duplication of packages and makes it easier to update
    - add space befroe a \
        - RUN apt-get update && apt-get install -y --no-install-recommends \
            bzr \
            cvs \
            git \
            mercurial \
            subverions \
            && rm -rf /var/lib/apt/lists/*
9. Leverage build cache
    - docker goes through the instructions in the dockerfile
    - docker checks if it can reuse the instruction from the build cache
    - using the build cache to your advantage can speed up build times
        - each instruction is a layer
        - if one layer changes all layers below it also get invalidated

10. Pin base image versions
    - image tags are mutable so they can be made to point to a different image
    - this helps update dependencies related to a particular version wihtout having
      to change the version in the client dockerfile
    - downside is it could possibly break builds
        - also we dont have a audit trail for version changes
    - we can pin the particular image to the version using its sha256 digest
        - can be tedious to look up the exact digest number everytime you want to change
    - Docker scout - Up to date image policy
        - checks if the base image version is the latest version
        - also checks if pinned digests in the dockerfile match the correct version
        - we can use policy compliance to check if the latest version has changed
          and update it
        - it also supports automatic updation if neede
        - it will raise a pr to change the digest if an update is availble
        - this solves the audit trail problem while also giving us security updates

11. Build and test images in CI
    - github actions or other Ci/CD pipline to auto build and tag a Docker image
      and test it

## Dockerfile Instructions
- FROM
    - use offical images as the basis of images

- LABEL
    - add metadata
        - organize images by project
        - record licensing information
    - labels store k v pairs
    - LABEL com.exapmle.version="0.0.1-beta"
    - LABEL vendor=ACME\ Incorporated \
        com.example.is-beta= \
        com.example.is-production="" \

- RUN
    - split long or complex statements
    - for ubuntu apt-get tool
        - combine apt-get update && apt-get install -y --no-install-recommends
        - make this a single run layer so that if more packages are added
        to the apt-get install the packages install will also be updated
        - which wont happen if the update was in a seperate line due to layer cache
    - using |
        - executed using /bin/sh -c interpreter
        - only evals exit code of the last command
        - we have to set -o pipefail && wget -O url | wc -l
            - this way if any command in the pipe fails it returns and error exit code

- CMD
    - run software inside the image
    - CMD ["executable", "param1", "param2"]
    - should uusually be used with a interactive shell like bash python or perl as exe
        - docker run -it python will then drop us into the shell

- EXPOSE
    - ports on which a container listens
    - only acts a indicator the -p flag with docker run does the external mapping

- ENV
    - update enviornment vairables
    - ENV ADMIN_USER="mark"
    - unsetting env variables in a future layer doesnt affect previous layers
        - to actually unset it set and unset should be done in a single RUN using &&

- ADD or COPY
    - copy files into the container from build context or from a stage in mulistage builds
    - ADD is same as COPY bu t aslo supports url fetching form HTTPs and Git URLS
        - extract tar files automatically when adding files from build context
    - for temporary addiontion of files to be used in RUN
        - --mount=type=bind, source=requirements.txt, target=/tmp/requirements.txt
        - more efficient than copies to include files from build context
        - dont persist in the final image
    - ADD is better than running wget and tar since it updates the build cache more
      precisely
        - also does checksum validation
        - protocol for parsing git branches tags and subdirecotories

- ENTRYPOINT
    - sets images mainj command
    - allows image to be run as that command
        - ENTRYPOINT ["s3cmd"]
        - CMD ["--help"]
        - equivalent to docker run s3cmd
    - usefule when we need to do some setup done is a script before running the image
        - like we need to setup some other processes that run only on certain conditions
        - we can make a bash script that checks and runs commands to run those conditiosn
        - then set the bashscript as the entrypoint
- VOLUME
    - expose db storage area, config storage or files created by container

- USER
    - change to non-root user
    - RUN groupadd -r postgres && useradd --no-log-init -r-g postgres postgres
        - create a user and group in the dockerfile
    - gosu can be used instead of sudo to avoid tty signal passing complicaitons

- WORKDIR
    - always use abs path for WORKDIR
    - use instead of doing RUN cd .. && do_something

- ONBUILD
    - executes after the current dockerfile completes
    - executes in any child image derived FROM current image
    - executes ONBUILD before any command in the child dokerfile
    - useful for giving instructions to images that are goin to be built FROM an image
    - ADD or COPY on ONBUILD images should be done carefully since it will fail
      if the new builds context doesnt have those files
    - should be LABEL as -onbuild eg. ruby:2.0-onbuild
    - useful to run images
