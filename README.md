### k8s-simple-file-server

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
