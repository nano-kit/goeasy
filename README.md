# goeasy

Build the web real-time messaging systems in lightning speed.

快速打造您的web实时通讯体系。

---

# Setup a SaaS

This project shows a detailed procedure on how to setup a SaaS system:

## You need a Cloud VM

The convenient way to get a Cloud VM is use [Droplets](https://m.do.co/c/9ad1e150c20e). Actually any cloud provider is okay, as we only need these cloud component:

* VM; this is where you deploy your service and application. 1C1G is enough.
* Firewall; to only allow access from TCP port 22, 80 and 443.
* Domain: brought from a domain name registrar. I use [Aliyun](https://wanwang.aliyun.com/domain), which may has good enough discount.

Any other cloud components are not going to be used. We will get the functionality of those either by open source component or by implement them in Golang ourselves.

The goal is to keep the cost minimum, as well as to keep the system simple enough. Total cost could be:

* VM; $5 per month (Droplets) or ¥99 per year
* Domain; ¥26 per year

## Be careful of security

Once you have a Cloud VM. The first thing is security provisioning. The Internet has full of hackers. your VM is vulnerable.

The first step is the create a normal user (non-root).

```
# adduser xxxx
New password: ****
```

Then, prohibit root from remote login by edit `/etc/ssh/sshd_config`

```
PermitRootLogin no
```

This is a basic step to prevent root password attack, as the username is unknown to attacker. It is simpler than setup a public key authentication. Keep in mind, the more your provisioning, the more management effort. If you do not want to manage the key pair, take my recommendation.

Go to the Cloud Firewall. Add a rule says only allow in-bound traffic from TCP port 22, 80 and 443. This is to narrow down the attack surface, so that you will not expose database or other critical service on the Internet accidentally.

Note: If you host your service on Aliyun VM, you should do ICP 备案 for your domain. Otherwise [TLS handshake could be failed](https://developer.aliyun.com/article/708243). Use Droplets to skip ICP 备案.

## Prepare a build environment

Why you need a build environment on VM?

Because we do not has a CI/CD process. The minimum alternative is:

1. pull your code from repository
1. build on the VM
1. start the service

So straightforward, right? And it is so easy to fix and run.

Let's say you are using the Ubuntu, to install Golang

```bash
sudo apt-get install golang-1.16
export PATH=$HOME/go/bin:/usr/lib/go-1.16/bin:$PATH
# add this if you are behind the GFW
export GOPROXY=https://goproxy.io,direct
```

If you are like me, using [go-micro](https://github.com/nano-kit/go-micro/) to develop services, you can setup additional tools following the instruction [Getting Started](https://nano-kit.github.io/go-micro-in-action/getting-started.html).

```
$ ls go/bin
micro  protoc-gen-go  protoc-gen-micro  protodoc
```

## Build the service

Now, it is the time for `goeasy`, aka this project, as well as a project template and best practice. The project directory is described by [ProjectLayout](ProjectLayout.md).

To generate protocol files and static assets bundle,

```
go generate -x ./...
```

Then build

```
go build
```

Please be aware what your build is an API service. The data transferred from front end to back end service should be minimal to save bandwidth charging. The front end is usually an SPA powered by [Vuetify](https://vuetifyjs.com/en/). You should deploy the front end to some sort of CDN, but not this VM. This is another story.

## Deploy the service

Use systemd to supervise and auto-start the service on OS boot.

Just put this file under `/etc/systemd/system/goeasy.service`

```
[Unit]
Description=goeasy
Requires=network-online.target
After=network-online.target

[Service]
Type=simple
User=root
Group=root
Restart=always
RestartSec=5
ExecStart=/opt/goeasy/bin/goeasy

[Install]
WantedBy=multi-user.target
```

### Configuration

Put `serverinit.yaml` under the same directory as goeasy, that is `/opt/goeasy/bin`. The initial config data includes but not limited to serving address, namespace and database address, which are parameters that should be known before server starts and can not be changed until server stops.

### Logging

Log files are under directory `/opt/goeasy/log`. They are rotated by logrotate.

### Metrics

Install Prometheus by put this file under `/etc/systemd/system/prometheus.service`

```
[Unit]
Description=Prometheus
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
User=huanghao
Group=huanghao
Restart=always
RestartSec=5
ExecReload=/bin/kill -HUP $MAINPID
ExecStart=/opt/prometheus/prometheus \
  --config.file=/opt/prometheus/prometheus.yml \
  --storage.tsdb.path=/opt/prometheus/data \
  --web.console.templates=/opt/prometheus/consoles \
  --web.console.libraries=/opt/prometheus/console_libraries
SyslogIdentifier=prometheus

[Install]
WantedBy=multi-user.target
```

## Play with the service
