# Github Container Image Uploads

## Github Container Registry
Store and manage Docker and OCI images

### Container Registry
- store conatiner images and associate them with a repo
- To authenticate
    - personal access token
    - GITHUB_TOKEN env variable
    - need to have read:packages scopre to install packages
      from other repos
    - Github actions:
        - use the GITHUB_TOKEN by adding {{ secrets.GITHUB_TOKEN}}
        - use it like a JWT token in the cli to POST to the API
            - add it as header authorization: Bearer ${{ secrets.GITHUB_ACTION}}
    - use the token to publish images
- pulling container images
    - docker inspect ghcr.io/NAMESPACE/IMAGE_NAME
    - docker pull image_url

### Docker Hub
```
name: Docker Hub
uses:  docker/login-action@v3
with:
    username: ${{ vars.DOCKERHUB_USSERNAME}}
    password: ${{ secrets.DOCKERHUB_TOKEN}}
```

### Github Container Registry
```
name: Login to Github Container Registry
uses: docker/login-action@v3
with:
    registry: ghcr.io
    username: ${{github.actor}}
    password: ${{ secrets.GITHUB_TOKEN}}
```
