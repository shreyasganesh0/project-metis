# Github Actions And Docker

## Github Actions
- Githubs CI CD platform
- create workflows for building and testing on every pull request or push (trigger)
- Event based. lets you run workflows when any event occurs on your repo
    - add labels on issue creation for example

- can have your own self hosted runners via cloud or data centers as well

### Components
- Workflow -> triggered by event in the repo.
    - events are pull requests, issues etc
- Workflow consits of jobs
    - run in sequential or parallel
    - each job runs in its own VM runner or container
        - has one or e more steps
            - steps run scripts or run actions (reusable extensions)
- Workflow
    - automated process
    - runs one or more jobs
    - defined by YAML file
        - present in .github/workflows dir
        - each repo can have multiple worflows
    - triggered on event, cron job or manually
    - uses
        - building and testing prs
        - deploying app on release creation
        - adding labels to issues
    - nested worflows possible reference worflows in other workflows (Reuse workflows)

- Events
    - activity that triggers worflow run
    - triggered when pr, open and issue pushes a commit
    - triggered on schedule or by posting to a REST API or manually

- Jobs
    - set of steps in a worflow executed on the same runner
    - either a shell script thats executed or an action
    - data can be shared between steps in a job
        - like build followed by test
    - configure job dependencies with other jobs
        - dependant jobs are sequential
        - default no dependencies all jobs run in parallel
    - matrix to run the same job multiple times
        - useful to try different combination of variables (like OS and Language version)

- Actions
    - custom application that performs complicated, frequently executed tasks
        - dont have to write repetitive code in the worflow file
    - acts like a plugin
        - action for pulling form github, setup toolchain for building or auth
    - can write custom actions

- Runners
    - server that runs worflows when triggered
    - each runner runs a single job at a time
    - github provides runners

## Github Docker Build Action
- Uses the Git context
    - dont have to use the actions/checkout action Buildkit will checkout automatically
    - reference will based on the even that triggered the workflow
```
name: Docker login
uses: docker/login-activity@v3
with:
    username: ${{ vars.DDCKERHUB_USERNAME}}
    password: ${{ secrets.DOCKERHUB_PASSWORD}}
-
 name: Setup Docker Buildx
 uses: docker/setup-buildx-action@v3
-
 name: Build and push
 uses: docker/build-push-action@v3
 with:
    push: true
    tags: user/app:latest
```
- setup-buildx is only required to build multiplatform images, export cache etc.
- any file mutations in the steps before the build step will be ignored
 name: Build and push
 uses: docker/build-push-action@v3
 with:
    push: true
    tags: user/app:latest
    secrets:
        GIT_AUTH_TOKEN=${{secrets.MYTOKEN}}
- need this is building from a different repo

