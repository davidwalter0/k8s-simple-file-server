### k8s-simple-file-server
[![build status](https://128.107.14.75/ci/projects/4/status.png?ref=master)](https://128.107.14.75/ci/projects/4?ref=master)

*trivial file server*

```
k8s-simple-file-server version built at: 2015.11.22.21.25.15.+00 commit: 0ef5559ash-2015-11-22T14:44:53-05:00
This can be configured with the following environment variables.
SIMPLE_FILE_SERVER_PATH to set the path to serve. The default path is /target.
SIMPLE_FILE_SERVER_PORT to set the port to serve on. The default port is 8080
SIMPLE_FILE_SERVER_HOST to set the host interface to serve. The default is all.
SIMPLE_FILE_SERVER_PATH not set or using .. to subvert paths, using /target instead
SIMPLE_FILE_SERVER_PORT not set, using 8080
SIMPLE_FILE_SERVER_HOST not set, default bind all
```

*Defaults*


```
PATH to serve  :/var/opt/simple-file-server/data
PORT on which  :8080
HOST interface :0.0.0.0
```

---
### Gitlab option tests to specify build targets 

*Build only target branches or tags*

```
  only:
    # v0.0.0
    - /^v[0-9]*\.[0-9]*\.[0-9]*$/
    # release-0.0
    - /^release-[0-9]*\.[0-9]$/
    # issue-0.0
    - /^issue-[0-9]*\.[0-9]$/
    # canary-0.0
    - /^canary-[0-9]*\.[0-9]$/
    # # wip-0.0
    # - /^wip-[0-9]*\.[0-9]$/
    - release
    - release/*
    - canary/*
    - canary
    # - wip/*
    # - wip

  except:
    - master
    - /^wip-[0-9]*\.[0-9]$/
    - wip/*
    - wip

```
