### simple-file-server
[![build status](https://128.107.14.75/ci/projects/4/status.png?ref=master)](https://128.107.14.75/ci/projects/4?ref=master)

*trivial file server*

```
simple-file-server version built at: 2015.11.22.21.25.15.+00 commit: 0ef5559ash-2015-11-22T14:44:53-05:00
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

- [x] Releases: tags of the form release-0.1 release-100.3
- [x] Prod: tags of the form prod-0.1 prod-100.3
- [x] Version tags of the form: v0.0.0 or v8.1.11
- [x] Issue tags of the form: 'issue-text-without-spaces-dash-hash-issue-number-#10'
- [x] Canary tags of the form: 'issue-text-without-spaces-dash-hash-issue-number-#10'
- [x] Branches or tags of the form: release, release/name, canary, canary/fix-x, prod, prod/2015-12.12
- [x] 

```
  only:
    # v0.0.0
    - /^v[0-9]*\.[0-9]*\.[0-9]*$/
    # release-0.0
    - /^release-[0-9]*\.[0-9]*$/
    # issue-any....text...-#[0-9]*
    - /^issue-.*-#[0-9]*$/
    # canary-0.0
    - /^canary-[0-9]*\.[0-9]*$/
    # # wip-0.0
    # - /^wip-[0-9]*\.[0-9]$/
    - release
    - release/*
    - canary/*
    - canary


  except:
    - master
    - /^wip-[0-9]*\.[0-9]$/
    - wip/*
    - wip

```
