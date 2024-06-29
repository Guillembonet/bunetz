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

As an additional (and totally unneeded) flex, I also set up a proxy using mitmproxy so that I am able to send requests from the Kubernetes nodes to update the domain IP in the DNS. Of course this is totally unnecessary since my VPS does not have a dynamic IP, but it was a fun exercise. It is important to not expose the mitmproxy to the internet, as bad actors could use it for malicious purposes. To avoid this, I just blocked the port in the firewall in my cloud provider dashboard, and that way it is only accessible through the tunnel from the nodes.

That's all the setup for the proxy server. Now we have a way to access the cluster from anywhere, and we can expose services to the internet without exposing our IP address.

### Setting up the nodes

<div>
    <a class="text-3xl my-3 block" id="ingress">Ingress</a>
</div>

## What's next?

```go
if err != nil {
    return fmt.Errorf("error: %w", err)
}
```

<div><img src="/blog/assets/kubernetes.png" alt="Kubernetes logo" title="a title" width="200" height="200" /></div>
