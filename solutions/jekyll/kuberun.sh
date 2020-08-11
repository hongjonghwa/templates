kubectl run -i --rm --tty jekyll --overrides='
{
  "apiVersion": "batch/v1",
  "spec": {
    "template": {
      "spec": {
        "containers": [
          {
            "name": "jekyll",
            "image": "jekyll:3.8",
            "args": [
              "bash"
            ],
            "stdin": true,
            "stdinOnce": true,
            "tty": true,
            "volumeMounts": [{
              "mountPath": "/Users/user/dev/hongjonghwa.github.io",
              "name": "store"
            }]
          }
        ],
        "volumes": [{
          "name":"store",
          "emptyDir":{}
        }]
      }
    }
  }
}
'  --image=jekyll:3.8 --restart=Never -- bash
