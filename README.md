# mattermost-plugin-sentry

[![Go](https://img.shields.io/github/go-mod/go-version/mizyind/mattermost-plugin-sentry?style=for-the-badge&label=&color=00add8&logo=go&logoColor=fff)](https://golang.org)
[![Mattermost](https://img.shields.io/badge/mattermost-0072c6?style=for-the-badge&logo=mattermost&logoColor=fff)](https://mattermost.com)
[![Sentry](https://img.shields.io/badge/sentry-584674?style=for-the-badge&logo=sentry&logoColor=fff)](https://sentry.io)

Sends Sentry notifications to Mattermost.

This plugin will decode the following notification payload from Sentry:
```json
{
  "project": "mattermost-plugin-sentry",
  "url": "http://sentry.example.com/mattermost-plugin-sentry/issues/1",
  "event": {
    "title": "Error: This is an error example",
    "location": "~/miZyind/mattermost-plugin-sentry/example.ts"
  }
}
```
And post the notification to the specific channel:
![Thumb-1](https://github.com/miZyind/mattermost-plugin-sentry/blob/master/images/thumb-1.png)

## 🔮 Usage

Add the callback url `https://<MATTERMOST-HOST>/plugins/com.github.mizyind.mattermost-plugin-sentry/webhook?channel=<MATTERMOST-CHANNEL-ID>` to Sentry's `WebHooks` plugin:
![Thumb-2](https://github.com/miZyind/mattermost-plugin-sentry/blob/master/images/thumb-2.png)

That's it!

## 🔨 Build

```bash
# Uses makefile to compile go binary file and generates a tar bundle of the plugin for install.
$ make
# Output "plugin.tar.gz" to the current working directory.
```

## 🖋 Author

miZyind <mizyind@gmail.com>

## 📇 License

Licensed under the [MIT](LICENSE) License.
