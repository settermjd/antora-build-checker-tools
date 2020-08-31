我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# antora-build-checker-tools

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/56e9bcca0698420ebb2b7b3383ec0c9b)](https://www.codacy.com/app/settermjd/antora-build-checker-tools?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=settermjd/antora-build-checker-tools&amp;utm_campaign=Badge_Grade)

This is a small Go app, designed to help with the Antora build process at ownCloud.

At this stage, the ambitions of the app have not been that well defined, however, a start has been made.
Currently, it only checks an AsciiDoc file for a limited amount of reStructuredText formatting.
Specifically, this is only headers.
Formatting such as source includes, lists etc, is not yet covered.

## Installing

To install the app, compile it for your platform.

## Usage

The script, currently, reads a filename from stdin.

### Example

```
echo ./docs/modules/ROOT/pages/configuration/server/antivirus_configuration.adoc | antora-build-checker-tools
```
