# antora-build-checker-tools

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/56e9bcca0698420ebb2b7b3383ec0c9b)](https://www.codacy.com/app/settermjd/antora-build-checker-tools?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=settermjd/antora-build-checker-tools&amp;utm_campaign=Badge_Grade)

This is a small Go app, designed to help with the Antora build process at ownCloud.

At this stage, the ambitions of the app have not been that well defined, however, a start has been made.
Currently, it only checks an AsciiDoc file for a limited amount of reStructuredText formatting.
Specifically, this is only headers.
Formatting such as source includes, lists etc, is not yet covered.

## Installing

To install the app, compile it for your platform.

## Usage

The script, currently, reads a filename from stdin.

### Example

```
echo ./docs/modules/ROOT/pages/configuration/server/antivirus_configuration.adoc | antora-build-checker-tools
```
