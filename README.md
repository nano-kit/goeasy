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
1.
