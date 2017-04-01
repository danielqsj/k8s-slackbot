# K8S Slackbot
[![Go Report Card](https://goreportcard.com/badge/github.com/danielqsj/k8s-slackbot)](https://goreportcard.com/report/github.com/danielqsj/k8s-slackbot)

A slack bot built to control kubernetes cluster.

Image
-------------
[![Docker Stars](https://img.shields.io/docker/stars/danielqsj/k8s-slackbot.svg?style=flat)](https://hub.docker.com/r/danielqsj/k8s-slackbot/)
[![Docker Pulls](https://img.shields.io/docker/pulls/danielqsj/k8s-slackbot.svg?style=flat)](https://hub.docker.com/r/danielqsj/k8s-slackbot/)
[![Docker Automated build](https://img.shields.io/docker/automated/danielqsj/k8s-slackbot.svg?style=flat)](https://hub.docker.com/r/danielqsj/k8s-slackbot/)

This image is based on Alpine Linux image, which is only a 5MB image.
Download size of this image is only:

[![](https://images.microbadger.com/badges/version/danielqsj/k8s-slackbot.svg)](https://microbadger.com/images/danielqsj/k8s-slackbot "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/image/danielqsj/k8s-slackbot.svg)](https://microbadger.com/images/danielqsj/k8s-slackbot "Get your own image badge on microbadger.com")

Arguments
-------------
- **kubecfg-file** (*string*): Location of kubecfg file for access to kubernetes master service; --kube-master-url overrides the URL part of this; if neither this nor --kube-master-url are provided, defaults to service account tokens
- **bot-token** (*string*): Token of slack bot to use
- **debug** (*boolean*): Whether enable debug log

Usage
-------------
```
$ docker pull danielqsj/k8s-slackbot
$ docker run -v ~/.kube/config:/etc/kubernetes/kubeconfig danielqsj/k8s-slackbot --kubecfg-file=/etc/kubernetes/kubeconfig --bot-token=$(bot-token)
```
Then you can talk to your slack bot via slack direct message.
The command is same as [kubectl](https://kubernetes.io/docs/user-guide/kubectl/) .
Such as ``` kubectl get nodes``` .

**Enjoy it.**