## Introduction

In this series I will talk about how I overengineered my Kubernetes cluster. I used to run some services on a single Raspberry Pi 4, but I was not very confident about the reliability or security of the setup.
For this reason, and out of pure interest in Kubernetes, I decided to move to a Kubernetes cluster.

Currently I have 2 Raspberry Pi (4 and 5) running a K3s cluster, with full support for metrics and alerts.
I will talk about the setup of the cluster, the services I run on it, and the monitoring and alerting setup. My goal with this post is to share my experience and maybe help someone who is thinking about setting up a similar cluster.
I have spent a lot of time trying to figure out many things, so I hope this post will save you some time or, at the very least, be an interesting read.

Everything I talk about in this series is available in **[this repository](https://github.com/guillembonet/home-k8s)**.

## Contents

<div class="mb-3">
    <a class="text-xl block mb-1" href="#the-cluster">1. The cluster setup</a>
    <a class="text-xl block mb-1" href="#ingress">2. Ingress</a>
</div>

### In the next posts...

#### 3. Storage

#### 4. Services

#### 5. Monitoring and alerting

#### 6. Backups and disaster recovery

<div>
    <a class="text-3xl my-3 block" id="the-cluster">The cluster setup</a>
</div>

As I mentioned, I decided to create a 2-node Kubernetes cluster using 2 Raspberry Pi. Given this starting point I had to make some decisions about the setup of the cluster. First of all, I decided from the very start to go with K3s, a lightweight Kubernetes distribution, which is a perfect fit. Secondly, I had to make this cluster accessible to the internet. For that, I wanted to avoid port-fowarding my router and exposing my IP address to the internet, so I decided to set-up a proxy server in the cloud. In order to avoid vendor lock-in, I decided to do it from scratch. I created a Wireguard VPN which connected the nodes to the proxy server. This way, I could access the cluster from anywhere, and I could also expose services to the internet without exposing my IP address. At this point we have 3 components: the proxy server and the 2 nodes, and the nodes are connected to the proxy server via a Wireguard tunnel.

### Setting up the proxy server

Now let's talk about the proxy server. The idea is very simple, we want to proxy HTTP, HTTPS and Kubernetes API connections. To create this proxy I used Traefik, and set it up to load-balance requests to the nodes for HTTP and HTTPS, and to direct requests to the Kubernetes API to the master node. The HTTPS and Kubernetes API proxies have to enable TLS passthrough, so that we are able to have TLS termination in the nodes rather than in the proxy. It is also important to enable the proxy protocol in Traefik, so we can get the real IP of the client in the logs.
You can find the specific configuration I used to do this in the repository I link to at the start of the post, inside of the proxy folder. 

As an additional (and totally unneeded) flex, I also set up a proxy using mitmproxy so that I am able to send requests from the Kubernetes nodes to update the domain's IP in the DNS with the proxy's IP rather than the actual node's IP.
Of course this is totally unnecessary since my VPS does not have a dynamic IP, but it was a fun exercise.
It is very important to not expose the mitmproxy to the internet, as bad actors could use it for malicious purposes. To avoid this, I just blocked the port in the firewall in my cloud provider dashboard, and that way it is only accessible through the tunnel from the nodes.

That's all the setup for the proxy server. Now we have a way to access the cluster from anywhere, and we can expose services to the internet without exposing our IP address.

### Setting up the nodes

The setup of the nodes is also pretty straight-forward. I just installed K3s on the master node with a few important configurations:

- 1. `--tls-san bunetz.dev`: This is the domain I use to access the cluster remotely, so I have to add it as a SAN in the TLS certificate.
- 2. `--tls-san 192.168.1.4`: This is the local IP (which I set statically) of the master node, so I can access the cluster from the local network.
- 3. `--disable=traefik`: Traefik is included in K3s by default, but I don't want to use it for reasons I will explain in the next section.
- 4. `--node-ip=10.1.10.1`: This is the local IP the node will advertise to the cluster. I set to use the Wireguard IP, so that the nodes communicate with each other through the tunnel.
- 5. `--flannel-iface=k8s-tunnel`: Flannel is in charge of redirecting traffic between nodes in K3s by default. With this flag we tell Flannel to use the Wireguard interface for this purpose so that all internal cluster traffic is end to end encrypted.

Of course, during this setup we will need to create the proper Wireguard configuration in the nodes to be able to communicate between them and to the proxy server. For this configuration it is important to activate keep alives in the Wireguard configuration in order to be able to communicate with the proxy server even if the nodes are behind a NAT (which is the case for most home routers) without having to set up port forwarding. A more detailed explanation of how to create and set up the Wireguard configuration can be found in the repository which I linked to at the start of the post.

<div>
    <a class="text-3xl my-3 block" id="ingress">Ingress</a>
</div>

The ingress is probably one of the most over-engineered parts of the whole setup. First of all, from the very start I decided it was very important to me to be able to see real client IPs in order to be more aware of what is going on.
This way I can have alerts if some particular IP is calling the cluster a lot in an effort to brute-force some password, for example. Another important thing for me was to have TLS encrypted traffic all the way to the nodes, so if someone could gain access to the proxy, he would still not be able to see my traffic which might contain passwords or other sensitive information.

For these reasons, I found I needed to disable Traefik, since TLS termination in Traefik happens in the end node, requests which arrived to one node and were forwarded to the other were losing their forwarded IP header because Flannel, which is in charge of redirecting traffic between nodes, does not support the proxy protocol.
Most likely, there is a way to replace Flannel with another CNI plugin which does support it or to have TLS termination as soon as traffic reaches any node, but since I struggled to find any documentation and I had experience with Nginx, I decided to go with it.

To install ingress-nginx as the cluster ingress I used Helm and applied some configurations. Here comes the over-engineering:

1. `enable-modsecurity: "true"`: I enabled ModSecurity to have a Web Application Firewall (WAF) in front of my services. It blocks weird requests which most likely come from bots that try to find credentials or vulnerabilities in the services, seems like people don't have a lot to do these days and like to do these things instead of just getting a job.

## What's next?

```go
if err != nil {
    return fmt.Errorf("error: %w", err)
}
```

<div><img src="/blog/assets/kubernetes.png" alt="Kubernetes logo" title="a title" width="200" height="200" /></div>
