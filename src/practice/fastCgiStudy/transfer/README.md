## fast-cgi的翻译
* 原文为 https://fast-cgi.github.io/spec

## 章节目录
- [ ] 1. Introduction
- [ ] 2. Initial process state
- [ ] 3. Protocol basics
- [ ] 4. Management record types
- [ ] 5. Application record types
- [ ] 6. Roles
- [ ] 7. Errors
- [ ] 8. Types and constants


## 译文内容
>
1. Introduction

FastCGI is an open extension to CGI that provides high performance for all Internet applications without the penalties of Web server APIs.

This specification has narrow goal: to specify, from an application perspective, the interface between a FastCGI application and a Web server that supports FastCGI. Many Web server features related to FastCGI, e.g. application management facilities, have nothing to do with the application to Web server interface, and are not described here.

This specification is for Unix (more precisely, for POSIX systems that support Berkeley Sockets). The bulk of the specification is a simple communications protocol that is independent of byte ordering and will extend to other systems.

We’ll introduce FastCGI by comparing it with conventional Unix implementations of CGI/1.1. FastCGI is designed to support long-lived application processes, i.e. application servers. That’s a major difference compared with conventional Unix implementations of CGI/1.1, which construct an application process, use it respond to one request, and have it exit.

The initial state of a FastCGI process is more spartan than the initial state of a CGI/1.1 process, because the FastCGI process doesn’t begin life connected to anything. It doesn’t have the conventional open files stdin, stdout, and stderr, and it doesn’t receive much information through environment variables. The key piece of initial state in a FastCGI process is a listening socket, through which it accepts connections from a Web server.

After a FastCGI process accepts a connection on its listening socket, the process executes a simple protocol to receive and send data. The protocol serves two purposes. First, the protocol multiplexes a single transport connection between several independent FastCGI requests. This supports applications that are able to process concurrent requests using event-driven or multi-threaded programming techniques. Second, within each request the protocol provides several independent data streams in each direction. This way, for instance, both stdout and stderr data pass over a single transport connection from the application to the Web server, rather than requiring separate pipes as with CGI/1.1.

A FastCGI application plays one of several well-defined roles. The most familiar is the Responder role, in which the application receives all the information associated with an HTTP request and generates an HTTP response; that’s the role CGI/1.1 programs play. A second role is Authorizer, in which the application receives all the information associated with an HTTP request and generates an authorized/unauthorized decision. A third role is Filter, in which the application receives all the information associated with an HTTP request, plus an extra stream of data from a file stored on the Web server, and generates a “filtered” version of the data stream as an HTTP response. The framework is extensible so that more FastCGI can be defined later.

In the remainder of this specification the terms “FastCGI application,” “application process,” or “application server” are abbreviated to “application” whenever that won’t cause confusion.


### 1.介绍
* `FastCGI` 是一个开放的 `CGI` 扩展，它为整个互联网应用提供了高性能而弥补了 Web 服务 APIs 的缺陷。
* 这个规范有很明确的目标：从应用程序的角度看，为了指导做一个支持 FastCGI 的 FastCGI 应用和一个 Web 服务。很多 Web 服务器和 FastCGI 有关联，
* 比如应用管理工具，没有好的方案来开发 Web 服务接口，以及一些其他此处没有描述的东西。
* 这个说规范是针对 Unix （更准确的说是针对支持`Berkeley Sockets`的 POSIX 系统）。这个规范的大部分是描述一个简单的通信协议，它不依赖于字节序，和系统的扩展。
* 我们将通过对比传统的 Unix 的 `CGI/1.1` 的实现来介绍 FastCGI 。 FastCGI 被设计成支持长连接，例如应用服务器。这是和传统的Unix 的 `CGI/1.1` 的实现的主要区别， `CGI/1.1` 创建一个进程，用它来响应一个请求，然后退出。
* 一个 FastCGI 的初始化状态比一个 `CGI/1.1` 进程的初始化状态更加简洁，因为 FastCGI 进程在声明周期的开始时，没有连接任何东西。它没有按照惯例那样打开 `stdin` ， `stdout` ， `stderr` 这3个文件描述符，并且它没有接受很多来自环境变量的信息。 FastCGI 进程在初始化状态时的关键一点是监听一个 `socket` ，通过它接收来自 Web 服务的连接。

