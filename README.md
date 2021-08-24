### Customized Chromium

The main goal of this project is to generate a shortcut of a Chromium instance to anonymify my presence on internet by using a custom user agent, disabling the [user agent client hint](https://www.chromium.org/updates/ua-ch) and using a proxy.

The supported customizations are:
  - [Proxy server](https://peter.sh/experiments/chromium-command-line-switches/#proxy-server)
  - [Proxy bypass list](https://peter.sh/experiments/chromium-command-line-switches/#proxy-bypass-list)
  - [Host resolver rules](https://peter.sh/experiments/chromium-command-line-switches/#host-resolver-rules)
  - [User agent](https://peter.sh/experiments/chromium-command-line-switches/#user-agent)
  - [Disable features](https://peter.sh/experiments/chromium-command-line-switches/#disable-features)

All of these Chromium-specific flags are customizable by editing a yaml configuration file.

#### Build
```shell
go build
```

#### How does it work
> It only works on Windows for now

There are no arguments you can use.

The first time you're using this software, a configuration file will be generated for you to customize the flags passed to Chromium.
Next, you'll have to specify the location of your Chromium executable in the aforementioned configuration file, along with the flags you wish to use.

#### Things to note
> It should also work with Chrome and all Chromium-based softwares
