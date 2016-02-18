# slack-storage

A simple API to interact with Slack to provide a key/value storage.

## Can I have a use case?

Yes, you can.

## Is it any good?

[Yes.](https://news.ycombinator.com/item?id=3067434)

## You've got me hooked. How do I install it?

You need to have [Go](https://golang.org/) and [GB](https://getgb.io/) installed. Once they're installed, it's simple.

```bash
gb build all
```

It'll create a binary inside a `bin` folder in the root of the repository.

You're done. Run that to run it.

## What about configuring it?

Configure away! All configuration is done by flags.

### Production

You're probably going to want to use this flag.

```bash
-production
```

### Slack Token

If you don't pass an token, anyone can use your server. Probably not the best idea.

```bash
-token SL4CKT0K3NG03SH3R3
```

### Host

The default host is `[::]` which will listen on both IPv4 and IPv6. What if you want to just listen on localhost though?

```bash
-host 127.0.0.1
```

### Port

The default port is `7043`. Why? I always keyboard smash to pick a port when starting a new project. Feel free to change it.

```bash
-port 3000
```
