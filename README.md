### Customized Chromium

The main goal of this project is to generate a shortcut of a Chromium instance to anonymify my presence on internet by using a custom user agent, disabling the [user agent client hint](https://www.chromium.org/updates/ua-ch) and using a proxy.

The supported customizations are:
  - [Proxy server](https://peter.sh/experiments/chromium-command-line-switches/#proxy-server)
  - [Proxy bypass list](https://peter.sh/experiments/chromium-command-line-switches/#proxy-bypass-list)
  - [Host resolver rules](https://peter.sh/experiments/chromium-command-line-switches/#host-resolver-rules)
  - [User agent](https://peter.sh/experiments/chromium-command-line-switches/#user-agent)
  - [Disable features](https://peter.sh/experiments/chromium-command-line-switches/#disable-features)

All of these Chromium-specific flags are customizable by editing a yaml configuration file.

#### Example of configuration
```yaml
--- 
chromium: 
  path: "C:\\Program Files\\Chromium\\Application\\chrome.exe"
  settings: 
    disable_features: 
      - UserAgentClientHint
    host_resolver_rules: "MAP * ~NOTFOUND , EXCLUDE ::1"
    proxy: "socks5://proxy.net:9000"
    proxy_bypass: 
      - "127.0.0.1:*"
    user_agent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.130 Safari/537.36"
```
This configuration will generate a shortcut with the following target
```
"C:\Program Files\Chromium\Application\chrome.exe" --proxy-server="socks5://proxy.net:9000" --proxy-bypass-list="127.0.0.1:*" --host-resolver-rules="MAP * ~NOTFOUND , EXCLUDE ::1" --user-agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.130 Safari/537.36" --disable-features="UserAgentClientHint"
```

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
